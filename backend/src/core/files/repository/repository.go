package files

import (
	files "backend/src/core/files/model"
	"backend/src/internal/metadb"
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type Repository struct {
	db *metadb.MetadataDatabase
}

func NewRepository(db *metadb.MetadataDatabase) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) SaveMetaData(meta files.MetaData, ctx context.Context) error {
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

// Helper method to save FilePart binary data.
func (repo *Repository) SaveFileData(
	basePath string,
	part *multipart.Part,
	filename string,
) error {
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return err
	}

	fileExtension := filepath.Ext(filename)
	base := strings.TrimSuffix(filename, fileExtension)

	tmp, err := os.CreateTemp(basePath, base+"-*"+fileExtension)
	if err != nil {
		return err
	}
	defer tmp.Close()

	if _, err := io.Copy(tmp, part); err != nil {
		return err
	}

	return nil
}
