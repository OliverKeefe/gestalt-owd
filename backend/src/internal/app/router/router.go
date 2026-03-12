package router

import (
	"backend/src/internal/auth"
	"backend/src/internal/db/metadb"
	"backend/src/internal/middleware"
	filesvc "backend/src/usecase/files"
	"context"
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
	repo := filesvc.NewRepository(db.Pool)
	svc := filesvc.NewService(repo)
	h := filesvc.NewHandler(svc)

	upload := route(a, h.Upload)
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
