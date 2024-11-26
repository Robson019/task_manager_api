package services

import (
	"github.com/google/uuid"
	"task_manager/src/core/domain/account"
	"task_manager/src/core/errors"
	"task_manager/src/core/errors/logger"
	"task_manager/src/core/interfaces/primary"
	"task_manager/src/core/interfaces/repository"
)

var _ primary.AccountManager = (*AccountServices)(nil)

type AccountServices struct {
	accountRepository repository.AccountLoader
	logger            logger.Logger
}

func (instance AccountServices) FetchProfileByID(accountID uuid.UUID) (*account.Account, errors.Error) {
	accountInstance, err := instance.accountRepository.FindProfileByID(accountID)
	if err != nil {
		instance.logger.Log(err)
		return nil, err
	}

	return accountInstance, nil
}

func NewAccountServices(AccountRepository repository.AccountLoader, logger logger.Logger) *AccountServices {
	return &AccountServices{
		accountRepository: AccountRepository,
		logger:            logger,
	}
}
