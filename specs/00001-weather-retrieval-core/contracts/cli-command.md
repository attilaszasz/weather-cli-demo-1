# CLI Command Contract

## Command Surface

| Interface | Syntax | Purpose | Output Boundary |
|-----------|--------|---------|-----------------|
| Primary command | `weather --latitude <float> --longitude <float>` | Execute the current-weather retrieval path | Internal normalized weather data returned to the command layer for downstream formatting |
| Help command | `weather --help` | Display baseline usage and explicit coordinate flags | Usage text only |

## Inputs

| Name | Type | Required | Rules |
|------|------|----------|-------|
| `--latitude` | float | yes | Range `-90` to `90` |
| `--longitude` | float | yes | Range `-180` to `180` |
| extra positional args | list | no | Must be empty for valid execution |

## Internal Success Payload

| Field | Type | Source | Notes |
|-------|------|--------|-------|
| `temperature` | float | Open-Meteo current temperature field | Internal normalized field only |
| `windSpeed` | float | Open-Meteo wind speed field | Internal normalized field only |
| `weatherCode` | int | Open-Meteo weather code field | Internal normalized field only |
| `observationTime` | string | Open-Meteo current time field | Preserve provider timestamp semantics for later contract formatting |

## Outbound Provider Contract

| Aspect | Value |
|--------|-------|
| Method | `GET` |
| Endpoint | `https://api.open-meteo.com/v1/forecast` |
| Required query params | `latitude`, `longitude`, minimum current-weather field selection |
| Timeout | `3s` |
| Retries | `none` |
| Transport | HTTPS only |

## Validation Contract

| Condition | Expected Behavior |
|-----------|-------------------|
| Missing latitude | Reject before request construction |
| Missing longitude | Reject before request construction |
| Latitude out of range | Reject before provider call |
| Longitude out of range | Reject before provider call |
| Extra positional args | Reject as invalid invocation shape |
