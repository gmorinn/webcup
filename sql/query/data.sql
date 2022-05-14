-- name: CreateData :one
INSERT INTO data (title, description, img, category, user_id) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateData :exec
UPDATE data
SET title = $1,
    description = $2,
    img = $3,
    category = $4,
    updated_at = NOW()
WHERE id = $5;

-- name: GetDataByID :one
SELECT * FROM data
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetDataByUserID :many
SELECT * FROM data
WHERE user_id = $1
AND deleted_at IS NULL;

-- name: ListDataMostRecent :many
SELECT * FROM data
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountData :one
SELECT COUNT(*) FROM data
WHERE deleted_at IS NULL;
