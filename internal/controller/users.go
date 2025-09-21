package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nicitapa/firstProgect/internal/errs"
	"github.com/nicitapa/firstProgect/internal/models"
	"net/http"
	"strconv"
)

// GetAllUsers
// @Summary Получение пользователей
// @Description Получение списка всех пользователей
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} CommonError
// @Router /Users [get]
func (ctrl *Controller) GetAllUsers(c *gin.Context) {
	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUsersByID
// @Summary Получить пользователя по ID
// @Description Получение пользователя по ID
// @Tags User
// @Produce json
// @Param id path int true "id Пользлвател"
// @Success 200 {object} models.User
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (ctrl *Controller) GetUsersByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUsersID)
		return
	}

	users, err := ctrl.service.GetUsersByID(id)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

type CreateUsersRequest struct {
	Name  string `json:"name" `
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// CreateUsers
// @Summary Добавление нового пользователя
// @Description Добавление нового пользователя
// @Tags Users
// @Consume json
// @Produce json
// @Param request_body body CreateUsersRequest true "информация о новом пользователе"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users [post]
func (ctrl *Controller) CreateUsers(c *gin.Context) {
	var users models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}
	if users.Name == "" || users.Email == "" || users.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	if err := ctrl.service.CreateUsers(users); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "Users created successfully"})
}

// UpdateUserByID
// @Summary Обновить пользователя по ID
// @Description Обновление пользователя по ID
// @Tags User
// @Consume json
// @Produce json
// @Param id path int true "id пользователя"
// @Param request_body body CreateUsersRequest true "информация о пользователе"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [put]
func (ctrl *Controller) UpdateUsersByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUsersID)
		return
	}

	var users models.User
	if err = c.ShouldBindJSON(&users); err != nil {
		ctrl.handleError(c, err)
		return
	}

	users.ID = id

	if err = ctrl.service.UpdateUsersByID(users); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Users updated successfully"})
}

// DeleteUsersByID
// @Summary Удалить пользователя по ID
// @Description Удаление пользователя по ID
// @Tags User
// @Produce json
// @Param id path int true "id пользователя"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [delete]
func (ctrl *Controller) DeleteUsersByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUsersID)
		return
	}

	if err = ctrl.service.DeleteUsersByID(id); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Users deleted successfully"})
}
