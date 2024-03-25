package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_studentship/model"
)

func (r *Repository) FindAll(c context.Context, page, pageSize int) []model.NewsStudentship {
	var news []model.NewsStudentship

	tx := r.DB.WithContext(c).Model(&model.NewsStudentship{})

	if pageSize > 0 {
		tx.Limit(pageSize)
	}

	tx.Find(&news)

	return news
}

func (r *Repository) FindById(c context.Context, id int) *model.NewsStudentship {
	var news model.NewsStudentship

	r.DB.WithContext(c).Model(&model.NewsStudentship{}).
		Where(&model.NewsStudentship{ID: id}).
		First(&news)

	return &news
}
