package ipfs

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"
)

type IPFSProvider interface {
	Authenticate(ctx context.Context, account string, did string) (bool, error)
	Upload(ctx context.Context, rdr io.Reader) (bool, cid.Cid, error)
	Remove(ctx context.Context) (bool, error)
	Audit(ctx context.Context, checksumVals []byte) ([]byte, bool, error)
}
