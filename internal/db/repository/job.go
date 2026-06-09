package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickishaf/job_scheduler/config"
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
