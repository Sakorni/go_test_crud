package models

import "time"

type User struct {
	ID uint `json:"id"`
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Email string `json:"email"`
	Age uint `json:"age"`
	Created time.Time `json:"created_at"`
}
