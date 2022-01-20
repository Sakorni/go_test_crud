package service

import (
	"gomasters/models"
	"gomasters/repository"
)

type User interface {
	GetUser(id int) (models.User, error)
	UpdateUser(models.User) error
	CreateUser(user models.User) (int, error)
}
type Service struct {
	User
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		NewUserService(*repository),
	}
}

