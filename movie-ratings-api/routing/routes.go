package routing

import (
	"movie-ratings-api/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/movies", handlers.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", handlers.GetMovieByID).Methods("GET")
	router.HandleFunc("/rating_by_movie", handlers.GetAverageRating).Methods("GET")
	router.HandleFunc("/movie", handlers.AddMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", handlers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{id}", handlers.DeleteMovie).Methods("DELETE")

	return router
}
