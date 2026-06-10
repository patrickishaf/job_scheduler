package app

import (
	"time"

	"github.com/google/uuid"
)

type JobType string
type ProcessingStatus string

const (
	JOB_TYPE_EMAIL_SEND          JobType          = "email.send"
	PROCESSING_STATUS_CANCELLED  ProcessingStatus = "cancelled"
	PROCESSING_STATUS_COMPLETED  ProcessingStatus = "completed"
	PROCESSING_STATUS_FAILED     ProcessingStatus = "failed"
	PROCESSING_STATUS_PENDING    ProcessingStatus = "pending"
	PROCESSING_STATUS_PROCESSING ProcessingStatus = "processing"
)

type CreateJobDTO struct {
	Type          JobType        `form:"type" json:"type" binding:"required"`
	Priority      int            `form:"priority" json:"priority" binding:"required,min=1,max=3"`
	Payload       map[string]any `form:"payload" json:"payload" binding:"required"`
	ScheduledTime time.Time      `form:"scheduled_time" json:"scheduled_time" binding:"omitempty"`
	Interval      *int           `form:"interval_seconds" json:"interval_seconds" binding:"omitempty"`
}

type GetJobsDTO struct {
	Search string           `json:"search" binding:"omitempty"`
	Status ProcessingStatus `json:"status" binding:"omitempty"`
}

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
