package files

import (
	model "backend/src/core/files/model"
	obj "backend/src/internal/cloud/objectstorage/storage"
	"backend/src/internal/middleware"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type DefaultStore struct {
	BasePath  string
	Path      string
	NameSpace string
}

type UploadFile struct {
	model   model.FileModel
	Storage *obj.Store
}

func NewUploadFile(repo FileRepository, storage *obj.Store) {
	return &UploadFile{
		Repo:    repo,
		Storage: storage,
	}
}



func (uc *UploadFile) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/files/upload",
		middleware.EnableCORS(http.HandlerFunc(uc.Api)))
}

type UploadRequest struct {
	Metadata []model.MetaData
	Files    []model.FileData
}

func parseUploadRequest(r *http.Request) (UploadRequest, error) {
	var req UploadRequest
	mr, err := r.MultipartReader()
	if err != nil {
		return req, err
	}

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return req, err
		}

		switch part.FormName() {
		case "metadata":
			if err := json.NewDecoder(part).Decode(&req.Metadata); err != nil {
				return req, err
			}

		case "files":
			if part.FileName() == "" {
				continue
			}

			req.Files = append(req.Files, model.FileData{
				Filename: part.FileName(),
				Reader:   part,
			})
		}
	}

	if len(req.Files) == 0 {
		return req, errors.New("no files uploaded")
	}

	return req, nil
}

func (uc *UploadFile) Api(w http.ResponseWriter, r *http.Request) {
	req, err := parseUploadRequest(r)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := uc.upload(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	saveTo, err := storage.Save(defaultStore, uploadedFile)
	if err != nil {
		return false, fmt.Errorf("unable to save file in obj %e", err)
	w.WriteHeader(http.StatusAccepted)
}

	}

	return saveTo, nil
}

func extractMetadata(path string, owner uuid.UUID) (*model.MetaData, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	var createdAt time.Time
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		createdAt = time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec)
	} else {
		createdAt = info.ModTime()
	}

	meta := &model.MetaData{
		FileName:   info.Name(),
		Path:       path,
		Size:       uint64(info.Size()),
		Mode:       info.Mode(),
		IsDir:      info.IsDir(),
		ModifiedAt: info.ModTime(),
		CreatedAt:  createdAt,
		Owner:      uuid.New(),
		AccessTo:   nil,
		Group:      nil,
		Links:      nil,
	}

	return meta, nil
}

func (uc UploadFile) repository() (bool, error) {
	return true, nil
}
