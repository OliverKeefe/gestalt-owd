package router

import (
	"backend/src/internal/auth"
	"backend/src/internal/db/metadb"
	"backend/src/internal/middleware"
	fileshandler "backend/src/usecase/files/handler"
	filesrepo "backend/src/usecase/files/repository"
	filesvc "backend/src/usecase/files/service"
	"net/http"
)

var (
	protected = func(a *auth.Authenticator, h http.HandlerFunc) http.Handler {
		return middleware.Protect(a, h)
	}
)

func RegisterFileRoutes(
	mux *http.ServeMux,
	a *auth.Authenticator,
	db *metadb.MetadataDatabase,
) {
	repo := filesrepo.NewRepository(db)
	svc := filesvc.NewService(repo)
	h := fileshandler.NewHandler(svc)

	upload := protected(a, h.Upload)
	mux.Handle(
		"POST /api/files/upload",
		upload,
	)
	findMetadata := protected(a, h.FindMetadata)
	mux.Handle(
		"POST /api/files/find",
		findMetadata,
	)
	getAllMetadata := protected(a, h.GetAll)
	mux.Handle(
		"POST /api/files/get-all",
		getAllMetadata,
	)
}
