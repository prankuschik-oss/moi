package contracts

import "github.com/nicitapa/firstProgect/internal/models"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceI interface {
	GetAllUsers() (users []models.Employees, err error)
	GetUsersByID(id int) (users models.Employees, err error)
	CreateUsers(users models.Employees) (err error)
	UpdateUsersByID(users models.Employees) (err error)
	DeleteUsersByID(id int) (err error)
}
