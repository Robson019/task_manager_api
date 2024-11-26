package dicontainer

import (
	"task_manager/src/core/interfaces/repository"
	"task_manager/src/infra/postgres"
	"task_manager/src/infra/redis"
)

func GetAuthRepository() repository.AuthLoader {
	return postgres.NewAuthPostgresRepository(GetPsqlConnectionManager())
}

func GetAccountRepository() repository.AccountLoader {
	return postgres.NewAccountPostgresRepository(GetPsqlConnectionManager())
}

func GetTaskRepository() repository.TaskLoader {
	return postgres.NewTaskPostgresRepository(GetPsqlConnectionManager())
}

func GetSessionRepository() repository.SessionLoader {
	return redis.NewSessionRepository(GetRedisConnectionManager())
}

func GetPsqlConnectionManager() *postgres.DatabaseConnectionManager {
	return &postgres.DatabaseConnectionManager{}
}

func GetRedisConnectionManager() *redis.RedisConnectionManager {
	return &redis.RedisConnectionManager{}
}
