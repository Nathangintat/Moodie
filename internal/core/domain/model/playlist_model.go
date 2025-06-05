package model

type Playlist struct {
	ID     int64  `gorm:"primary_key"`
	Name   string `gorm:"name"`
	UserID int64  `gorm:"user_id"`

	User User `gorm:"foreignkey:UserID;references:ID"`
}

func (Playlist) TableName() string {
	return "playlist"
}
