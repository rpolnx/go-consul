package controller

import (
	"github.com/labstack/echo/v4"
	"go-consul/internal/go-consul/config"
	"net/http"
)

type HealthcheckController interface {
	GetServerStatus(c echo.Context) error
}

type healthcheckController struct {
	cfg *config.Configuration
}

func (h *healthcheckController) GetServerStatus(c echo.Context) error {
	m := map[string]interface{}{"status": "OK"}

	return c.JSON(http.StatusOK, m)
}

func NewHealthcheckController(cfg *config.Configuration) *healthcheckController {
	return &healthcheckController{
		cfg: cfg,
	}
}
