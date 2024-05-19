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
	mux.HandleFunc("GET /v1/users", apiConf.getUserHandler)
	mux.HandleFunc("POST /v1/feeds", apiConf.createFeedHandler)
	mux.HandleFunc("GET /v1/feeds", apiConf.getAllFeedsHandler)
	mux.HandleFunc("POST /v1/feed_follows", apiConf.createFeedFollowHandler)
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiConf.deleteFeedFollowHandler)
	mux.HandleFunc("GET /v1/feed_follows", apiConf.getAllFeedFollowsForUserHandler)

	done := make(chan bool)
	go func() {
		if err := initServer(mux, port); err != nil {
			log.Panic(err)
		}
	}()
	// prevent the main goroutine from exiting
	<-done
}

func initServer(mux *http.ServeMux, port string) error {
	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}

	log.Printf("Server started on port %s", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
		return err
	}
	defer server.Shutdown(context.TODO())
	return nil
}
