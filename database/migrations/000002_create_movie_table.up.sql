-- users
CREATE TABLE IF NOT EXISTS movie (
    id SERIAL PRIMARY KEY,
    movie_name VARCHAR(255) NOT NULL,
    poster TEXT,
    overview TEXT,
    release_date DATE,
    runtime INT
);