package handlers

import (
	"encoding/json"
	"movie-ratings-api/models"
	"net/http"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
