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

func (this *JobRepository) DeleteJob(ctx context.Context, jobID uuid.UUID) error {
	return fmt.Errorf("method not implemented")
}

func (this *JobRepository) FindNextPendingScheduledJob(ctx context.Context, job *model.Job) ([]model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *JobRepository) FindNextPendingRecurringJob(ctx context.Context, job *model.Job) ([]model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *JobRepository) GetAllJobs(ctx context.Context) ([]model.Job, error) {
	logger := this.logger.With("caller", "JobRepository.GetAllJobs")
	query := `SELECT id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload FROM jobs`
	rows, err := this.connPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var jobs []model.Job
	for rows.Next() {
		var job model.Job
		err = rows.Scan(
			&job.ID,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.AttemptCount,
			&job.Error,
			&job.Interval,
			&job.LastAttemptAt,
			&job.Priority,
			&job.RetryCount,
			&job.ScheduledTime,
			&job.Status,
			&job.Type,
			&job.Payload,
		)
		if err != nil {
			logger.Error("failed to scan row into job", "err", err.Error())
			continue
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (this *JobRepository) MarkJobAsProcessing(ctx context.Context, jobID uuid.UUID, status common.ProcessingStatus, attemptCount int) (*model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *JobRepository) SaveDefaultJob(ctx context.Context, job *model.Job) (*model.Job, error) {
	logger := this.logger.With("caller", "JobRepository.SaveDefaultJob")
	query := `
	INSERT INTO jobs(
	interval,
	payload,
	priority,
	scheduled_time,
	type
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, status
	`
	err := this.connPool.QueryRow(
		ctx,
		query,
		&job.Interval,
		&job.Payload,
		&job.Priority,
		&job.ScheduledTime,
		&job.Type,
	).Scan(&job.ID, &job.Status)
	if err != nil {
		logger.Error("failed to save default job", "err", err.Error())
		return nil, err
	}
	return job, nil
}

func (this *JobRepository) UpdateJobStatus(ctx context.Context, jobID uuid.UUID, status common.ProcessingStatus) (*model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}
