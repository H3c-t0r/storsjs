// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package server

import (
	"io"
	"log"

	"github.com/zeebo/errs"
	"storj.io/storj/pkg/piecestore"
	"storj.io/storj/pkg/utils"
	pb "storj.io/storj/protos/piecestore"
)

// OK - Success!
const OK = "OK"

// StoreError is a type of error for failures in Server.Store()
var StoreError = errs.Class("store error")

// Store incoming data using piecestore
func (s *Server) Store(reqStream pb.PieceStoreRoutes_StoreServer) error {
	// Receive id/ttl
	recv, err := reqStream.Recv()
	if err != nil {
		return StoreError.Wrap(err)
	}
	if recv == nil {
		return StoreError.New("Error receiving Piece metadata")
	}

	pd := recv.GetPiecedata()
	if pd == nil {
		return StoreError.New("PieceStore message is nil")
	}

	log.Printf("Storing %s...", pd.GetId())

	if pd.GetId() == "" {
		return StoreError.New("Piece ID not specified")
	}

	// If we put in the database first then that checks if the data already exists
	if err = s.DB.AddTTLToDB(pd.GetId(), pd.GetTtl()); err != nil {
		return StoreError.New("Failed to write expiration data to database")
	}

	total, err := s.storeData(reqStream, pd.GetId())
	if err != nil {
		return err
	}

	log.Printf("Successfully stored %s.", pd.GetId())

	return reqStream.SendAndClose(&pb.PieceStoreSummary{Message: OK, TotalReceived: int64(total)})
}

func (s *Server) storeData(stream pb.PieceStoreRoutes_StoreServer, id string) (total int64, err error) {
	// Delete data if we error
	defer func(err error) {
		if err != nil && err != io.EOF {
			if err = s.deleteByID(id); err != nil {
				log.Printf("Failed on deleteByID in Store: %s", err.Error())
			}
		}
	}(err)

	// Initialize file for storing data
	storeFile, err := pstore.StoreWriter(id, s.DataDir)
	if err != nil {
		return 0, err
	}

	defer utils.Close(storeFile)

	reader := NewStreamReader(s, stream)
	total, err = io.Copy(storeFile, reader)

	defer func() {
		err := s.DB.WriteBandwidthAllocToDB(reader.bandwidthAllocation)
		if err != nil {
			log.Printf("WriteBandwidthAllocToDB Error: %s\n", err.Error())
		}
	}()

	if err != nil && err != io.EOF {
		return 0, err
	}

	return total, nil
}
