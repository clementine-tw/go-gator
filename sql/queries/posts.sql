-- name: CreatePost :one
INSERT INTO posts (title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.*, feeds.name as feed_name FROM posts
INNER JOIN feeds ON posts.feed_id = feeds.id
WHERE posts.feed_id IN
(
	SELECT feed_follows.feed_id FROM feed_follows
	WHERE feed_follows.user_id = $1
)
ORDER BY posts.published_at DESC
LIMIT $2;
