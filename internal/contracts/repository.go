package contracts

import "github.com/nicitapa/firstProgect/internal/models"

type RepositoryI interface {
	GetAllUsers() (users []models.User, err error)
	GetUsersByID(id int) (user models.User, err error)
	CreateUsers(users models.User) (err error)
	UpdateUsersByID(users models.User) (err error)
	DeleteUsersByID(id int) (err error)
}
