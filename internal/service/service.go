
package service

import (
	"context"
	"log"
	"time"

	"bytecounter/internal/collector"
)

type Service struct {
	collector collector.Collector
}

func New() *Service {
	return &Service{
		collector: collector.NewProcCollector(),
	}
}

func (s *Service) Run(ctx context.Context) error {

	log.Println("ByteCounter Service Started")

	heartbeat := time.NewTicker(5 * time.Second)
	worker := time.NewTicker(1 * time.Second)

	defer heartbeat.Stop()
	defer worker.Stop()

	for {

		select {

		case <-ctx.Done():
			log.Println("ByteCounter Service Stopped")
			return nil

		case <-heartbeat.C:
			log.Println("Heartbeat")

		case <-worker.C:
			s.DoWork()

		}

	}

}