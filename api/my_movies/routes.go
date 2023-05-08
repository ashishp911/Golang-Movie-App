package my_movies

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func Enter_routes() {
	my_router := mux.NewRouter()
	my_router.HandleFunc("/movies", GetMovies).Methods("GET")
	my_router.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	my_router.HandleFunc("/movies", CreateMovie).Methods("POST")
	my_router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	my_router.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000...\n")
	log.Fatal(http.ListenAndServe(":8000", my_router))
}
