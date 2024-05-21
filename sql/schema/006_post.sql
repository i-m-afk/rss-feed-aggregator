-- +goose Up
CREATE TABLE posts (
  id uuid PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  title Varchar(1024) not null, 
  url Varchar(256) not null unique, 
  description Varchar(30000),
  published_at Varchar(64),
  feed_id uuid References feeds(id) on delete cascade
);

-- +goose Down
DROP TABLE posts;
