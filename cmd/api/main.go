package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer pool.Close()

	h := &Handler{db: pool}
	// create a new ServeMux (router) and register the handler functions for each endpoint
	mux := http.NewServeMux()
	mux.HandleFunc("GET /events", h.handleEvents)
	mux.HandleFunc("GET /events/region/{region_code}", h.handleEventsByRegion)
	mux.HandleFunc("GET /stats", h.handleStats)

	// wrap the mux with CORS middleware to allow cross-origin requests from the frontend
	cors := corsMiddleware(mux)

	fmt.Println("Server running on :8000")
	http.ListenAndServe(":8000", cors)
}
