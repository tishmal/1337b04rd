DROP TABLE IF EXISTS comments CASCADE;

CREATE TABLE comments (
    id VARCHAR(36) PRIMARY KEY,
    content TEXT NOT NULL,
    post_id VARCHAR(36) REFERENCES posts(id),
    reply_to_id VARCHAR(36),
    user_id VARCHAR(36) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);