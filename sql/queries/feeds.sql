-- name: CreateFeed :one
Insert into feeds(id, created_at, updated_at, name, user_id, url)
VALUES ($1, $2, $3, $4, $5, $6)
returning *;
