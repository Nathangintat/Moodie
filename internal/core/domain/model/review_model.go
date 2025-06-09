package model

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	MovieID   int64     `gorm:"movie_id"`
	UserID    int64     `gorm:"user_id"`
	Headline  string    `gorm:"headline"`
	Content   string    `gorm:"content"`
	Rating    int64     `gorm:"rating"`
	Emoji     string    `gorm:"emoji"`
	CreatedAt time.Time `gorm:"created_at"`

	Movie Movie `gorm:"foreignkey:MovieID;references:ID"`
	User  User  `gorm:"foreignkey:UserID;references:ID"`
}

func (r *Review) TableName() string {
	return "review"
}
