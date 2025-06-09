package response

type PlaylistResponse struct {
	PlaylistID    int64  `json:"playlist_id"`
	Name          string `json:"name"`
	PlaylistImage string `json:"playlist_image"`
}

type PlaylistItemResponse struct {
	MovieID int64  `json:"movie_id"`
	Name    string `json:"name"`
	Poster  string `json:"poster"`
}
