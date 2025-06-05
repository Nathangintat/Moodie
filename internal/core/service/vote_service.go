package service

import (
	"context"

	"github.com/Nathangintat/Moodie/internal/adapter/repository"
	"github.com/gofiber/fiber/v2/log"
)

type VoteService interface {
	AddUpvote(ctx context.Context, reviewID, userID int64) error
	AddDownvote(ctx context.Context, reviewID, userID int64) error
}

type voteService struct {
	voteRepo repository.VoteRepository
}

func (v *voteService) AddUpvote(ctx context.Context, reviewID, userID int64) error {
	err := v.voteRepo.AddUpvote(ctx, reviewID, userID)
	if err != nil {
		code = "[SERVICE] AddUpvote - 1"
		log.Errorw(code, err)
		return err
	}
	return nil
}

func (v *voteService) AddDownvote(ctx context.Context, reviewID, userID int64) error {
	err := v.voteRepo.AddDownvote(ctx, reviewID, userID)
	if err != nil {
		code = "[SERVICE] AddDownvote - 1"
		log.Errorw(code, err)
		return err
	}
	return nil
}

func NewVoteService(voteRepo repository.VoteRepository) VoteService {
	return &voteService{
		voteRepo: voteRepo,
	}
}
