package router

import (
	"github.com/labstack/echo/v4"
	"task_manager/src/api/dicontainer"
)

func loadAccountRoutes(group *echo.Group) {
	accountGroup := group.Group("/account")
	accountHandlers := dicontainer.GetAccountHandlers()
	accountGroup.GET("/profile", accountHandlers.FindProfile)
}
