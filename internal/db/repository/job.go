package repository

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickishaf/job_scheduler/config"
	"github.com/patrickishaf/job_scheduler/internal/common"
	"github.com/patrickishaf/job_scheduler/internal/db/model"
)

type JobRepository struct {
	cfg      *config.AppConfig
	connPool *pgxpool.Pool
	logger   *slog.Logger
	mu       *sync.Mutex
}

func CreateJobRepository(cfg *config.AppConfig, pool *pgxpool.Pool, log *slog.Logger) *JobRepository {
	var mu sync.Mutex
	return &JobRepository{
		cfg:      cfg,
		connPool: pool,
		logger:   log,
		mu:       &mu,
	}
}

func (this *JobRepository) SaveOne(ctx context.Context, job *model.Job) (*model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *JobRepository) FindNextPendingScheduledJob(ctx context.Context, job *model.Job) ([]model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *JobRepository) FindNextPendingRecurringJob(ctx context.Context, job *model.Job) ([]model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *JobRepository) DeleteJob(ctx context.Context, jobID uuid.UUID) error {
	return fmt.Errorf("method not implemented")
}

func (this *JobRepository) MarkJobAsProcessing(ctx context.Context, jobID uuid.UUID, status common.ProcessingStatus, attemptCount int) (*model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *JobRepository) UpdateJobStatus(ctx context.Context, jobID uuid.UUID, status common.ProcessingStatus) (*model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}
