package app

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/patrickishaf/job_scheduler/internal/infra"
	"github.com/patrickishaf/job_scheduler/internal/net"
)

type service struct {
	cache           *infra.RedisCache
	logger          *slog.Logger
	socketConnStore net.SocketConnectionStore
}

func CreateService(cache *infra.RedisCache, socketConnStore *net.SocketConnectionStore, logger *slog.Logger) *service {
	l := logger.With("service", "service", "caller", "service")
	return &service{
		cache:           cache,
		logger:          l,
		socketConnStore: *socketConnStore,
	}
}

func (this *service) CreateJob(cto CreateJobDTO) {}

func (this *service) GetAllJobs(dto GetJobsDTO) {}

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
}
