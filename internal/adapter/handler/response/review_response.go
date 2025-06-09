package response

import "time"

type ReviewResponse struct {
	MovieID int64                `json:"movie_id"`
	Review  []ReviewItemResponse `json:"review"`
}

type ReviewsResponse struct {
	ReviewID      int64     `json:"review_id"`
	MovieID       int64     `json:"movie_id"`
	UserID        int64     `json:"user_id"`
	UserName      string    `json:"username"`
	Content       string    `json:"content"`
	Poster        string    `json:"poster"`
	Rating        int64     `json:"rating"`
	VoteCount     int64     `json:"vote_count"`
	DownvoteCount int64     `json:"downvote_count"`
	HasVoted      bool      `json:"has_voted"`
	HasDownvoted  bool      `json:"has_downvoted"`
	CreatedAt     time.Time `json:"created_at"`
}

type ReviewItemResponse struct {
	UserID        int64     `json:"user_id"`
	Content       string    `json:"content"`
	Rating        int64     `json:"rating"`
	CreatedAt     time.Time `json:"created_at"`
	VoteCount     int64     `json:"vote_count"`
	DownvoteCount int64     `json:"downvote_count"`
	HasVoted      bool      `json:"has_voted"`
	HasDownvoted  bool      `json:"has_downvoted"`
}
