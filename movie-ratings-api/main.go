package main

import (
	"log"
	"movie-ratings-api/routing"
	"net/http"
)

func main() {
	router := routing.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
