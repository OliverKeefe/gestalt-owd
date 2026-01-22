package files

import (
	data "backend/src/core/files/data"
	service "backend/src/core/files/service"
	"backend/src/internal/api/message"
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

	err := message.Response(w, "uploaded")
	if err != nil {
		log.Print(err)
		return
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	svc := h.svc

	userUUID := r.URL.Query().Get("user_uuid")
	files, err := svc.GetAll(userUUID, r.Context())
	if err != nil {
		log.Printf("couldn't get all user's files: %v", err)
		http.Error(w, "unable to get user's files", http.StatusInternalServerError)
	}

	if err = message.Response(w, "fetched user's files", files); err != nil {
		log.Printf("unable to return response, %v", err)
		return
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	svc := h.svc
	var request data.GetMetadataRequest

	if err := request.Bind(r); err != nil {
		log.Printf("unable to bind raw request to GetMetadataRequest, %v", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
	}

	files, err := svc.GetMetadata(r.Context(), request)
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
	panic("not implemented")
}
