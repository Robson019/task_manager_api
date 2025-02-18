package secondary

import (
	"task_manager/src/core/domain/account"
	"task_manager/src/core/domain/account/credentials"
	"task_manager/src/core/errors"
)

type AuthLoader interface {
	Login(credentials credentials.Credentials) (*account.Account, errors.Error)
}
