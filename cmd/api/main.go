package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// context.Background() is the root context — required by pgx for DB operations
	ctx := context.Background()

	// connect to DB once and reuse the connection for all requests
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	// close the DB connection when main() exits (i.e. when the server shuts down)
	defer conn.Close(ctx)

	// manual dependency injection — pass the DB connection into the Handler
	h := &Handler{db: conn}

	// register routes: "METHOD /path" syntax requires Go 1.22+
	http.HandleFunc("GET /events", h.handleEvents)
	http.HandleFunc("GET /events/region/{region_code}", h.handleEventsByRegion)
	http.HandleFunc("GET /stats", h.handleStats)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
