package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nicitapa/firstProgect/internal/contracts"
	"github.com/nicitapa/firstProgect/internal/errs"
	"net/http"
)

type Controller struct {
	router  *gin.Engine
	service contracts.ServiceI
}

func NewController(service contracts.ServiceI) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
	}
}

func (ctrl *Controller) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUsersNotfound) || errors.Is(err, errs.ErrNotfound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUsersID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue) || errors.Is(err, errs.ErrInvalidUserstName):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
