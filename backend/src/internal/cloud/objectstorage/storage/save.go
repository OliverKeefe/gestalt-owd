package objectstorage

import (
	model "backend/src/core/files/model"
	gestaltofs "backend/src/internal/os/filesystem"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func Save(file model.File) (bool, error) {
	basePath := "/home/oliver/Development/25-26_CE301_keefe_oliver_b/backend/"
	path := "tempfiles"

	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return false, fmt.Errorf("failed to create base path %e", err)
	}

	fullPath := filepath.Join(basePath, path)

	err := os.WriteFile(fullPath, file.FileData, 0644)
	if err != nil {
		return false, fmt.Errorf("failed to write file %e", err)
	}

	return true, nil
}

func checkBucketExists(bucketId uuid.UUID) (bool, error) {
	panic("Not implemented.")
}

func copyFileToBucket(bucketId uuid.UUID, remotePath string) (bool, error) {
	panic("Not implemented.")
}

func setRemotePath(meta *gestaltofs.MetaData) (string, error) {
	panic("Not implemented yet.")
}
