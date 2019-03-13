// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package pointerdb

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"storj.io/storj/pkg/identity"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storage/meta"
	"storj.io/storj/storage"
)

// Service structure
type Service struct {
	logger     *zap.Logger
	DB         storage.KeyValueStore
	allocation *AllocationSigner
	cache      *overlay.Cache
}

// NewService creates new pointerdb service
func NewService(logger *zap.Logger, db storage.KeyValueStore, allocation *AllocationSigner, cache *overlay.Cache) *Service {
	return &Service{logger: logger, DB: db, allocation: allocation, cache: cache}
}

// Put puts pointer to db under specific path
func (s *Service) Put(path string, pointer *pb.Pointer) (err error) {
	// Update the pointer with the creation date
	pointer.CreationDate = ptypes.TimestampNow()

	pointerBytes, err := proto.Marshal(pointer)
	if err != nil {
		return err
	}

	// TODO(kaloyan): make sure that we know we are overwriting the pointer!
	// In such case we should delete the pieces of the old segment if it was
	// a remote one.
	if err = s.DB.Put([]byte(path), pointerBytes); err != nil {
		return err
	}

	return nil
}

// Get gets pointer from db
func (s *Service) Get(path string) (pointer *pb.Pointer, err error) {
	pointerBytes, err := s.DB.Get([]byte(path))
	if err != nil {
		return nil, err
	}

	pointer = &pb.Pointer{}
	err = proto.Unmarshal(pointerBytes, pointer)
	if err != nil {
		return nil, errs.New("error unmarshaling pointer: %v", err)
	}

	return pointer, nil
}

// List returns all Path keys in the pointers bucket
func (s *Service) List(prefix string, startAfter string, endBefore string, recursive bool, limit int32,
	metaFlags uint32) (items []*pb.ListResponse_Item, more bool, err error) {

	var prefixKey storage.Key
	if prefix != "" {
		prefixKey = storage.Key(prefix)
		if prefix[len(prefix)-1] != storage.Delimiter {
			prefixKey = append(prefixKey, storage.Delimiter)
		}
	}

	rawItems, more, err := storage.ListV2(s.DB, storage.ListOptions{
		Prefix:       prefixKey,
		StartAfter:   storage.Key(startAfter),
		EndBefore:    storage.Key(endBefore),
		Recursive:    recursive,
		Limit:        int(limit),
		IncludeValue: metaFlags != meta.None,
	})
	if err != nil {
		return nil, false, err
	}

	for _, rawItem := range rawItems {
		items = append(items, s.createListItem(rawItem, metaFlags))
	}
	return items, more, nil
}

// createListItem creates a new list item with the given path. It also adds
// the metadata according to the given metaFlags.
func (s *Service) createListItem(rawItem storage.ListItem, metaFlags uint32) *pb.ListResponse_Item {
	item := &pb.ListResponse_Item{
		Path:     rawItem.Key.String(),
		IsPrefix: rawItem.IsPrefix,
	}
	if item.IsPrefix {
		return item
	}

	err := s.setMetadata(item, rawItem.Value, metaFlags)
	if err != nil {
		s.logger.Warn("err retrieving metadata", zap.Error(err))
	}
	return item
}

// getMetadata adds the metadata to the given item pointer according to the
// given metaFlags
func (s *Service) setMetadata(item *pb.ListResponse_Item, data []byte, metaFlags uint32) (err error) {
	if metaFlags == meta.None || len(data) == 0 {
		return nil
	}

	pr := &pb.Pointer{}
	err = proto.Unmarshal(data, pr)
	if err != nil {
		return err
	}

	// Start with an empty pointer to and add only what's requested in
	// metaFlags to safe to transfer payload
	item.Pointer = &pb.Pointer{}
	if metaFlags&meta.Modified != 0 {
		item.Pointer.CreationDate = pr.GetCreationDate()
	}
	if metaFlags&meta.Expiration != 0 {
		item.Pointer.ExpirationDate = pr.GetExpirationDate()
	}
	if metaFlags&meta.Size != 0 {
		item.Pointer.SegmentSize = pr.GetSegmentSize()
	}
	if metaFlags&meta.UserDefined != 0 {
		item.Pointer.Metadata = pr.GetMetadata()
	}

	return nil
}

// Delete deletes from item from db
func (s *Service) Delete(path string) (err error) {
	return s.DB.Delete([]byte(path))
}

// Iterate iterates over items in db
func (s *Service) Iterate(prefix string, first string, recurse bool, reverse bool, f func(it storage.Iterator) error) (err error) {
	opts := storage.IterateOptions{
		Prefix:  storage.Key(prefix),
		First:   storage.Key(first),
		Recurse: recurse,
		Reverse: reverse,
	}
	return s.DB.Iterate(opts, f)
}

// PayerBandwidthAllocation gets payer bandwidth allocation message
func (s *Service) PayerBandwidthAllocation(ctx context.Context, action pb.BandwidthAction) (resp *pb.OrderLimit, err error) {
	pi, err := identity.PeerIdentityFromContext(ctx)
	if err != nil {
		return nil, err
	}
	pba, err := s.allocation.PayerBandwidthAllocation(ctx, pi, action)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return pba, nil
}

// GetNodes returns the nodes associated with the pointer
func (s *Service) GetNodes(ctx context.Context, pointer *pb.Pointer) (nodes []*pb.Node) {
	for _, piece := range pointer.GetRemote().GetRemotePieces() {
		node, err := s.cache.Get(ctx, piece.NodeId)
		if err != nil {
			s.logger.Error("Error getting node from cache", zap.String("ID", piece.NodeId.String()), zap.Error(err))
			continue
		}
		nodes = append(nodes, node)
	}

	for _, v := range nodes {
		if v != nil {
			v.Type.DPanicOnInvalid("pdb server Get")
		}
	}
	return nodes
}
