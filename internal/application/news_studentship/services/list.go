package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_studentship/dto"
)

func (service *Service) List(c context.Context, req *dto.NewsStudentshipRequest) ([]dto.NewsStudentshipResponse, error) {
	headlines := service.Repository.FindAll(c, 0, 5)

	res := []dto.NewsStudentshipResponse{}
	for _, headline := range headlines {
		res = append(res, dto.NewsStudentshipResponse{
			ID:      headline.ID,
			Title:   headline.Title,
			Content: headline.Content,
			Image:   headline.Image,
		})
	}

	return res, nil
}
