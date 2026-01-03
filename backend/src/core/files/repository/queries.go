package files

import (
	files "backend/src/core/files/model"
	"backend/src/internal/metadb"
	"context"
)

type FileRepository struct {
	db *metadb.MetadataDatabase
}

func NewFileRepository(db *metadb.MetadataDatabase) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

func (repo *FileRepository) SaveMetaData(meta files.MetaData, ctx context.Context) error {
	const query = `
		INSERT INTO file_metadata (
		                      file_name,
		                      path,
		                      size,
		                      modified_at,
		                      version,
		                      owner
		) VALUES ($1, $2, $3, $4, $5, $6)
	`

	var db = repo.db.Pool
	_, err := db.Exec(
		ctx,
		query,
		meta.FileName,
		meta.Path,
		meta.Size,
		meta.ModifiedAt,
		meta.Version,
		meta.Owner,
	)

	return err
}
