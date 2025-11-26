package objectstorage

import (
	model "backend/src/core/files/model"
	"fmt"
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
