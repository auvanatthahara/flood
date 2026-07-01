CREATE TABLE flood_reports (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source        TEXT NOT NULL,
    source_id     TEXT,
    created_at    TIMESTAMPTZ NOT NULL,
    status        TEXT,
    flood_depth   INT,
    city          TEXT,
    region_code   TEXT,
    local_area_id TEXT,
    longitude     DOUBLE PRECISION,
    latitude      DOUBLE PRECISION,
    raw_text      TEXT,
    ingested_at   TIMESTAMPTZ DEFAULT NOW(),

    UNIQUE (source, source_id)
);
