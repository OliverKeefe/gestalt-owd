package files

import (
	"backend/src/internal/api/message"
	data "backend/src/usecase/files/data"
	service "backend/src/usecase/files/service"
	"log"
	"net/http"
)

type Handler struct {
	svc *service.Service
}

// Constructor
func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

// Upload Handler function for file upload.
// Endpoint api/files/upload
func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	svc := h.svc

	r.Body = http.MaxBytesReader(w, r.Body, 100<<20)

	if err := svc.Upload(r, r.Context()); err != nil {
		log.Printf("unable to save uploaded file: %v", err)
		http.Error(w, "could not save file", http.StatusInternalServerError)
	}

	err := message.Response(w, "uploaded", nil)
	if err != nil {
		log.Print(err)
		return
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	svc := h.svc

	if r.Body == nil || r.ContentLength <= 0 {
		http.Error(w, "invalid request", http.StatusBadRequest)
	}

	var request data.GetAllMetadataRequest
	err := request.Bind(r)
	if err != nil {
		log.Printf("couldn't map http request to dto: %v", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
	}

	files, err := svc.GetAll(r.Context(), request)
	if err != nil {
		log.Printf("couldn't get all user's files: %v", err)
		http.Error(w, "unable to get user's files", http.StatusInternalServerError)
	}

	if err = message.Response(w, "fetched user's files", files); err != nil {
		log.Printf("unable to return response, %v", err)
		return
	}
}

func (h *Handler) FindMetadata(w http.ResponseWriter, r *http.Request) {
	svc := h.svc
	var request data.FindMetadataRequest

	if err := request.Bind(r); err != nil {
		log.Printf("unable to bind raw request to FindMetadataRequest, %v", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
	}

	files, err := svc.FindMetadata(r.Context(), request)
	if err != nil {
		log.Printf("couldn't get all user's files: %v", err)
		http.Error(w, "unable to get user's files", http.StatusInternalServerError)
	}

	if err = message.Response(w, "fetched user's files", files); err != nil {
		log.Printf("unable to return response, %v", err)
		return
	}
}

func (h *Handler) Download(w http.ResponseWriter, r *http.Request) {
	panic("not implemented.")
}

func (h *Handler) UpdateMetadata(w http.ResponseWriter, r *http.Request) {
	panic("not implemented.")
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var request data.DeleteRequest

	if err := request.Bind(r); err != nil {
		log.Printf("unable to bind request to DeleteRequest %v", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
	}

	if err := h.svc.Delete(r.Context(), request); err != nil {
		log.Printf("unable to delete file metadata, %v", err)
		http.Error(w, "unable to delete file", http.StatusExpectationFailed)
	}
}

func (h *Handler) TempDelete(w http.ResponseWriter, r *http.Request) {
	var request data.DeleteRequest

	if err := request.Bind(r); err != nil {
		log.Printf("unable to bind request to DeleteRequest %v", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
	}

	if err := h.svc.MoveToRubbish(r.Context(), request); err != nil {
		log.Printf("unable to delete file metadata, %v", err)
		http.Error(w, "unable to delete file", http.StatusExpectationFailed)
	}
}
