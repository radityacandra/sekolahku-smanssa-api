package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/headline/model"
)

func (r *Repository) FindAll(c context.Context, page, pageSize int) []model.Headline {
	var headlines []model.Headline

	tx := r.DB.WithContext(c).Model(&model.Headline{})

	if pageSize > 0 {
		tx.Limit(pageSize)
	}

	if page > 0 {
		tx.Offset((page - 1) * pageSize)
	}

	tx.Find(&headlines)

	return headlines
}

func (r *Repository) FindById(c context.Context, id int) *model.Headline {
	var headline model.Headline

	r.DB.WithContext(c).Model(&model.Headline{}).
		Where(&model.Headline{ID: id}).
		First(&headline)

	return &headline
}
