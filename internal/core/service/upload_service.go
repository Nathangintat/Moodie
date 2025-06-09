package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type UploadService interface {
	SavePlaylistImage(file *multipart.FileHeader) (string, error)
	SaveUserProfileImage(file *multipart.FileHeader) (string, error)
}

type uploadService struct {
	uploadPath string
}

func (u *uploadService) SavePlaylistImage(file *multipart.FileHeader) (string, error) {
	return u.SaveImage(file, "playlists")
}

func (u *uploadService) SaveUserProfileImage(file *multipart.FileHeader) (string, error) {
	return u.SaveImage(file, "profiles")
}

func (u *uploadService) SaveImage(file *multipart.FileHeader, folder string) (string, error) {
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExts[ext] {
		code := "[SERVICE] upload - 1"
		err := errors.New("invalid file type: only jpg, jpeg, png allowed")
		log.Errorw(code, err)
		return "", err
	}

	const maxSize = 2 << 20 // 2MB
	if file.Size > maxSize {
		code := "[SERVICE] upload - 2"
		err := errors.New("file size exceeds 2MB limit")
		log.Errorw(code, err)
		return "", err
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	saveDir := filepath.Join(u.uploadPath, folder)
	err := os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		code := "[SERVICE] upload - 3"
		log.Errorw(code, err)
		return "", err
	}

	savePath := filepath.Join(saveDir, filename)

	err = saveFile(file, savePath)
	if err != nil {
		code := "[SERVICE] upload - 4"
		log.Errorw(code, err)
		return "", err
	}

	return filename, nil
}

func saveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		code = "[SERVICE] savefile - 1"
		log.Errorw(code, err)
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		code = "[SERVICE] savefile - 2"
		log.Errorw(code, err)
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(src)
	return err
}

func NewUploadService(uploadPath string) UploadService {
	return &uploadService{uploadPath}
}
