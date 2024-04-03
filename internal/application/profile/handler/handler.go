package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/repository"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/profile/services"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server/types"
	responsewrapper "github.com/radityacandra/sekolahku-smanssa-api/pkg/response_wrapper"
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

func (h *Handler) Get(c echo.Context) error {
	var req dto.DetailProfileRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responsewrapper.WrapperError(err, 422))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, responsewrapper.WrapperError(err, 400))
	}

	res, err := h.Service.Get(c.Request().Context(), &req)

	return c.JSON(http.StatusOK, responsewrapper.Wrapper(res, err, 200))
}
