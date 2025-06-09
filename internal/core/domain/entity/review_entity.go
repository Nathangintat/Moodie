package entity

import "time"

type ReviewEntity struct {
	ID        uint
	MovieID   int64
	UserID    int64
	Content   string
	Rating    int64
	CreatedAt time.Time
}

type ReviewsEntity struct {
	ID            uint
	MovieID       int64
	UserID        int64
	UserName      string
	Content       string
	Poster        string
	Rating        int64
	VoteCount     int64
	DownvoteCount int64
	HasVoted      bool
	HasDownvoted  bool
	CreatedAt     time.Time
}

type ReviewItemEntity struct {
	ID            uint
	MovieID       int64
	UserID        int64
	Content       string
	Rating        int64
	CreatedAt     time.Time
	VoteCount     int64
	DownvoteCount int64
	HasVoted      bool
	HasDownvoted  bool
}
