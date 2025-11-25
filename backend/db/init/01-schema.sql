type Bucket struct {
	ID              uuid.UUID
	TenantID        uuid.UUID
	Name            string
	Size            uint64
	Creator         uuid.UUID
	Owner           uuid.UUID
	GroupMembership []uuid.UUID
}

\connect appdb;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS bucket (
    id UUID PRIMARY KEY NOT NULL,
    tenant_id UUID NOT NULL,
    name VARCHAR(64) NOT NULL,
    size BIGINT NOT NULL,
    creator UUID NOT NULL,
    owner UUID NOT NULL
);

CREATE TABLE IF NOT EXISTS bucket_groups (
    bucket_id UUID NOT NULL,
    group_id UUID NOT NULL
);

CREATE TABLE IF NOT EXISTS tenant (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(64) NOT NULL
);