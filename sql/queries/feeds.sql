-- name: CreateFeeds :one
INSERT INTO Feeds (id, CreatedAt,  name, url, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM Feeds;

-- name: GetNextFeedToFetch :many
SELECT * FROM Feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE Feeds 
SET last_fetched_at = NOW(),
UpdatedAt = NOW()
WHERE id = $1
RETURNING *;

-- name: GetFollowedFeeds :many
SELECT FF.* 
FROM Feeds F 
JOIN Feed_Follows FF 
ON F.user_id = FF.user_id 
WHERE F.user_id = $1;

-- name: FollowFeed :one
INSERT INTO Feed_Follows(id, CreatedAt, user_id, feed_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows WHERE id = $1 AND user_id = $2;