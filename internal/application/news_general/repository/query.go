package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_general/model"
)

func (r *Repository) FindAll(c context.Context, page, pageSize int) []model.NewsGeneral {
	var news []model.NewsGeneral

	tx := r.DB.WithContext(c).Model(&model.NewsGeneral{})

	if pageSize > 0 {
		tx.Limit(pageSize)
	}

	tx.Find(&news)

	return news
}
