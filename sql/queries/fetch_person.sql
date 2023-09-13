-- name: FetchPerson :one
SELECT * FROM persons WHERE id=$1;