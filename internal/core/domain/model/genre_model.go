package model

type Genre struct {
	ID   int64  `gorm:"id"`
	Name string `gorm:"name"`
}

func (g *Genre) TableName() string {
	return "genre"
}
