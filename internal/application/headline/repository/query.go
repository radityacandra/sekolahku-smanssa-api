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

	tx.Find(&headlines)

	return headlines
}
