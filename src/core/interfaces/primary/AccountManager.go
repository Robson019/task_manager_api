package primary

import (
	"task_manager/src/core/domain/account"
	"task_manager/src/core/errors"

	"github.com/google/uuid"
)

type AccountManager interface {
	FetchProfileByID(accountID uuid.UUID) (*account.Account, errors.Error)
}
