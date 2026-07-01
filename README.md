# Jakarta Flood History

A full-stack web app that shows the historical flooding record for areas in DKI Jakarta Province. Users search or select an area on a map and see its flood history — when it flooded, how often, and at what depth.

## Architecture

```
PetaBencana API
      │
      ▼
  Go Ingester (cmd/ingest)
      │
      ▼
 PostgreSQL DB
      │
      ▼
  Airflow DAG (daily schedule)
      │
      ▼
  Go REST API (cmd/api)
      │
      ▼
 Vue + Leaflet Frontend
```

### Stack

| Layer | Technology |
|---|---|
| Ingestion | Go |
| Storage | PostgreSQL |
| Orchestration | Apache Airflow |
| Backend API | Go (chi + pgx) |
| Frontend | Vue 3 + Leaflet |
| AI enrichment | LLM-based location normalization |

## Project Structure

```
flood/
├── cmd/
│   ├── ingest/       # Ingestion binary — pulls PetaBencana API into DB
│   └── api/          # REST API server
├── internal/
│   ├── db/           # Shared DB connection and query helpers
│   └── models/       # Shared data structs
├── airflow/
│   └── dags/         # Airflow DAG definitions
├── frontend/         # Vue + Leaflet app
├── db/
│   └── migrations/   # SQL schema files
└── ai/               # LLM location normalization enrichment
```

## Data Source

Primary: [PetaBencana.id](https://petabencana.id) — a public flood reporting platform for Indonesian cities. Reports include flood depth (cm), coordinates, city/region, and timestamp.

## Status

Work in progress — Week 1 of 4.
