package server

import (
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server/types"
)

func InitServer(deps *types.Dependency) {
	e := echo.New()
	deps.Echo = e

	InitRoutes(deps)

	e.Logger.Fatal(e.Start(":" + deps.Config.ServerPort))
}
