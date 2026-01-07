package files

import (
	service "backend/src/core/files/service"
	"encoding/json"
	"log"
	"net/http"
)

// Handler function for file upload.
func UploadHandler(svc *service.UploadSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{
			if r.Method != http.MethodPost {
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}

			r.Body = http.MaxBytesReader(w, r.Body, 100<<20)

			if err := svc.Save(r, r.Context()); err != nil {
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
	}
}
