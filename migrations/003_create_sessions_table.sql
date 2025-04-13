DROP TABLE IF EXISTS sessions CASCADE;

CREATE TABLE sessions (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP DEFAULT NOW()
);