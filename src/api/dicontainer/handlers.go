package dicontainer

import (
	"task_manager/src/api/handlers"
)

func GetAuthHandlers() *handlers.AuthHandlers {
	return handlers.NewAuthHandlers(GetAuthServices())
}

func GetAccountHandlers() *handlers.AccountHandlers {
	return handlers.NewAccountHandlers(GetAccountServices())
}

func GetTaskHandlers() *handlers.TaskHandlers {
	return handlers.NewTaskHandlers(GetTaskServices())
}
