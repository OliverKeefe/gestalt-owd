package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

func IsUUIDUnique(ctx context.Context, id uuid.UUID, db pgx.Conn, tableName string) (bool, error) {
	allowedTables := map[string]bool{
		"object_storage_buckets":              true,
		"object_storage_hierarchical_buckets": true,
	}
	if !allowedTables[tableName] {
		return false, fmt.Errorf("invalid table name not allowed")
	}

	query := fmt.Sprintf(`SELECT id FROM %s WHERE id = $1`, tableName)

	var found uuid.UUID
	err := db.QueryRow(query, id, ctx).Scan(&found)

	if errors.Is(err, pgx.ErrNoRows) {
		return true, nil
	}

	return false, nil
}
