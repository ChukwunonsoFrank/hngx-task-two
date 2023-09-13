-- name: UpdatePerson :one
UPDATE persons
SET name = $1
WHERE id = $2
RETURNING *;