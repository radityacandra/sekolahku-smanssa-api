package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/gallery/dto"
)

func (service *Service) List(c context.Context, req *dto.GalleryRequest) ([]dto.GalleryResponse, error) {
	headlines := service.Repository.FindAll(c, req.Page, req.PageSize)

	res := []dto.GalleryResponse{}
	for _, headline := range headlines {
		res = append(res, dto.GalleryResponse{
			ID:    headline.ID,
			Title: headline.Title,
			Image: headline.Image,
		})
	}

	return res, nil
}
