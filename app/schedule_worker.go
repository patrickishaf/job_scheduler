package app

import (
	"log/slog"
	"time"

	"github.com/patrickishaf/job_scheduler/config"
)

type ScheduleWorker struct {
	cfg    *config.AppConfig
	logger *slog.Logger
	srv    *service
	ticker *time.Ticker
}

func CreateScheduledWorker(cfg *config.AppConfig, srv *service, logger *slog.Logger) *ScheduleWorker {
	return &ScheduleWorker{
		cfg:    cfg,
		logger: logger,
		srv:    srv,
		ticker: nil,
	}
}

func (this *ScheduleWorker) processJob() {
	/**
	 * An interval job is a job whose value of interval_minutes is not null
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

func (this *ScheduleWorker) Start() error {
	this.ticker = time.NewTicker(1 * time.Second / 2)
	for range this.ticker.C {
		go this.processJob()
	}
	return nil
}

func (this *ScheduleWorker) Stop() error {
	this.ticker.Stop()
	return nil
}
