// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

export class User {
    public id: string;
    public fullName: string;
    public shortName: string;
    public email: string;

    public constructor(fullName: string, shortName: string, email: string) {
        this.id = '';
        this.fullName = fullName;
        this.shortName = shortName;
        this.email = email;
    }
}

export class UpdatedUser {
    public fullName: string;
    public shortName: string;
}

// Used in users module to pass parameters to action
export class UpdatePasswordModel {
    public oldPassword: string;
    public newPassword: string;
}
