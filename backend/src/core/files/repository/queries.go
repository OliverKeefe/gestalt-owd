package files

import (
	files "backend/src/core/files/model"
	"backend/src/internal/metadb"
	"context"
	"log"
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
    	    id,
    	    file_name,
    	    path,
    	    size,
    	    file_type,
    	    modified_at,
    	    uploaded_at,
    	    version,
    	    owner
    	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
    `

	log.Printf("metadata=%+v", meta)

	var db = repo.db.Pool
	status, err := db.Exec(
		ctx,
		query,
		meta.ID,
		meta.FileName,
		meta.Path,
		meta.Size,
		meta.FileType,
		meta.ModifiedAt,
		meta.CreatedAt,
		meta.Version,
		meta.Owner,
	)

	log.Printf("DB Status: %s", status)

	return err
}
