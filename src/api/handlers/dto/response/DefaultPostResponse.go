package response

import (
	"github.com/google/uuid"
)

type DefaultPostResponse struct {
	ID uuid.UUID `json:"id"`
}
