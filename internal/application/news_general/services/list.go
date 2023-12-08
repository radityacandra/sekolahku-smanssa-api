package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_general/dto"
)

func (service *Service) List(c context.Context, req *dto.NewsGeneralRequest) ([]dto.NewsGeneralResponse, error) {
	headlines := service.Repository.FindAll(c, 0, 5)

	res := []dto.NewsGeneralResponse{}
	for _, headline := range headlines {
		res = append(res, dto.NewsGeneralResponse{
			ID:      headline.ID,
			Title:   headline.Title,
			Content: headline.Content,
			Image:   headline.Image,
		})
	}

	return res, nil
}
