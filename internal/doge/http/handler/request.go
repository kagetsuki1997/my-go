package httphandler

import (
	"my-go/internal/doge/model"

	"github.com/google/uuid"
)

type CreateDogeRequest struct {
	Name        string         `json:"name" validate:"required"`
	MagicNumber uint32         `json:"magicNumber" validate:"required"`
	Type        model.DogeType `json:"type" validate:"required"`
}

type GetDogeRequest struct {
	Id uuid.UUID `param:"id" validate:"required"`
}
