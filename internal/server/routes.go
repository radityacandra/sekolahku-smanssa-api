package server

import (
	headlineHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/headline/handler"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/user/handler"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server/types"
)

func InitRoutes(deps *types.Dependency) {
	v1 := deps.Echo.Group("/v1")

	// auth login group
	userHandler := handler.NewHandler(deps)
	v1.POST("/auth/login", userHandler.Auth)

	// headline
	headlineHandler := headlineHandler.NewHandler(deps)
	v1.GET("/headline", headlineHandler.List)
}
