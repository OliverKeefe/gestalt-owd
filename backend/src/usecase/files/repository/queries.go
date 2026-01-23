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
		WithUploadedAt(filter.UploadedAt),
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

// Option is function type that accepts the spec as a param.
type Option func(*GetMetadataSpec)

// buildGetMetadataSpec takes functional options as args and constructs
// a new GetMetadataSpec struct containing each field within the metadata
// relation that is to be queried. e.g. `ID`, `FileName`, `FileType` etc...
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
       uploaded_at, version, checksum, owner_id FROM file_metadata
 	`

	clauses := make([]string, len(spec.clauses))
	for i, c := range spec.clauses {
		clauses[i] = strings.Replace(c, "?", fmt.Sprintf("$%d", i+1), 1)
	}

	query := fmt.Sprintf("%s WHERE %s", baseQuery, strings.Join(clauses, " AND "))
	return query, spec.args, nil
}

func (spec *GetMetadataSpec) nextArg() int {
	return len(spec.args) + 1
}

func WithID(id *uuid.UUID) Option {
	if id == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.args = append(spec.args, *id)
		idx := len(spec.args)
		spec.clauses = append(
			spec.clauses,
			fmt.Sprintf("id = $%d", idx),
		)
	}
}

func WithPath(path *string) Option {
	if path == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.args = append(spec.args, *path)
		idx := len(spec.args)
		spec.clauses = append(
			spec.clauses,
			fmt.Sprintf("path = $%d", idx))
	}
}

func WithFileName(filename *string) Option {
	if filename == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.args = append(spec.args, *filename)
		idx := len(spec.args)
		spec.clauses = append(
			spec.clauses,
			fmt.Sprintf("file_name = $%d", idx))
	}
}

func WithSize(size *uint64) Option {
	if size == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.args = append(spec.args, *size)
		idx := len(spec.args)
		spec.clauses = append(
			spec.clauses,
			fmt.Sprintf("size = $%d", idx))
	}
}

// WithFileType is the functional option to append `file_type`
// to the query.
func WithFileType(fileType *string) Option {
	if fileType == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.args = append(spec.args, *fileType)
		idx := len(spec.args)
		spec.clauses = append(
			spec.clauses,
			fmt.Sprintf("file_type = $%d", idx))
	}
}

// WithModifiedAt is the functional option to append last time
// the file was modified at, to the query.
func WithModifiedAt(t *time.Time) Option {
	if t == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.args = append(spec.args, *t)
		idx := len(spec.args)
		spec.clauses = append(
			spec.clauses,
			fmt.Sprintf("modified_at = $%d", idx))
	}
}

// WithUploadedAt is the functional option to append the time the
// file was uploaded at, to the query.
func WithUploadedAt(t *time.Time) Option {
	if t == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.args = append(spec.args, *t)
		idx := len(spec.args)
		spec.clauses = append(
			spec.clauses,
			fmt.Sprintf("uploaded_at = $%d", idx))
	}
}

// WithOwner is the functional option to append file owner's
// uuid to the query.
func WithOwner(owner *uuid.UUID) Option {
	if owner == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.args = append(spec.args, *owner)
		idx := len(spec.args)
		spec.clauses = append(
			spec.clauses,
			fmt.Sprintf("owner_id = $%d", idx))
	}
}

// WithAccessTo is the functional option to append uuids of the
// user's with permissions to access the file to the query.
func WithAccessTo(ids *[]uuid.UUID) Option {
	if ids == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, `
			EXISTS (
				SELECT 1 FROM file_user_access fua 
				WHERE fua.file_id = file_metadata.id
					AND fua.user_id = ANY (?) 
			)
		`)
		spec.args = append(spec.args, pq.Array(*ids))
	}
}

// WithGroup is the functional option to append `groups` to the query.
func WithGroup(groups *[]uuid.UUID) Option {
	if groups == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.clauses = append(spec.clauses, `
			EXISTS (
				SELECT 1 FROM file_metadata_group_access fga
				WHERE fga.file_id = file_metadata.id
					AND fga.group_id = ANY (?)
			)
		`)
		spec.args = append(spec.args, pq.Array(*groups))
	}
}

// WithVersion is the functional option to append `version` to the query.
func WithVersion(version *time.Time) Option {
	if version == nil {
		return func(*GetMetadataSpec) {}
	}
	return func(spec *GetMetadataSpec) {
		spec.args = append(spec.args, *version)
		idx := len(spec.args)
		spec.clauses = append(
			spec.clauses,
			fmt.Sprintf("version = $%d", idx))
	}
}

// TODO: Fix this because it's never nil so this doesn't work.
func (spec *GetMetadataSpec) isEmpty() bool {
	return spec.clauses == nil &&
		spec.args == nil
}
