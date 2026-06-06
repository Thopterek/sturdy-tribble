-- +goose Up
CREATE TABLE users (
  id UUID PRIMARY KEY,

  email TEXT NOT NULL UNIQUE,
  google_id TEXT UNIQUE NOT NULL,

  name TEXT NOT NULL,
  avatar_url TEXT,

  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;
