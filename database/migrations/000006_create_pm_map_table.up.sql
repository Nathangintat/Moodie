CREATE TABLE IF NOT EXISTS pm_map (
    playlist_id INT NOT NULL,
    movie_id INT NOT NULL,
    PRIMARY KEY (playlist_id, movie_id),
    FOREIGN KEY (playlist_id) REFERENCES playlist(id) ON DELETE CASCADE,
    FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE
    );