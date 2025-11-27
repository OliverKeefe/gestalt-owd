package objectstorage

import (
	gestaltofs "backend/src/internal/os/filesystem"

	"github.com/google/uuid"
)

type Store struct {
	RemotePath string
}

func NewStore() {

}

type StoreObject interface {
	CheckBucketExists(bucketId uuid.UUID) (bool, error)
	CopyFileToBucket(bucketId uuid.UUID, remotePath string) (bool, error)
	RemotePath(meta *gestaltofs.MetaData)
}
