# Seat Forecast

Forecasts whether a closed Rutgers course section will open before add/drop ends, so a student can decide to wait or take a backup.

## Status

Planning. The spec is written; code starts with the poller and event logger.

## Planned stack

- Poller in Go that reads the public Schedule of Classes data on an interval and diffs section status
- Postgres event store holding one row per open/close transition, plus poll heartbeats
- Model in Go or a small Python script that computes open probability from historical rates by course, section and time in cycle
- API in Go that serves per-section probability and event history as JSON
- Plain HTML and JS web page that takes a section and shows probability and history
- Docker, GitHub Actions CI, and a small VM or free-tier host for deployment

## Docs

- [SPEC.md](SPEC.md)

## Not affiliated

Not affiliated with Rutgers University. This project sends no registrations and never logs in on a student's behalf.
