package primary

import (
	"task_manager/src/core/domain/account/credentials"
	"task_manager/src/core/domain/authorization"
	"task_manager/src/core/errors"

	"github.com/google/uuid"
)

type AuthManager interface {
	Login(credentials credentials.Credentials) (authorization.Authorization, errors.Error)
	Logout(accountID uuid.UUID) errors.Error
	Refresh(refreshToken string, accountID uuid.UUID) (authorization.Authorization, errors.Error)
	SessionExists(accountID uuid.UUID, token string) (bool, errors.Error)
}
