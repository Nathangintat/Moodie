package handler

import (
	"time"

	"github.com/Nathangintat/Moodie/internal/adapter/handler/request"
	"github.com/Nathangintat/Moodie/internal/adapter/handler/response"
	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/Nathangintat/Moodie/internal/core/service"
	"github.com/Nathangintat/Moodie/lib/conv"
	validatorLib "github.com/Nathangintat/Moodie/lib/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ReviewHandler interface {
	CreateReview(c *fiber.Ctx) error
	GetReviewByID(c *fiber.Ctx) error
	GetReviews(c *fiber.Ctx) error
}

type reviewHandler struct {
	reviewService service.ReviewService
}

func (rh *reviewHandler) CreateReview(c *fiber.Ctx) error {

	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] CreateReview - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	userID := claims.UserID
	var req request.ReviewRequest
	if err = c.BodyParser(&req); err != nil {
		code := "[HANDLER] CreateReview - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Invalid request body"

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	if err = validatorLib.ValidateStruct(&req); err != nil {
		code := "[HANDLER] CreateReview - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	reqEntity := entity.ReviewEntity{
		UserID:    int64(userID),
		MovieID:   req.MovieID,
		Headline:  req.Headline,
		Content:   req.Content,
		Rating:    req.Rating,
		Emoji:     req.Emoji,
		CreatedAt: time.Now(),
	}

	err = rh.reviewService.CreateReview(c.Context(), reqEntity)
	if err != nil {
		code := "[HANDLER] CreateReview - 4"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	resp := response.DefaultSuccessResponse{
		Meta: response.Meta{
			Status:  true,
			Message: "success",
		},
		Data: nil,
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (rh *reviewHandler) GetReviewByID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] GetReviewByID - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	userID := claims.UserID
	idParam := c.Params("movieID")
	movieID, err := conv.StringToInt64(idParam)
	if err != nil {
		code := "[HANDLER] GetReviewById - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	results, err := rh.reviewService.GetReviewByID(c.Context(), movieID, int64(userID))
	if err != nil {
		code := "[HANDLER] GetReviewByID - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success"

	var reviews []response.ReviewItemResponse
	for _, r := range results {
		reviews = append(reviews, response.ReviewItemResponse{
			UserID:        r.UserID,
			Username:      r.Username,
			ProfileImage:  r.ProfileImage,
			Content:       r.Content,
			Headline:      r.Headline,
			Rating:        r.Rating,
			Emoji:         r.Emoji,
			CreatedAt:     r.CreatedAt,
			VoteCount:     r.VoteCount,
			DownvoteCount: r.DownvoteCount,
			HasVoted:      r.HasVoted,
			HasDownvoted:  r.HasDownvoted,
		})
	}

	resp := response.ReviewResponse{
		MovieID: movieID,
		Review:  reviews,
	}

	defaultSuccessResponse.Data = resp
	return c.JSON(defaultSuccessResponse)
}

func (rh *reviewHandler) GetReviews(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] GetReviews - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	userID := claims.UserID
	results, err := rh.reviewService.GetReviews(c.Context(), int64(userID))
	if err != nil {
		code := "[HANDLER] GetReviews - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success"

	var resp []response.ReviewsResponse
	for _, r := range results {
		resp = append(resp, response.ReviewsResponse{
			ReviewID:      int64(r.ID),
			MovieID:       r.MovieID,
			MovieName:     r.MovieName,
			UserID:        r.UserID,
			ProfileImage:  r.ProfileImage,
			UserName:      r.UserName,
			Headline:      r.Headline,
			Content:       r.Content,
			Rating:        r.Rating,
			Emoji:         r.Emoji,
			Poster:        r.Poster,
			VoteCount:     r.VoteCount,
			DownvoteCount: r.DownvoteCount,
			HasVoted:      r.HasVoted,
			HasDownvoted:  r.HasDownvoted,
			CreatedAt:     r.CreatedAt,
		})
	}

	defaultSuccessResponse.Data = resp
	return c.JSON(defaultSuccessResponse)
}

func NewReviewHandler(reviewService service.ReviewService) ReviewHandler {
	return &reviewHandler{reviewService: reviewService}
}
