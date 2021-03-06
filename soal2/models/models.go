package models

type Response struct {
	Movie []Movie `json:"Search"`
}

type Movie struct {
	ImdbID string `json:"imdbID"`
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieDetail struct {
	ImdbID     string `json:"imdbID"`
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Type       string `json:"Type"`
	Poster     string `json:"Poster"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Country    string `json:"Country"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
}
