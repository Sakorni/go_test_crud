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
	query := "SELECT * FROM users WHERE id = $1"
	row := u.db.QueryRowx(query, id)
	var res models.User
	err := row.StructScan(&res)
	return res, err

}

func (u *UserPostgres) UpdateUser(user models.User) error {
	query := "UPDATE users SET firstname=:firstname, lastname=:lastname, email=:email, age=:age WHERE id=:id"
	_, err := u.db.NamedExec(query, user)
	if err != nil{
		return err
	}
	return nil

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

