package main

import (
	"api-server/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	router := initRouter()
	setRoutes(router)

	return router
}

func initRouter() *echo.Echo {
	router := echo.New()

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	return router
}

func setRoutes(router *echo.Echo) {
	router.GET("/", controllers.Hello)
}
