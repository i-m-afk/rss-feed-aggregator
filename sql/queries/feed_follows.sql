-- name: CreateFeedFollows :one
INSERT INTO feed_follow(ID, FEED_ID, USER_ID, CREATED_AT, UPDATED_AT)
VALUES ($1, $2, $3, $4, $5)
returning *;

-- name: DeleteFeedFollows :one
DELETE FROM feed_follow
WHERE id=$1
returning *;

-- name: GetFeedFollowsByUserID :many
SELECT * from feed_follow
WHERE user_id = $1
ORDER BY created_at DESC;
