package services

import (
	"task_manager/src/core/domain/account/credentials"
	"task_manager/src/core/domain/authorization"
	"task_manager/src/core/errors"
	"task_manager/src/core/errors/logger"
	"task_manager/src/core/interfaces/primary"
	"task_manager/src/core/interfaces/secondary"
	"task_manager/src/core/messages"
	"task_manager/src/utils/encrypt"

	"github.com/google/uuid"
)

var _ primary.AuthManager = (*AuthServices)(nil)

type AuthServices struct {
	accountRepository secondary.AccountLoader
	sessionRepository secondary.SessionLoader
	authRepository    secondary.AuthLoader
	logger            logger.Logger
}

func (instance AuthServices) Login(credentials credentials.Credentials) (authorization.Authorization, errors.Error) {
	account, err := instance.authRepository.Login(credentials)
	if err != nil {
		instance.logger.Log(err)
		return nil, err
	}

	passwordsMatch := encrypt.PasswordsMatch(credentials.Password(), account.Hash(), account.Password())
	if !passwordsMatch {
		return nil, errors.NewNotFoundError(messages.InvalidCredentialsErrorMessage, nil)
	}

	auth, authError := authorization.NewFromAccount(*account)
	if authError != nil {
		return nil, authError
	}

	err = instance.sessionRepository.Store(account.ID(), auth.Token(), auth.RefreshToken())
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (instance AuthServices) Refresh(refreshToken string, accountID uuid.UUID) (authorization.Authorization, errors.Error) {
	account, err := instance.accountRepository.FindProfileByID(accountID)
	if err != nil {
		instance.logger.Log(err)
		return nil, err
	}

	exists, err := instance.sessionRepository.RefreshTokenExists(account.ID(), refreshToken)
	if err != nil {
		return nil, err
	}

	if !exists {
		if err := instance.sessionRepository.Close(account.ID()); err != nil {
			return nil, err
		}
		return nil, errors.NewUnauthorizedError(messages.InvalidRefreshTokenErrorMessage, nil)
	}

	auth, authError := authorization.NewFromAccount(*account)
	if authError != nil {
		return nil, authError
	}

	err = instance.sessionRepository.Store(account.ID(), auth.Token(), auth.RefreshToken())
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (instance AuthServices) SessionExists(accountID uuid.UUID, token string) (bool, errors.Error) {
	sessionExists, err := instance.sessionRepository.SessionExists(accountID, token)
	if err != nil {
		instance.logger.Log(err)
		return false, err
	}

	if !sessionExists {
		return false, nil
	}

	return true, nil
}

func (instance AuthServices) Logout(accountID uuid.UUID) errors.Error {
	err := instance.sessionRepository.Close(accountID)
	if err != nil {
		instance.logger.Log(err)
		return err
	}

	return nil
}

func NewAuthServices(authRepository secondary.AuthLoader, accountRepository secondary.AccountLoader, sessionRepository secondary.SessionLoader, logger logger.Logger) *AuthServices {
	return &AuthServices{
		sessionRepository: sessionRepository,
		authRepository:    authRepository,
		accountRepository: accountRepository,
		logger:            logger,
	}
}
