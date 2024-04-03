package services

import (
	"context"
	"errors"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/model"
)

func (service *Service) Get(ctx context.Context, req *dto.DetailProfileRequest) (model.Profile, error) {
	profile := service.Repository.FindById(ctx, req.ID)
	if profile == nil {
		return model.Profile{}, errors.New("data not found")
	}

	return *profile, nil
}
