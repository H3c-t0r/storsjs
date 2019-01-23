// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleweb

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"path/filepath"

	"storj.io/storj/pkg/auth"

	"github.com/graphql-go/graphql"
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/storj/satellite/console"
	"storj.io/storj/satellite/console/consoleweb/consoleql"
)

const (
	authorization = "Authorization"
	contentType   = "Content-Type"

	authorizationBearer = "Bearer "

	applicationJSON    = "application/json"
	applicationGraphql = "application/graphql"
)

// Error is satellite console error type
var Error = errs.Class("satellite console error")

// Server represents console web server
type Server struct {
	logger *zap.Logger

	config   Config
	service  *console.Service
	listener net.Listener

	schema graphql.Schema
	server http.Server
}

// NewServer creates new instance of console server
func NewServer(logger *zap.Logger, config Config, service *console.Service, listener net.Listener) *Server {
	server := Server{
		logger:   logger,
		service:  service,
		config:   config,
		listener: listener,
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(server.config.StaticPath))

	mux.Handle("/api/graphql/v0", http.HandlerFunc(server.grapqlHandler))

	if server.config.StaticPath != "" {
		mux.Handle("/", http.HandlerFunc(server.appHandler))
		mux.Handle("/static/", http.StripPrefix("/static", fs))
	}

	server.server = http.Server{
		Handler: mux,
	}

	return &server
}

// appHandler is web app http handler function
func (s *Server) appHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, filepath.Join(s.config.StaticPath, "dist", "public", "index.html"))
}

// grapqlHandler is graphql endpoint http handler function
func (s *Server) grapqlHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(contentType, applicationJSON)

	token := getToken(req)
	query, err := getQuery(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := auth.WithAPIKey(context.Background(), []byte(token))
	auth, err := s.service.Authorize(ctx)
	if err != nil {
		ctx = console.WithAuthFailure(ctx, err)
	} else {
		ctx = console.WithAuth(ctx, auth)
	}

	result := graphql.Do(graphql.Params{
		Schema:         s.schema,
		Context:        ctx,
		RequestString:  query.Query,
		VariableValues: query.Variables,
		OperationName:  query.OperationName,
		RootObject:     make(map[string]interface{}),
	})

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}

	sugar := s.logger.Sugar()
	sugar.Debug(result)
}

// Run implements Responsibility interface
func (s *Server) Run(ctx context.Context) error {
	creator := consoleql.TypeCreator{}
	err := creator.Create(s.service)
	if err != nil {
		return Error.Wrap(err)
	}

	s.schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    creator.RootQuery(),
		Mutation: creator.RootMutation(),
	})

	if err != nil {
		return Error.Wrap(err)
	}

	return s.server.Serve(s.listener)
}

// Close closes server and underlying listener
func (s *Server) Close() error {
	return s.server.Close()
}
