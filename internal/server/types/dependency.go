package types

import (
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/sekolahku-smanssa-api/configs"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Dependency struct {
	DB     *gorm.DB
	Config *configs.Config
	Echo   *echo.Echo
	Logger *zap.Logger
}

func NewDependency(db *gorm.DB, config *configs.Config, logger *zap.Logger) *Dependency {
	return &Dependency{
		DB:     db,
		Config: config,
		Logger: logger,
	}
}
