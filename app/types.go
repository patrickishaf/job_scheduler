package app

type JobProcessingStatus string

const (
	PROCESSING_STATUS_CANCELLED  JobProcessingStatus = "cancelled"
	PROCESSING_STATUS_COMPLETED  JobProcessingStatus = "completed"
	PROCESSING_STATUS_FAILED     JobProcessingStatus = "failed"
	PROCESSING_STATUS_PENDING    JobProcessingStatus = "pending"
	PROCESSING_STATUS_PROCESSING JobProcessingStatus = "processing"
)
