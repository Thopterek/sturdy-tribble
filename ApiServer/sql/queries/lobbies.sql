-- name: CreateLobby :one
INSERT INTO lobbies (id, lobby_name, game_master, short_description, game_map, created_at, updated_at)
VALUES (
  gen_random_uuid(),
  $1,
  $2,
  $3,
  NULL,
  NOW(),
  NOW()
)
RETURNING *;

-- name: DeleteAllLobbies :exec
DELETE FROM lobbies;

-- name: GetLobbyIDByName :one
SELECT * FROM lobbies
WHERE lobby_name = $1;

-- name: GetLobbyById :one
SELECT * FROM lobbies
WHERE id = $1;

-- name: DeleteLobbyByName :exec
DELETE FROM lobbies
WHERE lobby_name = $1;
