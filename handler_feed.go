package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/i-m-afk/rss/internal/database"
)

type feedBody struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (cfg *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	var data feedBody
	if err := json.Unmarshal(body, &data); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	handler := func(w http.ResponseWriter, r *http.Request, u database.User) {
		feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      data.Name,
			Url:       data.Url,
			UserID:    uuid.NullUUID{UUID: u.ID, Valid: true},
		})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error creating feed")
			return
		}
		respondWithJSON(w, http.StatusCreated, cfg.databaseFeedToFeed(feed))
	}
	middleware := cfg.middlewareAuth(handler)
	middleware(w, r)
}
