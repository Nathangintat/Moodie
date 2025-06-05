package handler

import (
	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/Nathangintat/Moodie/internal/core/service"
	"github.com/Nathangintat/Moodie/lib/conv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type VoteHandler interface {
	AddUpvote(c *fiber.Ctx) error
	AddDownvote(c *fiber.Ctx) error
}

type voteHandler struct {
	voteService service.VoteService
}

func (v *voteHandler) AddUpvote(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] AddUpvote - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	idParam := c.Params("reviewID")
	reviewId, err := conv.StringToInt64(idParam)
	if err != nil {
		code := "[HANDLER] AddUpvote - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	err = v.voteService.AddUpvote(c.Context(), reviewId, int64(claims.UserID))
	if err != nil {
		code := "[HANDLER] AddUpvote - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Upvote Added successfully"
	defaultSuccessResponse.Data = nil

	return c.Status(fiber.StatusCreated).JSON(defaultSuccessResponse)
}

func (v *voteHandler) AddDownvote(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserID == 0 {
		code := "[HANDLER] AddDownvote - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	idParam := c.Params("reviewID")
	reviewId, err := conv.StringToInt64(idParam)
	if err != nil {
		code := "[HANDLER] AddDownvote - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	err = v.voteService.AddDownvote(c.Context(), reviewId, int64(claims.UserID))
	if err != nil {
		code := "[HANDLER] AddDownvote - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Downvote Added successfully"
	defaultSuccessResponse.Data = nil

	return c.Status(fiber.StatusCreated).JSON(defaultSuccessResponse)
}

func NewVoteHandler(voteService service.VoteService) VoteHandler {
	return &voteHandler{
		voteService: voteService,
	}
}
