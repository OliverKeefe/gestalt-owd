package files

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Metadata Model
type MetaData struct {
	ID         uuid.UUID
	FileName   string
	Path       string
	Size       uint64
	FileType   string
	ModifiedAt time.Time
	CreatedAt  time.Time
	Owner      uuid.UUID
	AccessTo   []uuid.UUID
	Group      []uuid.UUID
	CheckSum   string
	Version    time.Time
}

func (m *MetaData) ToResponse() MetaDataResponse {
	return MetaDataResponse{
		ID:         m.ID,
		FileName:   m.FileName,
		Path:       m.Path,
		Size:       m.Size,
		FileType:   m.FileType,
		ModifiedAt: m.ModifiedAt,
		CreatedAt:  m.CreatedAt,
		Owner:      m.Owner,
		AccessTo:   m.AccessTo,
		Group:      m.Group,
		CheckSum:   m.CheckSum,
		Version:    m.Version,
	}
}

type MetaDataResponse struct {
	ID         uuid.UUID   `json:"uuid"`
	FileName   string      `json:"file_name"`
	Path       string      `json:"path"`
	Size       uint64      `json:"size"`
	FileType   string      `json:"file_type"`
	ModifiedAt time.Time   `json:"modified_at"`
	CreatedAt  time.Time   `json:"created_at"`
	Owner      uuid.UUID   `json:"owner_id"`
	AccessTo   []uuid.UUID `json:"access_to"`
	Group      []uuid.UUID `json:"group_id"`
	CheckSum   string      `json:"checksum"`
	Version    time.Time   `json:"version"`
}

type GetAllMetadataRequest struct {
	UserID uuid.UUID       `json:"user_id"`
	Cursor *MetadataCursor `json:"cursor"`
	Limit  int             `json:"limit"`
}

func (req *GetAllMetadataRequest) Bind(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("invalid json request: %w", err)
	}
	return nil
}

type GetMetadataRequest struct {
	UserID uuid.UUID `json:"user_id"`
	FileID uuid.UUID `json:"file_id"`
}

type MetadataCursor struct {
	ModifiedAt time.Time `json:"modified_at"`
	ID         uuid.UUID `json:"id"`
}
