-- +goose Up
CREATE TABLE feed_follow(
    ID uuid PRIMARY KEY,
    FEED_ID uuid REFERENCES feeds(id) on delete cascade,
    USER_ID uuid REFERENCES users(id) on delete cascade,
    CREATED_AT TIMESTAMP NOT NULL DEFAULT NOW(),
    UPDATED_AT TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_feed_user_subscription UNIQUE (FEED_ID, USER_ID)
);

-- +goose Down
DROP TABLE feed_follow;
