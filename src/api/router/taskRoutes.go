package router

import (
	"task_manager/src/api/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadTaskRoutes(group *echo.Group) {
	taskGroup := group.Group("/tasks")

	taskHandlers := dicontainer.GetTaskHandlers()

	taskGroup.POST("", taskHandlers.CreateTask)
	taskGroup.GET("", taskHandlers.FindTasks)
	taskGroup.GET("/:id", taskHandlers.FindTaskByID)
	taskGroup.PUT("/:id", taskHandlers.UpdateTask)
	taskGroup.DELETE("/:id", taskHandlers.DeleteTask)
}
