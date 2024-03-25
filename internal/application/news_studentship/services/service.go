package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_studentship/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_studentship/repository"
)

type IService interface {
	List(c context.Context, req *dto.NewsStudentshipRequest) ([]dto.NewsStudentshipResponse, error)
	Get(c context.Context, req *dto.DetailNewsRequest) (dto.DetailNewsResponse, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
