package files

import (
	model "backend/src/core/files/model"
	repository "backend/src/core/files/repository"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UploadSvc struct {
	repo *repository.FileRepository
}

// MetaDataDTO is a Data Transfer Object for file Metadata.
// When the frontend sends a multipart form, metadata is stored
// as raw json,
type MetaDataDTO struct {
	Path             string `json:"path"`
	RelativePath     string `json:"relativePath"`
	LastModified     int64  `json:"lastModified"`
	LastModifiedDate string `json:"lastModifiedDate"`
	Size             uint64 `json:"size"`
	FileType         string `json:"fileType"`

	ID       string `json:"id"`
	OwnerID  string `json:"ownerId"`
	CheckSum string `json:"checkSum"`
}

// NewUploadService constructor for new UploadSvc (UploadService).
func NewUploadService(fileRepo *repository.FileRepository) *UploadSvc {
	return &UploadSvc{
		repo: fileRepo,
	}
}

func (svc *UploadSvc) Save(r *http.Request, ctx context.Context) error {
	mr, err := r.MultipartReader()
	if err != nil {
		return err
	}

	metadataByID := make(map[string]model.MetaData)

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		name := part.FormName()

		switch {
		// Handle Metadata.
		case strings.HasPrefix(name, "metadata-"):
			idStr := strings.TrimPrefix(name, "metadata-")

			// decode + build metadata
			var dto MetaDataDTO
			if err := json.NewDecoder(part).Decode(&dto); err != nil {
				return err
			}

			ownerID, err := uuid.Parse(dto.OwnerID)
			if err != nil {
				return err
			}

			metadataByID[idStr] = model.MetaData{
				ID:         uuid.MustParse(idStr),
				FileName:   dto.RelativePath,
				Path:       dto.Path,
				Size:       dto.Size,
				ModifiedAt: time.UnixMilli(dto.LastModified),
				CreatedAt:  time.Now(),
				Owner:      ownerID,
				Version:    time.Now(),
			}

		// Handle Part containing file's binary data.
		case strings.HasPrefix(name, "file-"):
			// File has to be saved here, if you try to pass this to another temp location
			// in memory then the data will be unusable.
			if err := svc.saveFileData(
				"/home/oliver/Development/25-26_CE301_keefe_oliver_b/backend/tempfiles",
				part,
				part.FileName(),
			); err != nil {
				return err
			}
		}
	}

	// Persist file metadata
	for _, md := range metadataByID {
		if err := svc.repo.SaveMetaData(md, ctx); err != nil {
			return err
		}
	}

	return nil
}

// Helper method to save FilePart binary data.
func (svc *UploadSvc) saveFileData(
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
