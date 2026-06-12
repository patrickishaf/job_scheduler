package app

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/patrickishaf/job_scheduler/internal/infra"
)

type PriorityQueue struct {
	cache          *infra.RedisCache
	collectionName string
	logger         *slog.Logger
}

func CreatePriorityQueue(cache *infra.RedisCache, logger *slog.Logger) *PriorityQueue {
	return &PriorityQueue{
		cache:          cache,
		collectionName: "priority_queue",
		logger:         logger,
	}
}

func (this *PriorityQueue) AddJob(ctx context.Context, job *Job) {
	logger := this.logger.With("caller", "PriorityQueue.AddJob")
	val, err := json.Marshal(job)
	if err != nil {
		logger.Error("failed to marshal job into json", "err", err.Error())
		return
	}
	if err := this.cache.Queue(ctx, this.collectionName, string(val)); err != nil {
		logger.Error("failed to queue job", "err", err.Error())
	}
}

func (this *PriorityQueue) DequeueJob(ctx context.Context) *Job {
	logger := this.logger.With("caller", "PriorityQueue.DequeueJob")
	data, err := this.cache.Dequeue(ctx, this.collectionName)
	if err != nil {
		logger.Error("failed to dequeue job", "err", err.Error())
		return nil
	}
	var job Job
	if err := json.Unmarshal(data, &job); err != nil {
		logger.Error("failed to unmarshal json in job", "err", err.Error())
		return nil
	}
	return &job
}

func (this *PriorityQueue) Size(ctx context.Context) int {
	r := this.cache.GetClient().LLen(ctx, this.collectionName).Val()
	return int(r)
}
