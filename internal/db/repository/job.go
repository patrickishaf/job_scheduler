package repository

import (
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickishaf/job_scheduler/config"
	"github.com/patrickishaf/job_scheduler/internal/db/model"
)

type JobRepository struct {
	cfg      *config.AppConfig
	connPool *pgxpool.Pool
	logger   *slog.Logger
}

func CreateJobRepository(cfg *config.AppConfig, pool *pgxpool.Pool, log *slog.Logger) *JobRepository {
	return &JobRepository{
		cfg:      cfg,
		connPool: pool,
		logger:   log,
	}
}

func (this *JobRepository) SaveOne(job *model.Job) (*model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *JobRepository) FindPendingJobs(job *model.Job) ([]model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}
