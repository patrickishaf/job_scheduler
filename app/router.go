package app

import "github.com/gin-gonic/gin"

type router interface {
	ConfigureRoutes(router *gin.RouterGroup)
}
