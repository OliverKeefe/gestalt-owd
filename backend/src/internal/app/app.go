package app

import (
	"backend/src/internal/app/router"
	"backend/src/internal/auth"
	"backend/src/internal/db/metadb"
	"backend/src/internal/middleware"
	midware "backend/src/internal/middleware"
	fileshandler "backend/src/usecase/files/handler"
	filesrepo "backend/src/usecase/files/repository"
	filesvc "backend/src/usecase/files/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Run() error {
	ctx := context.Background()

	a, err := registerKeycloakAuth(
		"http://127.0.0.1:8080/realms/gestalt",
		"http://127.0.0.1:8080/realms/gestalt/protocol/openid-connect/certs",
	)
	if err != nil {
		return err
	}

	db, err := metadb.New(ctx, "DATABASE_URL")
	if err != nil {
		return err
	}

	//TODO: Chain this instead, probably following functional options.
	handler := middleware.EnableCORS(registerRoutes(a, db))

	srv := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	log.Println("running on port: 8081...")
	return srv.ListenAndServe()
}

func registerRoutes(a *auth.Authenticator, metadataDB *metadb.MetadataDatabase) *http.ServeMux {
	filesRepo := filesrepo.NewRepository(metadataDB)
	filesService := filesvc.NewService(filesRepo)
	filesHandler := fileshandler.NewHandler(filesService)

	return router.New(
		router.Handle(
			"POST /api/files/upload",
			midware.Protect(a, http.HandlerFunc(filesHandler.Upload)),
		),
		router.Handle("POST /api/files/getall",
			http.HandlerFunc(filesHandler.GetAll),
		),
		router.Handle("POST /api/files/get",
			http.HandlerFunc(filesHandler.FindMetadata),
		),
	)
}

func registerKeycloakAuth(issuer string, jwksURL string) (*auth.Authenticator, error) {
	authenticator, err := auth.New(issuer, jwksURL)
	if err != nil {
		return nil, fmt.Errorf("unable to create new authenticator, %v", err)
	}
	return authenticator, nil
}

type Config struct {
	BaseURL          string
	HTTPClient       *http.Client
	MetadataDatabase *metadb.MetadataDatabase
	HTTPServer       *http.Server
	Multiplexer      *http.ServeMux
}

type Option func(*Config)

func NewConfig(opts ...Option) Config {
	cfg := Config{
		BaseURL: "",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		MetadataDatabase: &metadb.MetadataDatabase{
			Pool: nil,
		},
		HTTPServer:  nil,
		Multiplexer: nil,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

func WithBaseURL(baseUrl string) Option {
	return func(c *Config) {
		c.BaseURL = baseUrl
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Config) {
		c.HTTPClient = client
	}
}

func WithMetadataDB(mdb *metadb.MetadataDatabase) Option {
	return func(c *Config) {
		c.MetadataDatabase.Pool = mdb.Pool
	}
}

func WithHTTPServer(srv *http.Server) Option {
	return func(c *Config) {
		c.HTTPServer = srv
	}
}

func WithMux(mux *http.ServeMux) Option {
	return func(c *Config) {
		c.Multiplexer = mux
	}
}
