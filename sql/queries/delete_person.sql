-- name: DeletePerson :one
DELETE FROM persons
WHERE id=$1
RETURNING *;