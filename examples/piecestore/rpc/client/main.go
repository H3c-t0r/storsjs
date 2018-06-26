// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli"
	"github.com/zeebo/errs"
	"google.golang.org/grpc"

	"storj.io/storj/pkg/piecestore/rpc/client"
)

var argError = errs.Class("argError")

func main() {

	app := cli.NewApp()

	// Set up connection with rpc server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	psClient := client.NewPSClient(conn)

	app.Commands = []cli.Command{
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "upload data",
			Action: func(c *cli.Context) error {
				if c.Args().Get(0) == "" {
					return argError.New("No input file specified")
				}

				file, err := os.Open(c.Args().Get(0))
				if err != nil {
					return err
				}
				// Close the file when we are done
				defer file.Close()

				fileInfo, err := file.Stat()
				if err != nil {
					return err
				}

				if fileInfo.IsDir() {
					return argError.New(fmt.Sprintf("Path (%s) is a directory, not a file", c.Args().Get(0)))
				}

				var length = fileInfo.Size()
				var ttl = time.Now().Add(24 * time.Hour)

				// Created a section reader so that we can concurrently retrieve the same file.
				dataSection := io.NewSectionReader(file, 0, length)

				id := client.NewPieceID()

				err := psClient.Put(context.Background(), id, dataSection, ttl)
				if err != nil {
					fmt.Printf("Failed to Store data of id: %s\n", id)
					return err
				}

				return nil
			},
		},
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "download data",
			Action: func(c *cli.Context) error {
				if c.Args().Get(0) == "" {
					return argError.New("No id specified")
				}

				var id client.PieceID = client.PieceID(c.Args().Get(0))

				if c.Args().Get(1) == "" {
					return argError.New("No output file specified")
				}

				_, err := os.Stat(c.Args().Get(1))
				if err != nil && !os.IsNotExist(err) {
					return err
				}

				if err == nil {
					return argError.New("File already exists")
				}

				dataPath := c.Args().Get(1)

				if err = os.MkdirAll(filepath.Dir(dataPath), 0700); err != nil {
					return err
				}

				// Create File on file system
				dataFile, err := os.OpenFile(dataPath, os.O_RDWR|os.O_CREATE, 0755)
				if err != nil {
					return err
				}
				defer dataFile.Close()

				pieceInfo, err := psClient.Meta(context.Background(), id)
				if err != nil {
					os.Remove(dataPath)
					return err
				}

				ctx := context.Background()
				rr, err := psClient.Get(ctx, id, 0, pieceInfo.Size)
				if err != nil {
					fmt.Printf("Failed to retrieve file of id: %s\n", id)
					os.Remove(dataPath)
					return err
				}

				reader, err := rr.Range(ctx, 0, pieceInfo.Size)
				if err != nil {
					fmt.Printf("Failed to retrieve file of id: %s\n", id)
					os.Remove(dataPath)
					return err
				}

				_, err = io.Copy(dataFile, reader)
				if err != nil {
					fmt.Printf("Failed to retrieve file of id: %s\n", id)
					os.Remove(dataPath)
				} else {
					fmt.Printf("Successfully retrieved file of id: %s\n", id)
				}

				return reader.Close()
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"x"},
			Usage:   "delete data",
			Action: func(c *cli.Context) error {
				if c.Args().Get(0) == "" {
					return argError.New("Missing data Id")
				}
				err = psClient.Delete(context.Background(), client.PieceID(c.Args().Get(0)))

				return err
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
