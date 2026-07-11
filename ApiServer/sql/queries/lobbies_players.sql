-- name: JoinLobby :exec
INSERT INTO lobbies_players (lobby_id, player_id)
VALUES (
  $1,
  $2
);

-- name: LeaveLobby :exec
DELETE FROM lobbies_players
WHERE lobby_id = $1
AND player_id = $2;

-- name: GetLobbyPlayers :many
SELECT users.*
FROM users
JOIN lobbies_players
ON users.id = lobbies_players.player_id
WHERE lobbies_players.lobby_id = $1;

-- name: GetPlayerLobbies :many
SELECT lobbies.*
FROM lobbies
JOIN lobbies_players
ON lobbies.id = lobbies_players.lobby_id
WHERE lobbies_players.player_id = $1;
