// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

// Project is a type, used for creating new project in backend
declare type Project = {
	id: string,
	ownerName: string,

	name: string,
	description: string,
	companyName: string,
	isTermsAccepted: boolean,
	createdAt: string,

	isSelected: boolean,
}

// UpdateProjectModel is a type, used for updating project description
declare type UpdateProjectModel = {
	id: string,
	description: string,
}

// TeamMemberModel stores needed info about user info to show it on UI
declare type TeamMemberModel = {
	user: {
		id: string,
		email: string,
		firstName: string,
		lastName: string,
		company: {
			companyName: string,
		}
	}
	joinedAt: string,
}
