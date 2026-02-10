package files

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Metadata Model - need ContentCID string for IPFS
type MetaData struct {
	ID         uuid.UUID   `json:"uuid"`
	FileName   string      `json:"file_name"`
	Path       string      `json:"path"`
	Size       uint64      `json:"size"`
	FileType   string      `json:"file_type"`
	ModifiedAt time.Time   `json:"modified_at"`
	UploadedAt time.Time   `json:"created_at"`
	Owner      uuid.UUID   `json:"owner_id"`
	AccessTo   []uuid.UUID `json:"access_to"`
	Group      []uuid.UUID `json:"group_id"`
	CheckSum   []byte      `json:"checksum"`
	Version    time.Time   `json:"version"`
}

func (m *MetaData) ToResponse() MetaDataResponse {
	return MetaDataResponse{
		ID:         m.ID,
		FileName:   m.FileName,
		Path:       m.Path,
		Size:       m.Size,
		FileType:   m.FileType,
		ModifiedAt: m.ModifiedAt,
		UploadedAt: m.UploadedAt,
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
	UploadedAt time.Time   `json:"created_at"`
	Owner      uuid.UUID   `json:"owner_id"`
	AccessTo   []uuid.UUID `json:"access_to"`
	Group      []uuid.UUID `json:"group_id"`
	CheckSum   []byte      `json:"checksum"`
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

type FindMetadataRequest struct {
	ID         uuid.UUID   `json:"file_id"`
	FileName   string      `json:"file_name,omitempty"`
	Path       string      `json:"path,omitempty"`
	Size       uint64      `json:"size,omitempty"`
	FileType   string      `json:"file_type,omitempty"`
	ModifiedAt time.Time   `json:"modified_at,omitempty"`
	UploadedAt time.Time   `json:"uploaded_at,omitempty"`
	Owner      uuid.UUID   `json:"owner_id"`
	AccessTo   []uuid.UUID `json:"access_to,omitempty"`
	Group      []uuid.UUID `json:"group_id,omitempty"`
	CheckSum   []byte      `json:"checksum,omitempty"`
	Version    time.Time   `json:"version,omitempty"`
}

type MetadataCursor struct {
	ModifiedAt time.Time `json:"modified_at"`
	ID         uuid.UUID `json:"id"`
}

func (req *FindMetadataRequest) ToModel() MetaData {
	return MetaData{
		ID:         req.ID,
		FileName:   req.FileName,
		Path:       req.Path,
		Size:       req.Size,
		FileType:   req.FileType,
		ModifiedAt: req.ModifiedAt,
		UploadedAt: req.UploadedAt,
		Owner:      req.Owner,
		AccessTo:   req.AccessTo,
		Group:      req.Group,
		CheckSum:   req.CheckSum,
		Version:    req.Version,
	}
}

func (req *FindMetadataRequest) Bind(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	return nil
}

type DeleteRequest struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
}

func (req DeleteRequest) Bind(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	return nil
}
