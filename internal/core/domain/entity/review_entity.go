package entity

import "time"

type ReviewEntity struct {
	ID        uint
	MovieID   int64
	UserID    int64
	Headline  string
	Content   string
	Rating    int64
	Emoji     string
	CreatedAt time.Time
}

type ReviewsEntity struct {
	ID            uint
	MovieID       int64
	MovieName     string
	UserID        int64
	ProfileImage  string
	UserName      string
	Headline      string
	Content       string
	Poster        string
	Rating        int64
	Emoji         string
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
	ProfileImage  string
	Username      string
	Headline      string
	Content       string
	Rating        int64
	Emoji         string
	CreatedAt     time.Time
	VoteCount     int64
	DownvoteCount int64
	HasVoted      bool
	HasDownvoted  bool
}
