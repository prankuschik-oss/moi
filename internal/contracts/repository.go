package contracts

import "github.com/nicitapa/firstProgect/internal/models"

type RepositoryI interface {
	GetAllEmployees() (employees []models.Employees, err error)
	GetEmployeesByID(id int) (employees models.Employees, err error)
	CreateEmployees(employees models.Employees) (err error)
	UpdateEmployeesByID(employees models.Employees) (err error)
	DeleteEmployeesByID(id int) (err error)
}
