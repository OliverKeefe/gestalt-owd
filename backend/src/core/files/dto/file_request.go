package dto

import "github.com/google/uuid"

type GetAllFilesRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Index  uuid.UUID `json:"index"`
}
