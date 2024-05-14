package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/i-m-afk/rss/internal/database"
	"github.com/i-m-afk/rss/internal/utils"
)

type userBody struct {
	Name string `json:"name"`
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	var data userBody
	if err := json.Unmarshal(body, &data); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	name := data.Name
	if name == "" {
		respondWithError(w, http.StatusBadRequest, "Name is required")
		return
	}
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		CreatedAt: time.Now(),
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		Name:      name,
		ApiKey:    utils.GenSha(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}
	respondWithJSON(w, http.StatusCreated, cfg.databaseUserToUser(user))
}

func (cfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request, user database.User) {
		respondWithJSON(w, http.StatusOK, cfg.databaseUserToUser(user))
	}
	middleware := cfg.middlewareAuth(handler)
	middleware(w, r)
}
