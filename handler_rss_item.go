package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/i-m-afk/rss/internal/database"
)

func (conf *apiConfig) getRssItemHandler(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request, u database.User) {
		rssItems, err := conf.DB.GetRssItemsForUser(r.Context(), uuid.NullUUID{UUID: u.ID, Valid: true})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Something went wrong")
			return
		}
		respondWithJSON(w, http.StatusOK, conf.databaseRssItemsToRssItems(rssItems))
	}
	middleware := conf.middlewareAuth(handler)
	middleware(w, r)
}
