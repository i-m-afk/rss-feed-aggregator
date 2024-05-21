package main

import (
	"net/http"
	"strconv"

	"github.com/i-m-afk/rss/internal/database"
)

func (cfg *apiConfig) getPostByUser(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}
	l, err := strconv.Atoi(limit)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid limit")
	}

	handler := func(w http.ResponseWriter, r *http.Request, user database.User) {
		posts, err := cfg.DB.GetPosts(r.Context(), int32(l))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Something went wrong")
			return
		}
		respondWithJSON(w, http.StatusOK, cfg.databasePostsToPosts(posts))
	}
	middleware := cfg.middlewareAuth(handler)
	middleware(w, r)
}
