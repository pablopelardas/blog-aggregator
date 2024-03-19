-- name: CrateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: ListFeeds :many
SELECT * FROM feeds;


-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY (CASE WHEN last_fetched_at IS NULL THEN 1 ELSE 0 END) DESC, 
         last_fetched_at ASC LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds SET last_fetched_at = $1 WHERE id = $2 RETURNING *;
