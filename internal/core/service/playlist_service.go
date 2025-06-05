package service

import (
	"context"

	"github.com/Nathangintat/Moodie/internal/adapter/repository"
	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/gofiber/fiber/v2/log"
)

type PlaylistService interface {
	CreatePlaylist(ctx context.Context, req entity.PlaylistEntity) error
	GetPlaylistByID(ctx context.Context, userID int64) ([]entity.PlaylistEntity, error)
	InsertMovie(ctx context.Context, req *entity.PmMapEntity, userID int64) error
	GetPlaylistMovies(ctx context.Context, playlistID int64) ([]entity.MovieEntity, error)
}
type playlistService struct {
	playlistRepo repository.PlaylistRepository
}

func (p *playlistService) CreatePlaylist(ctx context.Context, req entity.PlaylistEntity) error {

	err = p.playlistRepo.CreatePlaylist(ctx, req)
	if err != nil {
		code = "[SERVICE] CreatePlaylist - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func (p *playlistService) GetPlaylistByID(ctx context.Context, userID int64) ([]entity.PlaylistEntity, error) {
	results, err := p.playlistRepo.GetPlaylistByID(ctx, userID)
	if err != nil {
		code = "[SERVICE] GetPlaylistByID - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return results, err
}

func (p *playlistService) InsertMovie(ctx context.Context, req *entity.PmMapEntity, userID int64) error {
	err := p.playlistRepo.InsertMovie(ctx, req, userID)
	if err != nil {
		code = "[SERVICE] InsertMovie - 1"
		log.Errorw(code, err)
		return err
	}
	return nil
}

func (p *playlistService) GetPlaylistMovies(ctx context.Context, playlistID int64) ([]entity.MovieEntity, error) {
	results, err := p.playlistRepo.GetPlaylistMovies(ctx, playlistID)
	if err != nil {
		code = "[SERVICE] GetPlaylistMovies - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return results, err
}

func NewPlaylistService(playlistRepo repository.PlaylistRepository) PlaylistService {
	return &playlistService{
		playlistRepo: playlistRepo,
	}
}
