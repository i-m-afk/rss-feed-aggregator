// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	UserID        uuid.NullUUID
	Url           string
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        uuid.UUID
	FeedID    uuid.NullUUID
	UserID    uuid.NullUUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt sql.NullString
	FeedID      uuid.NullUUID
}

type RssItem struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
	Title       string
	Url         string
	Author      sql.NullString
	Description sql.NullString
	PublishedAt sql.NullString
	PostID      uuid.NullUUID
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
