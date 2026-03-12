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
	route = func(a *auth.Authenticator, h http.HandlerFunc) http.Handler {
		return middleware.Protect(context.Background(), a, h)
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
	findMetadata := route(a, h.FindMetadata)
	mux.Handle(
		"POST /api/files/find",
		findMetadata,
	)
	getAllMetadata := route(a, h.GetAll)
	mux.Handle(
		"POST /api/files/get-all",
		getAllMetadata,
	)
}
