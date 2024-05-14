package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/i-m-afk/rss/internal/database"
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
	Name      string    `json:"name"`
	ID        uuid.UUID `json:"id"`
}

type Feed struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	Name      string    `json:"name"`
	UserID    uuid.UUID `json:"user_id"`
	ID        uuid.UUID `json:"id"`
}

func (cfg *apiConfig) databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

func (cfg *apiConfig) databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Name:      feed.Name,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.CreatedAt,
		Url:       feed.Url,
		UserID:    feed.UserID.UUID,
	}
}
