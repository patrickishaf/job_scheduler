package app

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID            uuid.UUID      `json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	AttemptCount  int32          `json:"-"`
	Error         *string        `json:"error"`
	Interval      *int           `json:"interval"`
	LastAttemptAt *time.Time     `json:"last_attempt_at"`
	Payload       map[string]any `json:"payload"`
	Priority      int            `json:"priority"`
	RetryCount    int            `json:"retry_count"`
	ScheduledTime time.Time      `json:"scheduled_time"`
	Status        string         `json:"status"`
	Type          string         `json:"type"`
}

func (j *Job) IsRecurring() bool {
	return *j.Interval > 0
}

func (j *Job) IsDue() bool {
	return !time.Now().Before(j.ScheduledTime)
}

func (j *Job) NextRun() time.Time {
	return time.Now().Add(time.Duration(*j.Interval))
}

func (j *Job) HasLessPriorityThan(b *Job) bool {
	if j.Priority != b.Priority {
		return j.Priority < b.Priority
	}
	if !j.ScheduledTime.Equal(b.ScheduledTime) {
		return j.ScheduledTime.Before(b.ScheduledTime)
	}
	return j.CreatedAt.Before(b.CreatedAt)
}
