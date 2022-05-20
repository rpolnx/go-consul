package main

import (
	"github.com/sirupsen/logrus"
	"go-consul/internal/go-consul/config"
	"go-consul/internal/go-consul/server"
)

func main() {
	logrus.Infof("Initializing server")

	loadConfig, err := config.LoadConfig()

	if err != nil {
		logrus.Fatal(err)
	}

	currentServer, err := server.LoadServer(loadConfig)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Starting server on port 8080")

	logrus.Fatal(currentServer.Start(":8080"))
}
