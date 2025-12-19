package files

import (
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"
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

func (m *FileModel) Save(ctx context.Context, filedata FileData, metadata MetaData) error {
	basePath := "/home/oliver/Development/25-26_CE301_keefe_oliver_b/backend/"
	tempDir := filepath.Join(basePath, "tempfiles")
	fullPath := filepath.Join(tempDir, metadata.FileName)

	if err := m.saveFileData(fullPath, filedata); err != nil {
		return err
	}

	if err := m.saveMetadata(ctx, metadata); err != nil {
		return err
	}

	return nil
}

func (m *FileModel) saveFileData(filepath string, fileData FileData) error {
	dest, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(dest *os.File) {
		err := dest.Close()
		if err != nil {

		}
	}(dest)

	_, err = io.Copy(dest, fileData.Reader)
	return err
}

func (m *FileModel) saveMetadata(ctx context.Context, metadata MetaData) error {
	const query = `
		INSERT INTO metadata (
		                      file_name,
		                      path,
		                      size,
		                      owner,
		                      created_at,
		                      modified_at
		) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := m.db.Exec(
		ctx,
		query,
		metadata.FileName,
		metadata.Path,
		metadata.Size,
		metadata.Owner,
		metadata.CreatedAt,
	)

	return err
}
