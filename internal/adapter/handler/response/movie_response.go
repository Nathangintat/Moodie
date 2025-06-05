package response

type MoviesResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Poster string `json:"poster"`
}

type MovieResponse struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Poster      string   `json:"poster"`
	Overview    string   `json:"overview"`
	ReleaseDate string   `json:"release_date"`
	Rating      float64  `json:"rating"`
	Genres      []string `json:"genres"`
}
