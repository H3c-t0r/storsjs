# satellite/admin

Satellite Admin package provides API endpoints for administrative tasks.

Requires setting `Authorization` header for requests.

<!-- Auto-generate this ToC with https://github.com/ycd/toc -->
<!-- toc -->
- [satellite/admin](#satelliteadmin)
    * [User Management](#user-management)
        * [POST /api/users](#post-apiusers)
        * [PUT /api/users/{user-email}](#put-apiusersuser-email)
        * [GET /api/users/{user-email}](#get-apiusersuser-email)
        * [DELETE /api/users/{user-email}](#delete-apiusersuser-email)
    * [Coupon Management](#coupon-management)
        * [POST /api/coupons](#post-apicoupons)
        * [GET /api/coupons/{coupon-id}](#get-apicouponscoupon-id)
        * [DELETE /api/coupons/{coupon-id}](#delete-apicouponscoupon-id)
    * [Project Management](#project-management)
        * [POST /api/projects](#post-apiprojects)
        * [GET /api/projects/{project-id}](#get-apiprojectsproject-id)
        * [PUT /api/projects/{project-id}](#put-apiprojectsproject-id)
        * [DELETE /api/projects/{project-id}](#delete-apiprojectsproject-id)
        * [GET /api/projects/{project}/apikeys](#get-apiprojectsprojectapikeys)
        * [POST /api/projects/{project}/apikeys](#post-apiprojectsprojectapikeys)
        * [DELETE /api/projects/{project}/apikeys/{name}](#delete-apiprojectsprojectapikeysname)
        * [GET /api/projects/{project-id}/usage](#get-apiprojectsproject-idusage)
        * [GET /api/projects/{project-id}/limit](#get-apiprojectsproject-idlimit)
        * [Update limits](#update-limits)
            * [POST /api/projects/{project-id}/limit?usage={value}](#post-apiprojectsproject-idlimitusagevalue)
            * [POST /api/projects/{project-id}/limit?bandwidth={value}](#post-apiprojectsproject-idlimitbandwidthvalue)
            * [POST /api/projects/{project-id}/limit?rate={value}](#post-apiprojectsproject-idlimitratevalue)
            * [POST /api/projects/{project-id}/limit?buckets={value}](#post-apiprojectsproject-idlimitbucketsvalue)
    * [APIKey Management](#apikey-management)
        * [DELETE /api/apikeys/{apikey}](#delete-apiapikeysapikey)

<!-- tocstop -->

## User Management

### POST /api/users

Adds a new user.

An example of a required request body:

```json
{
    "email": "alice@mail.test",
    "fullName": "Alice Test",
    "password": "password"
}
```

A successful response body:

```json
{
    "id":           "12345678-1234-1234-1234-123456789abc",
    "email":        "alice@mail.test",
    "fullName":     "Alice Test",
    "shortName":    "",
    "passwordHash": ""
}
```

### PUT /api/users/{user-email}

Updates the details of existing user found by its email.

Some example request bodies:

```json
{
    "email": "alice+2@mail.test"
}
```

```json
{
    "email": "alice+2@mail.test",
    "shortName": "myNickName"
}
```

```json
{
    "projectLimit": 200
}
```

### GET /api/users/{user-email}

This endpoint returns information about user and their projects.

A successful response body:

```json
{
    "user":{
        "id": "12345678-1234-1234-1234-123456789abc",
        "fullName": "Alice Bob",
        "email":"alice@example.test",
        "projectLimit": 10
    },
    "projects":[
        {
            "id": "abcabcab-1234-abcd-abcd-abecdefedcab",
            "name": "Project",
            "description": "Project to store data.",
            "ownerId": "12345678-1234-1234-1234-123456789abc"
        }
    ],
    "coupons": [
        {
            "id":          "2fcdbb8f-8d4d-4e6d-b6a7-8aaa1eba4c89",
            "userId":      "12345678-1234-1234-1234-123456789abc",
            "duration":    2,
            "amount":      3000,
            "description": "promotional coupon (valid for 2 billing cycles)",
            "type":        0,
            "status":      0,
            "created":     "2020-05-19T00:34:13.265761+02:00"
        }
    ]
}
```

### DELETE /api/users/{user-email}

Deletes the user.

## Coupon Management

The coupons have an amount and duration.
Amount is expressed in cents of USD dollars (e.g. 500 is $5)
Duration is expressed in billing periods, a billing period is a natural month.

### POST /api/coupons

Adds a coupon for specific user.

An example of a required request body:

```json
{
    "userId":      "12345678-1234-1234-1234-123456789abc",
    "duration":    2,
    "amount":      3000,
    "description": "promotional coupon (valid for 2 billing cycles)"
}
```

A successful response body:

```json
{
    "id": "2fcdbb8f-8d4d-4e6d-b6a7-8aaa1eba4c89"
}
```

### GET /api/coupons/{coupon-id}

Gets a coupon with the specified id.

A successful response body:

```json
{
    "id":          "2fcdbb8f-8d4d-4e6d-b6a7-8aaa1eba4c89",
    "userId":      "12345678-1234-1234-1234-123456789abc",
    "duration":    2,
    "amount":      3000,
    "description": "promotional coupon (valid for 2 billing cycles)",
    "type":        0,
    "status":      0,
    "created":     "2020-05-19T00:34:13.265761+02:00"
}
```

### DELETE /api/coupons/{coupon-id}

Deletes the specified coupon.

## Project Management

### POST /api/projects

Adds a project for specific user.

An example of a required request body:

```json
{
    "ownerId": "ca7aa0fb-442a-4d4e-aa36-a49abddae837",
    "projectName": "My Second Project"
}
```

A successful response body:

```json
{
    "projectId": "ca7aa0fb-442a-4d4e-aa36-a49abddae646"
}
```

### GET /api/projects/{project-id}

Gets the common information about a project.

### PUT /api/projects/{project-id}

Updates project name or description.

```json
{
    "projectName": "My new Project Name",
    "description": "My new awesome description!"
}
```

### DELETE /api/projects/{project-id}

Deletes the project.

### GET /api/projects/{project}/apikeys

Get the list of the API keys of a specific project.

A successful response body:

```json
[
    {
        "id": "b6988bd2-8d21-4bee-91ac-a3445bf38180",
        "ownerId": "ca7aa0fb-442a-4d4e-aa36-a49abddae837",
        "name": "mine",
        "partnerID": "a9d3b7ee-17da-4848-bb0e-1f64cf45af18",
        "createdAt": "2020-05-19T00:34:13.265761+02:00"
    },
    {
        "id": "f9f887c1-b178-4eb8-b669-14379c5a97ca",
        "ownerId": "3eb45ae9-822a-470e-a51a-9144dedda63e",
        "name": "family",
        "partnerID": "",
        "createdAt": "2020-02-20T15:34:24.265761+02:00"
    }
]
```

### POST /api/projects/{project}/apikeys

Adds an apikey for specific project.

An example of a required request body:

```json
{
    "name": "My first API Key"
}
```
**Note:** Additionally you can specify `partnerId` to associate it with the given apikey.
If you specify it, it has to be a valid uuid and not an empty string.

A successful response body:

```json
{
    "apikey": "13YqdMKxAVBamFsS6Mj3sCQ35HySoA254xmXCCQGJqffLnqrBaQDoTcCiCfbkaFPNewHT79rrFC5XRm4Z2PENtRSBDVNz8zcjS28W5v"
}
```

### DELETE /api/projects/{project}/apikeys/{name}

Deletes the given apikey by its name.

### GET /api/projects/{project-id}/usage

This endpoint returns whether the project has outstanding usage or not.

A project with not usage returns status code 200 and `{"result":"no project usage exist"}`.

### GET /api/projects/{project-id}/limit

This endpoint returns information about project limits.

A successful response body:

```json
{
  "usage": {
    "amount": "1.0 TB",
    "bytes": 1000000000000
  },
  "bandwidth": {
    "amount": "1.0 TB",
    "bytes": 1000000000000
  },
  "rate": {
    "rps": 0
  }
}
```

### Update limits

You can update the different limits with one single request just adding the
various query parameters (e.g. `usage=5000000&bandwidth=9000000`)

#### POST /api/projects/{project-id}/limit?usage={value}

Updates usage limit for a project. The value must be in bytes.

#### POST /api/projects/{project-id}/limit?bandwidth={value}

Updates bandwidth limit for a project. The value must be in bytes.

#### POST /api/projects/{project-id}/limit?rate={value}

Updates rate limit for a project.

#### POST /api/projects/{project-id}/limit?buckets={value}

Updates bucket limit for a project.

## APIKey Management

### DELETE /api/apikeys/{apikey}

Deletes the given apikey.
