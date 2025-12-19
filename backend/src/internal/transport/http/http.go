package http

import (
	"backend/src/internal/bootstrap"
	"bytes"
	"encoding/json"
	"net/http"
)

type AppConfig struct {
	Config *bootstrap.Config
}

func (app AppConfig) Post(endpointURL string, payload any) (*http.Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := app.buildURL(endpointURL)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := app.Config.HTTPClient

	return client.Do(req)
}

func (app AppConfig) Get() (http.Response, error) {
	panic("Not implemented.")
}

func (app AppConfig) Delete() (http.Response, error) {
	panic("Not implemented.")
}

func (app AppConfig) Put() (http.Response, error) {
	panic("Not implemented.")
}

func (app AppConfig) Trace() (http.Response, error) {
	panic("Not implemented.")
}

func (app AppConfig) Patch() (http.Response, error) {
	panic("Not implemented.")
}

func (app AppConfig) buildURL(endpointUrl string) string {
	return app.Config.BaseURL + endpointUrl
}
