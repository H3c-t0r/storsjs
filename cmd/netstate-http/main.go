// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"

	"storj.io/storj/netstate-http/routes"
	"storj.io/storj/storage/boltdb"
)

var (
	port   int
	dbPath string
	prod   bool
)

func initializeFlags() {
	flag.IntVar(&port, "port", 3000, "port")
	flag.StringVar(&dbPath, "db", "netstate.db", "db path")
	flag.BoolVar(&prod, "prod", false, "The environment this service is running in")
	flag.Parse()
}

func main() {
	err := Main()
	if err != nil {
		log.Fatalf("fatal error: %v", err)
		os.Exit(1)
	}
}

// Main allows simplified error handling
func Main() error {
	initializeFlags()

	// No err here because no vars passed into NewDevelopment().
	// The default won't return an error, but if args are passed in,
	// then there will need to be error handling.
	logger, _ := zap.NewDevelopment()
	if prod {
		logger, _ = zap.NewProduction()
	}
	defer logger.Sync()
	logger.Info(fmt.Sprintf("serving on %d", port))

	bdb, err := boltdb.New(logger, dbPath)
	if err != nil {
		return err
	}
	defer bdb.Close()

	ns := routes.NewNetStateRoutes(logger, bdb)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), routes.Start(ns))
}
