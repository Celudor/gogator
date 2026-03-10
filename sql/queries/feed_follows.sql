-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (
        id, created_at, updated_at, user_id, feed_id
    ) VALUES ( $1, $2, $3, $4, $5 )
    RETURNING *
)
SELECT inserted_feed_follows.*,
feeds.name AS feed_name,
users.name AS user_name
FROM inserted_feed_follows
INNER JOIN feeds ON inserted_feed_follows.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follows.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name as feed_name, feeds.url, users.name as user_name
from feed_follows ff 
inner join feeds on ff.feed_id = feeds.id
inner join users on ff.user_id = users.id
where users.name = $1;

-- name: Unfollow :exec
DELETE FROM feed_follows WHERE feed_id = $1 and user_id = $2;
