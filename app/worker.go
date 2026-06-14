package app

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/patrickishaf/job_scheduler/config"
	"github.com/patrickishaf/job_scheduler/internal/db/repository"
	"github.com/patrickishaf/job_scheduler/internal/net"
)

type Worker struct {
	cfg             *config.AppConfig
	logger          *slog.Logger
	heapScheduler   *HeapScheduler
	ipqScheduler    *IndexedPQScheduler
	jobsRepo        *repository.JobRepository
	socketConnStore *net.SocketConnectionStore
	ticker          *time.Ticker
}

func CreateWorker(
	cfg *config.AppConfig,
	heapScheduler *HeapScheduler,
	ipqScheduler *IndexedPQScheduler,
	jobsRepo *repository.JobRepository,
	socketStroe *net.SocketConnectionStore,
	logger *slog.Logger,
) *Worker {
	return &Worker{
		cfg:             cfg,
		logger:          logger,
		heapScheduler:   heapScheduler,
		ipqScheduler:    ipqScheduler,
		jobsRepo:        jobsRepo,
		socketConnStore: socketStroe,
		ticker:          nil,
	}
}

func (this *Worker) processJob(ctx context.Context) {
	logger := this.logger.With("caller", "Worker.processJob")

	jobModels, err := this.jobsRepo.FindAllPendingJobs(ctx)
	if err != nil {
		logger.Error("failed to get all pending jobs by status", "err", err.Error())
		return
	}

	if len(jobModels) == 0 {
		logger.Info("found no pending jobs")
		return
	}

	var jobs []*Job
	for _, j := range jobModels {
		job := CreateJobFromModel(&j)
		jobs = append(jobs, job)
	}

	jobIDMap := this.ipqScheduler.EnqueueDueJobs(jobs)

	if len(jobIDMap) == 0 {
		logger.Error("failed to queue all jobs")
		return
	}

	var wg sync.WaitGroup
	for _, job := range jobModels {
		wg.Go(func() {
			_, exists := jobIDMap[job.ID]
			if exists {
				logger.Info("broadcasting socket message")
				this.socketConnStore.BroadcastMessage(&net.SocketMessage{
					Event: net.SOCKET_EVENT_JOB_QUEUED,
				})
				_, err = this.jobsRepo.MarkJobAsProcessing(ctx, job.ID, int(job.AttemptCount)+1, job.RetryCount+1)
				if err != nil {
					logger.Info("failed to mark job as processing", "job_id", job.ID, "err", err.Error())
				}
			}
		})
	}
	wg.Wait()
}

func (this *Worker) Start(ctx context.Context) error {
	this.ticker = time.NewTicker(10 * time.Second)
	for range this.ticker.C {
		go this.processJob(ctx)
	}
	return nil
}

func (this *Worker) Stop() error {
	this.ticker.Stop()
	return nil
}
