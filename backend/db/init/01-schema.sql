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
    tenant_id UUID NOT NULL,
    name VARCHAR(64) NOT NULL
);

-- GROUPS
CREATE TABLE groups (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(64) NOT NULL
);

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