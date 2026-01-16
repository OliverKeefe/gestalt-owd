package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type HTTPError struct {
	Status int
	Err    error
}

func (e HTTPError) Error() string {
	return e.Err.Error()
}

type JSONHandlerConfig struct {
	Method        string `json:"method,omitempty"`
	Headers       map[string]string
	Body          any
	SuccessStatus int
}

func defaultJSONHandlerConfig() JSONHandlerConfig {
	return JSONHandlerConfig{
		Method: http.MethodPost,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:          nil,
		SuccessStatus: http.StatusOK,
	}
}

type JSONHandlerOption func(*JSONHandlerConfig)

func JSONHandler[TReq any, TResp any](
	fn func(context.Context, TReq) (TResp, error),
	opts ...JSONHandlerOption) http.HandlerFunc {

	cfg := defaultJSONHandlerConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != cfg.Method {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req TReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
		}

		resp, err := fn(r.Context(), req)
		if err != nil {
			var httpErr HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Error(), httpErr.Status)
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		//k: "Content-Type" v: "application/json"
		for k, v := range cfg.Headers {
			w.Header().Set(k, v)
		}

		w.WriteHeader(cfg.SuccessStatus)
		_ = json.NewEncoder(w).Encode(resp)
	}
}

func WithHeader(key, value string) JSONHandlerOption {
	return func(c *JSONHandlerConfig) {
		c.Headers[key] = value
	}
}

func WithMethod(method string) JSONHandlerOption {
	return func(c *JSONHandlerConfig) {
		c.Method = method
	}
}

func WithSuccessStatus(status int) JSONHandlerOption {
	return func(c *JSONHandlerConfig) {
		c.SuccessStatus = status
	}
}
