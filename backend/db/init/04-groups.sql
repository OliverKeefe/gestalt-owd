-- GROUPS
-- Groups map one-to-one with Keycloak groups. keycloak_group_id stores the
-- Keycloak group UUID so the backend can correlate JWT group membership.
-- name matches the Keycloak group name, the value in the JWT `groups` claim.
CREATE TABLE groups (
    id UUID PRIMARY KEY NOT NULL,
    org_id UUID NOT NULL REFERENCES orgs(id) ON DELETE CASCADE,
    name VARCHAR(64) NOT NULL,
    keycloak_group_id VARCHAR(64),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT uq_groups_org_name UNIQUE (org_id, name),
    CONSTRAINT uq_groups_keycloak_id UNIQUE (keycloak_group_id)
);

-- GROUP MEMBERSHIPS
CREATE TABLE group_memberships (
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT pk_group_memberships PRIMARY KEY (group_id, user_id)
);
