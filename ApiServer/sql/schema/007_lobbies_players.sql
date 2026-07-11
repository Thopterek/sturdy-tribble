-- +goose Up
CREATE TABLE lobbies_players (
    lobby_id UUID NOT NULL REFERENCES lobbies(id) ON DELETE CASCADE,
    player_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (lobby_id, player_id)
);

-- +goose Down
DROP TABLE lobbies_players;
