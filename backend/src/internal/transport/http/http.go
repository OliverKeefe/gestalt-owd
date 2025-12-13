package http

import (
	"backend/src/internal/bootstrap"
	"bytes"
	"encoding/json"
	"net/http"
)

type AppConfig struct {
	ApiConfig *bootstrap.Config
}

func (a AppConfig) Post(endpointURL string, payload any) (*http.Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := a.buildURL(endpointURL)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
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
