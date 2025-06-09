CREATE TABLE IF NOT EXISTS playlist (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    user_id INT NOT NULL,
    playlist_image VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES "users"(id) ON DELETE CASCADE
);