package main

import (
	"context"
	"fiber-wire-template/cmd/command/wire"
	"fiber-wire-template/internal/command"
	"fiber-wire-template/pkg/config"
	"fiber-wire-template/pkg/log"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"os"
)

func main() {
	configName := os.Getenv("APP_CONF")
	newConfig := config.NewConfig(configName)
	logger := log.NewLog(newConfig)
	app, err := wire.NewApp(logger, newConfig)
	if err != nil {
		logger.Error(err.Error())
	}
	app.Run(context.Background())
	defer func(app *command.Command) {
		err := app.Stop()
		if err != nil {
			logger.Error("app command stop err", zap.Error(err), zap.String("app", "command"))
		}
	}(app)
}
