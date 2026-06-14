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
	PROCESSING_STATUS_QUEUED     ProcessingStatus = "queued"
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

type jobsSummaryDTO struct {
	Queued     int `json:"queued"`
	Processing int `json:"processing"`
	Cancelled  int `json:"cancelled"`
	Completed  int `json:"completed"`
	Failed     int `json:"failed"`
	Pending    int `json:"pending"`
	Scheduled  int `json:"scheduled"`
}

type priorityQueueEntry struct {
	index int
	job   *Job
}

type requeueJobDTO struct {
	JobID string `uri:"id" binding:"required"`
}

type Scheduler interface {
	EnqueueDueJobs(jobs []*Job) map[uuid.UUID]bool
	Dequeue() (*Job, bool)
	Len() int
}
