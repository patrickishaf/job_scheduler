package app

import (
	"github.com/gin-gonic/gin"
)

type AppServer struct {
	Port    int
	Handler *gin.Engine
}

func CreateAppServer() *AppServer {
	return &AppServer{}
}

func (srv *AppServer) ConfigureRoutes(group *gin.RouterGroup) {}

func (srv *AppServer) GetAddr() string {
	return ""
}

func (srv *AppServer) Run() {}
