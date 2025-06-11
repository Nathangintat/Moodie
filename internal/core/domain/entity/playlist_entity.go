package entity

type PlaylistEntity struct {
	ID            int64
	Name          string
	UserID        int64
	PlaylistImage string
}

type MoviePlaylistEntity struct {
	ID            int64
	Name          string
	Poster        string
	PlaylistImage string
	PlaylistName  string
}
