-- name: CreateRssItem :one
INSERT INTO rss_items(id, created_at, updated_at, title, url, author, description, published_at, post_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
returning *;

-- name: GetRssItems :many
Select * from rss_items
ORDER BY published_at
LIMIT $1;

-- name: GetRssItemsForUser :many
SELECT 
    ri.id AS rss_item_id,
    ri.created_at AS rss_item_created_at,
    ri.updated_at AS rss_item_updated_at,
    ri.title AS rss_item_title,
    ri.url AS rss_item_url,
    ri.author AS rss_item_author,
    ri.description AS rss_item_description,
    ri.published_at AS rss_item_published_at,
    ri.post_id AS rss_item_post_id
FROM 
    rss_items ri
JOIN 
    posts p ON ri.post_id = p.id
JOIN 
    feeds f ON p.feed_id = f.id
JOIN 
    feed_follow ff ON f.id = ff.feed_id
WHERE 
    ff.user_id = $1;

