-- name: UpdateDescriptionUser :exec
UPDATE users
SET firstname = $1,
    lastname = $2,
    username = $3,
    email = $4,
    updated_at = NOW()
WHERE id = $5;

-- name: UpdateStock :exec
UPDATE users
SET stock = $1,
    updated_at = NOW()
WHERE id = $2;

-- name: UpdateAvatarUser :exec
UPDATE users
SET avatar = $1,
    updated_at = NOW()
WHERE id = $2;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetEmailByUserID :one
SELECT email FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetStockByUserID :one
SELECT stock FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserRandom :one
SELECT * FROM users
WHERE deleted_at IS NULL
LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteUserByID :exec
UPDATE
    users
SET
    deleted_at = NOW()
WHERE 
    id = $1;

-- name: CountUser :one
SELECT COUNT(*) FROM users
WHERE deleted_at IS NULL;

-- name: ListUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL
AND (firstname ILIKE $1 OR lastname ILIKE $1 OR email ILIKE $1 OR username ILIKE $1)
LIMIT 5;

-- name: ListUsersMostRecent :many
SELECT * FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;