# Decisions

One line per choice: what, and why. Newest at the bottom.

- Go and PostgreSQL for the poller and event store; no Redis in v1 because nothing needs a queue yet (see SPEC.md).
