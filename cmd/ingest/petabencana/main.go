package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type Tags struct {
	City               string  `json:"city"`
	RegionCode         string  `json:"region_code"`
	LocalAreaID        *string `json:"local_area_id"`
	InstanceRegionCode string  `json:"instance_region_code"`
}

type ReportData struct {
	FloodDepth *int `json:"flood_depth"`
}

type Properties struct {
	Pkey       string     `json:"pkey"`
	CreatedAt  string     `json:"created_at"`
	Status     string     `json:"status"`
	ReportData ReportData `json:"report_data"`
	Tags       Tags       `json:"tags"`
	Text       *string    `json:"text"`
}

type Geometry struct {
	Properties  Properties `json:"properties"`
	Coordinates []float64  `json:"coordinates"`
}

type APIResponse struct {
	StatusCode int `json:"statusCode"`
	Result     struct {
		Objects struct {
			Output struct {
				Geometries []Geometry `json:"geometries"`
			} `json:"output"`
		} `json:"objects"`
	} `json:"result"`
}

func main() {
	var ctx context.Context = context.Background()

	var dbURL string = os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgresql://flood:flood@localhost:5434/flood"
	}

	var conn, err = pgx.Connect(ctx, dbURL)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer conn.Close(ctx)

	var url string = "https://data.petabencana.id/reports?timeperiod=18748800"
	var resp, httpErr = http.Get(url)
	if httpErr != nil {
		fmt.Println("Error fetching API:", httpErr)
		return
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var inserted int = 0
	for _, report := range apiResp.Result.Objects.Output.Geometries {
		if report.Properties.Tags.InstanceRegionCode != "ID-JK" {
			continue
		}
		if insertReport(ctx, conn, report) {
			inserted++
		}
	}

	fmt.Printf("Done. Inserted %d Jakarta reports.\n", inserted)
}

func insertReport(ctx context.Context, conn *pgx.Conn, report Geometry) bool {
	_, err := conn.Exec(
		ctx,
		`INSERT INTO flood_reports (
			source, source_id, created_at, status,
			flood_depth, city, region_code, local_area_id,
			longitude, latitude, raw_text
		) VALUES (
			$1, $2, $3, $4,
			$5, $6, $7, $8,
			$9, $10, $11
		) ON CONFLICT (source, source_id) DO NOTHING`,
		"petabencana",
		report.Properties.Pkey,
		report.Properties.CreatedAt,
		report.Properties.Status,
		report.Properties.ReportData.FloodDepth,
		report.Properties.Tags.City,
		report.Properties.Tags.RegionCode,
		report.Properties.Tags.LocalAreaID,
		report.Coordinates[0],
		report.Coordinates[1],
		report.Properties.Text,
	)
	if err != nil {
		fmt.Println("Error inserting report:", report.Properties.Pkey, err)
		return false
	}

	return true
}
