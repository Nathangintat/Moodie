package service

import (
	"context"

	"github.com/Nathangintat/Moodie/internal/adapter/repository"
	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/gofiber/fiber/v2/log"
)

type MovieService interface {
	GetMovies(ctx context.Context, query entity.QueryString) ([]entity.MovieEntity, int64, int64, error)
	GetMovieByID(ctx context.Context, id int64) (*entity.MovieEntity, error)
	SearchMovie(ctx context.Context, query string) ([]entity.SearchMovie, error)
}

type movieService struct {
	movieRepo repository.MovieRepository
}

func (m *movieService) GetMovies(ctx context.Context, query entity.QueryString) ([]entity.MovieEntity, int64, int64, error) {
	results, totalData, totalPages, err := m.movieRepo.GetMovies(ctx, query)
	if err != nil {
		code = "[SERVICE] GetMovies - 1"
		log.Errorw(code, err)
		return nil, 0, 0, err
	}

	return results, totalData, totalPages, nil
}

func (m *movieService) GetMovieByID(ctx context.Context, id int64) (*entity.MovieEntity, error) {
	result, err := m.movieRepo.GetMovieByID(ctx, id)
	if err != nil {
		code := "[SERVICE] GetMovieByID - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return result, nil
}

func (m *movieService) SearchMovie(ctx context.Context, q string) ([]entity.SearchMovie, error) {
	result, err := m.movieRepo.SearchMovie(ctx, q)
	if err != nil {
		code := "[SERVICE] SearchMovie - 1"
		log.Errorw(code, err)
		return nil, err
	}

	return result, nil
}

func NewMovieService(movieRepo repository.MovieRepository) MovieService {
	return &movieService{movieRepo: movieRepo}
}
