package files

import (
	"backend/internal/auth"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrUnauthenticated = errors.New("unauthorized: no user id in context")
	ErrForbidden       = errors.New("forbidden: user does not have access to this file")
)

// groupOwnerRepo is the subset of repository capabilities needed to resolve
// group ownership during authorization.
type groupOwnerRepo interface {
	OwnerIsInUserGroups(ctx context.Context, ownerID uuid.UUID, groupNames []string) (bool, error)
}

// authorizeFile returns the authenticated user's UUID if they are allowed to
// access a file owned by ownerID. Access is granted when the user owns the
// file directly, or when the file is owned by a group the user belongs to
// (resolved from the JWT groups claim).
func authorizeFile(ctx context.Context, repo groupOwnerRepo, ownerID uuid.UUID) (uuid.UUID, error) {
	userID, ok := auth.UserIDFromCtx(ctx)
	if !ok {
		return uuid.Nil, ErrUnauthenticated
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user id in jwt claims: %w", err)
	}

	if ownerID == userUUID {
		return userUUID, nil
	}

	groupNames, ok := auth.GroupsFromCtx(ctx)
	if !ok {
		return uuid.Nil, ErrForbidden
	}

	inGroup, err := repo.OwnerIsInUserGroups(ctx, ownerID, groupNames)
	if err != nil {
		return uuid.Nil, err
	}
	if !inGroup {
		return uuid.Nil, ErrForbidden
	}

	return userUUID, nil
}
