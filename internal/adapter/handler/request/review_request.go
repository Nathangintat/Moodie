package request

import "time"

type ReviewRequest struct {
	MovieID   int64  `json:"movie_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	Rating    int64  `json:"rating"`
	CreatedAt time.Time
}
