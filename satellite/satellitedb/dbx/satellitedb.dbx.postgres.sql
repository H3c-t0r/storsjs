-- AUTOGENERATED BY gopkg.in/spacemonkeygo/dbx.v1
-- DO NOT EDIT
CREATE TABLE accounting_raws (
	id bigserial NOT NULL,
	node_id text NOT NULL,
	interval_end_time timestamp with time zone NOT NULL,
	data_total bigint NOT NULL,
	data_type integer NOT NULL,
	created_at timestamp with time zone NOT NULL,
	updated_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE accounting_rollups (
	id bigserial NOT NULL,
	node_id bytea NOT NULL,
	start_time timestamp with time zone NOT NULL,
	interval bigint NOT NULL,
	data_total bigint NOT NULL,
	data_type integer NOT NULL,
	created_at timestamp with time zone NOT NULL,
	updated_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE accounting_timestamps (
	name text NOT NULL,
	value timestamp with time zone NOT NULL,
	PRIMARY KEY ( name )
);
CREATE TABLE bwagreements (
	signature bytea NOT NULL,
	serialnum text NOT NULL,
	data bytea NOT NULL,
	created_at timestamp with time zone NOT NULL,
	expires_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( signature ),
	UNIQUE ( serialnum )
);
CREATE TABLE injuredsegments (
	id bigserial NOT NULL,
	info bytea NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE irreparabledbs (
	segmentpath bytea NOT NULL,
	segmentdetail bytea NOT NULL,
	pieces_lost_count bigint NOT NULL,
	seg_damaged_unix_sec bigint NOT NULL,
	repair_attempt_count bigint NOT NULL,
	PRIMARY KEY ( segmentpath )
);
CREATE TABLE nodes (
	id bytea NOT NULL,
	audit_success_count bigint NOT NULL,
	total_audit_count bigint NOT NULL,
	audit_success_ratio double precision NOT NULL,
	uptime_success_count bigint NOT NULL,
	total_uptime_count bigint NOT NULL,
	uptime_ratio double precision NOT NULL,
	created_at timestamp with time zone NOT NULL,
	updated_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE overlay_cache_nodes (
	node_id bytea NOT NULL,
	node_type integer NOT NULL,
	address text NOT NULL,
	protocol integer NOT NULL,
	operator_email text NOT NULL,
	operator_wallet text NOT NULL,
	free_bandwidth bigint NOT NULL,
	free_disk bigint NOT NULL,
	latency_90 bigint NOT NULL,
	audit_success_ratio double precision NOT NULL,
	audit_uptime_ratio double precision NOT NULL,
	audit_count bigint NOT NULL,
	audit_success_count bigint NOT NULL,
	uptime_count bigint NOT NULL,
	uptime_success_count bigint NOT NULL,
	PRIMARY KEY ( node_id ),
	UNIQUE ( node_id )
);
