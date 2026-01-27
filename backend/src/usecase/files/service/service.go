package files

import (
	data "backend/src/usecase/files/data"
	repository "backend/src/usecase/files/repository"
	"context"
	"encoding/json"
	"io"
	"log"
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

// multipartMetadata is a Data Transfer Object for multipart form for file uploads
// from the web gui frontend.
// When the web frontend sends a multipart form, metadata is stored
// as raw json.
type multipartMetadata struct {
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

func (svc *Service) Upload(r *http.Request, ctx context.Context) error {
	mr, err := r.MultipartReader()
	if err != nil {
		return err
	}

	metadataByID := make(map[string]data.MetaData)

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
			var decodedRequest multipartMetadata
			if err := json.NewDecoder(part).Decode(&decodedRequest); err != nil {
				return err
			}

			ownerID, err := uuid.Parse(decodedRequest.OwnerID)
			if err != nil {
				return err
			}

			metadataByID[idStr] = data.MetaData{
				ID:         uuid.MustParse(idStr),
				FileName:   decodedRequest.RelativePath,
				Path:       decodedRequest.Path,
				Size:       decodedRequest.Size,
				ModifiedAt: time.UnixMilli(decodedRequest.LastModified),
				UploadedAt: time.Now(),
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

func (svc *Service) GetAll(ctx context.Context, request data.GetAllMetadataRequest) ([]data.MetaDataResponse, error) {
	var (
		repo     = svc.repo
		files    []data.MetaData
		response []data.MetaDataResponse
	)

	files, err := repo.GetAllFiles(ctx, request)
	if err != nil {
		log.Printf("unable to get all files for user: %s, %v ", request.UserID, err)
	}

	for _, file := range files {
		file := file.ToResponse()
		if err != nil {
			log.Printf("unable to map file metadata: %s, to dto: %v", file, err)
		}
		response = append(response, file)
	}

	return response, nil
}

func (svc *Service) GetMetadata(ctx context.Context, request data.GetMetadataRequest) ([]data.MetaData, error) {
	var (
		repo  = svc.repo
		files []data.MetaData
	)

	model := request.ToModel()
	files, err := repo.GetFiles(ctx, model)
	if err != nil {
		log.Printf("unable to get file metadata: %v", err)
		return files, err
	}

	return files, nil
}

func (svc *Service) Delete(ctx context.Context, request data.DeleteRequest) error {

	err := svc.repo.DeleteMetadata(ctx, request.ID, request.OwnerID)
	if err != nil {
		log.Printf("could not delete file metadata, %v", err)
		return err
	}

	return nil
}
