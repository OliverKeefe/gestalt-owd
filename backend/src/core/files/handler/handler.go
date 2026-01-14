package files

import (
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
}
