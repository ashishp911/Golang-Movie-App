package my_movies

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-crud-movies/db"
	"go-crud-movies/models"
	"net/http"
)

var movies []models.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Connect to Databse
	my_db := db.Connect()
	movies = db.GetAllMovies(my_db, movies)
	// return the list of all the movies
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params to get the ID from the url
	params := mux.Vars(r)

	for _, item := range movies {
		if params["id"] == item.ID {
			json.NewEncoder(w).Encode(item)
			my_db := db.Connect()
			db.GetAMovie(my_db, item)
			return
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params to get the ID from the url
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// append all the data except the data with ID == params["ID"]
			movies = append(movies[:index], movies[index+1:]...)
			// Deleting from Database
			fmt.Println("Deleting from a database")
			// Connect to Databse
			my_db := db.Connect()
			db.DeleteFromDB(my_db, item.ID)
			break
		}
	}
	// after deleting the movie, show all the movies to the frontend
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

	// Connect to Databse
	my_db := db.Connect()
	// Adding the record to DB
	db.AddtoDB(my_db, movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// delete a movie and add a new movie
	w.Header().Set("Content-Type", "application/json")
	// params to get the ID from the url
	params := mux.Vars(r)

	for index, items := range movies {
		if items.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]

			// Connect to Databse
			my_db := db.Connect()
			db.UpdateInDB(my_db, movie)

			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
