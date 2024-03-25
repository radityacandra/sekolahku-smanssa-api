package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_student_council/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_student_council/repository"
)

type IService interface {
	List(c context.Context, req *dto.NewsStudentCouncilRequest) ([]dto.NewsStudentCouncilResponse, error)
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
