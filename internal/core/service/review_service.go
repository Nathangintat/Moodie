package service

import (
	"context"

	"github.com/Nathangintat/Moodie/internal/adapter/repository"
	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/gofiber/fiber/v2/log"
)

type ReviewService interface {
	CreateReview(ctx context.Context, req entity.ReviewEntity) error
	GetReviewByID(ctx context.Context, movieID, userID int64) ([]entity.ReviewItemEntity, error)
	GetReviews(ctx context.Context, userID int64) ([]entity.ReviewsEntity, error)
}
type reviewService struct {
	reviewRepo repository.ReviewRepository
}

func (r *reviewService) CreateReview(ctx context.Context, req entity.ReviewEntity) error {

	err = r.reviewRepo.CreateReview(ctx, req)
	if err != nil {
		code = "[SERVICE] CreateReview - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func (r *reviewService) GetReviews(ctx context.Context, userID int64) ([]entity.ReviewsEntity, error) {
	reviews, err := r.reviewRepo.GetReviews(ctx, userID)
	if err != nil {
		code = "[SERVICE] GetReviews - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return reviews, nil
}

func (r *reviewService) GetReviewByID(ctx context.Context, movieID, userID int64) ([]entity.ReviewItemEntity, error) {
	results, err := r.reviewRepo.GetReviewByID(ctx, movieID, userID)
	if err != nil {
		code = "[SERVICE] GetReviewByID - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return results, err
}

func NewReviewService(reviewRepo repository.ReviewRepository) ReviewService {
	return &reviewService{
		reviewRepo: reviewRepo,
	}
}
