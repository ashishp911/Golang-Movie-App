package db

import (
	"database/sql"
	"fmt"
	"log"
	"go-crud-movies/models"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB{
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

func AddtoDB(db *sql.DB, movie  models.Movie) {
	fmt.Println("Id is ", movie.ID)
	fmt.Println("isnb is ", movie.Isbn)
	fmt.Println("Title is", movie.Title)
	fmt.Println("Director f name  is", movie.Director.FirstName)
	fmt.Println("Director l name  is", movie.Director.LastName)

	result, err := db.Exec(`INSERT INTO movies (id, isbn, title, director_firstname, director_lastname) VALUES (?, ?, ?, ?, ?)`, movie.ID, movie.Isbn, movie.Title, movie.Director.FirstName, movie.Director.LastName)
	userID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println(userID)
	}
}
