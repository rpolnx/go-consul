package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-consul/internal/go-consul/config"
	"net/http"
	"os"
)

type EnvController interface {
	GetEnvByName(c echo.Context) error
}

type envController struct {
	cfg *config.Configuration
}

func (ctrl *envController) GetEnvByName(c echo.Context) error {
	envName := c.Param("env_name")
	envValue := os.Getenv(envName)

	m := map[string]interface{}{"name": envName, "value": envValue}

	logrus.Info(ctrl.cfg.App.Profile)

	return c.JSON(http.StatusOK, m)
}

func NewEnvController(cfg *config.Configuration) *envController {
	return &envController{
		cfg: cfg,
	}
}
