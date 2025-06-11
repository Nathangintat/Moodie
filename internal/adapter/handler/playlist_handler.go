package handler

import (
	"github.com/Nathangintat/Moodie/internal/adapter/handler/request"
	"github.com/Nathangintat/Moodie/internal/adapter/handler/response"
	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/Nathangintat/Moodie/internal/core/service"
	"github.com/Nathangintat/Moodie/lib/conv"
	validatorLib "github.com/Nathangintat/Moodie/lib/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type PlaylistHandler interface {
	CreatePlaylist(c *fiber.Ctx) error
	GetPlaylistByID(c *fiber.Ctx) error
	InsertMovie(c *fiber.Ctx) error
	GetPlaylistMovies(c *fiber.Ctx) error
}

type playlistHandler struct {
	playlistService service.PlaylistService
}

func (ph *playlistHandler) CreatePlaylist(c *fiber.Ctx) error {

	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] CreatePlaylist - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	userID := claims.UserID
	var req request.PlaylistRequest
	if err = c.BodyParser(&req); err != nil {
		code := "[HANDLER] CreatePlaylist - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Invalid request body"

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	result, err := c.FormFile("playlist_image")
	if err != nil {
		code := "[HANDLER] CreatePlaylist - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Invalid request body"

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	if err = validatorLib.ValidateStruct(&req); err != nil {
		code := "[HANDLER] CreatePlaylist - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	reqEntity := entity.PlaylistEntity{
		Name:   req.Name,
		UserID: int64(userID),
	}

	err = ph.playlistService.CreatePlaylist(c.Context(), reqEntity, result)
	if err != nil {
		code := "[HANDLER] CreatePlaylist - 4"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Playlist created successfully"
	defaultSuccessResponse.Data = nil

	return c.Status(fiber.StatusCreated).JSON(defaultSuccessResponse)
}

func (ph *playlistHandler) GetPlaylistByID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] GetPlaylistByID - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	userID := claims.UserID
	results, err := ph.playlistService.GetPlaylistByID(c.Context(), int64(userID))
	if err != nil {
		code := "[HANDLER] GetPlaylistByID - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success"

	resp := []response.PlaylistResponse{}
	for _, p := range results {
		playlist := response.PlaylistResponse{
			PlaylistID:    p.ID,
			Name:          p.Name,
			PlaylistImage: p.PlaylistImage,
		}

		resp = append(resp, playlist)
	}

	defaultSuccessResponse.Data = resp
	return c.JSON(defaultSuccessResponse)
}

func (ph *playlistHandler) InsertMovie(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] InsertMovie - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	userID := claims.UserID
	var req request.InsertMovieRequest
	if err = c.BodyParser(&req); err != nil {
		code := "[HANDLER] InsertMovie - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Invalid request body"

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	if err = validatorLib.ValidateStruct(&req); err != nil {
		code := "[HANDLER] InsertMovie - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	reqEntity := entity.PmMapEntity{
		PlaylistID: req.PlaylistID,
		MovieID:    req.MovieID,
	}

	err = ph.playlistService.InsertMovie(c.Context(), &reqEntity, int64(userID))
	if err != nil {
		code := "[HANDLER] InsertMovie - 4"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Content created successfully"
	defaultSuccessResponse.Data = nil

	return c.Status(fiber.StatusCreated).JSON(defaultSuccessResponse)

}

func (ph *playlistHandler) GetPlaylistMovies(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] GetPlaylistMovie - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	idParam := c.Params("playlistID")
	playlistID, err := conv.StringToInt64(idParam)
	if err != nil {
		code := "[HANDLER] GetMovieByID - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	results, err := ph.playlistService.GetPlaylistMovies(c.Context(), playlistID)
	if err != nil {
		code := "[HANDLER] GetPlaylistMovie - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success"

	resp := []response.PlaylistItemResponse{}
	for _, p := range results {
		playlist := response.PlaylistItemResponse{
			MovieID:       p.ID,
			Name:          p.Name,
			Poster:        p.Poster,
			PlaylistImage: p.PlaylistImage,
			PlaylistName:  p.PlaylistName,
		}

		resp = append(resp, playlist)
	}

	defaultSuccessResponse.Data = resp
	return c.JSON(defaultSuccessResponse)
}

func NewPlaylistHandler(playlistService service.PlaylistService) PlaylistHandler {
	return &playlistHandler{playlistService: playlistService}
}
