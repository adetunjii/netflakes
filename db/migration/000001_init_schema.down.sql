--- Drop indexes
DROP INDEX IF EXISTS idx_comment_movie_id;
DROP INDEX IF EXISTS idx_comment_sender_ip;
DROP INDEX IF EXISTS idx_comment_created_by;
DROP INDEX IF EXISTS idx_comment_created_at;

-- Drop Tables
DROP TABLE IF EXISTS "comments"; 