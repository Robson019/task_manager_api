package dicontainer

import (
	"os"
	"task_manager/src/core/errors/logger"
	"task_manager/src/core/interfaces/primary"
	"task_manager/src/core/services"
)

const API_TYPE = "API"

func GetAuthServices() primary.AuthManager {
	return services.NewAuthServices(GetAuthRepository(), GetAccountRepository(), GetSessionRepository(), GetLogger())
}

func GetAccountServices() primary.AccountManager {
	return services.NewAccountServices(GetAccountRepository(), GetLogger())
}

func GetTaskServices() primary.TaskManager {
	return services.NewTaskServices(GetTaskRepository(), GetLogger())
}

func GetLogger() logger.Logger {
	return logger.New()
}

func getAppType() string {
	return os.Getenv("APPLICATION_TYPE")
}

func appIsAPI() bool {
	return getAppType() == API_TYPE
}
