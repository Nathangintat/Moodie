package seeds

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Nathangintat/Moodie/internal/core/domain/model"
	"gorm.io/gorm"
)

func SeedGenres(db *gorm.DB, bearerToken string) error {
	url := "https://api.themoviedb.org/3/genre/movie/list?language=en-US"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var genreResp struct {
		Genres []model.Genre `json:"genres"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&genreResp); err != nil {
		return err
	}

	for _, g := range genreResp.Genres {
		db.FirstOrCreate(&g, model.Genre{ID: g.ID})
	}
	return nil
}

func SeedMovies(db *gorm.DB, bearerToken string) error {
	url := "https://api.themoviedb.org/3/movie/popular?language=en-US&page=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var movieResp struct {
		Results []struct {
			ID          int64   `json:"id"`
			Title       string  `json:"title"`
			Overview    string  `json:"overview"`
			PosterPath  string  `json:"poster_path"`
			ReleaseDate string  `json:"release_date"`
			GenreIDs    []int64 `json:"genre_ids"`
		} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&movieResp); err != nil {
		return err
	}

	for _, m := range movieResp.Results {
		rd, _ := time.Parse("2006-01-02", m.ReleaseDate)
		movie := model.Movie{
			ID:          m.ID,
			MovieName:   m.Title,
			Poster:      m.PosterPath,
			Overview:    m.Overview,
			ReleaseDate: rd,
		}
		db.FirstOrCreate(&movie, model.Movie{ID: m.ID})

		for _, gid := range m.GenreIDs {
			db.FirstOrCreate(&model.MgMap{}, model.MgMap{MovieID: m.ID, GenreID: gid})
		}
	}
	return nil
}
