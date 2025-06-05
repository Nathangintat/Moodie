package model

type MgMap struct {
	MovieID int64 `gorm:"movie_id"`
	GenreID int64 `gorm:"genre_id"`

	Movie Movie `gorm:"foreignkey:MovieID;references:ID"`
	Genre Genre `gorm:"foreignkey:GenreID;references:ID"`
}

func (m *MgMap) TableName() string {
	return "mg_map"
}
