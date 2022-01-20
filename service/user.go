package service

import (
	"gomasters/models"
	"gomasters/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService( rep repository.Repository) *UserService{
	return &UserService{repo: rep}
}

func (u *UserService) GetUser(id int) (models.User, error) {
	return u.repo.GetUser(id)
}

func (u *UserService) UpdateUser(user models.User) error {
	return u.repo.UpdateUser(user)
}

func (u *UserService) CreateUser(user models.User) (int, error) {
	return u.repo.CreateUser(user)
}

