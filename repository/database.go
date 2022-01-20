package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)


type Config struct{
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string

}

func InitDB(cfg Config) (*sqlx.DB, error){
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sqlx.Connect("postgres", config)
	if err != nil{
		return nil, err
	}
	return db, err
}



