-- migrations/000001_init_schema.up.sql
CREATE TABLE IF NOT EXISTS sessions (
    id SERIAL PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL UNIQUE,
    user_avatar_id INT NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    custom_name VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT sessions_session_id_key UNIQUE (session_id)
);

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    post_id VARCHAR(20) NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    image_path VARCHAR(255),
    session_id VARCHAR(255) NOT NULL,
    user_avatar_id INT NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_activity TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_archived BOOLEAN DEFAULT FALSE,
    CONSTRAINT posts_post_id_key UNIQUE (post_id),
    CONSTRAINT fk_session FOREIGN KEY (session_id) REFERENCES sessions(session_id)
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    comment_id VARCHAR(20) NOT NULL UNIQUE,
    post_id VARCHAR(20) NOT NULL,
    parent_id VARCHAR(20),
    content TEXT NOT NULL,
    image_path VARCHAR(255),
    session_id VARCHAR(255) NOT NULL,
    user_avatar_id INT NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT comments_comment_id_key UNIQUE (comment_id),
    CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
    CONSTRAINT fk_session FOREIGN KEY (session_id) REFERENCES sessions(session_id),
    CONSTRAINT fk_parent_comment FOREIGN KEY (parent_id) REFERENCES comments(comment_id) ON DELETE SET NULL
);

-- Создаем индексы для повышения производительности
CREATE INDEX idx_posts_session_id ON posts(session_id);
CREATE INDEX idx_comments_post_id ON comments(post_id);
CREATE INDEX idx_comments_session_id ON comments(session_id);
CREATE INDEX idx_comments_parent_id ON comments(parent_id);
CREATE INDEX idx_posts_last_activity ON posts(last_activity);
CREATE INDEX idx_posts_is_archived ON posts(is_archived);
CREATE INDEX idx_sessions_expires_at ON sessions(expires_at);

-- migrations/000001_init_schema.down.sql
DROP INDEX IF EXISTS idx_sessions_expires_at;
DROP INDEX IF EXISTS idx_posts_is_archived;
DROP INDEX IF EXISTS idx_posts_last_activity;
DROP INDEX IF EXISTS idx_comments_parent_id;
DROP INDEX IF EXISTS idx_comments_session_id;
DROP INDEX IF EXISTS idx_comments_post_id;
DROP INDEX IF EXISTS idx_posts_session_id;

DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS sessions;

-- migrations/000002_add_avatar_cache.up.sql
CREATE TABLE IF NOT EXISTS rick_morty_avatars (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    species VARCHAR(100),
    status VARCHAR(50),
    used_count INT DEFAULT 0,
    last_used TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_avatars_used_count ON rick_morty_avatars(used_count);
CREATE INDEX idx_avatars_last_used ON rick_morty_avatars(last_used);

-- migrations/000002_add_avatar_cache.down.sql
DROP INDEX IF EXISTS idx_avatars_last_used;
DROP INDEX IF EXISTS idx_avatars_used_count;
DROP TABLE IF EXISTS rick_morty_avatars;