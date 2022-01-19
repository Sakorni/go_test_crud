package models

import "time"

type User struct {
	ID uint
	Firstname string
	Lastname string
	Email string
	Age uint
	Created time.Time
}
