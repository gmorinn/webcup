-- name: UpdateUserBo :exec
UPDATE users
SET firstname = $1,
    lastname = $2,
    username = $3,
    email = $4,
    avatar = $5,
    role = $6,
    updated_at = NOW()
WHERE id = $7;

-- name: GetBoAllUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL
ORDER BY
  CASE WHEN sqlc.arg('firstname_asc')::bool THEN firstname END asc,
  CASE WHEN sqlc.arg('firstname_desc')::bool THEN firstname END desc,
  CASE WHEN sqlc.arg('username_asc')::bool THEN username END asc,
  CASE WHEN sqlc.arg('username_desc')::bool THEN username END desc,
  CASE WHEN sqlc.arg('lastname_asc')::bool THEN lastname END asc,
  CASE WHEN sqlc.arg('lastname_desc')::bool THEN lastname END desc,
  CASE WHEN sqlc.arg('email_asc')::bool THEN email END asc,
  CASE WHEN sqlc.arg('email_desc')::bool THEN email END desc,
  CASE WHEN sqlc.arg('role_asc')::bool THEN role END asc,
  CASE WHEN sqlc.arg('role_desc')::bool THEN role END desc
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
