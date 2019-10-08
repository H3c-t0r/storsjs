// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"

	"github.com/skyrings/skyring-common/tools/uuid"

	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

// customers is an implementation of payments.Accounts
type customers struct {
	db *dbx.DB
}

// Insert is a method for inserting stripe customer into the database.
func (customers *customers) Insert(ctx context.Context, userID uuid.UUID, customerID string) (err error) {
	defer mon.Task()(&ctx)(&err)

	_, err = customers.db.Create_StripeCustomers(
		ctx,
		dbx.StripeCustomers_UserId(userID[:]),
		dbx.StripeCustomers_CustomerId(customerID),
	)

	return err
}

// Insert is a method for inserting stripe customer into the database.
func (customers *customers) GetAllCustomerIDs(ctx context.Context) (ids []string, err error) {
	defer mon.Task()(&ctx)(&err)

	rows, err := customers.db.All_StripeCustomers_CustomerId_OrderBy_Asc_CreatedAt(ctx)
	if err != nil {
		return nil, err
	}

	return customers.fromDbxRowCustomerID(rows), nil
}

func (customers *customers) fromDbxRowCustomerID(rows []*dbx.CustomerId_Row) (ids []string) {
	for _, row := range rows {
		ids = append(ids, row.CustomerId)
	}

	return
}
