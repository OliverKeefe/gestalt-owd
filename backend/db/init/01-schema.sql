\connect appdb;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- USERS
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(64)
);

-- TENANTS
CREATE TABLE tenant (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(64) NOT NULL
);

-- GROUPS
CREATE TABLE groups (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(64) NOT NULL
);

-- BUCKETS
CREATE TABLE bucket (
    id UUID PRIMARY KEY NOT NULL,
    tenant_id UUID NOT NULL REFERENCES tenant(id) ON DELETE CASCADE,
    name VARCHAR(64) NOT NULL,
    size BIGINT NOT NULL,
    creator UUID NOT NULL REFERENCES users(id),
    owner UUID NOT NULL REFERENCES users(id)
);

-- BUCKET GROUP MEMBERSHIP
CREATE TABLE bucket_groups (
    bucket_id UUID NOT NULL REFERENCES bucket(id) ON DELETE CASCADE,
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    PRIMARY KEY (bucket_id, group_id)
);

-- FILE METADATA
CREATE TABLE file_metadata(
    id UUID PRIMARY KEY NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL,
    -- mode BIGINT NOT NULL,
    file_type VARCHAR(255),
    modified_at TIMESTAMP, --NOT NULL,
    uploaded_at TIMESTAMP, --NOT NULL,
    version TIMESTAMP, --NOT NULL,
    checksum BYTEA,
    owner UUID NOT NULL
    -- owner UUID NOT NULL REFERENCES users(id)
);

-- INDEX FOR FILE METADATA PAGINATION
CREATE INDEX CONCURRENTLY idx_file_owner_modified_desc
ON file_metadata (owner, modified_at DESC, id DESC);

-- DIRECTORY METADATA
CREATE TABLE dir_metadata(
    id UUID PRIMARY KEY NOT NULL,
    dir_name VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    uploaded_at TIMESTAMP NOT NULL,
    owner UUID NOT NULL REFERENCES users(id)
);

-- FILE ACCESS
CREATE TABLE file_metadata_access(
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    file_id UUID NOT NULL REFERENCES file_metadata(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, file_id)
);