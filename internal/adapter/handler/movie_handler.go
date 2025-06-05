package handler

import (
	"github.com/Nathangintat/Moodie/internal/adapter/handler/response"
	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/Nathangintat/Moodie/internal/core/service"
	"github.com/Nathangintat/Moodie/lib/conv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type MovieHandler interface {
	GetMovies(c *fiber.Ctx) error
	GetMovieByID(c *fiber.Ctx) error
}

type movieHandler struct {
	movieService service.MovieService
}

func (mh *movieHandler) GetMovies(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] GetMovies - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	page := 1
	if c.Query("page") != "" {
		page, err = conv.StringToInt(c.Query("page"))
		if err != nil {
			code := "[HANDLER] GetMovies - 2"
			log.Errorw(code, err)
			errorResp.Meta.Status = false
			errorResp.Meta.Message = "Invalid page number"

			return c.Status(fiber.StatusBadRequest).JSON(errorResp)
		}
	}

	limit := 10
	if c.Query("limit") != "" {
		limit, err = conv.StringToInt(c.Query("limit"))
		if err != nil {
			code := "[HANDLER] GetMovies - 3"
			log.Errorw(code, err)
			errorResp.Meta.Status = false
			errorResp.Meta.Message = "Invalid limit number"

			return c.Status(fiber.StatusBadRequest).JSON(errorResp)
		}
	}

	orderBy := "release_date"
	if c.Query("orderBy") != "" {
		orderBy = c.Query("orderBy")
	}

	orderType := "desc"
	if c.Query("orderType") != "" {
		orderType = c.Query("orderType")
	}

	search := ""
	if c.Query("search") != "" {
		search = c.Query("search")
	}

	reqEntity := entity.QueryString{
		Limit:     limit,
		Page:      page,
		OrderBy:   orderBy,
		OrderType: orderType,
		Search:    search,
	}

	results, totalData, totalPages, err := mh.movieService.GetMovies(c.Context(), reqEntity)
	if err != nil {
		code := "[HANDLER] GetMovies - 4"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success"

	respMoviess := []response.MoviesResponse{}
	for _, m := range results {
		respMovie := response.MoviesResponse{
			ID:     m.ID,
			Name:   m.Name,
			Poster: m.Poster,
		}

		respMoviess = append(respMoviess, respMovie)
	}

	defaultSuccessResponse.Data = respMoviess
	defaultSuccessResponse.Pagination = &response.PaginationResponse{
		TotalRecords: int(totalData),
		Page:         page,
		PerPage:      limit,
		TotalPages:   int(totalPages),
	}
	return c.JSON(defaultSuccessResponse)
}

func (mh *movieHandler) GetMovieByID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] GetMovieByID - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	println(int64(claims.UserID))

	idParam := c.Params("movieID")
	movieID, err := conv.StringToInt64(idParam)
	if err != nil {
		code := "[HANDLER] GetMovieByID - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	result, err := mh.movieService.GetMovieByID(c.Context(), movieID)
	if err != nil {
		code := "[HANDLER] GetMovieByID - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success"

	respMovie := response.MovieResponse{
		ID:          result.ID,
		Name:        result.Name,
		Poster:      result.Poster,
		Overview:    result.Overview,
		Rating:      result.Rating,
		ReleaseDate: result.ReleaseDate.Format("2006-01-02"),
		Genres:      result.Genres,
	}

	defaultSuccessResponse.Data = respMovie
	return c.JSON(defaultSuccessResponse)
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return &movieHandler{movieService: movieService}
}
