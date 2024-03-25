package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_student_council/model"
)

func (r *Repository) FindAll(c context.Context, page, pageSize int) []model.NewsStudentCouncil {
	var news []model.NewsStudentCouncil

	tx := r.DB.WithContext(c).Model(&model.NewsStudentCouncil{})

	if pageSize > 0 {
		tx.Limit(pageSize)
	}

	tx.Find(&news)

	return news
}

func (r *Repository) FindById(c context.Context, id int) *model.NewsStudentCouncil {
	var news model.NewsStudentCouncil

	r.DB.WithContext(c).Model(&model.NewsStudentCouncil{}).
		Where(&model.NewsStudentCouncil{ID: id}).
		First(&news)

	return &news
}
