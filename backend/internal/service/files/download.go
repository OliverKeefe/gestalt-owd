package files

import (
	"backend/internal/api/message"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"gocloud.dev/blob"
)

type downloadRepository interface {
	FindMetadataByID(ctx context.Context, id uuid.UUID) (FileMetadata, error)
	OwnerIsInUserGroups(ctx context.Context, ownerID uuid.UUID, groupNames []string) (bool, error)
}

type downloadBlob interface {
	Download(ctx context.Context, key string, writer io.Writer, opts *blob.ReaderOptions) error
}

type DownloadService struct {
	Db                downloadRepository
	BlobStorageClient downloadBlob
	BucketURL         string
}

func NewDownloadService(db downloadRepository, client downloadBlob, bucketUrl string) *DownloadService {
	return &DownloadService{
		Db:                db,
		BlobStorageClient: client,
		BucketURL:         bucketUrl,
	}
}

func (svc *DownloadService) Handle(w http.ResponseWriter, r *http.Request) {
	request, err := message.Bind[DownloadRequest](r)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := svc.execute(r.Context(), request, w); err != nil {
		if errors.Is(err, ErrForbidden) {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		slog.Error("download failed", "error", err)
		http.Error(w, "download failed", http.StatusInternalServerError)
		return
	}
}

func (svc *DownloadService) execute(ctx context.Context, req DownloadRequest, w http.ResponseWriter) error {
	metadata, err := svc.Db.FindMetadataByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("file not found: %w", err)
	}

	ownerID, err := authorizeFile(ctx, svc.Db, metadata.OwnerID)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, metadata.FileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", metadata.Size))

	return svc.download(ctx, ownerID, metadata.FileName, w)
}

func (svc *DownloadService) download(ctx context.Context, ownerID uuid.UUID, fileName string, w io.Writer) error {
	key := fmt.Sprintf("%s/%s", ownerID, url.PathEscape(fileName))
	return svc.BlobStorageClient.Download(ctx, key, w, nil)
}
