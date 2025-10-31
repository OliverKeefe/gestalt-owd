package files

import (
	model "backend/core/files/model"
	"net/http"
)

type UploadFile struct {
	file model.File
}

func (upload UploadFile) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/files/upload", upload.Api)
}
