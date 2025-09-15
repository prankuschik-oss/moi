package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *Controller) RegisterEndpoints() {

	ctrl.router.GET("/ping", ctrl.Ping)

	ctrl.router.GET("/users", ctrl.GetAllUsers)
	ctrl.router.GET("/users/:id", ctrl.GetUsersByID)
	ctrl.router.POST("/users", ctrl.CreateUsers)
	ctrl.router.PUT("/users/:id", ctrl.UpdateUsersByID)
	ctrl.router.DELETE("/users/:id", ctrl.DeleteUsersByID)
}

func (ctrl *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func (ctrl *Controller) RunServer(address string) error {
	// Регистрируем роуты
	ctrl.RegisterEndpoints()

	// Запускаем http-сервер
	if err := ctrl.router.Run(address); err != nil {
		return err
	}

	return nil
}
