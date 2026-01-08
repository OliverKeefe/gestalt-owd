package files

import (
	"backend/src/core/files/dto"
	model "backend/src/core/files/model"
	repository "backend/src/core/files/repository"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
}

// NewService constructor for new Service (UploadService).
func NewService(fileRepo *repository.Repository) *Service {
	return &Service{
		repo: fileRepo,
	}
}

func (svc *Service) Upload(r *http.Request, ctx context.Context) error {
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
			var decodedRequest dto.MetaDataDTO
			if err := json.NewDecoder(part).Decode(&decodedRequest); err != nil {
				return err
			}

			ownerID, err := uuid.Parse(decodedRequest.OwnerID)
			if err != nil {
				return err
			}

			metadataByID[idStr] = model.MetaData{
				ID:         uuid.MustParse(idStr),
				FileName:   decodedRequest.RelativePath,
				Path:       decodedRequest.Path,
				Size:       decodedRequest.Size,
				ModifiedAt: time.UnixMilli(decodedRequest.LastModified),
				CreatedAt:  time.Now(),
				Owner:      ownerID,
				Version:    time.Now(),
			}

		// Handle Part containing file's binary data.
		case strings.HasPrefix(name, "file-"):
			// File has to be saved here, if you try to pass this to another temp location
			// in memory then the data will be unusable.
			if err := svc.repo.SaveFileData(
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
