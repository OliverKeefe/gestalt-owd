package files

import (
	model "backend/src/core/files/model"
)

type FileRepository interface {
	SaveMetaData(meta *model.MetaData) (bool, error)
	GetMetaData(fileId *model.File) (model.MetaData, error)
}
