package repository

import (
	"github.com/jmoiron/sqlx"
	"gomasters/models"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}



func (u *UserPostgres) GetUser(id int) (models.User, error) {
	panic("implement me")
}

func (u *UserPostgres) UpdateUser(user models.User) error {
	panic("implement me")
}

func (u *UserPostgres) CreateUser(user models.User) (int, error) {
	query := "INSERT INTO users(firstname, lastname, email, age) values(:firstname,:lastname,:email,:age) RETURNING id"
	rows,err := u.db.NamedQuery(query,
		user)
	if err != nil{
		return 0, err
	}
	var id int
	rows.Next()
	if err = rows.Scan(&id); err != nil{
		return 0, err
	}
	return id, nil

}

