## Research Report

**Context**: Research focused on the MVP weather-retrieval foundation for a Go CLI using explicit coordinate flags and Open-Meteo as the initial provider. The goal was to strengthen scope boundaries, acceptance criteria, and edge cases for the E001 product specification.

## Open-Meteo request baseline
- **Key findings**: Open-Meteo exposes a `/v1/forecast` HTTPS endpoint that accepts `latitude` and `longitude` plus selected weather parameters and returns JSON. The public docs explicitly note that additional optional URL parameters may be added in the future while avoiding new required parameters, which supports building a stable provider adapter around a small required request shape.
- **Recommended**: Specify an adapter that sends only the minimum fields needed for current-weather retrieval, validates coordinates before the outbound call, and normalizes provider responses before handing data to downstream CLI contracts.
- **Avoid**: Binding the CLI public behavior directly to raw provider payload shape or assuming optional provider fields will always be present.
### Sources
- https://open-meteo.com/en/docs — Official provider documentation for forecast/current-weather request shape, JSON responses, and API stability notes.

## Go CLI flag handling baseline
- **Key findings**: Go’s standard `flag` package provides built-in named-flag parsing, usage output, and `FlagSet` support for isolated command surfaces. It accepts one- or two-dash flag forms and stops parsing at the first non-flag argument or `--`, which reinforces the need for explicit invocation-shape validation in the CLI layer.
- **Recommended**: Keep the MVP interface narrow with explicit named latitude and longitude flags, deterministic usage behavior, and validation that rejects missing required coordinates before any provider call.
- **Avoid**: Ambiguous positional inputs or mixed invocation styles that weaken discoverability and automation predictability.
### Sources
- https://pkg.go.dev/flag — Official Go package documentation for command-line flag parsing and flag syntax behavior.

### Summary
The spec should treat explicit coordinate flags, pre-request validation, and provider-response normalization as the core MVP flow. It should also preserve a clean boundary between Open-Meteo-specific request/response handling and the CLI-owned behavior that later epics will turn into stable public success and error contracts.

### Sources Index
| URL | Topic | Fetched |
|-----|-------|---------|
| https://open-meteo.com/en/docs | Open-Meteo request baseline | 2026-04-02 |
| https://pkg.go.dev/flag | Go CLI flag handling baseline | 2026-04-02 |
