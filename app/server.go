package app

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/job_scheduler/config"
)

type AppServer struct {
	cfg     *config.AppConfig
	handler *gin.Engine
}

func CreateAppServer(cfg *config.AppConfig) *AppServer {
	return &AppServer{
		cfg:     cfg,
		handler: gin.Default(),
	}
}

func (srv *AppServer) ConfigureRoutes(router *APIRouter) {
	srv.handler.Use(cors.Default())
	apiGroup := srv.handler.Group("/api")
	router.ConfigureRoutes(apiGroup)
}

func (srv *AppServer) GetAddr() string {
	return fmt.Sprintf(":%d", srv.cfg.Port)
}

func (srv *AppServer) Run() error {
	return srv.handler.Run(srv.GetAddr())
}
