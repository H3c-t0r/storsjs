// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package rewards

import (
	"context"
	"sort"

	"github.com/zeebo/errs"
)

// PartnersStaticDB implements partner lookup based on a static definition.
type PartnersStaticDB struct {
	list        *List
	byName      map[string]Partner
	byID        map[string]Partner
	byUserAgent map[string]Partner
}

var _ DB = (*PartnersStaticDB)(nil)

// NewPartnersStaticDB creates a new PartnersStaticDB.
func NewPartnersStaticDB(list *List) (*PartnersStaticDB, error) {
	db := &PartnersStaticDB{
		list:        list,
		byName:      map[string]Partner{},
		byID:        map[string]Partner{},
		byUserAgent: map[string]Partner{},
	}

	sort.Slice(list.Partners, func(i, k int) bool {
		return list.Partners[i].Name < list.Partners[k].Name
	})

	var errg errs.Group
	for _, p := range list.Partners {
		if _, exists := db.byName[p.Name]; exists {
			errg.Add(Error.New("name %q already exists", p.Name))
		} else {
			db.byName[p.Name] = p
		}

		if _, exists := db.byID[p.ID]; exists {
			errg.Add(Error.New("id %q already exists", p.ID))
		} else {
			db.byID[p.ID] = p
		}

		useragent := CanonicalUserAgentProduct(p.UserAgent())
		if _, exists := db.byUserAgent[useragent]; exists {
			errg.Add(Error.New("user agent %q already exists", useragent))
		} else {
			db.byUserAgent[useragent] = p
		}
	}

	return db, errg.Err()
}

// All returns all partners.
func (db *PartnersStaticDB) All(ctx context.Context) ([]Partner, error) {
	return append([]Partner{}, db.list.Partners...), nil
}

// ByName returns partner definitions for a given name.
func (db *PartnersStaticDB) ByName(ctx context.Context, name string) (Partner, error) {
	partner, ok := db.byName[name]
	if !ok {
		return Partner{}, ErrNotExist.New("%q", name)
	}
	return partner, nil
}

// ByID returns partner definition corresponding to an id.
func (db *PartnersStaticDB) ByID(ctx context.Context, id string) (Partner, error) {
	partner, ok := db.byID[id]
	if !ok {
		return Partner{}, ErrNotExist.New("%q", id)
	}
	return partner, nil
}

// ByUserAgent returns partner definition corresponding to an user agent product string.
func (db *PartnersStaticDB) ByUserAgent(ctx context.Context, agent string) (Partner, error) {
	partner, ok := db.byUserAgent[CanonicalUserAgentProduct(agent)]
	if !ok {
		return Partner{}, ErrNotExist.New("%q", agent)
	}
	return partner, nil
}
