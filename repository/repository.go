package repository

import (
	"github.com/jmoiron/sqlx"
	"gomasters/models"
)

type User interface {
	GetUser(id int) (models.User, error)
	UpdateUser(models.User) error
	CreateUser(user models.User) (int, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		NewUserPostgres(db),
	}
}