package app

import "github.com/gin-gonic/gin"

type APIRouter struct {
	httpHandler   *httpHandler
	socketHandler *socketHandler
}

func InitRouter(httpHandler *httpHandler, socketHandler *socketHandler) *APIRouter {
	return &APIRouter{
		httpHandler:   httpHandler,
		socketHandler: socketHandler,
	}
}

func (this *APIRouter) ConfigureRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/jobs", this.httpHandler.GetJobs)
	routerGroup.POST("/jobs", this.httpHandler.CreateJob)
	routerGroup.GET("/jobs/listen", this.socketHandler.handleConnection)
	routerGroup.GET("/jobs/summary", this.httpHandler.GetJobsSummary)
	routerGroup.GET("/jobs/:id/", this.httpHandler.GetSingleJob)
	routerGroup.PATCH("/jobs/:id/requeue", this.httpHandler.RequeueJob)
	routerGroup.GET("/dlq", this.httpHandler.GetDeadLetterQueueJobs)
}
