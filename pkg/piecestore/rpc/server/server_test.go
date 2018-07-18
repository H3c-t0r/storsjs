// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package server

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"storj.io/storj/pkg/piecestore"
	"storj.io/storj/pkg/piecestore/rpc/server/psdb"
	pb "storj.io/storj/protos/piecestore"
)

var db *sql.DB
var s Server
var c pb.PieceStoreRoutesClient

func TestPiece(t *testing.T) {
	var testID = "11111111111111111111"

	// simulate piece stored with farmer
	file, err := pstore.StoreWriter(testID, s.DataDir)
	if err != nil {
		return
	}

	// Close when finished
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader([]byte("butts")))

	if err != nil {
		t.Errorf("Error: %v\nCould not create test piece", err)
		return
	}

	defer pstore.Delete(testID, s.DataDir)

	// set up test cases
	tests := []struct {
		id         string
		size       int64
		expiration int64
		err        string
	}{
		{ // should successfully retrieve piece meta-data
			id:         "11111111111111111111",
			size:       5,
			expiration: 9999999999,
			err:        "",
		},
		{ // server should err with invalid id
			id:         "123",
			size:       5,
			expiration: 9999999999,
			err:        "rpc error: code = Unknown desc = argError: Invalid id length",
		},
		{ // server should err with nonexistent file
			id:         "22222222222222222222",
			size:       5,
			expiration: 9999999999,
			err:        fmt.Sprintf("rpc error: code = Unknown desc = stat %s: no such file or directory", path.Join(os.TempDir(), "/test-data/3000/22/22/2222222222222222")),
		},
	}

	for _, tt := range tests {
		t.Run("should return expected PieceSummary values", func(t *testing.T) {

			// simulate piece TTL entry
			_, err = db.Exec(fmt.Sprintf(`INSERT INTO ttl (id, created, expires) VALUES ("%s", "%d", "%d")`, tt.id, 1234567890, tt.expiration))
			if err != nil {
				t.Errorf("Error: %v\nCould not make TTL entry", err)
				return
			}

			defer db.Exec(fmt.Sprintf(`DELETE FROM ttl WHERE id="%s"`, tt.id))

			req := &pb.PieceId{Id: tt.id}
			resp, err := c.Piece(context.Background(), req)

			if len(tt.err) > 0 {

				if err != nil {
					if err.Error() == tt.err {
						return
					}
				}

				t.Errorf("\nExpected: %s\nGot: %v\n", tt.err, err)
				return
			}

			if err != nil && tt.err == "" {
				t.Errorf("\nExpected: %s\nGot: %v\n", tt.err, err)
				return
			}

			if resp.Id != tt.id || resp.Size != tt.size || resp.Expiration != tt.expiration {
				t.Errorf("Expected: %v, %v, %v\nGot: %v, %v, %v\n", tt.id, tt.size, tt.expiration, resp.Id, resp.Size, resp.Expiration)
				return
			}

			// clean up DB entry
			_, err = db.Exec(fmt.Sprintf(`DELETE FROM ttl WHERE id="%s"`, tt.id))
			if err != nil {
				t.Errorf("Error cleaning test DB entry")
				return
			}
		})
	}
}

func TestRetrieve(t *testing.T) {
	var testID = "11111111111111111111"

	// simulate piece stored with farmer
	file, err := pstore.StoreWriter(testID, s.DataDir)
	if err != nil {
		return
	}

	// Close when finished
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader([]byte("butts")))
	if err != nil {
		t.Errorf("Error: %v\nCould not create test piece", err)
		return
	}
	defer pstore.Delete(testID, s.DataDir)
	defer db.Exec(fmt.Sprintf(`DELETE FROM bandwidth_agreements WHERE size="%d"`, 5))

	// set up test cases
	tests := []struct {
		id       string
		reqSize  int64
		respSize int64
		offset   int64
		content  []byte
		err      string
	}{
		{ // should successfully retrieve data
			id:       testID,
			reqSize:  5,
			respSize: 5,
			offset:   0,
			content:  []byte("butts"),
			err:      "",
		},
		{ // should successfully retrieve data
			id:       "11111111111111111111",
			reqSize:  -1,
			respSize: 5,
			offset:   0,
			content:  []byte("butts"),
			err:      "",
		},
		{ // server should err with invalid id
			id:       "123",
			reqSize:  5,
			respSize: 5,
			offset:   0,
			content:  []byte("butts"),
			err:      "rpc error: code = Unknown desc = argError: Invalid id length",
		},
		{ // server should err with nonexistent file
			id:       "22222222222222222222",
			reqSize:  5,
			respSize: 5,
			offset:   0,
			content:  []byte("butts"),
			err:      fmt.Sprintf("rpc error: code = Unknown desc = stat %s: no such file or directory", path.Join(os.TempDir(), "/test-data/3000/22/22/2222222222222222")),
		},
		{ // server should return expected content and respSize with offset and excess reqSize
			id:       "11111111111111111111",
			reqSize:  5,
			respSize: 4,
			offset:   1,
			content:  []byte("utts"),
			err:      "",
		},
		{ // server should return expected content with reduced reqSize
			id:       "11111111111111111111",
			reqSize:  4,
			respSize: 4,
			offset:   0,
			content:  []byte("butt"),
			err:      "",
		},
	}

	for _, tt := range tests {
		t.Run("should return expected PieceRetrievalStream values", func(t *testing.T) {
			stream, err := c.Retrieve(context.Background())

			// send piece database
			stream.Send(&pb.PieceRetrieval{PieceData: &pb.PieceRetrieval_PieceData{Id: tt.id, Size: tt.reqSize, Offset: tt.offset}})
			if err != nil {
				t.Errorf("Unexpected error: %v\n", err)
				return
			}

			// Send bandwidth bandwidthAllocation
			baSize := tt.reqSize
			if baSize <= 0 {
				baSize = 1024 * 32
			}
			stream.Send(&pb.PieceRetrieval{Bandwidthallocation: &pb.BandwidthAllocation{Signature: []byte{'A', 'B'}, Data: &pb.BandwidthAllocation_Data{Payer: "payer-id", Renter: "renter-id", Size: baSize}}})
			if err != nil {
				t.Errorf("Unexpected error: %v\n", err)
				return
			}

			resp, err := stream.Recv()

			// Send bandwidth bandwidthAllocation
			stream.Send(&pb.PieceRetrieval{Bandwidthallocation: &pb.BandwidthAllocation{Signature: []byte{'A', 'B'}, Data: &pb.BandwidthAllocation_Data{Payer: "payer-id", Renter: "renter-id", Size: tt.reqSize}}})

			if len(tt.err) > 0 {
				if err != nil {
					if err.Error() == tt.err {
						return
					}
				}
				t.Errorf("\nExpected: %s\nGot: %v\n", tt.err, err)
				return
			}

			// send piece database
			stream.Send(&pb.PieceRetrieval{Bandwidthallocation: &pb.BandwidthAllocation{Signature: []byte{'A', 'B'}, Data: &pb.BandwidthAllocation_Data{Payer: "payer-id", Renter: "renter-id", Size: baSize * 2}}})
			if err != nil {
				t.Errorf("Unexpected error: %v\n", err)
				return
			}

			if err != nil && tt.err == "" {
				t.Errorf("\nExpected: %s\nGot: %v\n", tt.err, err)
				return
			}

			if resp.Size != tt.respSize || bytes.Equal(resp.Content, tt.content) != true {
				t.Errorf("Expected: %v, %v\nGot: %v, %v\n", tt.respSize, tt.content, resp.Size, resp.Content)
				return
			}
		})
	}
}

func TestStore(t *testing.T) {
	tests := []struct {
		id            string
		size          int64
		ttl           int64
		offset        int64
		content       []byte
		message       string
		totalReceived int64
		err           string
	}{
		{ // should successfully store data
			id:            "11111111111111111111",
			ttl:           9999999999,
			content:       []byte("butts"),
			message:       "OK",
			totalReceived: 5,
			err:           "",
		},
		{ // should err with invalid id length
			id:            "butts",
			ttl:           9999999999,
			content:       []byte("butts"),
			message:       "",
			totalReceived: 0,
			err:           "rpc error: code = Unknown desc = argError: Invalid id length",
		},
	}

	for _, tt := range tests {
		t.Run("should return expected PieceStoreSummary values", func(t *testing.T) {

			stream, err := c.Store(context.Background())
			if err != nil {
				t.Errorf("Unexpected error: %v\n", err)
				return
			}

			// Write the buffer to the stream we opened earlier
			if err = stream.Send(&pb.PieceStore{Piecedata: &pb.PieceStore_PieceData{Id: tt.id, Ttl: tt.ttl}}); err != nil {
				t.Errorf("Unexpected error: %v\n", err)
				return
			}

			// Send Bandwidth Allocation Data
			msg := &pb.PieceStore{
				Piecedata: &pb.PieceStore_PieceData{Content: tt.content},
				Bandwidthallocation: &pb.BandwidthAllocation{
					Signature: []byte{'A', 'B'},
					Data: &pb.BandwidthAllocation_Data{
						Payer: "payer-id", Renter: "renter-id", Size: int64(len(tt.content)),
					},
				},
			}

			// Write the buffer to the stream we opened earlier
			if err = stream.Send(msg); err != nil {
				t.Errorf("Unexpected error: %v\n", err)
				return
			}

			resp, err := stream.CloseAndRecv()

			defer db.Exec(fmt.Sprintf(`DELETE FROM ttl WHERE id="%s"`, tt.id))

			if len(tt.err) > 0 {
				if err != nil {
					if err.Error() == tt.err {
						return
					}
				}
				t.Errorf("\nExpected: %s\nGot: %v\n", tt.err, err)
				return
			}
			if err != nil && tt.err == "" {
				t.Errorf("\nExpected: %s\nGot: %v\n", tt.err, err)
				return
			}

			if resp.Message != tt.message || resp.TotalReceived != tt.totalReceived {
				t.Errorf("Expected: %v, %v\nGot: %v, %v\n", tt.message, tt.totalReceived, resp.Message, resp.TotalReceived)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	// set up test cases
	tests := []struct {
		id      string
		message string
		err     string
	}{
		{ // should successfully delete data
			id:      "11111111111111111111",
			message: "OK",
			err:     "",
		},
		{ // should err with invalid id length
			id:      "123",
			message: "rpc error: code = Unknown desc = argError: Invalid id length",
			err:     "rpc error: code = Unknown desc = argError: Invalid id length",
		},
		{ // should return OK with nonexistent file
			id:      "22222222222222222223",
			message: "OK",
			err:     "",
		},
	}

	for _, tt := range tests {
		t.Run("should return expected PieceDeleteSummary values", func(t *testing.T) {

			// simulate piece stored with farmer
			file, err := pstore.StoreWriter(tt.id, s.DataDir)
			if err != nil {
				return
			}

			// Close when finished
			defer file.Close()

			_, err = io.Copy(file, bytes.NewReader([]byte("butts")))
			if err != nil {
				t.Errorf("Error: %v\nCould not create test piece", err)
				return
			}

			// simulate piece TTL entry
			_, err = db.Exec(fmt.Sprintf(`INSERT INTO ttl (id, created, expires) VALUES ("%s", "%d", "%d")`, tt.id, 1234567890, 1234567890))
			if err != nil {
				t.Errorf("Error: %v\nCould not make TTL entry", err)
				return
			}

			defer db.Exec(fmt.Sprintf(`DELETE FROM ttl WHERE id="%s"`, tt.id))

			defer pstore.Delete(tt.id, s.DataDir)

			req := &pb.PieceDelete{Id: tt.id}
			resp, err := c.Delete(context.Background(), req)
			if len(tt.err) > 0 {
				if err != nil {
					if err.Error() == tt.err {
						return
					}
				}
				t.Errorf("\nExpected: %s\nGot: %v\n", tt.err, err)
				return
			}
			if err != nil && tt.err == "" {
				t.Errorf("\nExpected: %s\nGot: %v\n", tt.err, err)
				return
			}

			if resp.Message != tt.message {
				t.Errorf("Expected: %v\nGot: %v\n", tt.message, resp.Message)
				return
			}

			// if test passes, check if file was indeed deleted
			filePath, err := pstore.PathByID(tt.id, s.DataDir)
			if _, err = os.Stat(filePath); os.IsNotExist(err) != true {
				t.Errorf("File not deleted")
				return
			}
		})
	}
}

func StartServer() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcs := grpc.NewServer()
	pb.RegisterPieceStoreRoutesServer(grpcs, &s)
	if err := grpcs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestMain(m *testing.M) {
	go StartServer()

	// Set up a connection to the Server.
	const address = "localhost:3000"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c = pb.NewPieceStoreRoutesClient(conn)

	tempDBPath := path.Join(os.TempDir(), "test.db")

	psDB, err := psdb.NewPSDB(tempDBPath)
	if err != nil {
		log.Fatal(err)
	}

	tempDir := filepath.Join(os.TempDir(), "test-data", "3000")

	s = Server{DataDir: tempDir, DB: psDB}

	db = psDB.DB

	// clean up temp files
	defer os.RemoveAll(filepath.Join(os.TempDir(), "test-data"))
	defer os.Remove(tempDBPath)
	defer db.Close()

	m.Run()
}
