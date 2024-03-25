package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server/types"
)

func InitServer(deps *types.Dependency) {
	e := echo.New()
	e.Use(middleware.CORS())

	deps.Echo = e

	InitRoutes(deps)

	e.Logger.Fatal(e.Start(":" + deps.Config.ServerPort))
}
