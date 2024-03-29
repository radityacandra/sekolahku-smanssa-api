package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/organizational_structure/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/organizational_structure/repository"
)

type IService interface {
	List(c context.Context, req *dto.OrganizationStructureRequest) ([]dto.OrganizationStructureResponse, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
