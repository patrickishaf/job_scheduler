package app

import (
	"log/slog"
	"time"

	"github.com/patrickishaf/job_scheduler/config"
)

type Worker struct {
	cfg           *config.AppConfig
	logger        *slog.Logger
	srv           *service
	heapScheduler *HeapScheduler
	ipqScheduler  *IndexedPQScheduler
	ticker        *time.Ticker
}

func CreateWorker(
	cfg *config.AppConfig,
	srv *service,
	heapScheduler *HeapScheduler,
	ipqScheduler *IndexedPQScheduler,
	logger *slog.Logger,
) *Worker {
	return &Worker{
		cfg:           cfg,
		logger:        logger,
		srv:           srv,
		heapScheduler: heapScheduler,
		ipqScheduler:  ipqScheduler,
		ticker:        nil,
	}
}

func (this *Worker) processJob() {
	/**
	 * An interval job is a job whose value of interval is not null
		* Get the next interval job with status == pending and retry_count < cfg.MaxJobRetryCount and scheduled_time <= time.Now
		* update the job status to processing
		* update attempt_count to +1
		* if attempt_count > 0 update retry count to +1.
		* process processJob
		* if successful, update job status to successful
		* if failed and MaxRetryCount == cfg.MaxJobRetryCount, add to dead letter queue and delete job
		* if failed, increment retry count and
	*/
}

func (this *Worker) Start() error {
	this.ticker = time.NewTicker(1 * time.Second / 2)
	for range this.ticker.C {
		go this.processJob()
	}
	return nil
}

func (this *Worker) Stop() error {
	this.ticker.Stop()
	return nil
}
