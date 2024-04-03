package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/model"
)

func (r *Repository) FindById(ctx context.Context, id int) *model.Profile {
	var profile model.Profile

	r.DB.WithContext(ctx).Model(&model.Profile{}).
		Where(&model.Profile{ID: id}).
		First(&profile)

	return &profile
}
