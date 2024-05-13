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
	currTime := sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		CreatedAt: currTime,
		ID:        uuid.New(),
		UpdatedAt: currTime,
		Name:      name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}
	respondWithJSON(w, http.StatusCreated, cfg.databaseUserToUser(user))
}
