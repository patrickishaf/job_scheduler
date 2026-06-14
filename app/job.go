package app

import (
	"time"

	"github.com/google/uuid"
	"github.com/patrickishaf/job_scheduler/internal/db/model"
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

func CreateJobFromModel(j *model.Job) *Job {
	job := Job{
		ID:            j.ID,
		CreatedAt:     j.CreatedAt,
		UpdatedAt:     j.UpdatedAt,
		AttemptCount:  j.AttemptCount,
		Error:         j.Error,
		LastAttemptAt: j.LastAttemptAt,
		Payload:       j.Payload,
		Priority:      j.Priority,
		RetryCount:    j.RetryCount,
		ScheduledTime: j.ScheduledTime,
		Status:        j.Status,
		Type:          j.Type,
	}
	if j.Interval != nil {
		job.Interval = j.Interval
	}
	return &job
}

func (j *Job) effectivePriority() int {
	age := time.Since(j.CreatedAt)
	boost := int(age / starvationThreshold) // 1 boost per threshold exceeded
	p := j.Priority - boost                 // lower number = higher priority
	if p < 1 {
		p = 1 // cap at highest priority
	}
	return p
}

func (j *Job) ExhaustedRetries() bool {
	return int(j.AttemptCount) >= j.RetryCount
}

func (j *Job) IsDue() bool {
	return !time.Now().After(j.ScheduledTime)
}

func (j *Job) IsRecurring() bool {
	return j.Interval != nil && *j.Interval > 0
}

func (j *Job) NextScheduledTime() time.Time {
	return time.Now().Add(time.Duration(*j.Interval) * time.Second)
}

func (j *Job) HasLessPriorityThan(b *Job) bool {
	pa, pb := j.effectivePriority(), b.effectivePriority()
	if pa != pb {
		return pa < pb
	}
	if !j.ScheduledTime.Equal(b.ScheduledTime) {
		return j.ScheduledTime.Before(b.ScheduledTime)
	}
	return j.CreatedAt.Before(b.CreatedAt)
}
