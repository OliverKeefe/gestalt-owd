package blob

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Store struct {
	BasePath string
	Path     string
}

var defaultStore *Store

func NewBlobStore(basePath string, path string) *Store {
	return &Store{
		BasePath: basePath,
		Path:     path,
	}
}

func (blob Store) Save(file string) (bool, error) {
	panic("not implemented yet")
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
