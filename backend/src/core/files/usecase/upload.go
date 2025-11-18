package files

import (
	model "backend/src/core/files/model"
	"backend/src/internal/cloud/objectstorage/obj"
	"backend/src/internal/middleware"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

type UploadFile struct {
	file model.File
}

func (upload UploadFile) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/files/upload", middleware.EnableCORS(http.HandlerFunc(upload.Api)))
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

func (upload UploadFile) Service(request *http.Request) (bool, error) {
	err := request.ParseMultipartForm(15000)
	if err != nil {
		return false, fmt.Errorf("could not upload file %e", err)
	}

	file, header, err := request.FormFile("file")
	if err != nil {
		return false, fmt.Errorf("error getting file from request %e", err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			fmt.Errorf("unable to close file %e", err)
		}
	}(file)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return false, fmt.Errorf("error reading file %e", err)
	}

	meta := model.MetaData{
		Filename:     header.Filename,
		Size:         uint64(len(fileBytes)),
		Permissions:  0,
		LastModified: time.Time{},
		IsDirectory:  false,
	}

	uploadedFile := model.File{
		Metadata: meta,
		FileData: fileBytes,
	}

	var defaultStore = &gestaltoblob.Store{
		BasePath: "/home/oliver/Development/25-26_CE301_keefe_oliver_b",
		Path:     "/backend/src/cmd/gestalt/",
	}

	saveTo, err := gestaltoblob.Save(defaultStore, uploadedFile)
	if err != nil {
		return false, fmt.Errorf("unable to save file in obj %e", err)
	}

	return saveTo, nil
}

func (upload UploadFile) Database() (bool, error) {
	return true, nil
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse up to 10MB in memory - this'll need to be increased to 15gb at somepoint.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("Received file: %s\n", header.Filename)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "File uploaded successfully!")
}
