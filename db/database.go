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

func UpdateInDB(db *sql.DB, movie models.Movie) {
	_, err := db.Exec(`UPDATE movies SET id = ? , isbn = ?, title = ?, director_firstname = ?, director_lastname = ? WHERE id = ?;`, movie.ID, movie.Isbn, movie.Title, movie.Director.FirstName, movie.Director.LastName, movie.ID) // check err
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")
}

func GetAllMovies(db *sql.DB, movies []models.Movie) []models.Movie {

	rows, err := db.Query(`SELECT id, isbn, title, director_firstname, director_lastname FROM movies`) // check err
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Movie
		err := rows.Scan(&m.ID, &m.Isbn, &m.Title, &m.Director.FirstName, &m.Director.LastName) // check err
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, m)
	}
	return movies
}

func GetAMovie(db *sql.DB, movie models.Movie) {
	// Query the database and scan the values into out variables. Don't forget to check for errors.
	query := `SELECT id, isbn, title, director_firstname, director_lastname FROM movies WHERE id = ?`
	err := db.QueryRow(query, 1).Scan(&movie.ID, &movie.Isbn, &movie.Title, &movie.Director.FirstName, &movie.Director.LastName)
	if err != nil {
		log.Fatal(err)
	}
}
