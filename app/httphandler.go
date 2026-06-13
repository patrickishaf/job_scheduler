package app

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/patrickishaf/job_scheduler/internal/common"
)

type httpHandler struct {
	logger *slog.Logger
	svc    *service
}

func CreateHTTPHandler(svc *service, logger *slog.Logger) *httpHandler {
	return &httpHandler{
		logger: logger,
		svc:    svc,
	}
}

func (this *httpHandler) CreateJob(c *gin.Context) {
	requestID := uuid.NewString()
	logger := this.logger.With("caller", "httpHandler.CreateJob", common.CTX_KEY_REQUEST_ID, requestID)

	var dto createJobDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		logger.Error("failed to create job", "err", err.Error())
		c.IndentedJSON(400, common.ErrorResponse(common.ErrBadRequest))
		return
	}
	sc, data, err := this.svc.CreateJob(context.WithValue(c.Request.Context(), common.CTX_KEY_REQUEST_ID, requestID), dto)
	if err != nil {
		logger.Error("failed to create job", "err", err.Error())
		c.IndentedJSON(sc, common.ErrorResponse(err.Error()))
		return
	}
	c.IndentedJSON(sc, common.SuccessResponse(data, "created jobs successfully"))
}

func (this *httpHandler) GetJobs(c *gin.Context) {
	var dto getJobsDTO
	if err := c.ShouldBindQuery(&dto); err != nil {
		c.IndentedJSON(400, common.ErrorResponse(common.ErrBadRequest))
		return
	}
	sc, data, err := this.svc.GetJobs(context.WithValue(c.Request.Context(), common.CTX_KEY_REQUEST_ID, uuid.NewString()), dto)
	if err != nil {
		c.IndentedJSON(sc, common.ErrorResponse(err.Error()))
		return
	}
	c.IndentedJSON(sc, common.SuccessResponse(data, "fetched jobs successfully"))
}

func (this *httpHandler) GetDeadLetterQueueJobs(c *gin.Context) {
	sc, data, err := this.svc.GetDeadLetterQueueJobs(context.WithValue(c.Request.Context(), common.CTX_KEY_REQUEST_ID, uuid.NewString()))
	if err != nil {
		c.IndentedJSON(sc, common.ErrorResponse(err.Error()))
		return
	}
	c.IndentedJSON(sc, common.SuccessResponse(data, "fetched DLQ jobs successfully"))
}

func (this *httpHandler) GetJobsSummary(c *gin.Context) {
	sc, data, err := this.svc.GetJobsSummary(context.WithValue(c.Request.Context(), common.CTX_KEY_REQUEST_ID, uuid.NewString()))
	if err != nil {
		c.IndentedJSON(sc, common.ErrorResponse(err.Error()))
		return
	}
	c.IndentedJSON(sc, common.SuccessResponse(data, "fetched jobs summary successfully"))
}

func (this *httpHandler) GetSingleJob(c *gin.Context) {
	requestID := uuid.NewString()
	logger := this.logger.With("caller", "httpHandler.GetSingleJob", common.CTX_KEY_REQUEST_ID, requestID)

	var params getSingleJobReqParams
	if err := c.ShouldBindUri(&params); err != nil {
		logger.Error("failed to get single job", "err", err.Error())
		c.IndentedJSON(400, common.ErrorResponse(common.ErrBadRequest))
		return
	}
	jobID, err := uuid.Parse(params.JobID)
	if err != nil {
		logger.Error("invalid job id. not a UUID", "job_id", params.JobID)
		c.IndentedJSON(400, common.ErrorResponse(common.ErrBadRequest))
		return
	}

	sc, data, err := this.svc.GetSingleJob(context.WithValue(c.Request.Context(), common.CTX_KEY_REQUEST_ID, requestID), jobID)
	if err != nil {
		c.IndentedJSON(sc, common.ErrorResponse(err.Error()))
		return
	}
	c.IndentedJSON(sc, common.SuccessResponse(data, "fetched single job successfully"))
}

func (this *httpHandler) RequeueJob(c *gin.Context) {
	requestID := uuid.NewString()
	logger := this.logger.With("caller", "httpHandler.RequeueJob", common.CTX_KEY_REQUEST_ID, requestID)

	var params requeueJobDTO
	if err := c.ShouldBindUri(&params); err != nil {
		logger.Error("failed to get single job", "err", err.Error())
		c.IndentedJSON(400, common.ErrorResponse(common.ErrBadRequest))
		return
	}

	jobID, err := uuid.Parse(params.JobID)
	if err != nil {
		logger.Error("invalid job id. not a UUID", "job_id", params.JobID)
		c.IndentedJSON(400, common.ErrorResponse(common.ErrBadRequest))
		return
	}

	sc, data, err := this.svc.RequeueJob(context.WithValue(c.Request.Context(), common.CTX_KEY_REQUEST_ID, uuid.NewString()), jobID)
	if err != nil {
		c.IndentedJSON(sc, common.ErrorResponse(err.Error()))
		return
	}
	c.IndentedJSON(sc, common.SuccessResponse(data, "requeued job successfully"))
}
