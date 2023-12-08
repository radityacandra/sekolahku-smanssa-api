package services

import (
	"context"

	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/user/dto"
)

func (service *Service) Authentication(c context.Context, req *dto.AuthenticationRequest) (*dto.AuthenticationResponse, error) {
	return &dto.AuthenticationResponse{}, nil
}
