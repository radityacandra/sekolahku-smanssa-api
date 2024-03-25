package services

import (
	"context"
	"errors"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_curriculum/dto"
)

func (service *Service) Get(c context.Context, req *dto.DetailNewsRequest) (dto.DetailNewsResponse, error) {
	news := service.Repository.FindById(c, req.ID)
	if news == nil {
		return dto.DetailNewsResponse{}, errors.New("data not found")
	}

	return dto.DetailNewsResponse{
		ID:          news.ID,
		Title:       news.Title,
		Content:     news.Content,
		Image:       news.Image,
		PublishedAt: news.Date,
		Author:      news.Author,
	}, nil
}
