package files

import (
	model "backend/src/core/files/model"
	"backend/src/internal/middleware"
	"net/http"
)

type ViewFiles struct {
	files []model.File
}

func (view ViewFiles) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/files/upload", middleware.EnableCORS(http.HandlerFunc(v.Api)))
}

func (view ViewFiles) Api(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "unable to view files", http.StatusMethodNotAllowed)
		return
	}

	viewFiles, err := view.Service(request)
	if err != nil {
		http.Error()
	}
}

func (view ViewFiles) Service(request *http.Request) (bool, error) {

}
