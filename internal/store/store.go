// Package store reads and writes section events in Postgres.
package store

// Schema creates the events table. Run once against a fresh database.
const Schema = `
CREATE TABLE IF NOT EXISTS section_events (
	id            BIGSERIAL PRIMARY KEY,
	section_index TEXT        NOT NULL,
	course        TEXT        NOT NULL,
	status        TEXT        NOT NULL CHECK (status IN ('open', 'closed')),
	observed_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS section_events_section_idx
	ON section_events (section_index, observed_at);
`
