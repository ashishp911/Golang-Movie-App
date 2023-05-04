package routes

import(
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	// "encoding/json"
	"log"
	"go-crud-movies/bizlogic"
)


func All_routes(){
	my_router := mux.NewRouter()
	bizlogic.All_bizz_logic()
	my_router.HandleFunc("/movies", bizlogic.GetMovies).Methods("GET")
	my_router.HandleFunc("/movies/{id}", bizlogic.GetMovie).Methods("GET")
	my_router.HandleFunc("/movies", bizlogic.CreateMovie).Methods("POST")
	my_router.HandleFunc("/movies/{id}", bizlogic.UpdateMovie).Methods("PUT")
	my_router.HandleFunc("/movies/{id}", bizlogic.DeleteMovie).Methods("DELETE")

		
	fmt.Printf("Starting server at port 8000...\n")
	log.Fatal(http.ListenAndServe(":8000", my_router))
}