# Weather CLI

Weather CLI is a small Go command-line application that fetches current weather data for a latitude/longitude pair and prints a stable JSON response for automation-friendly use.

## What it does

- accepts coordinates through explicit flags
- validates input before making a network request
- fetches current weather from Open-Meteo
- prints canonical JSON to `stdout` on success
- prints structured error JSON to `stderr` on failure
- supports automated cross-platform releases through GitHub Actions and GoReleaser

## Requirements

- Go `1.24`

## Project structure

- `src/cmd/weather` — CLI entrypoint and command behavior
- `src/internal/provider/openmeteo` — Open-Meteo integration
- `src/internal/service` — weather service layer
- `src/internal/contract` — success and error JSON contracts
- `.github/workflows` — CI and release automation
- `.goreleaser.yaml` — release packaging configuration

## Run locally

From the repository root:

```powershell
go run ./src/cmd/weather --latitude 47.4979 --longitude 19.0402
```

## Usage

```text
Usage: weather --latitude <float> --longitude <float>

Options:
  --latitude float
  --longitude float
  --help
```

## Success output

On success, the CLI writes JSON to `stdout`.

Example:

```json
{
  "temperature": 20.5,
  "windSpeed": 4.1,
  "weatherCode": 2,
  "observationTime": "2026-04-02T10:00"
}
```

## Error output

On failure, the CLI writes structured JSON to `stderr` and exits with a non-zero exit code.

Example:

```json
{
  "category": "validation",
  "code": "invalid_input",
  "message": "latitude is required",
  "exitCode": 2
}
```

### Exit codes

- `0` — success
- `2` — validation / usage failure
- `3` — downstream transport or provider failure
- `4` — internal failure

## Development

Run the main validation commands from the repository root:

```powershell
go test ./...
go build ./...
```

Optional coverage command:

```powershell
go test -coverprofile coverage.out ./...
```

## Release automation

The repository includes:

- `.github/workflows/ci.yml` — runs test and build on pushes and pull requests
- `.github/workflows/release.yml` — runs on tags matching `v*`
- `.goreleaser.yaml` — builds Linux, macOS, and Windows artifacts plus checksums

To trigger a release workflow, push a semantic version tag such as:

```powershell
git tag v0.0.4
git push origin v0.0.4
```

## Notes

- success payloads are CLI-owned and intentionally independent from raw provider JSON
- failure payloads are separated onto `stderr` so scripts can safely parse `stdout`
- the application currently uses Open-Meteo as its weather provider
