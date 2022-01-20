package models

import "time"

type User struct {
	ID uint `json:"id" db:"id"`
	Firstname string `json:"first_name" db:"firstname"`
	Lastname string `json:"last_name" db:"lastname"`
	Email string `json:"email" db:"email"`
	Age uint `json:"age" db:"age"`
	Created time.Time `json:"created_at" db:"created_at"`
}
