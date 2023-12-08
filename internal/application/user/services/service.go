package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/user/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/user/repository"
)

type IService interface {
	Authentication(c context.Context, req *dto.AuthenticationRequest) (*dto.AuthenticationResponse, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
