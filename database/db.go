package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("mysql", "shiki:1234@tcp(localhost:3306)/booklist")
	if err != nil {
		log.Fatal(err)
	}
}
