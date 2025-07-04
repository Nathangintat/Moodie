package repository

import (
	"context"

	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/Nathangintat/Moodie/internal/core/domain/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, req entity.ReviewEntity) error
	GetReviewByID(ctx context.Context, movieID, userID int64) ([]entity.ReviewItemEntity, error)
	GetReviews(ctx context.Context, userID int64) ([]entity.ReviewsEntity, error)
}
type reviewRepository struct {
	db *gorm.DB
}

func (r *reviewRepository) CreateReview(ctx context.Context, req entity.ReviewEntity) error {

	review := model.Review{
		ID:        req.ID,
		MovieID:   req.MovieID,
		UserID:    req.UserID,
		Headline:  req.Headline,
		Content:   req.Content,
		Rating:    req.Rating,
		Emoji:     req.Emoji,
		CreatedAt: req.CreatedAt,
	}

	err := r.db.WithContext(ctx).Create(&review).Error
	if err != nil {
		code = "[REPOSITORY] CreateReview - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func (r *reviewRepository) GetReviewByID(ctx context.Context, movieID, userID int64) ([]entity.ReviewItemEntity, error) {
	var reviews []model.Review
	err := r.db.WithContext(ctx).Preload(clause.Associations).Where("movie_id = ?", movieID).Find(&reviews).Error
	if err != nil {
		code = "[REPOSITORY] GetReviewByID - 1"
		log.Errorw(code, err)
		return nil, err
	}

	resp := []entity.ReviewItemEntity{}
	for _, review := range reviews {

		var voteCount int64
		var downvoteCount int64
		var hasVoted bool
		var hasDownvoted bool

		r.db.Model(&model.UpvoteReview{}).Where("review_id = ?", review.ID).Count(&voteCount)
		r.db.Model(&model.DownvoteReview{}).Where("review_id = ?", review.ID).Count(&downvoteCount)

		var tmp int64
		r.db.Model(&model.UpvoteReview{}).Where("review_id = ? AND user_id = ?", review.ID, userID).Count(&tmp)
		hasVoted = tmp > 0

		r.db.Model(&model.DownvoteReview{}).Where("review_id = ? AND user_id = ?", review.ID, userID).Count(&tmp)
		hasDownvoted = tmp > 0

		reviewItem := entity.ReviewItemEntity{
			ID:            review.ID,
			MovieID:       review.MovieID,
			UserID:        review.UserID,
			Username:      review.User.Username,
			ProfileImage:  review.User.ProfileImage,
			Headline:      review.Headline,
			Content:       review.Content,
			Rating:        review.Rating,
			Emoji:         review.Emoji,
			CreatedAt:     review.CreatedAt,
			VoteCount:     voteCount,
			DownvoteCount: downvoteCount,
			HasVoted:      hasVoted,
			HasDownvoted:  hasDownvoted,
		}

		resp = append(resp, reviewItem)
	}

	return resp, err
}

func (r *reviewRepository) GetReviews(ctx context.Context, userID int64) ([]entity.ReviewsEntity, error) {
	var reviews []model.Review
	err := r.db.WithContext(ctx).Preload(clause.Associations).Find(&reviews).Error
	if err != nil {
		code = "[REPOSITORY] GetReviews - 1"
		log.Errorw(code, err)
		return nil, err
	}

	resp := []entity.ReviewsEntity{}
	for _, review := range reviews {
		var voteCount int64
		var downvoteCount int64
		var hasVoted bool
		var hasDownvoted bool

		r.db.Model(&model.UpvoteReview{}).Where("review_id = ?", review.ID).Count(&voteCount)
		r.db.Model(&model.DownvoteReview{}).Where("review_id = ?", review.ID).Count(&downvoteCount)

		var tmp int64
		r.db.Model(&model.UpvoteReview{}).Where("review_id = ? AND user_id = ?", review.ID, userID).Count(&tmp)
		hasVoted = tmp > 0

		r.db.Model(&model.DownvoteReview{}).Where("review_id = ? AND user_id = ?", review.ID, userID).Count(&tmp)
		hasDownvoted = tmp > 0

		reviewItem := entity.ReviewsEntity{
			ID:            review.ID,
			MovieID:       review.MovieID,
			MovieName:     review.Movie.MovieName,
			UserID:        review.UserID,
			ProfileImage:  review.User.ProfileImage,
			UserName:      review.User.Username,
			Headline:      review.Headline,
			Content:       review.Content,
			Poster:        review.Movie.Poster,
			Rating:        review.Rating,
			Emoji:         review.Emoji,
			VoteCount:     voteCount,
			DownvoteCount: downvoteCount,
			HasVoted:      hasVoted,
			HasDownvoted:  hasDownvoted,
			CreatedAt:     review.CreatedAt,
		}
		resp = append(resp, reviewItem)
	}

	return resp, nil
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{
		db: db,
	}
}
