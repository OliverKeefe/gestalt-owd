package files

import (
	"github.com/google/uuid"
	"io/fs"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FileModel struct {
	db *pgxpool.Pool
}

type File struct {
	Metadata MetaData
	FileData FileData
}

type FileData struct {
	Filename string
	Reader   io.Reader
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
