package service

import "github.com/nicitapa/firstProgect/internal/repository"

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}
