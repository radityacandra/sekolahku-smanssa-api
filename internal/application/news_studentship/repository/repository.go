package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_studentship/model"
	"gorm.io/gorm"
)

type IRepository interface {
	FindAll(c context.Context, page, pageSize int) []model.NewsStudentship
	FindById(c context.Context, id int) *model.NewsStudentship
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
