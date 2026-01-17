package app

import (
	fileshandler "backend/src/core/files/handler"
	filesrepo "backend/src/core/files/repository"
	filesvc "backend/src/core/files/service"
	"backend/src/internal/app/router"
	metadb "backend/src/internal/metadb"
	"backend/src/internal/middleware"
	"context"
	"log"
	"net/http"
	"time"
)

func Run() error {
	ctx := context.Background()

	db, err := metadb.New(ctx, "DATABASE_URL")
	if err != nil {
		return err
	}

	//TODO: Chain this instead, probably following functional options.
	handler := middleware.EnableCORS(registerRoutes(db))

	srv := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	log.Println("running on port: 8081...")
	return srv.ListenAndServe()
}

func registerRoutes(metadataDB *metadb.MetadataDatabase) *http.ServeMux {
	filesRepo := filesrepo.NewRepository(metadataDB)
	filesService := filesvc.NewService(filesRepo)
	filesHandler := fileshandler.NewHandler(filesService)

	return router.New(
		router.Handle(
			"POST /api/files/upload",
			http.HandlerFunc(filesHandler.Upload),
		),
		router.Handle("POST /api/files/get-all",
			http.HandlerFunc(filesHandler.GetAll),
		),
	)
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
