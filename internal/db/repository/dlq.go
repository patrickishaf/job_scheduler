package repository

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickishaf/job_scheduler/config"
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
	return nil, fmt.Errorf("method not implemented")
}

func (this *DLQRepository) FindByJobID(ctx context.Context, jobID uuid.UUID) ([]model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *DLQRepository) FindAll(ctx context.Context, job *model.Job) ([]model.Job, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (this *DLQRepository) DeleteJob(ctx context.Context, jobID uuid.UUID) error {
	return fmt.Errorf("method not implemented")
}
