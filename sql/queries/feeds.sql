-- name: CreateFeeds :one
INSERT INTO Feeds (id, CreatedAt, UpdatedAt, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
