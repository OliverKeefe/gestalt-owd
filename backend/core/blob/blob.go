package blob

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Store struct {
	BasePath string
}

func NewBlobStore(basePath string) *Store {
	return &Store{BasePath: basePath}
}
