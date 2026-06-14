package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/patrickishaf/job_scheduler/config"
	"github.com/patrickishaf/job_scheduler/internal/net"
)

type QueueWorker struct {
	cfg             *config.AppConfig
	heapScheduler   *HeapScheduler
	ipqScheduler    *IndexedPQScheduler
	logger          *slog.Logger
	socketConnStore *net.SocketConnectionStore
	ticker          *time.Ticker
}

func CreateQueueWorker(cfg *config.AppConfig, heapScheduler *HeapScheduler, ipqScheduler *IndexedPQScheduler, socketStore *net.SocketConnectionStore, logger *slog.Logger) *QueueWorker {
	return &QueueWorker{
		cfg:             cfg,
		heapScheduler:   heapScheduler,
		ipqScheduler:    ipqScheduler,
		logger:          logger,
		socketConnStore: socketStore,
	}
}

func (this *QueueWorker) processJob(_ context.Context) {
	logger := this.logger.With("caller", "QueueWorker.processJob")

	job, ok := this.heapScheduler.Dequeue()
	if !ok {
		logger.Error("failed to dequeue job")
		return
	}

	this.socketConnStore.BroadcastMessage(&net.SocketMessage{
		Event: net.SOCKET_EVENT_JOB_RUNNING,
	})
	logger.Info("dequeued job", "job_id", job.ID)
}

func (this *QueueWorker) Start(ctx context.Context) error {
	this.ticker = time.NewTicker(10 * time.Second)
	for range this.ticker.C {
		go this.processJob(ctx)
	}
	return nil
}

func (this *QueueWorker) Stop() error {
	this.ticker.Stop()
	return nil
}
