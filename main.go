package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/i-m-afk/rss/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env must be set")
	}
	// database connection string
	dbUrl := os.Getenv("CONN")
	if dbUrl == "" {
		log.Fatal("CONN env must be set")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiConf := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", readinessHandler)
	mux.HandleFunc("GET /v1/err", errHandler)
	mux.HandleFunc("POST /v1/users", apiConf.createUserHandler)

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
	defer server.Shutdown(context.TODO())

	log.Printf("Server started on port %s", port)
}
