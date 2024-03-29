package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/organizational_structure/dto"
)

func (service *Service) List(c context.Context, req *dto.OrganizationStructureRequest) ([]dto.OrganizationStructureResponse, error) {
	stakeholders := service.Repository.FindAll(c, req.Page, req.PageSize)

	res := []dto.OrganizationStructureResponse{}
	for _, stakeholder := range stakeholders {
		res = append(res, dto.OrganizationStructureResponse{
			ID:    stakeholder.ID,
			Name:  stakeholder.Name,
			Title: stakeholder.Title,
			Image: stakeholder.Image,
		})
	}

	return res, nil
}
