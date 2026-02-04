package files

import (
	"backend/src/internal/db/metadb"
	data "backend/src/usecase/files/data"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
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
    	INSERT INTO file_metadata (id, file_name, path, size, file_type, modified_at,
    	    uploaded_at, version, owner_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);
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
func (repo *Repository) SaveFileData(basePath string, part *multipart.Part, filename string) error {
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

func (repo *Repository) GetAllFiles(
	ctx context.Context,
	req data.GetAllMetadataRequest,
) ([]data.MetaData, error) {

	var (
		rows pgx.Rows
		err  error
	)

	if req.Cursor == nil || req.Cursor.ID == uuid.Nil || req.Cursor.ModifiedAt.IsZero() {
		rows, err = repo.db.Pool.Query(ctx, `
			SELECT id, file_name, path, size, file_type, modified_at,
				uploaded_at, owner_id, checksum, version 
			FROM file_metadata
			WHERE owner_id = $1
			ORDER BY modified_at DESC, id DESC
			LIMIT $2;
		`, req.UserID, req.Limit)

	} else {
		rows, err = repo.db.Pool.Query(ctx, `
			SELECT id, file_name, path, size, file_type, modified_at,
				uploaded_at, owner_id, checksum, version
			FROM file_metadata
			WHERE owner_id = $1
				AND (modified_at, id) < ($2, $3)
			ORDER BY modified_at DESC, id DESC
			LIMIT $4;
		`, req.UserID, req.Cursor.ModifiedAt, req.Cursor.ID, req.Limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]data.MetaData, 0, req.Limit)

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
			//&model.AccessTo,
			//&model.Group,
			&model.CheckSum,
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

func (repo *Repository) FindMetadata(ctx context.Context, model data.MetaData) ([]data.MetaData, error) {
	var result []data.MetaData

	query, args := FindMetadataQuery(model)

	rows, err := repo.db.Pool.Query(ctx, query, args...)
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

func (repo *Repository) DeleteMetadata(ctx context.Context, id uuid.UUID, ownerId uuid.UUID) error {
	const query = `DELETE FROM file_metadata WHERE id = $1 AND owner_id = $2;`

	status, err := repo.db.Pool.Exec(ctx, query, id, ownerId)
	if err != nil {
		return fmt.Errorf(
			"status: %s, could not delete file metadata, %w",
			status,
			err,
		)
	}

	rows := status.RowsAffected()

	if rows == 0 {
		return errors.New("no record found")
	}

	return nil
}

func (repo *Repository) Modify(ctx context.Context, model data.MetaData) (error, data.MetaData) {
	panic("not implemented")
}

func (repo *Repository) MarkForDeletion(ctx context.Context, id uuid.UUID, id2 uuid.UUID) error {
	panic("not implemented")
	// Should add flag to relation in metadata to delete x days (depending on user policy
	// default = 30d), then need cleaner func to bulk cleanse metadata db and before this,
	// remove file data from both virtual disc and IPFS node.
}
