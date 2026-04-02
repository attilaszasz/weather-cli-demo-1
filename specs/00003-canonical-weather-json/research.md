# Research: Canonical weather JSON

## Topic: Canonical CLI success contract
- Recommendation: define a CLI-owned success contract in a dedicated internal contract package and emit it with the Go standard library JSON encoder.
- Why it matters: this decouples the public automation surface from provider-specific structs and enforces stable stdout behavior.
- Source 1: ADR-003 in `specs/sad.md`.
- Source 2: Existing success output path in `src/cmd/weather/run.go`.

## Topic: Stable field mapping
- Recommendation: map internal normalized weather data into a canonical payload with explicit JSON field names and avoid exposing provider package types directly from the command layer.
- Why it matters: this keeps downstream scripts insulated from future provider changes and makes contract tests deterministic.
- Source 1: `src/internal/provider/provider.go` current internal weather model.
- Source 2: `specs/project-plan.md` E002 acceptance criteria.

## Topic: Success-path verification
- Recommendation: add command tests that assert stdout is valid JSON, stderr stays empty on success, and the payload fields match the expected canonical schema.
- Why it matters: E002 folds contract quality into the feature itself, so parseability and mapping stability must be verified directly.
- Source 1: existing command tests in `src/cmd/weather/main_test.go`.
- Source 2: project instruction requiring test-backed delivery.
