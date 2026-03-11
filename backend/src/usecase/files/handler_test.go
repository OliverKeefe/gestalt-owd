package files

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
)

type mockService struct{}

var (
	mockSvc = &mockService{}
)

func (m *mockService) Upload(r *http.Request, ctx context.Context) ([]MetaData, error) {
	return []MetaData{}, nil
}

func (m *mockService) FindMetadata(ctx context.Context, request FindMetadataRequest) ([]MetaData, error) {
	return []MetaData{}, nil
}

func (m *mockService) Delete(ctx context.Context, request DeleteRequest) error {
	return nil
}

func (m *mockService) MoveToRubbish(ctx context.Context, request DeleteRequest) error {
	return nil
}

func (m *mockService) GetAll(ctx context.Context, request GetAllMetadataRequest) ([]MetaDataResponse, error) {
	return []MetaDataResponse{}, nil
}

func TestHandler_GetAllInvalidRequest(t *testing.T) {
	h := Handler{svc: mockSvc}
	req := httptest.NewRequest(http.MethodPost, "/api/files/get-all", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	recorder := httptest.NewRecorder()

	h.GetAll(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status: 400, got %d", recorder.Code)
	}
}

func TestHandler_GetAll(t *testing.T) {
	h := &Handler{svc: mockSvc}

	payload := GetAllMetadataRequest{
		UserID: uuid.New(),
		Cursor: &MetadataCursor{
			ModifiedAt: time.Now(),
			ID:         uuid.New(),
		},
		Limit: 20,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("invalid test json payload")
	}

	var req = httptest.NewRequest(
		http.MethodPost,
		"/api/files/get-all",
		bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	recorder := httptest.NewRecorder()

	h.GetAll(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got: %v", recorder.Code)
	}
}

func TestHandler_Delete(t *testing.T) {
	h := Handler{svc: mockSvc}

	payload := DeleteRequest{ID: uuid.New()}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("invalid test json payload")
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/files/delete",
		bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	recorder := httptest.NewRecorder()

	h.Delete(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got: %v", recorder.Code)
	}
}

func TestHandler_TempDelete(t *testing.T) {
	h := Handler{svc: mockSvc}
	payload := DeleteRequest{ID: uuid.New()}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("invalid test json payload")
	}

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/files/temp-delete",
		bytes.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	recorder := httptest.NewRecorder()

	h.TempDelete(recorder, req)
	if recorder.Code != 204 {
		t.Fatalf("expected 204, got %v", recorder.Code)
	}
}
