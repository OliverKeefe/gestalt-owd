package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONHandler_ReceivePost(t *testing.T) {
	type request struct {
		Name string `json:"name"`
	}

	type response struct {
		Message string `json:"message"`
	}

	handler := JSONHandler[request, response](
		func(ctx context.Context, req request) (response, error) {
			if req.Name != "test" {
				t.Fatalf("unexpected request payload: %+v", req)
			}
			return response{Message: "ok"}, nil
		},
		WithMethod(http.MethodPost),
	)

	body := bytes.NewBufferString(`{"name":"test"}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/test",
		body,
	)
	recorder := httptest.NewRecorder()
	handler(recorder, req)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 but got: %v", resp.StatusCode)
	}

	var got response
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode dto: %v", err)
	}

	if got.Message != "ok" {
		t.Fatalf("unexpected dto: %+v", got)
	}
}
