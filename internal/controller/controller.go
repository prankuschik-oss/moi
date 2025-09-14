package controller

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
