package repository

import (
	"github.com/jmoiron/sqlx"
	"gomasters/models"
)

type User interface {
	GetUser() (models.User, error)
	UpdateUser(models.User) error
	CreateUser(user models.User) (models.User, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{}
}