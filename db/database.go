package db

import (
	"database/sql"
	"fmt"
	"go-crud-movies/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
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

func AddtoDB(db *sql.DB, movie models.Movie) {
	result, err := db.Exec(`INSERT INTO movies (id, isbn, title, director_firstname, director_lastname) VALUES (?, ?, ?, ?, ?)`, movie.ID, movie.Isbn, movie.Title, movie.Director.FirstName, movie.Director.LastName)
	userID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(userID)
	}
	fmt.Println("Record added to Datbase successfully.")
}

func DeleteFromDB(db *sql.DB, id string) {
	_, err := db.Exec(`DELETE FROM movies WHERE id = ?`, id) // check err
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Record deleted from a Datbase successfully.")
}
