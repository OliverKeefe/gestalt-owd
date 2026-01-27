package objectstorage

import (
	"context"
	"io"

	"github.com/google/uuid"
)

type ObjectStore interface {
	PutStream(
		ctx context.Context,
		key string,
		reader io.Reader,
		size uint64,
		contentType string,
		metadata map[string]string,
	) error

	GetStream(
		ctx context.Context,
		key string,
		start, end *uint64,
	) (io.ReadCloser, error)

	Head(
		ctx context.Context,
		key string,
	) (*ObjectInfo, error)

	Delete(
		ctx context.Context,
		key string,
	) error

	CheckObjectStoreExists(
		ctx context.Context,
		key string,
		ID uuid.UUID,
	) bool

	GetRemotePath(ctx context.Context,
		path string,
		key string,
		ID uuid.UUID,
	) (string, error)
}

type ObjectInfo struct {
	Size        int64
	ETag        string
	ContentType string
	Metadata    map[string]string
}
