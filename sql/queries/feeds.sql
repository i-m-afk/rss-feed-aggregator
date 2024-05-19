-- name: CreateFeed :one
Insert into feeds(id, created_at, updated_at, name, user_id, url)
VALUES ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetFeeds :many
SELECT * FROM feeds ORDER BY created_at;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at LIMIT $1;

-- name: UpdateFeed :exec
UPDATE feeds
  SET last_fetched_at = $1;

-- name: MarkFeedAsFetched :exec
UPDATE feeds
  SET updated_at = now(), last_fetched_at = $1
  WHERE id = $2;
