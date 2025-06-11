package response

import "time"

type ReviewResponse struct {
	MovieID int64                `json:"movie_id"`
	Review  []ReviewItemResponse `json:"review"`
}

type ReviewsResponse struct {
	ReviewID      int64     `json:"review_id"`
	MovieID       int64     `json:"movie_id"`
	MovieName     string    `json:"movie_name"`
	UserID        int64     `json:"user_id"`
	ProfileImage  string    `json:"profile_image"`
	UserName      string    `json:"username"`
	Headline      string    `json:"headline"`
	Content       string    `json:"content"`
	Poster        string    `json:"poster"`
	Rating        int64     `json:"rating"`
	Emoji         string    `json:"emoji"`
	VoteCount     int64     `json:"vote_count"`
	DownvoteCount int64     `json:"downvote_count"`
	HasVoted      bool      `json:"has_voted"`
	HasDownvoted  bool      `json:"has_downvoted"`
	CreatedAt     time.Time `json:"created_at"`
}

type ReviewItemResponse struct {
	UserID        int64     `json:"user_id"`
	Username      string    `json:"username"`
	ProfileImage  string    `json:"profile_image"`
	Headline      string    `json:"headline"`
	Content       string    `json:"content"`
	Rating        int64     `json:"rating"`
	Emoji         string    `json:"emoji"`
	CreatedAt     time.Time `json:"created_at"`
	VoteCount     int64     `json:"vote_count"`
	DownvoteCount int64     `json:"downvote_count"`
	HasVoted      bool      `json:"has_voted"`
	HasDownvoted  bool      `json:"has_downvoted"`
}
