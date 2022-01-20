package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
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
	_, tableCheck := db.Query("Select * from users")
	if tableCheck != nil{
		err = runMigrationLocal(db)
	}
	return db, err
}

func runMigrationLocal(db *sqlx.DB) error{
	file, err := ioutil.ReadFile("./schema/000001_init.up.sql")
	if err != nil{
		return err
	}
	_, err = db.Exec(string(file))
	return err

}


