package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/model"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/repository"
)

type IService interface {
	Get(ctx context.Context, req *dto.DetailProfileRequest) (model.Profile, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
