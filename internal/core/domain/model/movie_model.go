package model

import "time"

type Movie struct {
	ID          int64     `gorm:"id"`
	MovieName   string    `gorm:"movie_name"`
	Poster      string    `gorm:"poster"`
	Overview    string    `gorm:"overview"`
	ReleaseDate time.Time `gorm:"release_date"`
	Runtime     int64     `gorm:"runtime"`
}

func (m *Movie) TableName() string {
	return "movie"
}
