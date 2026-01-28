package files

import (
	model "backend/src/usecase/files/data"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FileRepository struct {
	db *pgxpool.Pool
}

func (r *FileRepository) SaveMetadata(meta *model.MetaData) (bool, error) {
	// SQL
}
