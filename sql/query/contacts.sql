-- name: GetMessagesFiltered :many
SELECT * FROM contacts
WHERE deleted_at IS NULL
ORDER BY
  CASE WHEN sqlc.arg('email_asc')::bool THEN email END asc,
  CASE WHEN sqlc.arg('email_desc')::bool THEN email END desc
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetMessageByID :one
SELECT * FROM contacts
WHERE deleted_at IS NULL
AND id = $1;

-- name: DeleteMessageByID :exec
UPDATE
    contacts
SET
    deleted_at = NOW()
WHERE 
    id = $1;

-- name: InsertMessage :exec
INSERT INTO contacts (user_id, msg, email) 
VALUES ($1, $2, $3);

-- name: CountMessage :one
SELECT COUNT(*) FROM contacts
WHERE deleted_at IS NULL;
