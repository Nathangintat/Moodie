package repository

import (
	"context"
	"fmt"
	"math"

	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/Nathangintat/Moodie/internal/core/domain/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MovieRepository interface {
	GetMovies(ctx context.Context, query entity.QueryString) ([]entity.MovieEntity, int64, int64, error)
	GetMovieByID(ctx context.Context, id int64) (*entity.MovieEntity, error)
	SearchMovie(ctx context.Context, q string) ([]entity.SearchMovie, error)
}

type movieRepository struct {
	db *gorm.DB
}

func (m *movieRepository) GetMovies(ctx context.Context, query entity.QueryString) ([]entity.MovieEntity, int64, int64, error) {
	var modelContents []model.Movie
	var countData int64

	order := fmt.Sprintf("%s %s", query.OrderBy, query.OrderType)
	offset := (query.Page - 1) * query.Limit

	sqlMain := m.db.Preload(clause.Associations).
		Where("movie_name ilike ? OR overview ilike ? ", "%"+query.Search+"%", "%"+query.Search+"%")

	err = sqlMain.Model(&modelContents).Count(&countData).Error
	if err != nil {
		code = "[REPOSITORY] GetMovies - 1"
		log.Errorw(code, err)
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(countData) / float64(query.Limit)))

	err = sqlMain.
		Order(order).
		Limit(query.Limit).
		Offset(offset).
		Find(&modelContents).Error
	if err != nil {
		code = "[REPOSITORY] GetMovies - 2"
		log.Errorw(code, err)
		return nil, 0, 0, err
	}

	resps := []entity.MovieEntity{}
	for _, val := range modelContents {
		resp := entity.MovieEntity{
			ID:     val.ID,
			Name:   val.MovieName,
			Poster: val.Poster,
		}

		resps = append(resps, resp)
	}
	return resps, countData, int64(totalPages), nil
}

func (m *movieRepository) GetMovieByID(ctx context.Context, id int64) (*entity.MovieEntity, error) {
	var modelMovie model.Movie
	err = m.db.Where("id = ?", id).First(&modelMovie).Error
	if err != nil {
		code := "[REPOSITORY] GetMovieByID - 1"
		log.Errorw(code, err)
		return nil, err
	}

	var avgRating float64
	err = m.db.Model(&model.Review{}).
		Where("movie_id = ?", id).
		Select("COALESCE(AVG(rating), 0)").Scan(&avgRating).Error
	if err != nil {
		code := "[REPOSITORY] GetMovieByID - 2"
		log.Errorw(code, err)
		return nil, err
	}

	var modelGenre []model.MgMap
	err = m.db.Model(&model.MgMap{}).Preload("Genre").
		Where("movie_id = ?", id).Find(&modelGenre).Error
	if err != nil {
		code = "[REPOSITORY] GetMovieByID - 3"
		log.Errorw(code, err)
		return nil, err
	}

	var genres []string
	for _, g := range modelGenre {
		genres = append(genres, g.Genre.Name)
	}

	resp := entity.MovieEntity{
		ID:          modelMovie.ID,
		Name:        modelMovie.MovieName,
		Poster:      modelMovie.Poster,
		Overview:    modelMovie.Overview,
		Rating:      avgRating,
		ReleaseDate: modelMovie.ReleaseDate,
		Genres:      genres,
	}

	return &resp, nil
}

func (m *movieRepository) SearchMovie(ctx context.Context, q string) ([]entity.SearchMovie, error) {
	var movieModels []model.Movie

	err := m.db.Where("movie_name ILIKE ?", "%"+q+"%").Find(&movieModels).Error

	if err != nil {
		code := "[REPOSITORY] SearchMovie - 1"
		log.Errorw(code, err)
		return nil, err
	}

	var resp []entity.SearchMovie
	for _, m := range movieModels {

		resp = append(resp, entity.SearchMovie{
			ID:     m.ID,
			Name:   m.MovieName,
			Poster: m.Poster,
		})
	}

	return resp, nil
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}
