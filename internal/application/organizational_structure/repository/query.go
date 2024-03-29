package repository

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/organizational_structure/model"
)

func (r *Repository) FindAll(c context.Context, page, pageSize int) []model.OrganizationalStructure {
	var structures []model.OrganizationalStructure

	tx := r.DB.WithContext(c).Model(&model.OrganizationalStructure{})

	if pageSize > 0 {
		tx.Limit(pageSize)
	}
	tx.Order("kd COLLATE \"num_insensitive\" ASC")

	tx.Find(&structures)

	return structures
}
