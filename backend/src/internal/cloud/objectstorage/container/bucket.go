package objectstorage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Bucket is a struct containing the metadata of an object bucket.
// object storage buckets contain two components, the metadata as
// stored here, and the minimal distro-less container that isolates
// the bucket's file system.
type Bucket struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	Name      string
	Size      uint64
	CreatorID uuid.UUID
	OwnerID   uuid.UUID
	Groups    []uuid.UUID
}

type BucketService struct {
	db *pgxpool.Pool
}

type CreateBucketInput struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	Name      string
	Size      uint64
	CreatorID uuid.UUID
	OwnerID   uuid.UUID
	Groups    []uuid.UUID
}

func (svc *BucketService) CreateBucket(ctx context.Context, in CreateBucketInput) (bool, error) {
	panic("Not implemented yet.")
}

func (svc *BucketService) DeleteBucket(bucketId uuid.UUID) (bool, error) {
	panic("Not implemented yet.")
}

type UpdateBucketInput struct {
	ID      uuid.UUID
	Name    *string
	Size    *uint64
	OwnerID *uuid.UUID
	Groups  *[]uuid.UUID
}

func (svc *BucketService) UpdateBucket(bucketId uuid.UUID) (bool, error) {
	panic("Not implemented yet.")
}

func (svc *BucketService) HibernateBucket(bucketId uuid.UUID) (bool, error) {
	panic("Not implemented yet.")
}

func (svc *BucketService) RunBucket(bucketId uuid.UUID) (bool, error) {
	panic("Not implemented yet.")
}
