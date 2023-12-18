package main

import (
	"fiber-wire-template/cmd/app/wire"
	"fiber-wire-template/pkg/config"
	"fiber-wire-template/pkg/log"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	configName := os.Getenv("APP_CONF")
	newConfig := config.NewConfig(configName)
	logger := log.NewLog(newConfig)
	app, err := wire.NewApp(newConfig, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	if err := app.Start(newConfig.ServerCnf.Port); err != nil {
		logger.Fatal(err.Error())
	}
}
