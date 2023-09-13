-- name: CreatePerson :one
INSERT INTO persons (id, name)
VALUES ($1, $2)
RETURNING *;