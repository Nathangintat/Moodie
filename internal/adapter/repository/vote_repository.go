package repository

import (
	"context"

	"github.com/Nathangintat/Moodie/internal/core/domain/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type VoteRepository interface {
	AddUpvote(ctx context.Context, reviewID, userID int64) error
	AddDownvote(ctx context.Context, reviewID, userID int64) error
}

type voteRepository struct {
	db *gorm.DB
}

func (v *voteRepository) AddUpvote(ctx context.Context, reviewID, userID int64) error {
	return v.db.WithContext(ctx).Transaction(func(db *gorm.DB) error {

		db.Where("review_id = ? AND user_id = ?", reviewID, userID).Delete(&model.DownvoteReview{})

		var exists bool
		err := db.Model(&model.UpvoteReview{}).
			Select("count(*) > 0").
			Where("review_id = ? AND user_id = ?", reviewID, userID).
			Find(&exists).Error

		if err != nil {
			code = "[REPOSITORY] AddUpvote - 1"
			log.Errorw(code, err)
			return err
		}

		if !exists {
			err := db.Create(&model.UpvoteReview{
				ReviewID: reviewID,
				UserID:   userID,
			}).Error
			if err != nil {
				code = "[REPOSITORY] AddUpvote - 2"
				log.Errorw(code, err)
				return err
			}
		} else {
			err := db.Where("review_id = ? AND user_id = ?", reviewID, userID).Delete(&model.UpvoteReview{}).Error
			if err != nil {
				code = "[REPOSITORY] AddUpvote - 3"
				log.Errorw(code, err)
				return err
			}
		}
		return err
	})
}

func (v *voteRepository) AddDownvote(ctx context.Context, reviewID, userID int64) error {
	return v.db.WithContext(ctx).Transaction(func(db *gorm.DB) error {

		db.Where("review_id = ? AND user_id = ?", reviewID, userID).Delete(&model.UpvoteReview{})

		var exists bool
		err := db.Model(&model.DownvoteReview{}).
			Select("count(*) > 0").
			Where("review_id = ? AND user_id = ?", reviewID, userID).
			Find(&exists).Error

		if err != nil {
			code = "[REPOSITORY] AddDownvote - 1"
			log.Errorw(code, err)
			return err
		}

		if !exists {
			err := db.Create(&model.DownvoteReview{
				ReviewID: reviewID,
				UserID:   userID,
			}).Error
			if err != nil {
				code = "[REPOSITORY] AddDownvote - 2"
				log.Errorw(code, err)
				return err
			}
		} else {
			err := db.Where("review_id = ? AND user_id = ?", reviewID, userID).Delete(&model.DownvoteReview{}).Error
			if err != nil {
				code = "[REPOSITORY] AddDownvote - 3"
				log.Errorw(code, err)
				return err
			}
		}
		return err
	})
}

func NewVoteRepository(db *gorm.DB) VoteRepository {
	return &voteRepository{
		db: db,
	}
}
