package service

import (
	"context"
	"mime/multipart"

	"github.com/Nathangintat/Moodie/internal/adapter/repository"
	"github.com/Nathangintat/Moodie/internal/core/domain/entity"
	"github.com/Nathangintat/Moodie/lib/conv"

	"github.com/gofiber/fiber/v2/log"
)

type UserService interface {
	UpdatePassword(ctx context.Context, newPass string, id int64) error
	GetUserByID(ctx context.Context, id int64) (*entity.UserEntity, error)
	Register(ctx context.Context, req entity.UserEntity) error
	ChangeProfileImage(ctx context.Context, image *multipart.FileHeader, id int64) error
}

type userService struct {
	userRepo      repository.UserRepository
	uploadService UploadService
}

func (u *userService) Register(ctx context.Context, req entity.UserEntity) error {
	req.Password, err = conv.HashPassword(req.Password)
	if err != nil {
		code := "[SERVICE] Register - 1"
		log.Errorw(code, err)
		return err
	}

	err := u.userRepo.Register(ctx, req)
	if err != nil {
		code := "[SERVICE] Register - 2"
		log.Errorw(code, err)
		return err
	}
	return nil
}

func (u *userService) GetUserByID(ctx context.Context, id int64) (*entity.UserEntity, error) {
	result, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		code := "[SERVICE] GetUserByID - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return result, nil
}

func (u *userService) UpdatePassword(ctx context.Context, newPass string, id int64) error {
	password, err := conv.HashPassword(newPass)
	if err != nil {
		code := "[SERVICE] UpdatePassword - 1"
		log.Errorw(code, err)
		return err
	}

	err = u.userRepo.UpdatePassword(ctx, password, id)
	if err != nil {
		code := "[SERVICE] UpdatePassword - 2"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func (u *userService) ChangeProfileImage(ctx context.Context, image *multipart.FileHeader, id int64) error {

	filename, err := u.uploadService.SaveUserProfileImage(image)
	if err != nil {
		code = "[SERVICE] ChangeProfileImage - 1"
		log.Errorw(code, err)
		return err
	}

	err = u.userRepo.ChangeProfileImage(ctx, filename, id)
	if err != nil {
		code := "[SERVICE] ChangeProfileImage - 2"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func NewUserService(userRepo repository.UserRepository, service UploadService) UserService {
	return &userService{
		userRepo:      userRepo,
		uploadService: service,
	}
}
