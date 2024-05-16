package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/i-m-afk/rss/internal/database"
)

type ffReqBody struct {
	FeedId uuid.UUID `json:"feed_id"`
}

func (conf *apiConfig) createFeedFollowHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	var data ffReqBody
	if err := json.Unmarshal(body, &data); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request structure")
		return
	}

	handler := func(w http.ResponseWriter, r *http.Request, u database.User) {
		ff, err := conf.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
			ID:        uuid.New(),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
			FeedID:    uuid.NullUUID{UUID: data.FeedId, Valid: true},
			UserID:    uuid.NullUUID{UUID: u.ID, Valid: true},
		})
		if err != nil {
			// catching duplicate key error (there can be better ways)
			if err.Error()[:23] == "pq: duplicate key value" {
				respondWithError(w, http.StatusConflict, "Feed already followed")
				return
			}
			respondWithError(w, http.StatusInternalServerError, "Something went wrong")
			return
		}
		respondWithJSON(w, http.StatusCreated, conf.databaseFeedFollowToFeedFollow(ff))
	}
	middleware := conf.middlewareAuth(handler)
	middleware(w, r)
}

func (conf *apiConfig) deleteFeedFollowHandler(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request, u database.User) {
		ffid, err := uuid.Parse(r.URL.Path[len("/v1/feed_follows/feedFollowID="):])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid feedFollowID")
			return
		}
		_, err = conf.DB.DeleteFeedFollows(r.Context(), ffid)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				respondWithError(w, http.StatusNotFound, "Feed follow not found")
			default:
				respondWithError(w, http.StatusInternalServerError, "Something went wrong")
			}
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Feed follow deleted"})
	}
	middleware := conf.middlewareAuth(handler)
	middleware(w, r)
}

func (conf *apiConfig) getAllFeedFollowsForUserHandler(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request, u database.User) {
		feedfollows, err := conf.DB.GetFeedFollowsByUserID(r.Context(), uuid.NullUUID{UUID: u.ID, Valid: true})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Something went wrong")
			return
		}
		respondWithJSON(w, http.StatusOK, conf.databaseFeedsFollowsToFeedsFollows(feedfollows))
	}
	middleware := conf.middlewareAuth(handler)
	middleware(w, r)
}
