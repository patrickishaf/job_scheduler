package model

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	AttemptCount  int32
	Error         *string
	Interval      *int
	LastAttemptAt *time.Time
	Payload       map[string]any
	Priority      int
	RetryCount    int
	ScheduledTime time.Time
	Status        string
	Type          string
}
