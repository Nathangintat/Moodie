package model

type User struct {
	ID           int64  `gorm:"id"`
	Username     string `gorm:"username"`
	Email        string `gorm:"email"`
	Password     string `gorm:"password"`
	ProfileImage string `gorm:"profile_image"`
}
