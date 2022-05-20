package routes

import (
	"github.com/labstack/echo/v4"
	"go-consul/internal/go-consul/controller"
)

func NewHealthcheckRoute(e *echo.Echo, ctrl controller.HealthcheckController) {
	e.GET("/healthcheck", ctrl.GetServerStatus)

}

func NewEnvRoute(e *echo.Echo, ctrl controller.EnvController) {
	e.GET("/env", ctrl.GetEnvByName)
}
