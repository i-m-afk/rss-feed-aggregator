-- +goose Up
CREATE TABLE feeds(
  id UUID PRIMARY KEY,
  created_at timestamp not null default Now(),
  updated_at timestamp not null default Now(),
  name Varchar(64) not null,
  user_id  uuid REFERENCES users(id) on delete cascade,
  url Varchar(256) not null unique
);

-- +goose Down
DROP table feeds;
