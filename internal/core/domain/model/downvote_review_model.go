package model

type DownvoteReview struct {
	ReviewID int64 `gorm:"review_id"`
	UserID   int64 `gorm:"user_id"`

	Review Review `gorm:"foreignkey:ReviewID;references:ID"`
	User   User   `gorm:"foreignkey:UserID;references:ID"`
}

func (DownvoteReview) TableName() string {
	return "downvote_review"
}
