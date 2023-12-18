//go:build wireinject
// +build wireinject

package wire

import (
	"fiber-wire-template/internal/job"
	"fiber-wire-template/pkg/log"
	"github.com/google/wire"
)

var jobSet = wire.NewSet(job.NewJob)

func NewApp(logger *log.Logger) (*job.Job, error) {
	panic(wire.Build(
		jobSet,
	))
	return &job.Job{}, nil
}
