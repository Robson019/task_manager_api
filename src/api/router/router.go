package router

import (
	_ "task_manager/src/api/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router interface {
	Load(*echo.Group)
}

type router struct {
}

func New() Router {
	return &router{}
}

func (instance *router) Load(group *echo.Group) {
	group.GET("/docs/*", echoSwagger.WrapHandler)

	loadAuthRoutes(group)
	loadAccountRoutes(group)
	loadTaskRoutes(group)
}
