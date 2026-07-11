-- name: CreateUser :one
INSERT INTO users (id, email, google_id, username, hashed_password, avatar_url, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    null,
    $2,
    $3,
    null,
    NOW(),
    NOW()
)
RETURNING *;

-- name: CreateGoogleUser :one
INSERT INTO users (id, email, google_id, username, avatar_url, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    NOW(),
    NOW()
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
