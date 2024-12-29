package api

import (
	"log"

	"github.com/go-co-op/gocron/v2"
)

const (
	CronEverydayAtMidnight = "0 0 * * *"
	CronEveryHour          = "0 * * * *"
)

type jobServer struct {
	scheduler gocron.Scheduler
	running   map[string]struct{}
}

func NewJobServer() (*jobServer, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &jobServer{
		scheduler: s,
		running:   map[string]struct{}{},
	}, nil
}

func (s *jobServer) runTask(cron, name string, task func()) error {
	_, err := s.scheduler.NewJob(gocron.CronJob(cron, false), gocron.NewTask(func() {
		_, ok := s.running[name]
		if ok {
			return
		}

		s.running[name] = struct{}{}
		task()
		delete(s.running, name)
	}), gocron.WithName(name))
	if err != nil {
		log.Printf("Can't start '%s' job\n", name)
	}

	return err
}

func (s *jobServer) Start() error {
	log.Print("TODO JOB")
	return nil
}
