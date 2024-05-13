package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/i-m-afk/rss/internal/database"
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ID        uuid.UUID `json:"id"`
}

func (cfg *apiConfig) databaseUserToUser(user database.User) User {
	t := time.Time{}
	if user.UpdatedAt.Valid {
		t = user.UpdatedAt.Time
	}
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: t,
		UpdatedAt: t,
	}
}
