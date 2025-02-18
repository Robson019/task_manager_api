package secondary

import (
	"task_manager/src/core/domain/account"
	"task_manager/src/core/errors"

	"github.com/google/uuid"
)

type AccountLoader interface {
	FindProfileByID(accountID uuid.UUID) (*account.Account, errors.Error)
}
