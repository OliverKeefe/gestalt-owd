package files

import (
	files "backend/src/usecase/files/data"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type GetMetadataFilter struct {
	ID         *uuid.UUID
	FileName   *string
	Path       *string
	Size       *uint64
	FileType   *string
	ModifiedAt *time.Time
	UploadedAt *time.Time
	Owner      *uuid.UUID
	AccessTo   *[]uuid.UUID
	Group      *[]uuid.UUID
	Version    *time.Time
}

func buildMetadataFilter(model files.MetaData) GetMetadataFilter {
	return GetMetadataFilter{
		ID:         &model.ID,
		FileName:   &model.FileName,
		Path:       &model.Path,
		Size:       &model.Size,
		FileType:   &model.FileType,
		ModifiedAt: &model.ModifiedAt,
		UploadedAt: &model.UploadedAt,
		Owner:      &model.Owner,
		AccessTo:   &model.AccessTo,
		Group:      &model.Group,
		Version:    &model.Version,
	}
}

type GetMetadataSpec struct {
	clauses []string
	args    []any
}

func GetMetadataQuery(model files.MetaData) (string, []any, error) {
	filter := buildMetadataFilter(model)

	spec := buildGetMetadataSpec(
		WithID(filter.ID),
		WithFileName(filter.FileName),
		WithPath(filter.Path),
		WithSize(filter.Size),
		WithFileType(filter.FileType),
		WithModifiedAt(filter.ModifiedAt),
		WithCreatedAt(filter.UploadedAt),
		WithOwner(filter.Owner),
		WithAccessTo(filter.AccessTo),
		WithGroup(filter.Group),
		WithVersion(filter.Version),
	)

	query, args, err := buildGetMetadataQuery(spec)
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

type Option func(*GetMetadataSpec)

func buildGetMetadataSpec(opts ...Option) GetMetadataSpec {
	q := GetMetadataSpec{}
	for _, opt := range opts {
		opt(&q)
	}
	return q
}

func buildGetMetadataQuery(spec GetMetadataSpec) (string, []any, error) {
	if spec.isEmpty() {
		return "", nil, errors.New("unable to build query as MetadataSpec fields are all nil")
	}

	const baseQuery = `SELECT id, file_name, path, size, file_type, modified_at, 
       uploaded_at, version, checksum, owner FROM file_metadata
 	`

	clauses := make([]string, len(spec.clauses))
	for i, c := range spec.clauses {
		clauses[i] = strings.Replace(c, "?", fmt.Sprintf("$%d", i+1), 1)
	}

	query := fmt.Sprintf("%s WHERE %s", baseQuery, strings.Join(clauses, " AND "))
	return query, spec.args, nil
}

func WithID(id *uuid.UUID) Option {
	if id == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "id = ?")
		spec.args = append(spec.args, *id)
	}
}

func WithPath(path *string) Option {
	if path == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "path = ?")
		spec.args = append(spec.args, *path)
	}
}

func WithFileName(filename *string) Option {
	if filename == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "file_name = ?")
		spec.args = append(spec.args, *filename)
	}
}

func WithSize(size *uint64) Option {
	if size == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "size = ?")
		spec.args = append(spec.args, *size)
	}
}

func WithFileType(fileType *string) Option {
	if fileType == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "file_type = ?")
		spec.args = append(spec.args, *fileType)
	}
}

func WithModifiedAt(t *time.Time) Option {
	if t == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "modified_at = ?")
		spec.args = append(spec.args, *t)
	}
}

func WithCreatedAt(t *time.Time) Option {
	if t == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "uploaded_at = ?")
		spec.args = append(spec.args, *t)
	}
}

func WithOwner(owner *uuid.UUID) Option {
	if owner == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "owner = ?")
		spec.args = append(spec.args, *owner)
	}
}

func WithAccessTo(ids *[]uuid.UUID) Option {
	if ids == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "access_to @> ?")
		spec.args = append(spec.args, pq.Array(*ids))
	}
}

func WithGroup(groups *[]uuid.UUID) Option {
	if groups == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "groups @> ?")
		spec.args = append(spec.args, pq.Array(*groups))
	}
}

func WithVersion(version *time.Time) Option {
	if version == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, "version = ?")
		spec.args = append(spec.args, *version)
	}
}

func (spec *GetMetadataSpec) isEmpty() bool {
	return spec.clauses == nil &&
		spec.args == nil
}
