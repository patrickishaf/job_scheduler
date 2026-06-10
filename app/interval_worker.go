package app

import (
	"log/slog"
	"time"

	"github.com/patrickishaf/job_scheduler/config"
)

type IntervalWorker struct {
	cfg    *config.AppConfig
	logger *slog.Logger
	srv    *service
	ticker *time.Ticker
}

func CreateIntervalWorker(cfg *config.AppConfig, srv *service, logger *slog.Logger) *IntervalWorker {
	return &IntervalWorker{
		cfg:    cfg,
		logger: logger,
		srv:    srv,
		ticker: nil,
	}
}

/**
 * This worker picks only one job at a time to prevent multiple workers from picking the same job
 */
func (this *IntervalWorker) processJob() {
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

func (this *IntervalWorker) Start() error {
	this.ticker = time.NewTicker(time.Second)
	for range this.ticker.C {
		go this.processJob()
	}
	return nil
}

func (this *IntervalWorker) Stop() error {
	this.ticker.Stop()
	return nil
}
