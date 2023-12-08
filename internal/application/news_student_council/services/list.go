package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_student_council/dto"
)

func (service *Service) List(c context.Context, req *dto.NewsStudentCouncilRequest) ([]dto.NewsStudentCouncilResponse, error) {
	headlines := service.Repository.FindAll(c, 0, 5)

	res := []dto.NewsStudentCouncilResponse{}
	for _, headline := range headlines {
		res = append(res, dto.NewsStudentCouncilResponse{
			ID:      headline.ID,
			Title:   headline.Title,
			Content: headline.Content,
			Image:   headline.Image,
		})
	}

	return res, nil
}
