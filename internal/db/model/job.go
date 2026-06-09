package model

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID              uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AttemptCount    int32
	Error           string
	IntervalMinutes int
	LastAttemptAt   time.Time
	Priority        int32
	RetryCount      int32
	ScheduledTime   time.Time
	Status          string
	Type            string
}
