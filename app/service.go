package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/patrickishaf/job_scheduler/internal/common"
	"github.com/patrickishaf/job_scheduler/internal/db/model"
	"github.com/patrickishaf/job_scheduler/internal/db/repository"
	"github.com/patrickishaf/job_scheduler/internal/infra"
	"github.com/patrickishaf/job_scheduler/internal/net"
)

type service struct {
	cache           *infra.RedisCache
	dlqRepo         *repository.DLQRepository
	jobRepo         *repository.JobRepository
	logger          *slog.Logger
	socketConnStore *net.SocketConnectionStore
}

func CreateService(
	cache *infra.RedisCache,
	dlqRepo *repository.DLQRepository,
	jobRepo *repository.JobRepository,
	socketConnStore *net.SocketConnectionStore,
	logger *slog.Logger,
) *service {
	return &service{
		cache:           cache,
		dlqRepo:         dlqRepo,
		jobRepo:         jobRepo,
		logger:          logger,
		socketConnStore: socketConnStore,
	}
}

func (this *service) convertJobToModel(job *Job) *model.Job {
	return &model.Job{
		ID:            job.ID,
		CreatedAt:     job.CreatedAt,
		UpdatedAt:     job.UpdatedAt,
		AttemptCount:  job.AttemptCount,
		Error:         job.Error,
		Interval:      job.Interval,
		LastAttemptAt: job.LastAttemptAt,
		Payload:       job.Payload,
		Priority:      job.Priority,
		RetryCount:    job.RetryCount,
		ScheduledTime: job.ScheduledTime,
		Status:        job.Status,
		Type:          job.Type,
	}
}

func (this *service) convertModelToJob(j *model.Job) *Job {
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

func (this *service) CreateJob(ctx context.Context, dto createJobDTO) (int, *Job, error) {
	jobData := &model.Job{
		Payload:       dto.Payload,
		Priority:      dto.Priority,
		ScheduledTime: dto.ScheduledTime,
		Type:          string(dto.Type),
	}
	if dto.Interval != nil {
		jobData.Interval = dto.Interval
	}
	jobModel, err := this.jobRepo.SaveDefaultJob(ctx, jobData)
	if err != nil {
		this.logger.Error("failed to save job", "err", err.Error(), "caller", "service.CreateJob")
		return 500, nil, err
	}
	job := this.convertModelToJob(jobModel)
	this.socketConnStore.BroadcastMessage(&net.SocketMessage{
		Event:   net.SOCKET_EVENT_JOB_CREATED,
		Payload: job,
	})
	return 201, job, nil
}

func (this *service) GetJobs(ctx context.Context, dto getJobsDTO) (int, []Job, error) {
	jobModels, err := this.jobRepo.GetAllJobs(ctx)
	if err != nil {
		this.logger.Error("failed to get all jobs", "err", err.Error(), "caller", "service.GetJobs")
		return 500, nil, err
	}
	jobs := make([]Job, 0)
	for _, j := range jobModels {
		job := this.convertModelToJob(&j)
		jobs = append(jobs, *job)
	}
	return 200, jobs, nil
}

func (this *service) GetDeadLetterQueueJobs(ctx context.Context) (int, []Job, error) {
	logger := this.logger.With("caller", "service.GetDeadLetterQueueJobs", "request_id", ctx.Value(common.CTX_KEY_REQUEST_ID))

	jobModels, err := this.dlqRepo.FindAll(ctx)
	if err != nil {
		logger.Error("failed to get all dead letter queue jobs", "err", err.Error())
		return 500, nil, err
	}
	if len(jobModels) == 0 {
		return 200, make([]Job, 0), nil
	}
	jobs := make([]Job, 0)
	for _, j := range jobModels {
		job := this.convertModelToJob(&j)
		jobs = append(jobs, *job)
	}
	return 200, jobs, nil
}

func (this *service) GetJobsSummary(ctx context.Context) (int, *jobsSummaryDTO, error) {
	logger := this.logger.With("caller", "service.GetJobs", "request_id", ctx.Value(common.CTX_KEY_REQUEST_ID))

	var summary jobsSummaryDTO
	jobCounts, err := this.jobRepo.CountJobsByStatus(ctx)
	if err != nil {
		logger.Error("failed to GetJobs", "err", err.Error())
		return 500, nil, err
	}
	for _, v := range jobCounts {
		if v.Status == string(PROCESSING_STATUS_CANCELLED) {
			summary.Cancelled = v.Count
		} else if v.Status == string(PROCESSING_STATUS_COMPLETED) {
			summary.Completed = v.Count
		} else if v.Status == string(PROCESSING_STATUS_FAILED) {
			summary.Failed = v.Count
		} else if v.Status == string(PROCESSING_STATUS_PENDING) {
			summary.Pending = v.Count
		} else if v.Status == string(PROCESSING_STATUS_PROCESSING) {
			summary.Processing = v.Count
		}
	}
	// TODO: Fetch queued job count but checking the size of the queue
	return 200, &summary, nil
}

func (this *service) GetSingleJob(ctx context.Context, jobID uuid.UUID) (int, *Job, error) {
	logger := this.logger.With("caller", "service.GetSingleJob", "request_id", ctx.Value(common.CTX_KEY_REQUEST_ID))

	jobFromDB, err := this.jobRepo.GetJobByID(ctx, jobID)
	if err != nil {
		logger.Error("failed to get single job", "err", err.Error())
		return 500, nil, err
	}

	job := this.convertModelToJob(jobFromDB)
	return 500, job, nil
}

func (this *service) RequeueJob(ctx context.Context, jobID uuid.UUID) (int, any, error) {
	/**
	 * steps
		* ensure the job exists in the dead letter queue
		* add the job to the db for the worker to pick it up
		* delete the job from the dead letter queue
		* return a response
	*/
	return 500, nil, fmt.Errorf(common.ErrMethodNotImplemented)
}

func (this *service) StoreSocketConn(conn *websocket.Conn) {
	/**
	 * steps
		* write the socket connection to the socket connection store
	*/
	this.socketConnStore.Add(conn)
	net.SendSocketMessage(conn, &net.SocketMessage{
		Event: net.SOCKET_EVENT_ACK,
	})
}
