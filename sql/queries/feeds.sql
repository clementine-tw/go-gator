-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFeedsWithUserName :many
SELECT feeds.name, feeds.url, users.name AS user_name FROM feeds
INNER JOIN users ON feeds.user_id = users.id;

-- name: GetFeedIDByUrl :one
SELECT id FROM feeds WHERE url = $1;
