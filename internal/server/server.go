package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server/types"
	"go.uber.org/zap"
)

func InitServer(deps *types.Dependency) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORS())
	e.Use(nrecho.Middleware(deps.NewRelic))

	deps.Echo = e

	InitRoutes(deps)

	deps.Logger.Info("Server ready", zap.Int("port", 8080))

	e.Logger.Fatal(e.Start(":" + deps.Config.ServerPort))
}
