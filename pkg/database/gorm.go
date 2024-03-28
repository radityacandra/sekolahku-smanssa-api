package database

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
)

func NewDatabase(connection string, logger *zap.Logger) (*gorm.DB, error) {
	gormLogger := New(logger)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connection,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: gormLogger.LogMode(gl.Info),
	})
	if err != nil {
		return nil, err
	}

	return db.Debug(), nil
}
