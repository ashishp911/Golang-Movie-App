package my_movies

import (
	"encoding/json"
	"fmt"
	"go-crud-movies/db"
	"go-crud-movies/models"
	"net/http"

	"github.com/gorilla/mux"
)


var movies []models.Movie
func All_bizz_logic() {
	// dummy data 
	movies = append(movies, models.Movie{
		ID:"1",
		Isbn: "43811ab",
		Title: "Titanic",
		Director: &models.Director{
			FirstName: "John",
			LastName: "Mayers",
		},
	})
	movies = append(movies, models.Movie{
		ID:"2",
		Isbn: "43699yz",
		Title: "Shawshank Redemption",
		Director: &models.Director{
			FirstName: "David",
			LastName: "Guttenberg",
		},
	})
	movies = append(movies, models.Movie{
		ID:"3",
		Isbn: "43322fy",
		Title: "Forrest Gump",
		Director: &models.Director{
			FirstName: "Michael",
			LastName: "Scott",
		},
	})
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// return the list of all the movies
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// params to get the ID from the url
	params := mux.Vars(r)	
	
	for _, item := range movies{
		if params["id"] == item.ID{
			json.NewEncoder(w).Encode(item)
			return
		}
	} 
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params to get the ID from the url
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			// append all the data except the data with ID == params["ID"]
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// after deleting the movie, show all the movies to the frontend
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	
	// Connect to Databse
	my_db := db.Connect()
	fmt.Println(my_db)
	// Adding the record to DB
	db.AddtoDB(my_db, movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request){
	// delete a movie and add a new movie
	w.Header().Set("Content-Type", "application/json")
	// params to get the ID from the url
	params := mux.Vars(r)	

	for index, items := range movies{
		if items.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie models.Movie 
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)	
			json.NewEncoder(w).Encode(movie)
			return 	
		}
	}	
}

