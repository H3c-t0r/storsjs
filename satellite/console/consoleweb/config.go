// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleweb

import (
	"context"
	"storj.io/storj/pkg/utils"
	"storj.io/storj/satellite/satellitedb"

	"github.com/graphql-go/graphql"
	"go.uber.org/zap"

	"storj.io/storj/pkg/provider"
	"storj.io/storj/satellite/console"
	"storj.io/storj/satellite/console/consoleauth"
	"storj.io/storj/satellite/console/consoleweb/consoleql"
)

// Config contains info needed for satellite account related services
type Config struct {
	GatewayConfig
	SatelliteAddr string `help:"satellite main endpoint" default:""`
	DatabaseURL   string `help:"" default:"sqlite3://$CONFDIR/satellitedb.db"`
}

// Run implements Responsibility interface
func (c Config) Run(ctx context.Context, server *provider.Provider) error {
	log := zap.NewExample()

	// Create satellite DB
	driver, source, err := utils.SplitDBURL(c.DatabaseURL)
	if err != nil {
		return err
	}

	db, err := satellitedb.NewConsoleDB(driver, source)
	if err != nil {
		return err
	}

	err = db.CreateTables()
	if err != nil {
		log.Error(err.Error())
	}

	service, err := console.NewService(
		log,
		&consoleauth.Hmac{Secret: []byte("my-suppa-secret-key")},
		db,
	)

	if err != nil {
		return err
	}

	creator := consoleql.TypeCreator{}
	err = creator.Create(service)
	if err != nil {
		return err
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    creator.RootQuery(),
		Mutation: creator.RootMutation(),
	})

	if err != nil {
		return err
	}

	go (&gateway{
		log:     log,
		schema:  schema,
		service: service,
		config:  c.GatewayConfig,
	}).run()

	return server.Run(ctx)
}
