# Campaign Events Tracking Backend

A small Go HTTP service for ingesting campaign events and reading deduplicated campaign statistics.

## Stack and structure

The project uses Go, `net/http`, `encoding/json`, and an in-memory synchronized store. It has no third-party dependencies.

```text
models.go       event validation and API models
store.go        synchronized idempotent storage and aggregation
handlers.go     HTTP routing and JSON responses
main.go         server entry point
main_test.go    API and concurrency tests
starter/        representative seed data
debugging/      fixed JSONL debugging exercise
NOTES.md        decisions and scale/marketing analysis
BUGS.md         debugging findings
AI_USAGE.md     honest tool-use record
```

## Prerequisites and installation

Install Go 1.22 or newer and ensure `go` is on `PATH`. From the project root:

```text
go test ./...
go run .
```

The server listens on `http://localhost:8080`.

## API

### `POST /events`

Accepts a JSON array. The batch is atomic. Identical replays are counted as duplicates; conflicting reuse of an ID returns `409`.

```json
[
  {"event_id":"evt_00042","campaign_id":"cmp_summer_sale","contact_id":"ct_017","type":"opened","timestamp":"2026-08-10T09:15:04Z","metadata":{"url":"https://example.com/offer"}}
]
```

Successful response (`201`): `{"accepted":1,"duplicates":0}`. Invalid JSON or events return `400`; conflicting IDs return `409`.

### `GET /campaigns/{campaign_id}/stats`

Example: `GET /campaigns/cmp_summer_sale/stats`

```json
{"campaign_id":"cmp_summer_sale","sent":1,"delivered":1,"opened":1,"clicked":1}
```

### `GET /campaigns/{campaign_id}/events`

Optional listing included. Use `limit` (1-1000, default 100) and `offset` (default 0), for example `/campaigns/cmp_summer_sale/events?limit=20&offset=0`. Results are sorted by timestamp ascending.

## Tests and seed data

Run `go test ./...` and `go test -race ./...`. The seed file is `starter/seed/events.json`; post its array to `/events` with a JSON-capable client.

The debugging program is run from its directory:

```text
cd debugging
go run -race . events.jsonl
go run . events.jsonl
```

Both commands should match `expected_output.txt` exactly.

## Verification note

The implementation was written in this empty workspace, but Go was not installed in the execution environment, so the commands above could not be run here. `AI_USAGE.md` and `NOTES.md` record this limitation precisely.