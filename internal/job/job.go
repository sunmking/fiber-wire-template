package job

import (
	"context"
	"fiber-wire-template/internal/job/task"
	"fiber-wire-template/pkg/log"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"
	"time"
)

type Job struct {
	log       *log.Logger
	scheduler gocron.Scheduler
	// 任务
	jobTask task.JobTask
}

func NewJob(logger *log.Logger, jobTask task.JobTask) *Job {
	return &Job{
		log:     logger,
		jobTask: jobTask,
	}
}

func (j *Job) Run(ctx context.Context) {
	j.scheduler, _ = gocron.NewScheduler()
	defer func() { _ = j.scheduler.Shutdown() }()

	_, err := j.scheduler.NewJob(
		gocron.CronJob("0/3 * * * * *", true),
		gocron.NewTask(j.jobTask.Run),
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

func (j *Job) Stop() error {
	err := j.scheduler.StopJobs()
	if err != nil {
		j.log.Error("job stop err", zap.Error(err))
		return err
	}
	return nil
}
