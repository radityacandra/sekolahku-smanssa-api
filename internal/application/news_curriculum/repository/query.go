package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_curriculum/model"
)

func (r *Repository) FindAll(c context.Context, page, pageSize int) []model.NewsCurriculum {
	var news []model.NewsCurriculum

	tx := r.DB.WithContext(c).Model(&model.NewsCurriculum{})

	if pageSize > 0 {
		tx.Limit(pageSize)
	}

	tx.Find(&news)

	return news
}

func (r *Repository) FindById(c context.Context, id int) *model.NewsCurriculum {
	var news model.NewsCurriculum

	r.DB.WithContext(c).Model(&model.NewsCurriculum{}).
		Where(&model.NewsCurriculum{ID: id}).
		First(&news)

	return &news
}
