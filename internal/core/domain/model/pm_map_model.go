package model

type PmMap struct {
	PlaylistID int64 `gorm:"playlist_id"`
	MovieID    int64 `gorm:"movie_id"`

	Playlist Playlist `gorm:"foreignkey:PlaylistID;references:ID"`
	Movie    Movie    `gorm:"foreignkey:MovieID;references:ID"`
}

func (PmMap) TableName() string {
	return "pm_map"
}
