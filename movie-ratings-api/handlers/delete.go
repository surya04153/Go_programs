package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	for index, movie := range movies {

		if movie.ID == key {
			index++
		}

	}
}
