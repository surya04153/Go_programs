package handlers

import (
	"encoding/json"
	"movie-ratings-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = []models.Movie{
	{ID: "1", Title: "Daiba Daudi", Rating: 8.8},
	{ID: "2", Title: "Jaga Hatare Pagha", Rating: 8.6},
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	for _, movie := range movies {

		if movie.ID == key {
			json.NewEncoder(w).Encode(movie)
		}

	}
}

func GetAverageRating(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
