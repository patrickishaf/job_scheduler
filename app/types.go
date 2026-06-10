package app

import "time"

type JobProcessingStatus string

const (
	PROCESSING_STATUS_CANCELLED  JobProcessingStatus = "cancelled"
	PROCESSING_STATUS_COMPLETED  JobProcessingStatus = "completed"
	PROCESSING_STATUS_FAILED     JobProcessingStatus = "failed"
	PROCESSING_STATUS_PENDING    JobProcessingStatus = "pending"
	PROCESSING_STATUS_PROCESSING JobProcessingStatus = "processing"
)

type CreateJobDTO struct {
	Type           string         `json:"type" binding:"required"`
	Priority       int            `json:"priority" binding:"required,min=1,max=3"`
	Payload        map[string]any `json:"payload" binding:"required"`
	ScheduledTime  time.Time      `json:"scheduled_time" binding:"required"`
	IntervaSeconds int            `json:"interval_seconds" binding:"omitempty"`
}

type GetJobsDTO struct {
	Search string              `json:"search" binding:"omitempty"`
	Status JobProcessingStatus `json:"status" binding:"omitempty"`
}
