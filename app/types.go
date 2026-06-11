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

type createJobDTO struct {
	Type          JobType        `form:"type" json:"type" binding:"required"`
	Priority      int            `form:"priority" json:"priority" binding:"required,min=1,max=3"`
	Payload       map[string]any `form:"payload" json:"payload" binding:"required"`
	ScheduledTime time.Time      `form:"scheduled_time" json:"scheduled_time" binding:"omitempty"`
	Interval      *int           `form:"interval_seconds" json:"interval_seconds" binding:"omitempty"`
}

type getJobsDTO struct {
	Search string           `form:"search" binding:"omitempty"`
	Status ProcessingStatus `form:"status" binding:"omitempty"`
}

type getSingleJobReqParams struct {
	JobID string `uri:"id" binding:"required"`
}

type job struct {
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

type jobsSummaryDTO struct {
	Queued     int `json:"queued"`
	Processing int `json:"processing"`
	Cancelled  int `json:"cancelled"`
	Completed  int `json:"completed"`
	Failed     int `json:"failed"`
	Pending    int `json:"pending"`
	Scheduled  int `json:"scheduled"`
}
