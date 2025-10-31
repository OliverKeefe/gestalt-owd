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

func (blob Store) Save(file string) (bool, error) {
	panic("not implemented yet")
}

func (blob Store) Get(filename string) { panic("not implemented yet") }

func (blob Store) Delete() (bool, error) { panic("not implemented yet") }
