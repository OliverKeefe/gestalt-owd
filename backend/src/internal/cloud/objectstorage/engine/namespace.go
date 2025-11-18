package objectstorage

import (
	model "backend/src/core/files/model"
	"backend/src/internal/util"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"golang.org/x/sys/unix"
	"syscall"
)

/*

   Mount
   PID
   Network
   Cgroup
   IPC
   Time
   UTS
   User

*/

// create namespace / cgroup container for object storage.

type Namespace interface{}

type Bucket struct {
	ID              uuid.UUID
	Name            string
	Size            uint64
	Creator         uuid.UUID
	Owner           uuid.UUID
	GroupMembership []uuid.UUID
}

// TODO: make params functional options pattern (options ... func(*Bucket)) *Bucket or config struct
func (Bucket) NewObjectBucket(ctx context.Context, conn pgx.Conn, name string, size uint64,
	creator uuid.UUID, owner uuid.UUID, groupMembership []uuid.UUID) (Bucket, error) {

	var id uuid.UUID
	var err error

	for {
		id = uuid.New()
		unique, err := util.IsUUIDUnique(ctx, id, conn, "object_storage_buckets")
		if err != nil {
			return Bucket{}, err
		}
		if unique {
			break
		}
	}

	bucket := Bucket{
		ID:              id,
		Name:            name,
		Size:            size,
		Creator:         creator,
		Owner:           owner,
		GroupMembership: groupMembership,
	}

	return bucket, nil
}

type Mount struct{}

type PID struct{}

type Network struct{}

type CGroup struct{}

type IPC struct{}

type Time struct{}

type UTS struct{}

type User struct{}

func (mnt Mount) createMount(mount bool, fork bool, pid uint64, proc bool, path string) (unix.TIPCServiceName, error) {
}

func must(err error) {
	if err != nil {
		panic("AAhhh")
	}
}
