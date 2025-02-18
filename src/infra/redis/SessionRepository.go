package redis

import (
	"fmt"
	"task_manager/src/core/domain/authorization"
	"task_manager/src/core/errors"
	"task_manager/src/core/interfaces/secondary"
	"task_manager/src/core/messages"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var _ secondary.SessionLoader = (*redisSessionRepository)(nil)

type redisSessionRepository struct {
	connectorManager
}

func NewSessionRepository(manager connectorManager) *redisSessionRepository {
	return &redisSessionRepository{connectorManager: manager}
}

func (repo redisSessionRepository) Store(uID uuid.UUID, accessToken string, refreshToken string) errors.Error {
	conn, err := repo.getConnection()
	if err != nil {
		return errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer repo.closeConnection(conn)

	err = conn.Set(fmt.Sprintf("access:%s", uID), accessToken, authorization.AccessTokenTimeout).Err()
	if err != nil {
		return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	err = conn.Set(fmt.Sprintf("refresh:%s", uID), refreshToken, authorization.RefreshTokenTimeout).Err()
	if err != nil {
		return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	return nil
}

func (repo redisSessionRepository) SessionExists(uID uuid.UUID, handledToken string) (bool, errors.Error) {
	conn, err := repo.getConnection()
	if err != nil {
		return false, errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer repo.closeConnection(conn)

	aKey := fmt.Sprintf("access:%s", uID)
	keyExists, err := conn.Exists(aKey).Result()
	if err != nil {
		log.Error().Msg(err.Error())
		return false, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, err)
	}

	if keyExists <= 0 {
		return false, nil
	}

	storedValue, err := conn.Get(aKey).Result()
	if err != nil {
		log.Error().Msg(err.Error())
		return false, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, err)
	}

	return storedValue == handledToken, nil
}

func (repo redisSessionRepository) RefreshTokenExists(uID uuid.UUID, handledToken string) (bool, errors.Error) {
	conn, err := repo.getConnection()
	if err != nil {
		return false, errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer repo.closeConnection(conn)

	rKey := fmt.Sprintf("refresh:%s", uID)
	keyExists, err := conn.Exists(rKey).Result()
	if err != nil {
		log.Error().Msg(err.Error())
		return false, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, err)
	}

	if keyExists <= 0 {
		return false, nil
	}

	storedValue, err := conn.Get(rKey).Result()
	if err != nil {
		log.Error().Msg(err.Error())
		return false, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, err)
	}

	return storedValue == handledToken, nil
}

func (repo redisSessionRepository) Close(uID uuid.UUID) errors.Error {
	conn, err := repo.getConnection()
	if err != nil {
		return errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer repo.closeConnection(conn)

	aKey := fmt.Sprintf("access:%s", uID)
	rKey := fmt.Sprintf("refresh:%s", uID)

	result := conn.Del(aKey, rKey)
	if result.Err() != nil {
		return errors.NewUnexpectedError(messages.InsertingDataErrorMessage, result.Err())
	}

	return nil
}
