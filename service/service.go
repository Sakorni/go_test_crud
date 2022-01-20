package service

import (
	"gomasters/models"
	"gomasters/repository"
)

type User interface {
	GetUser() (models.User, error)
	UpdateUser(models.User) error
	CreateUser(user models.User) (models.User, error)
}
type Service struct {
	User
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}

