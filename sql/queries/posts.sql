-- name: CreatePosts :one
INSERT INTO Posts
(
    id,
    created_at,
    updated_at,
    title,
    description,
    published_at,
    url,
    feed_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: GetPostsForUser :many
SELECT P.* 
FROM Posts P
JOIN Feed_Follows FF
ON FF.feed_id = P.feed_id
WHERE FF.user_id = $1
ORDER BY P.published_at DESC
LIMIT $2;