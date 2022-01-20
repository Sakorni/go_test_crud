package database

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB



func init(){
	var err error
	databaseurl := "user=postgres password=admin dbname=gomasters sslmode=disable"
	db, err = sql.Open("pgx", databaseurl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}


