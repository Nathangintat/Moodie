CREATE TABLE IF NOT EXISTS upvote_review (
    review_id INT NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY (review_id, user_id),
    FOREIGN KEY (review_id) REFERENCES review(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES "users"(id) ON DELETE CASCADE
    );