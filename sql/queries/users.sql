-- name: CreateUser :one
INSERT INTO Users (id, CreatedAt, UpdatedAt, name)
VALUES ($1, $2, $3, $4)
RETURNING *;