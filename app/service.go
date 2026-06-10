package app

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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
	socketConnStore net.SocketConnectionStore
}

func CreateService(
	cache *infra.RedisCache,
	dlqRepo *repository.DLQRepository,
	jobRepo *repository.JobRepository,
	socketConnStore *net.SocketConnectionStore,
	logger *slog.Logger,
) *service {
	l := logger.With("service", "service", "caller", "service")
	return &service{
		cache:           cache,
		dlqRepo:         dlqRepo,
		jobRepo:         jobRepo,
		logger:          l,
		socketConnStore: *socketConnStore,
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

func (this *service) CreateJob(ctx context.Context, dto CreateJobDTO) (int, *Job, error) {
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

func (this *service) GetAllJobs(ctx context.Context, dto GetJobsDTO) (int, []Job, error) {
	jobModels, err := this.jobRepo.GetAllJobs(ctx)
	if err != nil {
		this.logger.Error("failed to get all jobs", "err", err.Error(), "caller", "service.CreateJob")
		return 500, nil, err
	}
	var jobs []Job
	for _, j := range jobModels {
		job := this.convertModelToJob(&j)
		jobs = append(jobs, *job)
	}
	return 200, jobs, nil
}

func (this *service) GetDeadLetterQueueJobs() {}

func (this *service) GetJobsSummary() {}

func (this *service) GetSingleJob(jobID uuid.UUID) {}

func (this *service) RequeueJob(jobID uuid.UUID) {
	/**
	 * steps
		* ensure the job exists in the dead letter queue
		* add the job to the db for the worker to pick it up
		* delete the job from the dead letter queue
		* return a response
	*/
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
