package files

import (
	service "backend/src/core/files/service"
	"log"
	"net/http"
)

func UploadHandler(svc *service.UploadSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{
			if r.Method != http.MethodPost {
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}

			r.Body = http.MaxBytesReader(w, r.Body, 100<<20)

			//TODO: make sure request body for metadata matches schema, if not, SQLSTATE 22007 error.
			if err := svc.Save(r, r.Context()); err != nil {
				log.Printf("unable to save uploaded file: %v", err)
				http.Error(w, "could not save file", http.StatusInternalServerError)
			}
		}
	}
}
