//go:build wireinject
// +build wireinject

package wire

import (
	"fiber-wire-template/internal/job"
	"fiber-wire-template/internal/job/task"
	"fiber-wire-template/internal/repository"
	"fiber-wire-template/pkg/config"
	"fiber-wire-template/pkg/gredis"
	"fiber-wire-template/pkg/log"
	"fiber-wire-template/pkg/ozzodb"
	"github.com/google/wire"
)

var jobSet = wire.NewSet(
	job.NewJob,
)
var RepositorySet = wire.NewSet(
	ozzodb.NewDb,
	gredis.NewRedis,
	repository.NewRepository,
)

func NewApp(logger *log.Logger, config *config.Config) (*job.Job, error) {
	panic(wire.Build(
		RepositorySet,
		task.NewJobTask,
		jobSet,
	))
	return &job.Job{}, nil
}
