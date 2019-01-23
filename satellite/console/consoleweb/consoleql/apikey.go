// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleql

import (
	"github.com/graphql-go/graphql"

	"storj.io/storj/satellite/console"
)

const (
	// APIKeyInfoType is graphql type name for api key
	APIKeyInfoType = "keyInfo"
	// CreateAPIKeyType is graphql type name for createAPIKey struct
	// which incapsulates the actual key and it's info
	CreateAPIKeyType = "graphqlCreateAPIKey"
	// FieldKey is field name for the actual key in createAPIKey
	FieldKey = "key"
)

// graphqlAPIKeyInfo creates satellite.APIKeyInfo graphql object
func graphqlAPIKeyInfo() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: APIKeyInfoType,
		Fields: graphql.Fields{
			FieldID: &graphql.Field{
				Type: graphql.String,
			},
			FieldProjectID: &graphql.Field{
				Type: graphql.String,
			},
			FieldName: &graphql.Field{
				Type: graphql.String,
			},
			FieldCreatedAt: &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})
}

// graphqlCreateAPIKey creates createAPIKey graphql object
func graphqlCreateAPIKey(types Types) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: CreateAPIKeyType,
		Fields: graphql.Fields{
			FieldKey: &graphql.Field{
				Type: graphql.String,
			},
			APIKeyInfoType: &graphql.Field{
				Type: types.APIKeyInfo(),
			},
		},
	})
}

// createAPIKey holds satellite.APIKey and satellite.APIKeyInfo
type createAPIKey struct {
	Key     *console.APIKey
	KeyInfo *console.APIKeyInfo
}
