// Code generated by lockedgen using 'go generate'. DO NOT EDIT.

// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"sync"
	"time"

	"github.com/skyrings/skyring-common/tools/uuid"

	"storj.io/storj/pkg/accounting"
	"storj.io/storj/pkg/bwagreement"
	"storj.io/storj/pkg/datarepair/irreparable"
	"storj.io/storj/pkg/datarepair/queue"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/statdb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/console"
)

// locked implements a locking wrapper around satellite.DB.
type locked struct {
	sync.Locker
	db satellite.DB
}

// newLocked returns database wrapped with locker.
func newLocked(db satellite.DB) satellite.DB {
	return &locked{&sync.Mutex{}, db}
}

// Accounting returns database for storing information about data use
func (m *locked) Accounting() accounting.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedAccounting{m.Locker, m.db.Accounting()}
}

// lockedAccounting implements locking wrapper for accounting.DB
type lockedAccounting struct {
	sync.Locker
	db accounting.DB
}

// GetRaw retrieves all raw tallies
func (m *lockedAccounting) GetRaw(ctx context.Context) ([]*accounting.Raw, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetRaw(ctx)
}

// GetRawSince r retrieves all raw tallies sinces
func (m *lockedAccounting) GetRawSince(ctx context.Context, latestRollup time.Time) ([]*accounting.Raw, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetRawSince(ctx, latestRollup)
}

// LastTimestamp records the latest last tallied time.
func (m *lockedAccounting) LastTimestamp(ctx context.Context, timestampType string) (time.Time, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.LastTimestamp(ctx, timestampType)
}

// QueryPaymentInfo queries StatDB, Accounting Rollup on nodeID
func (m *lockedAccounting) QueryPaymentInfo(ctx context.Context, start time.Time, end time.Time) ([]*accounting.CSVRow, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.QueryPaymentInfo(ctx, start, end)
}

// SaveAtRestRaw records raw tallies of at-rest-data.
func (m *lockedAccounting) SaveAtRestRaw(ctx context.Context, latestTally time.Time, nodeData map[storj.NodeID]float64) error {
	m.Lock()
	defer m.Unlock()
	return m.db.SaveAtRestRaw(ctx, latestTally, nodeData)
}

// SaveBWRaw records raw sums of agreement values to the database and updates the LastTimestamp.
func (m *lockedAccounting) SaveBWRaw(ctx context.Context, tallyEnd time.Time, bwTotals map[storj.NodeID][]int64) error {
	m.Lock()
	defer m.Unlock()
	return m.db.SaveBWRaw(ctx, tallyEnd, bwTotals)
}

// SaveRollup records raw tallies of at rest data to the database
func (m *lockedAccounting) SaveRollup(ctx context.Context, latestTally time.Time, stats accounting.RollupStats) error {
	m.Lock()
	defer m.Unlock()
	return m.db.SaveRollup(ctx, latestTally, stats)
}

// BandwidthAgreement returns database for storing bandwidth agreements
func (m *locked) BandwidthAgreement() bwagreement.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedBandwidthAgreement{m.Locker, m.db.BandwidthAgreement()}
}

// lockedBandwidthAgreement implements locking wrapper for bwagreement.DB
type lockedBandwidthAgreement struct {
	sync.Locker
	db bwagreement.DB
}

// CreateAgreement adds a new bandwidth agreement.
func (m *lockedBandwidthAgreement) CreateAgreement(ctx context.Context, a1 *pb.RenterBandwidthAllocation) error {
	m.Lock()
	defer m.Unlock()
	return m.db.CreateAgreement(ctx, a1)
}

// GetTotalsSince returns the sum of each bandwidth type after (exluding) a given date range
func (m *lockedBandwidthAgreement) GetTotals(ctx context.Context, a1 time.Time, a2 time.Time) (map[storj.NodeID][]int64, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetTotals(ctx, a1, a2)
}

// GetTotals returns stats about an uplink
func (m *lockedBandwidthAgreement) GetUplinkStats(ctx context.Context, a1 time.Time, a2 time.Time) ([]bwagreement.UplinkStat, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetUplinkStats(ctx, a1, a2)
}

// Close closes the database
func (m *locked) Close() error {
	m.Lock()
	defer m.Unlock()
	return m.db.Close()
}

// Console returns database for satellite console
func (m *locked) Console() console.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedConsole{m.Locker, m.db.Console()}
}

// lockedConsole implements locking wrapper for console.DB
type lockedConsole struct {
	sync.Locker
	db console.DB
}

// APIKeys is a getter for APIKeys repository
func (m *lockedConsole) APIKeys() console.APIKeys {
	m.Lock()
	defer m.Unlock()
	return &lockedAPIKeys{m.Locker, m.db.APIKeys()}
}

// lockedAPIKeys implements locking wrapper for console.APIKeys
type lockedAPIKeys struct {
	sync.Locker
	db console.APIKeys
}

// Create creates and stores new APIKeyInfo
func (m *lockedAPIKeys) Create(ctx context.Context, key console.APIKey, info console.APIKeyInfo) (*console.APIKeyInfo, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Create(ctx, key, info)
}

// Delete deletes APIKeyInfo from store
func (m *lockedAPIKeys) Delete(ctx context.Context, id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Delete(ctx, id)
}

// Get retrieves APIKeyInfo with given ID
func (m *lockedAPIKeys) Get(ctx context.Context, id uuid.UUID) (*console.APIKeyInfo, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Get(ctx, id)
}

// GetByKey retrieves APIKeyInfo for given key
func (m *lockedAPIKeys) GetByKey(ctx context.Context, key console.APIKey) (*console.APIKeyInfo, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetByKey(ctx, key)
}

// GetByProjectID retrieves list of APIKeys for given projectID
func (m *lockedAPIKeys) GetByProjectID(ctx context.Context, projectID uuid.UUID) ([]console.APIKeyInfo, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetByProjectID(ctx, projectID)
}

// Update updates APIKeyInfo in store
func (m *lockedAPIKeys) Update(ctx context.Context, key console.APIKeyInfo) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Update(ctx, key)
}

// Buckets is a getter for Buckets repository
func (m *lockedConsole) Buckets() console.Buckets {
	m.Lock()
	defer m.Unlock()
	return &lockedBuckets{m.Locker, m.db.Buckets()}
}

// lockedBuckets implements locking wrapper for console.Buckets
type lockedBuckets struct {
	sync.Locker
	db console.Buckets
}

// AttachBucket attaches a bucket to a project
func (m *lockedBuckets) AttachBucket(ctx context.Context, name string, projectID uuid.UUID) (*console.Bucket, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.AttachBucket(ctx, name, projectID)
}

// DeattachBucket deletes bucket info for a bucket by name
func (m *lockedBuckets) DeattachBucket(ctx context.Context, name string) error {
	m.Lock()
	defer m.Unlock()
	return m.db.DeattachBucket(ctx, name)
}

// GetBucket retrieves bucket info of bucket with given name
func (m *lockedBuckets) GetBucket(ctx context.Context, name string) (*console.Bucket, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetBucket(ctx, name)
}

// ListBuckets returns bucket list of a given project
func (m *lockedBuckets) ListBuckets(ctx context.Context, projectID uuid.UUID) ([]console.Bucket, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.ListBuckets(ctx, projectID)
}

// Close is used to close db connection
func (m *lockedConsole) Close() error {
	m.Lock()
	defer m.Unlock()
	return m.db.Close()
}

// CreateTables is a method for creating all tables for satellitedb
func (m *lockedConsole) CreateTables() error {
	m.Lock()
	defer m.Unlock()
	return m.db.CreateTables()
}

// ProjectMembers is a getter for ProjectMembers repository
func (m *lockedConsole) ProjectMembers() console.ProjectMembers {
	m.Lock()
	defer m.Unlock()
	return &lockedProjectMembers{m.Locker, m.db.ProjectMembers()}
}

// lockedProjectMembers implements locking wrapper for console.ProjectMembers
type lockedProjectMembers struct {
	sync.Locker
	db console.ProjectMembers
}

// Delete is a method for deleting project member by memberID and projectID from the database.
func (m *lockedProjectMembers) Delete(ctx context.Context, memberID uuid.UUID, projectID uuid.UUID) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Delete(ctx, memberID, projectID)
}

// GetByMemberID is a method for querying project members from the database by memberID.
func (m *lockedProjectMembers) GetByMemberID(ctx context.Context, memberID uuid.UUID) ([]console.ProjectMember, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetByMemberID(ctx, memberID)
}

// GetByProjectID is a method for querying project members from the database by projectID, offset and limit.
func (m *lockedProjectMembers) GetByProjectID(ctx context.Context, projectID uuid.UUID, pagination console.Pagination) ([]console.ProjectMember, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetByProjectID(ctx, projectID, pagination)
}

// Insert is a method for inserting project member into the database.
func (m *lockedProjectMembers) Insert(ctx context.Context, memberID uuid.UUID, projectID uuid.UUID) (*console.ProjectMember, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Insert(ctx, memberID, projectID)
}

// Projects is a getter for Projects repository
func (m *lockedConsole) Projects() console.Projects {
	m.Lock()
	defer m.Unlock()
	return &lockedProjects{m.Locker, m.db.Projects()}
}

// lockedProjects implements locking wrapper for console.Projects
type lockedProjects struct {
	sync.Locker
	db console.Projects
}

// Delete is a method for deleting project by Id from the database.
func (m *lockedProjects) Delete(ctx context.Context, id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Delete(ctx, id)
}

// Get is a method for querying project from the database by id.
func (m *lockedProjects) Get(ctx context.Context, id uuid.UUID) (*console.Project, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Get(ctx, id)
}

// GetAll is a method for querying all projects from the database.
func (m *lockedProjects) GetAll(ctx context.Context) ([]console.Project, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetAll(ctx)
}

// GetByUserID is a method for querying all projects from the database by userID.
func (m *lockedProjects) GetByUserID(ctx context.Context, userID uuid.UUID) ([]console.Project, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetByUserID(ctx, userID)
}

// Insert is a method for inserting project into the database.
func (m *lockedProjects) Insert(ctx context.Context, project *console.Project) (*console.Project, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Insert(ctx, project)
}

// Update is a method for updating project entity.
func (m *lockedProjects) Update(ctx context.Context, project *console.Project) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Update(ctx, project)
}

// Users is a getter for Users repository
func (m *lockedConsole) Users() console.Users {
	m.Lock()
	defer m.Unlock()
	return &lockedUsers{m.Locker, m.db.Users()}
}

// lockedUsers implements locking wrapper for console.Users
type lockedUsers struct {
	sync.Locker
	db console.Users
}

// Delete is a method for deleting user by Id from the database.
func (m *lockedUsers) Delete(ctx context.Context, id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Delete(ctx, id)
}

// Get is a method for querying user from the database by id.
func (m *lockedUsers) Get(ctx context.Context, id uuid.UUID) (*console.User, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Get(ctx, id)
}

// GetByEmail is a method for querying user by email from the database.
func (m *lockedUsers) GetByEmail(ctx context.Context, email string) (*console.User, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetByEmail(ctx, email)
}

// Insert is a method for inserting user into the database.
func (m *lockedUsers) Insert(ctx context.Context, user *console.User) (*console.User, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Insert(ctx, user)
}

// Update is a method for updating user entity.
func (m *lockedUsers) Update(ctx context.Context, user *console.User) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Update(ctx, user)
}

// CreateTables initializes the database
func (m *locked) CreateTables() error {
	m.Lock()
	defer m.Unlock()
	return m.db.CreateTables()
}

// Irreparable returns database for failed repairs
func (m *locked) Irreparable() irreparable.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedIrreparable{m.Locker, m.db.Irreparable()}
}

// lockedIrreparable implements locking wrapper for irreparable.DB
type lockedIrreparable struct {
	sync.Locker
	db irreparable.DB
}

// Delete removes irreparable segment info based on segmentPath.
func (m *lockedIrreparable) Delete(ctx context.Context, segmentPath []byte) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Delete(ctx, segmentPath)
}

// Get returns irreparable segment info based on segmentPath.
func (m *lockedIrreparable) Get(ctx context.Context, segmentPath []byte) (*irreparable.RemoteSegmentInfo, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Get(ctx, segmentPath)
}

// IncrementRepairAttempts increments the repair attempts.
func (m *lockedIrreparable) IncrementRepairAttempts(ctx context.Context, segmentInfo *irreparable.RemoteSegmentInfo) error {
	m.Lock()
	defer m.Unlock()
	return m.db.IncrementRepairAttempts(ctx, segmentInfo)
}

// OverlayCache returns database for caching overlay information
func (m *locked) OverlayCache() overlay.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedOverlayCache{m.Locker, m.db.OverlayCache()}
}

// lockedOverlayCache implements locking wrapper for overlay.DB
type lockedOverlayCache struct {
	sync.Locker
	db overlay.DB
}

// Delete deletes node based on id
func (m *lockedOverlayCache) Delete(ctx context.Context, id storj.NodeID) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Delete(ctx, id)
}

// FilterNodes looks up nodes based on reputation requirements
func (m *lockedOverlayCache) FilterNodes(ctx context.Context, filterNodesRequest *overlay.FilterNodesRequest) ([]*pb.Node, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.FilterNodes(ctx, filterNodesRequest)
}

// Get looks up the node by nodeID
func (m *lockedOverlayCache) Get(ctx context.Context, nodeID storj.NodeID) (*pb.Node, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Get(ctx, nodeID)
}

// GetAll looks up nodes based on the ids from the overlay cache
func (m *lockedOverlayCache) GetAll(ctx context.Context, nodeIDs storj.NodeIDList) ([]*pb.Node, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetAll(ctx, nodeIDs)
}

// GetWalletAddress gets the node's wallet address
func (m *lockedOverlayCache) GetWalletAddress(ctx context.Context, id storj.NodeID) (string, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetWalletAddress(ctx, id)
}

// List lists nodes starting from cursor
func (m *lockedOverlayCache) List(ctx context.Context, cursor storj.NodeID, limit int) ([]*pb.Node, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.List(ctx, cursor, limit)
}

// Paginate will page through the database nodes
func (m *lockedOverlayCache) Paginate(ctx context.Context, offset int64, limit int) ([]*pb.Node, bool, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Paginate(ctx, offset, limit)
}

// Update updates node information
func (m *lockedOverlayCache) Update(ctx context.Context, value *pb.Node) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Update(ctx, value)
}

// RepairQueue returns queue for segments that need repairing
func (m *locked) RepairQueue() queue.RepairQueue {
	m.Lock()
	defer m.Unlock()
	return &lockedRepairQueue{m.Locker, m.db.RepairQueue()}
}

// lockedRepairQueue implements locking wrapper for queue.RepairQueue
type lockedRepairQueue struct {
	sync.Locker
	db queue.RepairQueue
}

// Dequeue removes an injured segment.
func (m *lockedRepairQueue) Dequeue(ctx context.Context) (pb.InjuredSegment, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Dequeue(ctx)
}

// Enqueue adds an injured segment.
func (m *lockedRepairQueue) Enqueue(ctx context.Context, qi *pb.InjuredSegment) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Enqueue(ctx, qi)
}

// Peekqueue lists limit amount of injured segments.
func (m *lockedRepairQueue) Peekqueue(ctx context.Context, limit int) ([]pb.InjuredSegment, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Peekqueue(ctx, limit)
}

// StatDB returns database for storing node statistics
func (m *locked) StatDB() statdb.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedStatDB{m.Locker, m.db.StatDB()}
}

// lockedStatDB implements locking wrapper for statdb.DB
type lockedStatDB struct {
	sync.Locker
	db statdb.DB
}

// Create adds a new stats entry for node.
func (m *lockedStatDB) Create(ctx context.Context, nodeID storj.NodeID, initial *statdb.NodeStats) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Create(ctx, nodeID, initial)
}

// CreateEntryIfNotExists creates a node stats entry if it didn't already exist.
func (m *lockedStatDB) CreateEntryIfNotExists(ctx context.Context, nodeID storj.NodeID) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.CreateEntryIfNotExists(ctx, nodeID)
}

// FindInvalidNodes finds a subset of storagenodes that have stats below provided reputation requirements.
func (m *lockedStatDB) FindInvalidNodes(ctx context.Context, nodeIDs storj.NodeIDList, maxStats *statdb.NodeStats) (invalid storj.NodeIDList, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.FindInvalidNodes(ctx, nodeIDs, maxStats)
}

// Get returns node stats.
func (m *lockedStatDB) Get(ctx context.Context, nodeID storj.NodeID) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Get(ctx, nodeID)
}

// Update all parts of single storagenode's stats.
func (m *lockedStatDB) Update(ctx context.Context, request *statdb.UpdateRequest) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Update(ctx, request)
}

// UpdateAuditSuccess updates a single storagenode's audit stats.
func (m *lockedStatDB) UpdateAuditSuccess(ctx context.Context, nodeID storj.NodeID, auditSuccess bool) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.UpdateAuditSuccess(ctx, nodeID, auditSuccess)
}

// UpdateBatch for updating multiple storage nodes' stats.
func (m *lockedStatDB) UpdateBatch(ctx context.Context, requests []*statdb.UpdateRequest) (statslist []*statdb.NodeStats, failed []*statdb.UpdateRequest, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.UpdateBatch(ctx, requests)
}

// UpdateUptime updates a single storagenode's uptime stats.
func (m *lockedStatDB) UpdateUptime(ctx context.Context, nodeID storj.NodeID, isUp bool) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.UpdateUptime(ctx, nodeID, isUp)
}
