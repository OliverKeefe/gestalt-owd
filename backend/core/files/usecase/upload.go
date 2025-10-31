package files

import (
	model "backend/core/files/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UploadFile struct {
	file model.File
}

func (upload UploadFile) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/files/upload", upload.Api)
}

func (upload UploadFile) Api(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPut {
		http.Error(writer, "unable to upload file", http.StatusMethodNotAllowed)
		return
	}

	uploaded, err := upload.Service(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Fatal(fmt.Errorf("unable to upload file, %e", err))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	if uploaded == true {
		writer.WriteHeader(http.StatusAccepted)
		err := json.NewEncoder(writer).Encode("File successfully uploaded")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			log.Fatal(fmt.Errorf("file uploaded but could not send response, %e", err))
			return
		}
		return
	} else {
		writer.WriteHeader(http.StatusAccepted)
		err := json.NewEncoder(writer).Encode("File could not be uploaded.")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			log.Fatal(fmt.Errorf("file not uploaded and could not send response, %e", err))
			return
		}
		return
	}
}
