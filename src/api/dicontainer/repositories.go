package dicontainer

import (
	"task_manager/src/core/interfaces/secondary"
	"task_manager/src/infra/postgres"
	"task_manager/src/infra/redis"
)

func GetAuthRepository() secondary.AuthLoader {
	return postgres.NewAuthPostgresRepository(GetPsqlConnectionManager())
}

func GetAccountRepository() secondary.AccountLoader {
	return postgres.NewAccountPostgresRepository(GetPsqlConnectionManager())
}

func GetTaskRepository() secondary.TaskLoader {
	return postgres.NewTaskPostgresRepository(GetPsqlConnectionManager())
}

func GetSessionRepository() secondary.SessionLoader {
	return redis.NewSessionRepository(GetRedisConnectionManager())
}

func GetPsqlConnectionManager() *postgres.DatabaseConnectionManager {
	return &postgres.DatabaseConnectionManager{}
}

func GetRedisConnectionManager() *redis.RedisConnectionManager {
	return &redis.RedisConnectionManager{}
}
