package request

type PlaylistRequest struct {
	Name string `json:"name"`
}

type InsertMovieRequest struct {
	PlaylistID int64 `json:"playlist_id"`
	MovieID    int64 `json:"movie_id"`
}
