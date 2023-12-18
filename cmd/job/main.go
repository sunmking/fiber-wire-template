package main

import (
	"fiber-wire-template/cmd/job/wire"
	"fiber-wire-template/pkg/config"
	"fiber-wire-template/pkg/log"
	"os"
)

func main() {
	configName := os.Getenv("APP_CONF")
	newConfig := config.NewConfig(configName)
	logger := log.NewLog(newConfig)
	app, err := wire.NewApp(logger)
	if err != nil {
		logger.Error(err.Error())
	}
	app.Run()
}
