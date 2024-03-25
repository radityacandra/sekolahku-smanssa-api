package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/gallery/model"
)

func (r *Repository) FindAll(c context.Context, page, pageSize int) []model.Gallery {
	var news []model.Gallery

	tx := r.DB.WithContext(c).Model(&model.Gallery{})

	if pageSize > 0 {
		tx.Limit(pageSize)
	}

	tx.Find(&news)

	return news
}
