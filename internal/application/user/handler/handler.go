package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/user/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/user/repository"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/user/services"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server/types"
)

type Handler struct {
	Service services.IService
}

func NewHandler(deps *types.Dependency) *Handler {
	repository := repository.NewRepository(deps.DB)
	service := services.NewService(repository)

	return &Handler{
		Service: service,
	}
}

func (h *Handler) Auth(c echo.Context) error {
	var req dto.AuthenticationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(req); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := h.Service.Authentication(c.Request().Context(), nil)

	if err == nil {
		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusBadRequest, errors.New("failed to authenticate"))
}
