-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: ListFeeds :many
SELECT f.name, f.url, u.name as username
FROM feeds f
LEFT JOIN users u ON f.user_id = u.id;

-- name: GetFeedByURL :one
SELECT * from feeds WHERE url = $1;
