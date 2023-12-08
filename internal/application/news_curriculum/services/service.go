package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_curriculum/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_curriculum/repository"
)

type IService interface {
	List(c context.Context, req *dto.NewsCurriculumRequest) ([]dto.NewsCurriculumResponse, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
