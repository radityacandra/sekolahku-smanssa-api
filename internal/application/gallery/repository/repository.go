package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/gallery/model"
	"gorm.io/gorm"
)

type IRepository interface {
	FindAll(c context.Context, page, pageSize int) []model.Gallery
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
