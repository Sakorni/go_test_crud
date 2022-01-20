package models

import "time"

type User struct {
	ID        uint      `json:"id" db:"id"`
	Firstname string    `json:"first_name" db:"firstname" binding:"required" `
	Lastname  string    `json:"last_name" db:"lastname" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required"`
	Age       uint      `json:"age" db:"age" binding:"required"`
	Created   time.Time `json:"created_at" db:"created_at"`
}
