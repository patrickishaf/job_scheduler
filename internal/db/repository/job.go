package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickishaf/job_scheduler/config"
	"github.com/patrickishaf/job_scheduler/internal/common"
	"github.com/patrickishaf/job_scheduler/internal/db"
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

func (this *JobRepository) CountJobsByStatus(ctx context.Context) ([]db.JobStatusCount, error) {
	logger := this.logger.With("caller", "DLQRepository.CountJobsByStatus", "request_id", ctx.Value(common.CTX_KEY_REQUEST_ID))
	query := `SELECT status, COUNT(*) FROM jobs GROUP BY status ORDER BY status`
	rows, err := this.connPool.Query(ctx, query)
	if err != nil {
		logger.Error("failed to count jobs by status", "err", err.Error())
		return nil, errors.New(common.ErrDBOperationFailed)
	}
	defer rows.Close()
	var entries []db.JobStatusCount
	for rows.Next() {
		var entry db.JobStatusCount
		if err = rows.Scan(&entry.Status, &entry.Count); err != nil {
			logger.Error("failed to scan job status count into row", "err", err.Error())
			return nil, errors.New(common.ErrDBOperationFailed)
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (this *JobRepository) DeleteJob(ctx context.Context, jobID uuid.UUID) error {
	logger := this.logger.With("caller", "DLQRepository.DeleteJob")
	tag, err := this.connPool.Exec(ctx, `DELETE FROM jobs WHERE id = $1`, jobID)
	if err != nil {
		logger.Error("failed to delete job", "err", err.Error())
		return err
	}
	if tag.RowsAffected() == 0 {
		logger.Error("no job found with id", "job_id", jobID)
		return errors.New(common.ErrDBResourceNotFound)
	}
	return nil
}

func (this *JobRepository) GetAllJobs(ctx context.Context) ([]model.Job, error) {
	logger := this.logger.With("caller", "JobRepository.GetAllJobs")
	query := `SELECT id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload FROM jobs`
	rows, err := this.connPool.Query(ctx, query)
	if err != nil {
		return nil, errors.New(common.ErrDBOperationFailed)
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

func (this *JobRepository) GetAllJobsByStatus(ctx context.Context, status common.ProcessingStatus) ([]model.Job, error) {
	logger := this.logger.With("caller", "JobRepository.GetAllJobs")
	query := `SELECT id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload FROM jobs WHERE status = $1`
	rows, err := this.connPool.Query(ctx, query, status)
	if err != nil {
		return nil, errors.New(common.ErrDBOperationFailed)
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

func (this *JobRepository) GetAllJobsLikeSearch(ctx context.Context, search string) ([]model.Job, error) {
	logger := this.logger.With("caller", "JobRepository.GetAllJobs")
	query := `SELECT id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload FROM jobs WHERE type LIKE %$1%`
	rows, err := this.connPool.Query(ctx, query, search)
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

func (this *JobRepository) GetAllJobsByStatusAndSearch(ctx context.Context, status string, search string) ([]model.Job, error) {
	logger := this.logger.With("caller", "JobRepository.GetAllJobs")
	query := `SELECT id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload FROM jobs WHERE type LIKE %$1% AND status = $2`
	rows, err := this.connPool.Query(ctx, query, search)
	if err != nil {
		return nil, errors.New(common.ErrDBOperationFailed)
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

func (this *JobRepository) GetJobByID(ctx context.Context, jobID uuid.UUID) (*model.Job, error) {
	logger := this.logger.With("caller", "DLQRepository.GetJobByID", "request_id", ctx.Value(common.CTX_KEY_REQUEST_ID))
	var job model.Job
	query := `SELECT id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload FROM jobs WHERE id = $1`
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
		logger.Error("failed to find job by id", "err", err.Error(), "job_id", jobID)
		return nil, errors.New(common.ErrDBResourceNotFound)
	}
	return &job, nil
}

func (this *JobRepository) GetJobByIDInTx(ctx context.Context, jobID uuid.UUID, tx *pgx.Tx) (*model.Job, error) {
	logger := this.logger.With("caller", "DLQRepository.GetJobByIDInTx", "request_id", ctx.Value(common.CTX_KEY_REQUEST_ID))
	var job model.Job
	query := `SELECT id, created_at, updated_at, attempt_count, error, interval, last_attempt_at, priority, retry_count, scheduled_time, status, type, payload FROM jobs WHERE id = $1`
	err := (*tx).QueryRow(ctx, query, jobID).Scan(
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
		logger.Error("failed to find job by id", "err", err.Error(), "job_id", jobID)
		return nil, errors.New(common.ErrDBResourceNotFound)
	}
	return &job, nil
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
		return nil, errors.New(common.ErrDBOperationFailed)
	}
	return job, nil
}

func (this *JobRepository) UpdateJobStatus(ctx context.Context, jobID uuid.UUID, status common.ProcessingStatus) (*model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}
