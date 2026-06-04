-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2
)
RETURNING *;

-- name: DeleteAllChirps :exec
DELETE FROM chirps;

-- name: GetChirps :many
SELECT id, created_at, updated_at, body, user_id
FROM chirps
WHERE ($1 = '00000000-0000-0000-0000-000000000000'::uuid OR user_id = $1)
ORDER BY created_at;

-- name: GetSingleChirp :one
SELECT id, created_at, updated_at, body, user_id
  FROM chirps
  WHERE id = $1;

-- name: DeleteSingleChirp :exec
DELETE FROM chirps
  WHERE id = $1;
