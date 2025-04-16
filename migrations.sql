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