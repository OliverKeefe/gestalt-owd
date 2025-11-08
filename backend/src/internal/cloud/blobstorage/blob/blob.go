package gestaltoblob

import (
	model "backend/src/core/files/model"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	BasePath string
	Path     string
}

type BlobMetaData struct {
	Key        string
	Size       uint64
	CreatedAt  time.Time
	ModifiedAt time.Time
	MetaData   map[string]string
}

var defaultStore *Store

func NewBlobStore(basePath string, path string) *Store {
	return &Store{
		BasePath: basePath,
		Path:     path,
	}
}

func Init() {
	defaultStore = NewBlobStore("./blob", "/tmp/")
}

//func (blob Store) Get(filename string) { panic("not implemented yet") }
//
//func (blob Store) Delete() (bool, error) { panic("not implemented yet") }

func Save(blobstore Store, file model.File) (bool, error) {
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
