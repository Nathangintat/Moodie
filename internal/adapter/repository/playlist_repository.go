package repository

import (
	"context"
	"errors"

	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/Nathangintat/Moodie/internal/core/domain/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlaylistRepository interface {
	CreatePlaylist(ctx context.Context, req entity.PlaylistEntity) error
	GetPlaylistByID(ctx context.Context, userID int64) ([]entity.PlaylistEntity, error)
	InsertMovie(ctx context.Context, req *entity.PmMapEntity, userId int64) error
	GetPlaylistMovies(ctx context.Context, playlistID int64) ([]entity.MoviePlaylistEntity, error)
}
type playlistRepository struct {
	db *gorm.DB
}

func (r *playlistRepository) CreatePlaylist(ctx context.Context, req entity.PlaylistEntity) error {

	playlist := model.Playlist{
		ID:            req.ID,
		Name:          req.Name,
		UserID:        req.UserID,
		PlaylistImage: req.PlaylistImage,
	}

	err := r.db.WithContext(ctx).Create(&playlist).Error
	if err != nil {
		code = "[REPOSITORY] CreatePlaylist - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func (r *playlistRepository) GetPlaylistByID(ctx context.Context, userID int64) ([]entity.PlaylistEntity, error) {
	var playlists []model.Playlist
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&playlists).Error
	if err != nil {
		code = "[REPOSITORY] GetPlaylists - 1"
		log.Errorw(code, err)
		return nil, err
	}
	resp := []entity.PlaylistEntity{}
	for _, p := range playlists {
		playlist := entity.PlaylistEntity{
			ID:            p.ID,
			Name:          p.Name,
			UserID:        p.UserID,
			PlaylistImage: p.PlaylistImage,
		}
		resp = append(resp, playlist)
	}
	return resp, nil
}

func (r *playlistRepository) InsertMovie(ctx context.Context, req *entity.PmMapEntity, userID int64) error {

	var playlist model.Playlist
	err = r.db.Where("id = ? AND user_id = ?", req.PlaylistID, userID).First(&playlist).Error
	if err != nil {
		return errors.New("unauthorized or playlist not found")
	}

	var exists bool
	err := r.db.Model(&model.PmMap{}).
		Select("count(*) > 0").
		Where("playlist_id = ? AND movie_id = ?", req.PlaylistID, req.MovieID).
		Find(&exists).Error

	if err != nil {
		code = "[REPOSITORY] InsertMovie - 1"
		log.Errorw(code, err)
		return err
	}

	if !exists {
		err := r.db.Create(&model.PmMap{
			PlaylistID: req.PlaylistID,
			MovieID:    req.MovieID,
		}).Error
		if err != nil {
			code = "[REPOSITORY] InsertMovie - 2"
			log.Errorw(code, err)
			return err
		}
	} else {
		return errors.New("movie already exists")
	}

	return nil
}

func (r *playlistRepository) GetPlaylistMovies(ctx context.Context, playlistID int64) ([]entity.MoviePlaylistEntity, error) {
	var playlists []model.PmMap
	err := r.db.WithContext(ctx).Preload(clause.Associations).
		Where("playlist_id = ?", playlistID).
		Find(&playlists).Error
	if err != nil {
		code = "[REPOSITORY] GetPlaylistMovie - 1"
		log.Errorw(code, err)
		return nil, err
	}

	resp := []entity.MoviePlaylistEntity{}
	for _, p := range playlists {
		playlist := entity.MoviePlaylistEntity{
			ID:            p.Movie.ID,
			Name:          p.Movie.MovieName,
			Poster:        p.Movie.Poster,
			PlaylistImage: p.Playlist.PlaylistImage,
			PlaylistName:  p.Playlist.Name,
		}
		resp = append(resp, playlist)
	}
	return resp, nil
}

func NewPlaylistRepository(db *gorm.DB) PlaylistRepository {
	return &playlistRepository{
		db: db,
	}
}
