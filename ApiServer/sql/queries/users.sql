-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUserIDByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UserUpdatePassword :exec
UPDATE users
SET
  updated_at = NOW(),
  hashed_password = $2
WHERE id = $1;

-- name: UserUpgradeToChirpRed :exec
UPDATE users
SET
  updated_at = NOW(),
  is_chirpy_red = TRUE
WHERE id = $1;
