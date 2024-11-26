package postgres

import (
	"context"
	"strings"
	"task_manager/src/core/domain/account"
	"task_manager/src/core/domain/account/credentials"
	"task_manager/src/core/domain/account/role"
	"task_manager/src/core/errors"
	"task_manager/src/core/interfaces/repository"
	"task_manager/src/core/messages"
	"task_manager/src/infra/postgres/bridge"
	postgresmsgs "task_manager/src/infra/postgres/messages"
)

var _ repository.AuthLoader = &AuthPostgresRepository{}

type AuthPostgresRepository struct {
	connectorManager
}

func (instance AuthPostgresRepository) Login(credentials credentials.Credentials) (*account.Account, errors.Error) {
	conn, err := instance.getConnection()
	if err != nil {
		return nil, errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer instance.closeConnection(conn)

	query := bridge.New(conn)
	row, err := query.SelectAccountByEmail(context.Background(), credentials.Email())
	if err != nil {
		return nil, instance.handleError(err)
	}

	roleBuilder := role.NewBuilder()
	roleBuilder.WithID(row.RoleID).WithName(row.RoleName)
	roleInstance, validationError := roleBuilder.Build()

	if validationError != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, validationError)
	}
	accountBuilder := account.NewBuilder()
	accountBuilder.WithID(row.AccountID).WithEmail(row.AccountEmail).WithPassword(row.AccountPassword)
	accountBuilder.WithHash(row.AccountHash).WithRole(*roleInstance)
	accountInstance, validationError := accountBuilder.Build()
	if validationError != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, validationError)
	}

	return accountInstance, nil
}

func (instance AuthPostgresRepository) handleError(err error) errors.Error {
	msg := err.Error()

	if strings.Contains(msg, postgresmsgs.SQLNoRowsInResultSet) {
		return errors.NewNotFoundError(messages.InvalidCredentialsErrorMessage, err)
	}

	return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
}

func NewAuthPostgresRepository(manager connectorManager) repository.AuthLoader {
	return &AuthPostgresRepository{manager}
}
