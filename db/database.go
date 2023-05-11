package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-crud-movies/models"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/movies_db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is connected successfully")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
