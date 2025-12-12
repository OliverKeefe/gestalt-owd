package http

import (
	"backend/src/internal/bootstrap"
	"net/http"
)

type AppConfig struct {
	ApiConfig *bootstrap.Config
}

func (a AppConfig) Post() (http.Response, error) {
	panic("Not implemented.")
}

func (a AppConfig) Get() (http.Response, error) {
	panic("Not implemented.")
}

func (a AppConfig) Delete() (http.Response, error) {
	panic("Not implemented.")
}

func (a AppConfig) Put() (http.Response, error) {
	panic("Not implemented.")
}

func (a AppConfig) Trace() (http.Response, error) {
	panic("Not implemented.")
}

func (a AppConfig) Patch() (http.Response, error) {
	panic("Not implemented.")
}
