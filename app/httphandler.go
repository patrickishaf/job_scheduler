package app

import (
	"log/slog"

	"github.com/gin-gonic/gin"
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
	logger := this.logger.With("caller", "httpHandler.CreateJob")
	var dto CreateJobDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		this.logger.Error("failed to create job", "err", err.Error())
		c.IndentedJSON(400, common.ErrorResponse(common.ErrBadRequest))
		return
	}
	sc, data, err := this.svc.CreateJob(c.Request.Context(), dto)
	if err != nil {
		logger.Error("failed to create job", "err", err.Error())
		c.IndentedJSON(sc, common.ErrorResponse(err.Error()))
		return
	}
	c.IndentedJSON(sc, common.SuccessResponse(data, "fetched jobs successfully"))
}

func (this *httpHandler) GetJobs(c *gin.Context) {
	sc, data, err := this.svc.GetAllJobs(c.Request.Context(), GetJobsDTO{})
	if err != nil {
		c.IndentedJSON(sc, common.ErrorResponse(err.Error()))
		return
	}
	c.IndentedJSON(sc, common.SuccessResponse(data, "fetched jobs successfully"))
}
