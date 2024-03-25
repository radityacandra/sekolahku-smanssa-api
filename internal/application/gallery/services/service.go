package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/gallery/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/gallery/repository"
)

type IService interface {
	List(c context.Context, req *dto.GalleryRequest) ([]dto.GalleryResponse, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
