package entity

import "time"

type MovieEntity struct {
	ID          int64
	Name        string
	Poster      string
	Overview    string
	ReleaseDate time.Time
	Rating      float64
	Genres      []string
}

type QueryString struct {
	Limit     int
	Page      int
	OrderBy   string
	OrderType string
	Search    string
}

type SearchMovie struct {
	ID     int64
	Name   string
	Poster string
}
