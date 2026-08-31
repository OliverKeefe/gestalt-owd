package auth

import (
	"context"
	"errors"
	"slices"
	"strings"

	"github.com/google/uuid"
)

// UserIDFromCtx returns the Keycloak user ID stored in the context by ValidateJWT.
func UserIDFromCtx(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}

// ClaimsFromCtx returns the raw JWT claims map stored in the context.
func ClaimsFromCtx(ctx context.Context) (map[string]interface{}, bool) {
	claims, ok := ctx.Value(claimsKey).(map[string]interface{})
	return claims, ok
}

// typedClaimsFromCtx returns the parsed KeycloakClaims stored in the context.
func typedClaimsFromCtx(ctx context.Context) (*KeycloakClaims, bool) {
	claims, ok := ctx.Value(typedClaimsKey).(*KeycloakClaims)
	return claims, ok
}

// GroupsFromCtx returns the Keycloak group names assigned to the authenticated user.
func GroupsFromCtx(ctx context.Context) ([]string, bool) {
	claims, ok := typedClaimsFromCtx(ctx)
	if !ok || claims == nil {
		return nil, false
	}
	if len(claims.Groups) == 0 {
		return nil, false
	}
	return claims.Groups, true
}

// RealmRolesFromCtx returns the Keycloak realm roles assigned to the authenticated user.
func RealmRolesFromCtx(ctx context.Context) ([]string, bool) {
	claims, ok := typedClaimsFromCtx(ctx)
	if !ok || claims == nil {
		return nil, false
	}
	if len(claims.RealmAccess.Roles) == 0 {
		return nil, false
	}
	return claims.RealmAccess.Roles, true
}

// ClientRolesFromCtx returns the Keycloak client roles for the given client assigned to the authenticated user.
func ClientRolesFromCtx(ctx context.Context, clientID string) ([]string, bool) {
	claims, ok := typedClaimsFromCtx(ctx)
	if !ok || claims == nil {
		return nil, false
	}
	clientAccess, exists := claims.ResourceAccess[clientID]
	if !exists || len(clientAccess.Roles) == 0 {
		return nil, false
	}
	return clientAccess.Roles, true
}

// PreferredUsernameFromCtx returns the preferred username from the authenticated user's JWT claims.
func PreferredUsernameFromCtx(ctx context.Context) (string, bool) {
	claims, ok := typedClaimsFromCtx(ctx)
	if !ok || claims == nil {
		return "", false
	}
	if claims.PreferredUsername == "" {
		return "", false
	}
	return claims.PreferredUsername, true
}

// EmailFromCtx returns the email address from the authenticated user's JWT claims.
func EmailFromCtx(ctx context.Context) (string, bool) {
	claims, ok := typedClaimsFromCtx(ctx)
	if !ok || claims == nil {
		return "", false
	}
	if claims.Email == "" {
		return "", false
	}
	return claims.Email, true
}

// HasRealmRole reports whether the authenticated user holds the given Keycloak realm role.
func HasRealmRole(ctx context.Context, role string) bool {
	roles, ok := RealmRolesFromCtx(ctx)
	if !ok {
		return false
	}
	return slices.Contains(roles, role)
}

// HasClientRole reports whether the authenticated user holds the given role for the specified Keycloak client.
func HasClientRole(ctx context.Context, clientID, role string) bool {
	roles, ok := ClientRolesFromCtx(ctx, clientID)
	if !ok {
		return false
	}
	return slices.Contains(roles, role)
}

// HasGroup reports whether the authenticated user belongs to the given Keycloak group.
func HasGroup(ctx context.Context, group string) bool {
	groups, ok := GroupsFromCtx(ctx)
	if !ok {
		return false
	}
	return slices.Contains(groups, group)
}

// HasAnyGroup reports whether the authenticated user belongs to at least one of the given Keycloak groups.
func HasAnyGroup(ctx context.Context, groupNames []string) bool {
	groups, ok := GroupsFromCtx(ctx)
	if !ok {
		return false
	}
	groupSet := make(map[string]struct{}, len(groups))
	for _, g := range groups {
		groupSet[g] = struct{}{}
	}
	for _, name := range groupNames {
		if _, exists := groupSet[name]; exists {
			return true
		}
	}
	return false
}

// HasClaim verifies that the user ID in the context matches the provided ID and that the token contains the required scope claim.
func HasClaim(ctx context.Context, userID uuid.UUID, requiredClaim string) (bool, error) {
	userIDStoredInCtx, ok := UserIDFromCtx(ctx)
	if !ok {
		return false, errors.New("unauthenticated: no user ID found in context")
	}

	if userIDStoredInCtx != userID.String() {
		return false, errors.New("unauthorized: userID mismatch")
	}

	claims, ok := ClaimsFromCtx(ctx)
	if !ok {
		return false, errors.New("unable to get claims from context")
	}

	scopes, present := claims["scopes"]
	if !present {
		return false, errors.New("no scope claims found in token")
	}

	scopesStr, ok := scopes.(string)
	if !ok {
		return false, errors.New("couldn't convert scope claim to string")
	}

	hasClaim := slices.Contains(strings.Fields(scopesStr), requiredClaim)

	if !hasClaim {
		return false, errors.New("unauthorized: missing required scope")
	}

	return true, nil
}
