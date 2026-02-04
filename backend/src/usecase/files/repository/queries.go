package files

import (
	query "backend/src/internal/api/query"
	files "backend/src/usecase/files/data"
	"fmt"
	"log"
	"strings"
)

func PersistMetadataQuery(m files.MetaData) (string, []any) {
	const sql = `
    	INSERT INTO file_metadata (id, file_name, path, size, file_type, modified_at,
    	    uploaded_at, version, owner_id, checksum) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);
    `

	args := []any{
		m.ID,
		m.FileName,
		m.Path,
		m.Size,
		m.FileType,
		m.ModifiedAt,
		m.UploadedAt,
		m.Version,
		m.Owner,
		m.CheckSum,
	}

	log.Printf("checksum len=%d hex=%x", len(m.CheckSum), m.CheckSum)
	return sql, args
}

func FindMetadataQuery(m files.MetaData) (string, []any) {
	const baseQuery = `SELECT id, file_name, path, size, file_type, modified_at, 
       	uploaded_at, version, checksum, owner_id 
		FROM file_metadata
	`

	var b query.Builder

	b.Equal("id", m.ID)
	b.Equal("file_name", m.FileName)
	b.Equal("path", m.Path)
	b.Equal("size", m.Size)
	b.Equal("file_type", m.FileType)
	b.Equal("modified_at", m.ModifiedAt)
	b.Equal("uploaded_at", m.UploadedAt)
	b.Equal("owner_id", m.Owner)
	b.Equal("version", m.Version)
	if len(m.Group) > 0 {
		const selectGroups = `
			EXISTS (
				SELECT * FROM file_metadata_group_access fga 
				WHERE fga.file_id = file_metadata.id
					AND fga.user_id = ANY (?) 
			)`
		b.Raw(selectGroups, m.Group)
	}
	if len(m.AccessTo) > 0 {
		const accessTo = `
			EXISTS (
				SELECT * FROM file_user_access fua
				WHERE fua.file_id = file_medata.id
					AND fua.user_id = ANY (?)
			)`
		b.Raw(accessTo, m.AccessTo)
	}

	sql := fmt.Sprintf("%s WHERE %s", baseQuery, strings.Join(b.Clauses, " AND "))
	args := b.Args

	return sql, args
}
