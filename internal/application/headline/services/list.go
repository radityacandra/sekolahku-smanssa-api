package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/headline/dto"
)

func (service *Service) List(c context.Context, req *dto.HeadlineListRequest) ([]dto.HeadlineListResponse, error) {
	headlines := service.Repository.FindAll(c, req.Page, req.PageSize)

	res := []dto.HeadlineListResponse{}
	for _, headline := range headlines {
		res = append(res, dto.HeadlineListResponse{
			ID:      headline.ID,
			Title:   headline.Title,
			Content: headline.Content,
			Image:   headline.Image,
		})
	}

	return res, nil
}
