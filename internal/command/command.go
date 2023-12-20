package command

import (
	"context"
	"fiber-wire-template/internal/command/task"
	"fiber-wire-template/pkg/log"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"
	"time"
)

type Command struct {
	log       *log.Logger
	scheduler gocron.Scheduler
	// 任务
	demoTask task.DemoTask
}

func NewCommand(logger *log.Logger, demoTask task.DemoTask) *Command {
	return &Command{
		log:      logger,
		demoTask: demoTask,
	}
}

func (j *Command) Run(ctx context.Context) {
	j.scheduler, _ = gocron.NewScheduler()
	defer func() { _ = j.scheduler.Shutdown() }()

	_, err := j.scheduler.NewJob(
		gocron.CronJob("0/3 * * * * *", true),
		gocron.NewTask(j.demoTask.Run),
	)
	if err != nil {
		j.log.Error("Task1 error", zap.Error(err))
	}
	_, err = j.scheduler.NewJob(
		gocron.CronJob("0/1 * * * * *", true),
		gocron.NewTask(func() {
			j.log.Info("I'm a Task2.")
		}),
	)
	if err != nil {
		j.log.Error("Task2 error", zap.Error(err))
	}

	_, _ = j.scheduler.NewJob(
		gocron.DurationJob(
			10*time.Millisecond,
		),
		gocron.NewTask(
			func(one string, two int) {
				fmt.Printf("%s, %d\n %d", one, two, time.Now().UnixNano())
			},
			"one", 2,
		),
		gocron.WithLimitedRuns(200),
	)
	j.scheduler.Start()
	select {}
}

func (j *Command) Stop() error {
	err := j.scheduler.StopJobs()
	if err != nil {
		j.log.Error("command stop err", zap.Error(err))
		return err
	}
	return nil
}
