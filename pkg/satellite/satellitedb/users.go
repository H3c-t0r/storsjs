// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"

	"github.com/zeebo/errs"

	"storj.io/storj/pkg/satellite"

	"github.com/skyrings/skyring-common/tools/uuid"
	"storj.io/storj/pkg/satellite/satellitedb/dbx"
)

// implementation of Users interface repository using spacemonkeygo/dbx orm
type users struct {
	db *dbx.DB
}

// Get is a method for querying user from the database by id
func (users *users) Get(ctx context.Context, id uuid.UUID) (*satellite.User, error) {

	userID := dbx.User_Id([]byte(id.String()))

	user, err := users.db.Get_User_By_Id(ctx, userID)

	if err != nil {
		return nil, err
	}

	return userFromDBX(user)
}

// GetByCredentials is a method for querying user by credentials from the database.
func (users *users) GetByCredentials(ctx context.Context, password []byte, email string) (*satellite.User, error) {

	userEmail := dbx.User_Email(email)
	userPassword := dbx.User_PasswordHash(password)

	user, err := users.db.Get_User_By_Email_And_PasswordHash(ctx, userEmail, userPassword)

	if err != nil {
		return nil, err
	}

	return userFromDBX(user)
}

// Insert is a method for inserting user into the database
func (users *users) Insert(ctx context.Context, user *satellite.User) (*satellite.User, error) {

	userID, err := uuid.New()
	if err != nil {
		return nil, err
	}

	createdUser, err := users.db.Create_User(ctx,
		dbx.User_Id([]byte(userID.String())),
		dbx.User_FirstName(user.FirstName),
		dbx.User_LastName(user.LastName),
		dbx.User_Email(user.Email),
		dbx.User_PasswordHash(user.PasswordHash))

	if err != nil {
		return nil, err
	}

	return userFromDBX(createdUser)
}

// Delete is a method for deleting user by Id from the database.
func (users *users) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := users.db.Delete_User_By_Id(ctx, dbx.User_Id([]byte(id.String())))

	return err
}

// Update is a method for updating user entity
func (users *users) Update(ctx context.Context, user *satellite.User) error {
	_, err := users.db.Update_User_By_Id(ctx,
		dbx.User_Id([]byte(user.ID.String())),
		dbx.User_Update_Fields{
			FirstName:    dbx.User_FirstName(user.FirstName),
			LastName:     dbx.User_LastName(user.LastName),
			Email:        dbx.User_Email(user.Email),
			PasswordHash: dbx.User_PasswordHash(user.PasswordHash),
		})

	return err
}

// userFromDBX is used for creating User entity from autogenerated dbx.User struct
// TODO: move error strings to better place
func userFromDBX(user *dbx.User) (*satellite.User, error) {
	if user == nil {
		return nil, errs.New("user parameter is nil")
	}

	id, err := uuid.Parse(string(user.Id))
	if err != nil {
		return nil, errs.New("Id in not valid UUID string")
	}

	u := &satellite.User{
		ID:           *id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
	}

	return u, nil
}
