package files

import (
	model "backend/src/core/files/model"
	repository "backend/src/core/files/repository"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
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

	var (
		metadata model.MetaData
		fileData model.FileData
	)

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch part.FormName() {
		case "metadata":
			if err := json.NewDecoder(part).Decode(&metadata); err != nil {
				return err
			}
		case "file":
			fileData = model.FileData{
				Filename: part.FileName(),
				Reader:   part,
			}
		}
	}

	if err := svc.repo.SaveMetaData(metadata, ctx); err != nil {
		return err
	}

	if err := svc.saveFileData("backend/tempfiles", fileData); err != nil {
		return err
	}

	return nil
}

func (svc *UploadSvc) saveFileData(path string, data model.FileData) error {
	err := os.MkdirAll(path, 0655)
	if err != nil {
		return err
	}

	tmp, err := os.CreateTemp(path, data.Filename+"-*")
	if err != nil {
		return err
	}
	defer func(tmp *os.File) {
		err := tmp.Close()
		if err != nil {
			return
		}
	}(tmp)

	if _, err := io.Copy(tmp, data.Reader); err != nil {
		return err
	}

	return nil
}
