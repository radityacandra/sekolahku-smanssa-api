package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/model"
	"gorm.io/gorm"
)

type IRepository interface {
	FindById(ctx context.Context, id int) *model.Profile
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
