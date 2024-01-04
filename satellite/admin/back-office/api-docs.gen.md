# API Docs

**Version:** `v1`

<h2 id='list-of-endpoints'>List of Endpoints</h2>

* Settings
  * [Get settings](#settings-get-settings)
* PlacementManagement
  * [Get placements](#placementmanagement-get-placements)
* UserManagement
  * [Get user](#usermanagement-get-user)

<h3 id='settings-get-settings'>Get settings (<a href='#list-of-endpoints'>go to full list</a>)</h3>

Gets the settings of the service and relevant Storj services settings

`GET /back-office/api/v1/settings/`

**Response body:**

```typescript
{
	admin: 	{
		features: 		{
			account: 			{
				create: boolean
				delete: boolean
				history: boolean
				list: boolean
				projects: boolean
				suspend: boolean
				unsusped: boolean
				resetMFA: boolean
				update: boolean
				view: boolean
			}

			project: 			{
				create: boolean
				delete: boolean
				history: boolean
				list: boolean
				update: boolean
				view: boolean
				memberList: boolean
				memberAdd: boolean
				memberRemove: boolean
			}

			bucket: 			{
				create: boolean
				delete: boolean
				history: boolean
				list: boolean
				update: boolean
				view: boolean
			}

			dashboard: boolean
			operator: boolean
			singOut: boolean
			switchSatellite: boolean
		}

	}

}

```

<h3 id='placementmanagement-get-placements'>Get placements (<a href='#list-of-endpoints'>go to full list</a>)</h3>

Gets placement rule IDs and their locations

`GET /back-office/api/v1/placements/`

**Response body:**

```typescript
[
	{
		id: number
		location: string
	}

]

```

<h3 id='usermanagement-get-user'>Get user (<a href='#list-of-endpoints'>go to full list</a>)</h3>

Gets user by email address

`GET /back-office/api/v1/users/{email}`

**Path Params:**

| name | type | elaboration |
|---|---|---|
| `email` | `string` |  |

**Response body:**

```typescript
{
	id: string // UUID formatted as `00000000-0000-0000-0000-000000000000`
	fullName: string
	email: string
	paidTier: boolean
	createdAt: string // Date timestamp formatted as `2006-01-02T15:00:00Z`
	status: string
	userAgent: string
	defaultPlacement: number
	projectUsageLimits: 	[
		{
			id: string // UUID formatted as `00000000-0000-0000-0000-000000000000`
			name: string
			storageLimit: number
			storageUsed: number
			bandwidthLimit: number
			bandwidthUsed: number
			segmentLimit: number
			segmentUsed: number
		}

	]

}

```

