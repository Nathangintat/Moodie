package request

import "time"

type ReviewRequest struct {
	MovieID   int64  `json:"movie_id"`
	UserID    int64  `json:"user_id"`
	Headline  string `json:"headline"`
	Content   string `json:"content"`
	Rating    int64  `json:"rating"`
	Emoji     string `json:"emoji"`
	CreatedAt time.Time
}
