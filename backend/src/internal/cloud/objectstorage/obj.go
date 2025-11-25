package objectstorage

import (
	model "backend/src/core/files/model"
	"backend/src/internal/util"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"os"
	"path/filepath"
	"time"
)

type ObjectStore struct {
	NameSpace string
	BasePath  string
	Path      string
}

type ObjectMetaData struct {
	Key        string
	Size       uint64
	CreatedAt  time.Time
	ModifiedAt time.Time
	Metadata   map[string]string
}

/*
 - Store an obj with key (path / id) and metadata

 - Retrieve a blos data and metadata

 - Delete a obj

 - Check if a obj exists

 - List blobs

 - move or rename a obj

 - retrieve metadata without reading a full obj
*/

type Bucket struct {
	ID              uuid.UUID
	Name            string
	Size            uint64
	Creator         uuid.UUID
	Owner           uuid.UUID
	GroupMembership []uuid.UUID
}

// TODO: make params functional options pattern (options ... func(*Bucket)) *Bucket or config struct
func (Bucket) NewObjectBucket(ctx context.Context, conn pgx.Conn, name string, size uint64,
	creator uuid.UUID, owner uuid.UUID, groupMembership []uuid.UUID) (Bucket, error) {

	var id uuid.UUID

	for {
		id = uuid.New()
		unique, err := util.IsUUIDUnique(ctx, id, conn, "object_storage_buckets")
		if err != nil {
			return Bucket{}, err
		}
		if unique {
			break
		}
	}

	bucket := Bucket{
		ID:              id,
		Name:            name,
		Size:            size,
		Creator:         creator,
		Owner:           owner,
		GroupMembership: groupMembership,
	}

	return bucket, nil
}

var defaultStore *ObjectStore

func NewObjStore(basePath string, path string, ns string) *ObjectStore {
	return &ObjectStore{
		NameSpace: ns,
		BasePath:  basePath,
		Path:      path,
	}
}

func init() {
	defaultStore = NewObjStore("./obj", "/tmp/", "fileplace")
}

func (obj *ObjectStore) GetObjStore(filename string) {
	panic("not implemented yet")

}

func (obj *ObjectStore) Delete() (bool, error) {
	panic("not implemented yet")
}

func Save(blobstore *ObjectStore, file model.File) (bool, error) {
	if err := os.MkdirAll(blobstore.BasePath, os.ModePerm); err != nil {
		return false, fmt.Errorf("failed to create base path %e", err)
	}

	fullPath := filepath.Join(blobstore.BasePath, blobstore.Path, file.Metadata.Filename)

	err := os.WriteFile(fullPath, file.FileData, 0644)
	if err != nil {
		return false, fmt.Errorf("failed to write file %e", err)
	}

	return true, nil
}
