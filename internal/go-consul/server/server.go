package server

import (
	"github.com/labstack/echo/v4"
	"go-consul/internal/go-consul/config"
	"go-consul/internal/go-consul/controller"
	"go-consul/internal/go-consul/routes"
	"log"
)

func LoadServer(cfg *config.Configuration) (*echo.Echo, error) {
	log.Println("Initializing dependencies")

	e := echo.New()

	envController := controller.NewEnvController(cfg)
	healthcheckController := controller.NewHealthcheckController(cfg)

	routes.NewEnvRoute(e, envController)
	routes.NewHealthcheckRoute(e, healthcheckController)

	return e, nil
}
