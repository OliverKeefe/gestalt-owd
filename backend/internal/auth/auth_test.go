package auth

import (
	"context"
	"testing"
)

func TestUserIDFromCtx(t *testing.T) {
	tests := []struct {
		name           string
		ctx            context.Context
		expectedUserID string
		found          bool
	}{
		{
			name:           "success, user ID exists as string in context",
			ctx:            context.WithValue(context.Background(), userIDKey, "username"),
			expectedUserID: "username",
			found:          true,
		},
		{
			name:           "failure, key missing from context",
			ctx:            context.Background(),
			expectedUserID: "",
			found:          false,
		},
		{
			name:           "failure, key stored as incorrect type in context",
			ctx:            context.WithValue(context.Background(), userIDKey, 11111111),
			expectedUserID: "",
			found:          false,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			id, ok := UserIDFromCtx(testcase.ctx)
			if id != testcase.expectedUserID {
				t.Errorf("got id %q, want %q", id, testcase.expectedUserID)
			}
			if ok != testcase.found {
				t.Errorf("got ok %t, want %t", ok, testcase.found)
			}
		})
	}
}

// TestGroupsFromCtx verifies that group names stored in the JWT claims are returned with a positive presence flag.
func TestGroupsFromCtx(t *testing.T) {
	a := &Authenticator{}
	ctx := a.InjectKeycloakClaims(context.Background(), &KeycloakClaims{
		Groups: []string{"org-gestalt", "team-core"},
	})

	groups, ok := GroupsFromCtx(ctx)
	if !ok {
		t.Fatal("expected groups to be present")
	}
	if len(groups) != 2 || groups[0] != "org-gestalt" || groups[1] != "team-core" {
		t.Errorf("got groups %v, want [org-gestalt team-core]", groups)
	}
}

// TestGroupsFromCtx_Missing verifies that an empty context reports group absence.
func TestGroupsFromCtx_Missing(t *testing.T) {
	ctx := context.Background()
	if _, ok := GroupsFromCtx(ctx); ok {
		t.Error("expected groups to be absent")
	}
}

// TestRealmRolesFromCtx verifies that realm roles stored in the JWT claims are returned with a positive presence flag.
func TestRealmRolesFromCtx(t *testing.T) {
	a := &Authenticator{}
	ctx := a.InjectKeycloakClaims(context.Background(), &KeycloakClaims{
		RealmAccess: RealmAccess{Roles: []string{"admin", "user"}},
	})

	roles, ok := RealmRolesFromCtx(ctx)
	if !ok {
		t.Fatal("expected realm roles to be present")
	}
	if len(roles) != 2 || roles[0] != "admin" {
		t.Errorf("got roles %v, want [admin user]", roles)
	}
}

// TestClientRolesFromCtx verifies that client roles are returned for a known client and reported absent for an unknown one.
func TestClientRolesFromCtx(t *testing.T) {
	a := &Authenticator{}
	ctx := a.InjectKeycloakClaims(context.Background(), &KeycloakClaims{
		ResourceAccess: map[string]ClientAccess{
			"open-web-drive": {Roles: []string{"files:upload", "files:delete"}},
		},
	})

	roles, ok := ClientRolesFromCtx(ctx, "open-web-drive")
	if !ok {
		t.Fatal("expected client roles to be present")
	}
	if len(roles) != 2 || roles[0] != "files:upload" {
		t.Errorf("got roles %v, want [files:upload files:delete]", roles)
	}

	if _, ok := ClientRolesFromCtx(ctx, "other-client"); ok {
		t.Error("expected no roles for unknown client")
	}
}

// TestPreferredUsernameFromCtx verifies the preferred username is read from the JWT claims.
func TestPreferredUsernameFromCtx(t *testing.T) {
	a := &Authenticator{}
	ctx := a.InjectKeycloakClaims(context.Background(), &KeycloakClaims{
		PreferredUsername: "alice",
	})

	name, ok := PreferredUsernameFromCtx(ctx)
	if !ok {
		t.Fatal("expected preferred_username to be present")
	}
	if name != "alice" {
		t.Errorf("got %q, want alice", name)
	}
}

// TestEmailFromCtx verifies the email address is read from the JWT claims.
func TestEmailFromCtx(t *testing.T) {
	a := &Authenticator{}
	ctx := a.InjectKeycloakClaims(context.Background(), &KeycloakClaims{
		Email: "alice@example.com",
	})

	email, ok := EmailFromCtx(ctx)
	if !ok {
		t.Fatal("expected email to be present")
	}
	if email != "alice@example.com" {
		t.Errorf("got %q, want alice@example.com", email)
	}
}

// TestHasRealmRole verifies realm role membership matches only an assigned role.
func TestHasRealmRole(t *testing.T) {
	a := &Authenticator{}
	ctx := a.InjectKeycloakClaims(context.Background(), &KeycloakClaims{
		RealmAccess: RealmAccess{Roles: []string{"admin"}},
	})

	if !HasRealmRole(ctx, "admin") {
		t.Error("expected admin role to match")
	}
	if HasRealmRole(ctx, "user") {
		t.Error("did not expect user role to match")
	}
}

// TestHasClientRole verifies client role membership matches only an assigned role.
func TestHasClientRole(t *testing.T) {
	a := &Authenticator{}
	ctx := a.InjectKeycloakClaims(context.Background(), &KeycloakClaims{
		ResourceAccess: map[string]ClientAccess{
			"open-web-drive": {Roles: []string{"files:upload"}},
		},
	})

	if !HasClientRole(ctx, "open-web-drive", "files:upload") {
		t.Error("expected files:upload role to match")
	}
	if HasClientRole(ctx, "open-web-drive", "files:delete") {
		t.Error("did not expect files:delete to match")
	}
}

// TestHasGroup verifies group membership matches only an assigned group.
func TestHasGroup(t *testing.T) {
	a := &Authenticator{}
	ctx := a.InjectKeycloakClaims(context.Background(), &KeycloakClaims{
		Groups: []string{"org-gestalt"},
	})

	if !HasGroup(ctx, "org-gestalt") {
		t.Error("expected group to match")
	}
	if HasGroup(ctx, "other-group") {
		t.Error("did not expect other-group to match")
	}
}

// TestHasAnyGroup verifies that any group match succeeds and no match fails.
func TestHasAnyGroup(t *testing.T) {
	a := &Authenticator{}
	ctx := a.InjectKeycloakClaims(context.Background(), &KeycloakClaims{
		Groups: []string{"org-gestalt", "team-core"},
	})

	if !HasAnyGroup(ctx, []string{"missing", "team-core"}) {
		t.Error("expected at least one group to match")
	}
	if HasAnyGroup(ctx, []string{"missing", "elsewhere"}) {
		t.Error("did not expect any group to match")
	}
}
