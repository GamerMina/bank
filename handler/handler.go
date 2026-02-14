package handler

import (
	"bank/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine  *gin.Engine
	Service *service.Services
}

func NewHandler(engine *gin.Engine, services *service.Services) *Handler {
	return &Handler{
		Engine:  engine,
		Service: services,
	}
}
