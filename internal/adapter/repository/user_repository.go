package repository

import (
	"context"

	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/Nathangintat/Moodie/internal/core/domain/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type UserRepository interface {
	UpdatePassword(ctx context.Context, newPass string, id int64) error
	GetUserByID(ctx context.Context, id int64) (*entity.UserEntity, error)
	Register(ctx context.Context, req entity.UserEntity) error
	ChangeProfileImage(ctx context.Context, image string, id int64) error
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Register(ctx context.Context, req entity.UserEntity) error {

	modelUser := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err = u.db.Create(&modelUser).Error
	if err != nil {
		code := "[REPOSITORY] Register - 2"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func (u *userRepository) GetUserByID(ctx context.Context, id int64) (*entity.UserEntity, error) {
	var modelUser model.User
	err = u.db.Where("id = ?", id).First(&modelUser).Error
	if err != nil {
		code := "[REPOSITORY] GetUserByID - 1"
		log.Errorw(code, err)
		return nil, err
	}

	return &entity.UserEntity{
		ID:           id,
		Username:     modelUser.Username,
		Email:        modelUser.Email,
		ProfileImage: modelUser.ProfileImage,
	}, nil
}

func (u *userRepository) UpdatePassword(ctx context.Context, newPass string, id int64) error {
	err = u.db.Model(&model.User{}).Where("id = ?", id).Update("password", newPass).Error
	if err != nil {
		code := "[REPOSITORY] UpdatePassword - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func (u *userRepository) ChangeProfileImage(ctx context.Context, image string, id int64) error {
	err = u.db.Model(&model.User{}).Where("id = ?", id).Update("profile_image", image).Error
	if err != nil {
		code := "[REPOSITORY] ChangeProfileImage - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
