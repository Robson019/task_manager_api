package secondary

import (
	"task_manager/src/core/errors"

	"github.com/google/uuid"
)

type SessionLoader interface {
	Store(accountID uuid.UUID, accessToken string, refreshToken string) errors.Error
	Close(uID uuid.UUID) errors.Error
	SessionExists(uID uuid.UUID, token string) (bool, errors.Error)
	RefreshTokenExists(uID uuid.UUID, token string) (bool, errors.Error)
}
