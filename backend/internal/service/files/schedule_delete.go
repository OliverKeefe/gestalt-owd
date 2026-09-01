package files

import (
	"backend/internal/api/message"
	"backend/internal/platform"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type scheduleDeleteRepository interface {
	FindMetadataByID(ctx context.Context, id uuid.UUID) (FileMetadata, error)
	OwnerIsInUserGroups(ctx context.Context, ownerID uuid.UUID, groupNames []string) (bool, error)
	MarkForDeletion(ctx context.Context, id []uuid.UUID, ownerId uuid.UUID) error
}

type ScheduleDeleteService struct {
	db                scheduleDeleteRepository
	BlobStorageClient *platform.BlobStorageClient
	BucketURL         string
}

type scheduleDeleteRequest struct {
	ID []uuid.UUID `json:"id"`
}

func (svc *ScheduleDeleteService) Handle(w http.ResponseWriter, r *http.Request) {
	request, err := message.Bind[scheduleDeleteRequest](r)
	if err != nil {
		log.Printf("unable to bind request to DeleteRequest %v", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := svc.execute(r.Context(), request); err != nil {
		if errors.Is(err, ErrForbidden) {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		log.Printf("unable to delete file metadata, %v", err)
		http.Error(w, "unable to delete file", http.StatusExpectationFailed)
		return
	}
}

func (svc *ScheduleDeleteService) execute(ctx context.Context, request scheduleDeleteRequest) error {
	var ownerID uuid.UUID
	for i, id := range request.ID {
		metadata, err := svc.db.FindMetadataByID(ctx, id)
		if err != nil {
			return fmt.Errorf("file not found: %w", err)
		}

		authedOwner, err := authorizeFile(ctx, svc.db, metadata.OwnerID)
		if err != nil {
			return err
		}
		if i == 0 {
			ownerID = authedOwner
		} else if ownerID != authedOwner {
			return ErrForbidden
		}
	}

	if err := svc.db.MarkForDeletion(ctx, request.ID, ownerID); err != nil {
		log.Printf("unable to move file or metadata to rubbish bin, %v", err)
		return err
	}

	return nil
}
