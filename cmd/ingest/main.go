package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	url := "https://data.petabencana.id/reports?timeperiod=604800"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		log.Fatal(err)
	}

	reports := apiResp.Result.Objects.Output.Geometries

	fmt.Printf("Total reports fetched: %d\n\n", len(reports))

	jakartaCount := 0
	for _, report := range reports {
		if report.Properties.Tags.InstanceRegionCode != "ID-JK" {
			continue
		}

		depth := "unknown"
		if report.Properties.ReportData.FloodDepth != nil {
			depth = fmt.Sprintf("%d cm", *report.Properties.ReportData.FloodDepth)
		}

		fmt.Printf("ID:    %s\n", report.Properties.Pkey)
		fmt.Printf("Time:  %s\n", report.Properties.CreatedAt)
		fmt.Printf("City:  %s\n", report.Properties.Tags.City)
		fmt.Printf("Depth: %s\n", depth)
		fmt.Printf("Loc:   %.5f, %.5f\n", report.Coordinates[0], report.Coordinates[1])
		fmt.Println("---")
		jakartaCount++
	}

	fmt.Printf("Jakarta reports: %d\n", jakartaCount)
}
