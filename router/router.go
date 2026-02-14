package router

import (
	"bank/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, H *handler.Handler) {
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})

}
