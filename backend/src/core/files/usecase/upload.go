package files

import (
	model "backend/src/core/files/model"
	"backend/src/internal/cloud/objectstorage"
	"backend/src/internal/middleware"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"syscall"
	"time"
)

type DefaultStore struct {
	BasePath  string
	Path      string
	NameSpace string
}

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

	uploaded, err := upload.service(request)
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

func (upload UploadFile) service(request *http.Request) (bool, error) {
	err := request.ParseMultipartForm(1500 << 1500)
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

	var filepath = file.Name
	metadata := extractMetadata()

	uploadedFile := model.File{
		Metadata: metadata,
		FileData: fileBytes,
	}

	var defaultStore = &objectstorage.ObjectStore{
		BasePath:  "/home/oliver/Development/25-26_CE301_keefe_oliver_b",
		Path:      "/backend/tempfiles",
		NameSpace: "hello",
	}

	saveTo, err := objectstorage.Save(defaultStore, uploadedFile)
	if err != nil {
		return false, fmt.Errorf("unable to save file in obj %e", err)
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

func (upload UploadFile) repository() (bool, error) {
	return true, nil
}
