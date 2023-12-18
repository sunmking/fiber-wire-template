package job

import (
	"fiber-wire-template/pkg/log"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"time"
)

type Job struct {
	log *log.Logger
}

func NewJob(logger *log.Logger) *Job {
	return &Job{
		log: logger,
	}
}

func (j *Job) Run() {
	s, err := gocron.NewScheduler()
	if err != nil {
		// handle error
	}

	// add a job to the scheduler
	job, err := s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func(a string, b int) {
				// do things
				j.log.Info(a)
			},
			"hello",
			1,
		),
	)
	if err != nil {
		// handle error
	}
	// each job has a unique id
	fmt.Println(job.ID())

	// start the scheduler
	s.Start()

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
	}
}
