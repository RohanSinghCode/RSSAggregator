-- name: CreateUser :one
INSERT INTO Users (id, CreatedAt, UpdatedAt, name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM Users WHERE api_key = $1;