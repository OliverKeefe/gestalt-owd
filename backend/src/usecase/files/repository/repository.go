package files

import (
	"backend/src/internal/db/metadb"
	data "backend/src/usecase/files/data"
	"context"
	"fmt"
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

// TODO: add the checksum field
func (repo *Repository) SaveMetaData(meta data.MetaData, ctx context.Context) error {
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
    	    owner_id
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
		meta.UploadedAt,
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

func (repo *Repository) GetAllFiles(ctx context.Context, req data.GetAllMetadataRequest) ([]data.MetaData, error) {
	var db = repo.db.Pool

	const query = `
		SELECT
			id,
			file_name,
			path,
			size,
			file_type,
			modified_at,
			uploaded_at,
			version,
			checksum,
			owner_id
		FROM file_metadata
		WHERE owner_id = $1
  			AND (modified_at, id) < ($2, $3)
		ORDER BY modified_at, id 
		LIMIT 20;
	`

	var result []data.MetaData

	rows, err := db.Query(
		ctx,
		query,
		req.UserID,
		req.Cursor.ModifiedAt,
		req.Cursor.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var model data.MetaData
		if err := rows.Scan(
			&model.ID,
			&model.FileName,
			&model.Path,
			&model.Size,
			&model.FileType,
			&model.ModifiedAt,
			&model.UploadedAt,
			&model.Owner,
			&model.AccessTo,
			&model.Group,
			&model.Version,
		); err != nil {
			return nil, err
		}

		result = append(result, model)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *Repository) GetFiles(ctx context.Context, model data.MetaData) ([]data.MetaData, error) {
	db := repo.db.Pool
	var result []data.MetaData

	query, args, err := GetMetadataQuery(model)
	if err != nil {
		fmt.Errorf("unable to build GetMetadataQuery, %v", err)
		return nil, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var model data.MetaData
		if err := rows.Scan(
			&model.ID,
			&model.FileName,
			&model.Path,
			&model.Size,
			&model.FileType,
			&model.ModifiedAt,
			&model.UploadedAt,
			&model.Owner,
			&model.AccessTo,
			&model.Group,
			&model.Version,
		); err != nil {
			return nil, err
		}

		result = append(result, model)
	}

	return result, nil
}
