// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"

	"storj.io/storj/pkg/satellite"
	"storj.io/storj/pkg/satellite/satellitedb/dbx"
	"storj.io/storj/pkg/utils"
)

// ProjectMembers exposes methods to manage ProjectMembers table in database.
type projectMembers struct {
	db dbx.Methods
}

// GetByMemberID is a method for querying project member from the database by memberID.
func (pm *projectMembers) GetByMemberID(ctx context.Context, memberID uuid.UUID) (*satellite.ProjectMember, error) {
	projectMemberDbx, err := pm.db.Get_ProjectMember_By_MemberId(ctx, dbx.ProjectMember_MemberId(memberID[:]))
	if err != nil {
		return nil, err
	}

	return projectMemberFromDBX(projectMemberDbx)
}

// GetByProjectID is a method for querying project members from the database by projectID.
func (pm *projectMembers) GetByProjectID(ctx context.Context, projectID uuid.UUID) ([]satellite.ProjectMember, error) {
	projectMembersDbx, err := pm.db.All_ProjectMember_By_ProjectId(ctx, dbx.ProjectMember_ProjectId(projectID[:]))
	if err != nil {
		return nil, err
	}

	return projectMembersFromDbxSlice(projectMembersDbx)
}

// GetPaged is a method for querying project members from the database by projectID, offset and limit.
func (pm *projectMembers) GetPaged(ctx context.Context, projectID uuid.UUID, limit, offset int64) ([]satellite.ProjectMember, error) {
	projectMembersDbx, err := pm.db.Limited_ProjectMember_By_ProjectId(
		ctx,
		dbx.ProjectMember_ProjectId(projectID[:]),
		int(limit),
		offset)

	if err != nil {
		return nil, err
	}

	return projectMembersFromDbxSlice(projectMembersDbx)
}

// Insert is a method for inserting project member into the database.
func (pm *projectMembers) Insert(ctx context.Context, memberID, projectID uuid.UUID) (*satellite.ProjectMember, error) {
	createdProjectMember, err := pm.db.Create_ProjectMember(ctx,
		dbx.ProjectMember_MemberId(memberID[:]),
		dbx.ProjectMember_ProjectId(projectID[:]))
	if err != nil {
		return nil, err
	}

	return projectMemberFromDBX(createdProjectMember)
}

// Delete is a method for deleting project member by memberID and projectID from the database.
func (pm *projectMembers) Delete(ctx context.Context, memberID, projectID uuid.UUID) error {
	_, err := pm.db.Delete_ProjectMember_By_MemberId_And_ProjectId(
		ctx,
		dbx.ProjectMember_MemberId(memberID[:]),
		dbx.ProjectMember_ProjectId(projectID[:]),
	)

	return err
}

// projectMemberFromDBX is used for creating ProjectMember entity from autogenerated dbx.ProjectMember struct
func projectMemberFromDBX(projectMember *dbx.ProjectMember) (*satellite.ProjectMember, error) {
	if projectMember == nil {
		return nil, errs.New("projectMember parameter is nil")
	}

	memberID, err := bytesToUUID(projectMember.MemberId)
	if err != nil {
		return nil, err
	}

	projectID, err := bytesToUUID(projectMember.ProjectId)
	if err != nil {
		return nil, err
	}

	return &satellite.ProjectMember{
		MemberID:  memberID,
		ProjectID: projectID,
		CreatedAt: projectMember.CreatedAt,
	}, nil
}

// projectMembersFromDbxSlice is used for creating []ProjectMember entities from autogenerated []*dbx.ProjectMember struct
func projectMembersFromDbxSlice(projectMembersDbx []*dbx.ProjectMember) ([]satellite.ProjectMember, error) {
	var projectMembers []satellite.ProjectMember
	var errors []error

	// Generating []dbo from []dbx and collecting all errors
	for _, projectMemberDbx := range projectMembersDbx {
		projectMember, err := projectMemberFromDBX(projectMemberDbx)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		projectMembers = append(projectMembers, *projectMember)
	}

	return projectMembers, utils.CombineErrors(errors...)
}
