-- +goose Up
create table rss_items (
  id uuid primary key,
  created_at timestamp not null default now(),
  updated_at timestamp,
  title varchar(1024) not null,
  url varchar(256) not null unique,
  author varchar(256),
  description varchar(30000),
  published_at varchar(64),
  post_id uuid references posts(id) on delete cascade
);

-- +goose Down
drop table rss_items;
