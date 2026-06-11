package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickishaf/job_scheduler/config"
	"github.com/patrickishaf/job_scheduler/internal/common"
	"github.com/patrickishaf/job_scheduler/internal/db/model"
)

type DLQRepository struct {
	cfg      *config.AppConfig
	connPool *pgxpool.Pool
	logger   *slog.Logger
	mu       *sync.Mutex
}

func CreateDLQRepository(cfg *config.AppConfig, pool *pgxpool.Pool, log *slog.Logger) *DLQRepository {
	var mu sync.Mutex
	return &DLQRepository{
		cfg:      cfg,
		connPool: pool,
		logger:   log,
		mu:       &mu,
	}
}

func (this *DLQRepository) SaveOne(ctx context.Context, job *model.Job) (*model.Job, error) {
	logger := this.logger.With("caller", "DLQRepository.SaveOne")
	query := `
	INSERT INTO dead_letter_queue(
	id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	tag, err := this.connPool.Exec(
		ctx,
		query,
		&job.ID,
		&job.CreatedAt,
	)
	if err != nil {
		logger.Error("failed to save job to dead letter queue", "err", err.Error())
		return nil, fmt.Errorf(common.ErrDBOperationFailed)
	}
	if tag.RowsAffected() == 0 {
		logger.Error("failed to save job to dead letter queue. rows affected is zero")
		return nil, errors.New(common.ErrDBOperationFailed)
	}
	return job, nil
}

func (this *DLQRepository) FindByJobID(ctx context.Context, jobID uuid.UUID) (*model.Job, error) {
	logger := this.logger.With("caller", "DLQRepository.FindByJobID")
	var job model.Job
	query := `SELECT id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload FROM dead_letter_queue WHERE id = $1`
	err := this.connPool.QueryRow(ctx, query, jobID).Scan(
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
		logger.Error("failed to find job by id", "err", err.Error())
		return nil, fmt.Errorf(common.ErrDBResourceNotFound)
	}
	return &job, nil
}

func (this *DLQRepository) FindAll(ctx context.Context) ([]model.Job, error) {
	logger := this.logger.With("caller", "DLQRepository.FindAll")
	query := `SELECT id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload FROM dead_letter_queue`
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

func (this *DLQRepository) DeleteJob(ctx context.Context, jobID uuid.UUID) error {
	logger := this.logger.With("caller", "DLQRepository.DeleteJob")
	tag, err := this.connPool.Exec(ctx, `DELETE FROM dead_letter_queue WHERE id = $1`, jobID)
	if err != nil {
		logger.Error("failed to delete dead letter queue entry", "err", err.Error())
		return err
	}
	if tag.RowsAffected() == 0 {
		logger.Error("no job found with id", "job_id", jobID)
		return errors.New(common.ErrDBResourceNotFound)
	}
	return nil
}
