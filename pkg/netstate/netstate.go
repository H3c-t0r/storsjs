// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package netstate

import (
	"context"

	"go.uber.org/zap"

	proto "storj.io/storj/protos/netstate"
	"storj.io/storj/storage/boltdb"
)

// Server implements the network state RPC service
type Server struct {
	DB     DB
	logger *zap.Logger
}

// NewServer creates instance of Server
func NewServer(db DB, logger *zap.Logger) *Server {
	return &Server{
		DB:     db,
		logger: logger,
	}
}

// DB interface allows more modular unit testing
// and makes it easier in the future to substitute
// db clients other than bolt
type DB interface {
	Put(boltdb.File) error
	Get([]byte) (boltdb.File, error)
	List() ([][]byte, error)
	Delete([]byte) error
}

// Put formats and hands off a file path to be saved to boltdb
func (s *Server) Put(ctx context.Context, filepath *proto.FilePath) (*proto.PutResponse, error) {
	s.logger.Debug("entering NetState.Put(...)")

	file := boltdb.File{
		Path:  []byte(filepath.Path),
		Value: []byte(filepath.SmallValue),
	}

	if err := s.DB.Put(file); err != nil {
		s.logger.Error("err putting file", zap.Error(err))
		return nil, err
	}
	s.logger.Debug("the file was put to the db")

	return &proto.PutResponse{
		Confirmation: "success",
	}, nil
}

// Get formats and hands off a file path to get from boltdb
func (s *Server) Get(ctx context.Context, filepath *proto.FilePath) (*proto.GetResponse, error) {
	s.logger.Debug("entering NetState.Get(...)")

	fileInfo, err := s.DB.Get(filepath.Path)
	if err != nil {
		s.logger.Error("err getting file", zap.Error(err))
		return nil, err
	}

	return &proto.GetResponse{
		Content: fileInfo.Value,
	}, nil
}

// List calls the bolt client's List function and returns all file paths
func (s *Server) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	s.logger.Debug("entering NetState.List(...)")

	filePaths, err := s.DB.List()
	if err != nil {
		s.logger.Error("err listing file paths", zap.Error(err))
		return nil, err
	}

	s.logger.Debug("file paths retrieved")
	return &proto.ListResponse{
		// filePaths is an array of byte arrays
		Filepaths: filePaths,
	}, nil
}

// Delete formats and hands off a file path to delete from boltdb
func (s *Server) Delete(ctx context.Context, filepath *proto.FilePath) (*proto.DeleteResponse, error) {
	s.logger.Debug("entering NetState.Delete(...)")

	err := s.DB.Delete([]byte(filepath.Path))
	if err != nil {
		s.logger.Error("err deleting file", zap.Error(err))
		return nil, err
	}
	s.logger.Debug("file deleted")
	return &proto.DeleteResponse{
		Confirmation: "success",
	}, nil
}
