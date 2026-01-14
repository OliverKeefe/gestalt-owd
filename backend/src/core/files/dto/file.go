package dto

import (
	files "backend/src/core/files/model"
	"time"

	"github.com/google/uuid"
)

type FileDTO struct {
	ID       uuid.UUID `json:"uuid"`
	FileName string    `json:"file_name"`
	Path     string    `json:"path"`
	Size     uint64    `json:"size"`
	FileType string    `json:"file_type"`
	//Mode       fs.FileMode
	//IsDir      bool
	ModifiedAt time.Time   `json:"modified_at"`
	CreatedAt  time.Time   `json:"created_at"`
	Owner      uuid.UUID   `json:"owner_id"`
	AccessTo   []uuid.UUID `json:"access_to"`
	Group      []uuid.UUID `json:"group_id"`
	//Links      *uint64
	Version time.Time `json:"version"`
}

func MapTo(m files.MetaData) FileDTO {
	return FileDTO{
		ID:         m.ID,
		FileName:   m.FileName,
		Path:       m.Path,
		Size:       m.Size,
		FileType:   m.FileType,
		CreatedAt:  m.CreatedAt,
		ModifiedAt: m.ModifiedAt,
		Owner:      m.Owner,
		AccessTo:   m.AccessTo,
		Group:      m.Group,
		Version:    m.Version,
	}
}
