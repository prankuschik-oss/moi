package controller

import "net/http"

func (ctrl *Controller) RegisterEndpoints() {

	ctrl.router.GET("/ping", ctrl.Ping)

	ctrl.router.GET("/products", ctrl.GetAllProducts)
	ctrl.router.GET("/products/:id", ctrl.GetProductByID)
	ctrl.router.POST("/products", ctrl.CreateProduct)
	ctrl.router.PUT("/products/:id", ctrl.UpdateProductByID)
	ctrl.router.DELETE("/products/:id", ctrl.DeleteProductByID)
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
