package types

import (
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/sekolahku-smanssa-api/configs"
	"gorm.io/gorm"
)

type Dependency struct {
	DB     *gorm.DB
	Config *configs.Config
	Echo   *echo.Echo
}

func NewDependency(db *gorm.DB, config *configs.Config) *Dependency {
	return &Dependency{
		DB:     db,
		Config: config,
	}
}
