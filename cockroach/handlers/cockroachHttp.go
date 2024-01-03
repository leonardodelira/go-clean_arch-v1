package handlers

import (
	"leonardodelira/gocleanarch/cockroach/models"
	"leonardodelira/gocleanarch/cockroach/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type cockroachHttp struct {
	usecase usecases.CockroachUsecase
}

func NewCockroachHttp(cockroachUsecase usecases.CockroachUsecase) CockroachHandler {
	return &cockroachHttp{
		usecase: cockroachUsecase,
	}
}

func (h *cockroachHttp) DetectCockroach(c echo.Context) error {
	reqBody := new(models.AddCockroachData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}

	if err := h.usecase.CockroachDataProcessing(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing data failed")
	}

	return response(c, http.StatusOK, "Success 🪳🪳🪳")
}
