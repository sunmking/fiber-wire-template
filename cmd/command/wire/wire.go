//go:build wireinject
// +build wireinject

package wire

import (
	"fiber-wire-template/internal/command"
	"fiber-wire-template/internal/command/task"
	"fiber-wire-template/internal/repository"
	"fiber-wire-template/pkg/config"
	"fiber-wire-template/pkg/gredis"
	"fiber-wire-template/pkg/log"
	"fiber-wire-template/pkg/ozzodb"
	"github.com/google/wire"
)

var jobSet = wire.NewSet(
	command.NewCommand,
)
var RepositorySet = wire.NewSet(
	ozzodb.NewDb,
	gredis.NewRedis,
	repository.NewRepository,
)

func NewApp(logger *log.Logger, config *config.Config) (*command.Command, error) {
	panic(wire.Build(
		RepositorySet,
		task.NewDemoTask,
		jobSet,
	))
	return &command.Command{}, nil
}
