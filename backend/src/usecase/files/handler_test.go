package files

import (
	"context"
	"net/http"
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
