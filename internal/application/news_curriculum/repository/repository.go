package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_curriculum/model"
	"gorm.io/gorm"
)

type IRepository interface {
	FindAll(c context.Context, page, pageSize int) []model.NewsCurriculum
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
