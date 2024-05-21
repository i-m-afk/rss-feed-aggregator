-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
returning *;

-- name: GetPosts :many
Select * from posts
ORDER BY published_at
LIMIT $1;

