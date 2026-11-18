# Ready-to-run examples

[README](../README.md) · [Methods](../docs/endpoints.md) · [Response example](../docs/responses.md)

Each example makes one data call. Set `UPTIMESCOUT_LICENSE_KEY` privately first. Clone this repository and run commands from its root. `--help` makes no API call.

| Example | Run from repository root |
| --- | --- |
| [Search the service catalog](services/main.go) | `go run ./examples/services --region GB --lang en` |
| [Read one regional signal](status/main.go) | `go run ./examples/status --region GB --lang en` |
| [Check several services in one call](statuses/main.go) | `go run ./examples/statuses --region GB --lang en` |
| [Summarize selected services](summary/main.go) | `go run ./examples/summary --region GB --lang en` |
| [Read observation history](history/main.go) | `go run ./examples/history --region GB --lang en` |
| [Inspect observation statistics](stats/main.go) | `go run ./examples/stats --region GB --lang en` |
| [Find fresh elevated signals](incidents/main.go) | `go run ./examples/incidents --region GB --lang en` |

For a different service, add `--service cloudflare` to a single-service example. Batch and summary examples accept `--services github,cloudflare`. History examples accept `--hours 24`, `--limit 50` and `--offset 0`.

Read [the response example](../docs/responses.md) alongside each request. The optional `UPTIMESCOUT_BASE_URL` variable is for a trusted development/mock server; do not send a real key to an untrusted host.

## Responses for every method

Public demonstration snapshots captured on **2026-11-18**. Each file includes its exact request arguments and complete response. These historical signals are examples, not current status claims. They contain no license key or customer usage fields.

| Method | Request and response JSON |
| --- | --- |
| `services` | [services.json](responses/services.json) |
| `status` | [status.json](responses/status.json) |
| `statuses` | [statuses.json](responses/statuses.json) |
| `summary` | [summary.json](responses/summary.json) |
| `history` | [history.json](responses/history.json) |
| `stats` | [stats.json](responses/stats.json) |
| `incidents` | [incidents.json](responses/incidents.json) |
