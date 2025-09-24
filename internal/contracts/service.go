package contracts

import "github.com/nicitapa/firstProgect/internal/models"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceI interface {
	GetAllEmployees() (users []models.Employees, err error)
	GetEmployeesByID(id int) (users models.Employees, err error)
	CreateEmployees(users models.Employees) (err error)
	UpdateEmployeesByID(users models.Employees) (err error)
	DeleteEmployeesByID(id int) (err error)
}
