package models

type Movie struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
}
