# QC Report: Structured CLI errors

**Feature**: `specs/00004-structured-cli-errors/spec.md`  
**Plan**: `specs/00004-structured-cli-errors/plan.md`  
**Tasks**: `specs/00004-structured-cli-errors/tasks.md`  
**Status**: Passed
**Date**: 2026-04-02

## Summary

The Weather CLI now emits canonical structured error JSON on stderr with deterministic exit code mapping for validation, downstream transport/provider, and internal failures. Failure-path tests verify clean stdout separation and stable category/exit-code behavior.

## Checks Run

| Check | Result | Evidence |
|---|---|---|
| `go test ./...` | PASS | command and package tests passed across the repository |
| `go build ./...` | PASS | repository builds successfully |
| `go test -coverprofile coverage.out ./...` | PASS | command package coverage reached 82.9% and the repository coverage command completed successfully |

## Artifacts Verified

- `src/internal/contract/error.go`
- `src/cmd/weather/main.go`
- `src/cmd/weather/run.go`
- `src/cmd/weather/main_test.go`
- `specs/00004-structured-cli-errors/spec.md`
- `specs/00004-structured-cli-errors/plan.md`
- `specs/00004-structured-cli-errors/tasks.md`

## Exit Condition

QC passed for E003.
