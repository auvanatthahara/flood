package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

// FloodReport is the DTO — defines what the API returns as JSON
// *int and *string are pointers, meaning the field can be nil (like Java's Integer vs int)
// the backtick tags tell the JSON encoder what key name to use (like @JsonProperty)
type FloodReport struct {
	ID          string    `json:"id"`
	Source      string    `json:"source"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string  `json:"status"`
	FloodDepth  *int    `json:"flood_depth"`
	City        string  `json:"city"`
	RegionCode  string  `json:"region_code"`
	LocalAreaID *string `json:"local_area_id"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	RawText     *string `json:"raw_text"`
}

type RegionCount struct {
	RegionCode string `json:"region_code"`
	Count      int    `json:"count"`
}

type Stats struct {
	TotalReports   int           `json:"total_reports"`
	EarliestReport time.Time     `json:"earliest_report"`
	LatestReport   time.Time     `json:"latest_report"`
	TopRegions     []RegionCount `json:"top_regions"`
}

// Handler holds shared dependencies (the DB connection) for all endpoint functions
// equivalent to a Spring @Controller that has a @Repository injected into it
type Handler struct {
	db *pgx.Conn
}

// handleEvents handles GET /events — returns all flood reports ordered by date
// w is where you write the response, r is the incoming request
func (h *Handler) handleEvents(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query(r.Context(), `
		SELECT id, source, created_at, status, flood_depth,
		       city, region_code, local_area_id, longitude, latitude, raw_text
		FROM flood_reports
		ORDER BY created_at DESC
	`)
	if err != nil {
		// http.Error writes the message + status code and returns — like throwing an exception
		http.Error(w, "Failed to query database", http.StatusInternalServerError)
		return
	}
	// always close the result set when done, same as closing a JDBC ResultSet
	defer rows.Close()

	var reports []FloodReport
	for rows.Next() {
		var rep FloodReport
		// Scan reads one row into the struct fields — order must match the SELECT columns
		// & means "address of" — you pass a pointer so Scan can write into the variable
		err := rows.Scan(
			&rep.ID, &rep.Source, &rep.CreatedAt, &rep.Status, &rep.FloodDepth,
			&rep.City, &rep.RegionCode, &rep.LocalAreaID, &rep.Longitude, &rep.Latitude,
			&rep.RawText,
		)
		if err != nil {
			http.Error(w, "Failed to read row", http.StatusInternalServerError)
			return
		}
		// append adds to the slice — equivalent to List.add() in Java
		reports = append(reports, rep)
	}

	w.Header().Set("Content-Type", "application/json")
	// serialize the slice to JSON and write it directly to the response
	json.NewEncoder(w).Encode(reports)
}

func (h *Handler) handleEventsByRegion(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query(r.Context(), `
		SELECT id, source, created_at, status, flood_depth,
		       city, region_code, local_area_id, longitude, latitude, raw_text
		FROM flood_reports
		WHERE region_code = $1
		ORDER BY created_at DESC
	`, r.PathValue("region_code"))
	if err != nil {
		http.Error(w, "Failed to query database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reports []FloodReport
	for rows.Next() {
		var rep FloodReport
		err := rows.Scan(
			&rep.ID, &rep.Source, &rep.CreatedAt, &rep.Status, &rep.FloodDepth,
			&rep.City, &rep.RegionCode, &rep.LocalAreaID, &rep.Longitude, &rep.Latitude,
			&rep.RawText,
		)
		if err != nil {
			http.Error(w, "Failed to read row", http.StatusInternalServerError)
			return
		}
		reports = append(reports, rep)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}

func (h *Handler) handleStats(w http.ResponseWriter, r *http.Request) {
	var stats Stats

	// query 1: total count and date range — returns a single row
	err := h.db.QueryRow(r.Context(), `
		SELECT COUNT(*), MIN(created_at), MAX(created_at)
		FROM flood_reports
	`).Scan(&stats.TotalReports, &stats.EarliestReport, &stats.LatestReport)
	if err != nil {
		http.Error(w, "Failed to query stats", http.StatusInternalServerError)
		return
	}

	// query 2: top 5 regions by report count
	rows, err := h.db.Query(r.Context(), `
		SELECT region_code, COUNT(*) as count
		FROM flood_reports
		GROUP BY region_code
		ORDER BY count DESC
		LIMIT 5
	`)
	if err != nil {
		http.Error(w, "Failed to query top regions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rc RegionCount
		if err := rows.Scan(&rc.RegionCode, &rc.Count); err != nil {
			http.Error(w, "Failed to read row", http.StatusInternalServerError)
			return
		}
		stats.TopRegions = append(stats.TopRegions, rc)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
