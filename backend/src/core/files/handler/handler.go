package files

import (
	service "backend/src/core/files/service"
	"encoding/json"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{
		"status": "uploaded",
	})
	if err != nil {
		return
	}
}
