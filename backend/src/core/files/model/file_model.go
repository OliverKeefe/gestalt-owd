package files

import (
	"github.com/google/uuid"
	"io/fs"
	"time"
)

type File struct {
	Metadata MetaData
	FileData []byte
}

type MetaData struct {
	FileName   string
	Path       string
	Size       uint64
	Mode       fs.FileMode
	IsDir      bool
	ModifiedAt time.Time
	CreatedAt  time.Time
	Owner      uuid.UUID
	AccessTo   []uuid.UUID
	Group      []uuid.UUID
	Links      *uint64
}
