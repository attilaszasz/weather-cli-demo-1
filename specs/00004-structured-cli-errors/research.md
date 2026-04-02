# Research: Structured CLI errors

## Topic: Canonical error contract
- Recommendation: define a CLI-owned error response type in `src/internal/contract` with explicit category, message, and exit code metadata controlled by the command layer.
- Why it matters: this keeps stderr stable for scripts and avoids leaking raw provider or internal error details directly.
- Source 1: Failure-path requirements in `specs/project-plan.md` E003.
- Source 2: ADR-003 and failure-path guidance in `specs/sad.md`.

## Topic: Error classification strategy
- Recommendation: classify failures into validation, downstream transport, downstream provider, and internal categories by mapping known command and provider error surfaces into a small deterministic set.
- Why it matters: stable exit codes and machine-readable categories are the public automation contract for failures.
- Source 1: validation package errors in `src/internal/validation/*`.
- Source 2: provider failure surfaces in `src/internal/provider/openmeteo/client.go`.

## Topic: Failure-path verification
- Recommendation: extend command tests to assert stdout stays clean, stderr contains valid JSON, and error categories/exit codes are deterministic for validation and provider failures.
- Why it matters: E003 acceptance requires contract verification for failure paths, not just raw error returns.
- Source 1: current command tests in `src/cmd/weather/main_test.go`.
- Source 2: main entrypoint behavior in `src/cmd/weather/main.go`.
