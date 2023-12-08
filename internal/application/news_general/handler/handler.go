package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_general/dto"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_general/repository"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_general/services"
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

func (h *Handler) List(c echo.Context) error {
	var req dto.NewsGeneralRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responsewrapper.WrapperError(err, 422))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, responsewrapper.WrapperError(err, 400))
	}

	res, err := h.Service.List(c.Request().Context(), &req)

	return c.JSON(http.StatusOK, responsewrapper.Wrapper(res, err, 200))
}
