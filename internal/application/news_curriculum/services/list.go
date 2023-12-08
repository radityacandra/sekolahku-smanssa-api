package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_curriculum/dto"
)

func (service *Service) List(c context.Context, req *dto.NewsCurriculumRequest) ([]dto.NewsCurriculumResponse, error) {
	headlines := service.Repository.FindAll(c, 0, 5)

	res := []dto.NewsCurriculumResponse{}
	for _, headline := range headlines {
		res = append(res, dto.NewsCurriculumResponse{
			ID:      headline.ID,
			Title:   headline.Title,
			Content: headline.Content,
			Image:   headline.Image,
		})
	}

	return res, nil
}
