-- +goose Up
CREATE TABLE lobbies (
  id UUID PRIMARY KEY,

  lobby_name TEXT NOT NULL UNIQUE,
  game_master UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

  short_description TEXT,
  game_map TEXT,

  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE lobbies;
