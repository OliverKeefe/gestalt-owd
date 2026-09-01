package router

import (
	"backend/internal/auth"
	"backend/internal/database"
	"backend/internal/middleware"
	"backend/internal/platform"
	f "backend/internal/service/files"
	"context"
	"fmt"
	"net/http"
	"os"
)

var (
	route = func(a *auth.Authenticator, h http.HandlerFunc) http.Handler {
		return middleware.Protect(a, h)
	}
)

const clientID = "open-web-drive"

func routeWithClientRole(a *auth.Authenticator, role string, h http.HandlerFunc) http.Handler {
	return middleware.Protect(a, middleware.RequireClientRole(a, clientID, role, h))
}

func RegisterFileRoutes(mux *http.ServeMux, a *auth.Authenticator, db *database.MetadataDatabase) error {
	bucketURL := os.Getenv("BLOB_BUCKET_URL")
	if bucketURL == "" {
		return fmt.Errorf("BLOB_BUCKET_URL is required")
	}

	client, err := platform.NewBlobStorageClient(context.Background(), bucketURL)
	if err != nil {
		panic(err)
	}

	repository := f.NewFileRepository(db.Pool)

	uploadSvc := f.NewUploadService(repository, client, bucketURL)
	downloadSvc := f.NewDownloadService(repository, client, bucketURL)
	deleteSvc := f.NewDeleteService(repository, client, bucketURL)
	updateSvc := f.NewUpdateMetadataService(repository, client, bucketURL)
	searchSvc := f.NewSearchService(repository, client, bucketURL)

	uploadEndpoint := routeWithClientRole(a, "files:upload", uploadSvc.Handle)
	mux.Handle(
		"POST /api/files/upload",
		uploadEndpoint,
	)
	downloadRoute := routeWithClientRole(a, "files:download", downloadSvc.Handle)
	mux.Handle(
		"POST /api/files/download",
		downloadRoute,
	)
	deleteRoute := routeWithClientRole(a, "files:delete", deleteSvc.Handle)
	mux.Handle(
		"DELETE /api/files/delete",
		deleteRoute,
	)
	updateMetadataRoute := route(a, updateSvc.Handle)
	mux.Handle(
		"PUT /api/files/metadata",
		updateMetadataRoute,
	)
	searchRoute := route(a, searchSvc.Handle)
	mux.Handle(
		"POST /api/files/get-all",
		searchRoute,
	)

	return nil
}
