package router

import (
	"task_manager/src/api/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadAuthRoutes(group *echo.Group) {
	accountGroup := group.Group("/auth")
	authHandlers := dicontainer.GetAuthHandlers()
	accountGroup.POST("/login", authHandlers.Login)
	accountGroup.POST("/refresh", authHandlers.Refresh)
	accountGroup.DELETE("/logout", authHandlers.Logout)
}
