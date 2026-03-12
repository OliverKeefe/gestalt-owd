package files

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

type mockRepository struct{}

var (
	mockRepo = &mockRepository{}
	svc      = Service{repo: mockRepo}
)

func (r *mockRepository) SaveMetaData(ctx context.Context, meta MetaData) (MetaData, error) {
	return MetaData{}, nil
}
func (r *mockRepository) SaveFileData(basePath string, rdr io.Reader, filename string) error {
	return nil
}
func (r *mockRepository) FindMetadata(ctx context.Context, model MetaData) ([]MetaData, error) {
	return []MetaData{}, nil
}
func (r *mockRepository) DeleteMetadata(ctx context.Context, id uuid.UUID, ownerId uuid.UUID) error {
	return nil
}
func (r *mockRepository) FindAllMetadata(ctx context.Context, req GetAllMetadataRequest) ([]MetaData, error) {
	return []MetaData{}, nil
}
func (r *mockRepository) MarkForDeletion(ctx context.Context, id uuid.UUID, id2 uuid.UUID) error {
	return nil
}

func TestService_Upload(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "test.txt")
	if err != nil {
		t.Fatal(err)
	}

	fileData := []byte("Test file content.")
	_, err = part.Write(fileData)
	if err != nil {
		t.Fatal(err)
	}

	_ = writer.WriteField("ID", uuid.New().String())
	writer.Close()

	r := httptest.NewRequest(
		http.MethodPost,
		"/api/files/upload",
		body,
	)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	r.Header.Set("Authorization", "Bearer token")

	files, err := svc.Upload(r.Context(), r)
	if err != nil {
		t.Fatalf("unexpected error, got: %v", err)
	}

	if reflect.TypeOf(files) != reflect.TypeOf([]MetaData{}) {
		t.Fatalf("expected type []MetaData, got: %T", files)
	}

	if len(files) != 0 {
		t.Fatalf("expected empty slice of Metadata, got: %v", files)
	}

}
