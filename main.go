package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env must be set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", readinessHandler)
	mux.HandleFunc("GET /v1/err", errHandler)

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
