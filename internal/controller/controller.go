package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nicitapa/firstProgect/internal/service"
)

type Controller struct {
	router  *gin.Engine
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
	}
}
