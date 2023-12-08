package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/headline/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/headline/repository"
)

type IService interface {
	List(c context.Context, req *dto.HeadlineListRequest) ([]dto.HeadlineListResponse, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
