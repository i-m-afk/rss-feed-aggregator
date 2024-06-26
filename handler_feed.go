package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/i-m-afk/rss/internal/database"
	"github.com/lib/pq"
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
			if perr, ok := err.(*pq.Error); ok && perr.Code.Name() == "unique_violation" {
				respondWithError(w, http.StatusConflict, "Feed already exists")
				return
			}
			respondWithError(w, http.StatusInternalServerError, "Error creating feed")
			return
		}
		feedFollow, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
			ID:        uuid.New(),
			FeedID:    uuid.NullUUID{UUID: feed.ID, Valid: true},
			UserID:    uuid.NullUUID{UUID: u.ID, Valid: true},
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error creating feedfollow")
			return
		}
		respondWithJSON(w, http.StatusCreated, cfg.databaseFeedNFollowsToFeedNFollows(feed, feedFollow))
	}
	middleware := cfg.middlewareAuth(handler)
	middleware(w, r)
}

func (cfg *apiConfig) getAllFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrive feeds")
		return
	}
	respondWithJSON(w, http.StatusOK, cfg.databaseFeedsToFeeds(feeds))
}
