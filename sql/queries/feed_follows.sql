-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
	INSERT INTO feed_follows (feed_id, user_id)
	VALUES ($1, $2)
	RETURNING *
)
SELECT
inserted_feed_follow.*,
users.name as user_name,
feeds.name as feed_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFollowingFeedsByUserID :many
SELECT
feed_follows.*,
feeds.name as feed_name
FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;
