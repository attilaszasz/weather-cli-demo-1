# Task List: Weather retrieval core

**Feature**: `specs/00001-weather-retrieval-core/spec.md`  
**Plan**: `specs/00001-weather-retrieval-core/plan.md`

## Project Mode

- Greenfield

## Epic / Capability Map

- **US1**: Retrieve weather with explicit coordinates
- **US2**: Fail fast on invalid coordinate invocation
- **US3**: Maintain a swappable provider seam

## Phase 1: Foundational

- [X] T001 {FR-001,FR-006} Initialize Go module and command/provider/service package scaffolding in `go.mod`, `/src/cmd/weather/main.go`, `/src/internal/provider/provider.go`, and `/src/internal/service/weather.go`
- [X] T002 [P] {FR-008} Define internal normalized weather models and provider request/response types in `/src/internal/service/weather.go`, `/src/internal/provider/openmeteo/request.go`, and `/src/internal/provider/openmeteo/response.go`
- [X] T003 [P] {FR-007} Implement shared HTTP client configuration with HTTPS endpoint constants and 3-second no-retry transport policy in `/src/internal/provider/openmeteo/client.go`

## Phase 2: 🎯 MVP User Story 1 - Retrieve weather with explicit coordinates

- [X] T004 [US1] {FR-009} Add failing command and provider success-path tests in `/src/cmd/weather/main_test.go`, `/src/internal/provider/openmeteo/client_test.go`, and `/src/internal/service/weather_test.go`
- [X] T005 [US1] {FR-001,FR-002} Implement CLI entrypoint, explicit `--latitude`/`--longitude` flags, and baseline help output in `/src/cmd/weather/main.go` and `/src/cmd/weather/usage.go`
- [X] T006 [US1] {FR-006,FR-008} Implement weather service orchestration and normalized success-path mapping in `/src/internal/service/weather.go`
- [X] T007 [US1] {FR-006,FR-007,FR-008} Implement Open-Meteo request building, HTTPS execution, and response parsing in `/src/internal/provider/openmeteo/client.go`
- [X] T008 [US1] {FR-001,FR-002,FR-006,FR-008} Wire command execution to the service/provider flow in `/src/cmd/weather/run.go`

## Phase 3: 🎯 MVP User Story 2 - Fail fast on invalid coordinate invocation

- [X] T009 [US2] {FR-009} Add failing validation-boundary and invocation-shape tests in `/src/internal/validation/coordinates_test.go` and `/src/cmd/weather/main_test.go`
- [X] T010 [US2] {FR-003,FR-004,FR-005} Implement coordinate presence, range, and extra-argument validation in `/src/internal/validation/coordinates.go` and `/src/internal/validation/args.go`
- [X] T011 [US2] {FR-003,FR-004,FR-005} Integrate fail-fast validation into the command flow before provider invocation in `/src/cmd/weather/run.go`

## Phase 4: User Story 3 - Maintain a swappable provider seam

- [X] T012 [US3] {FR-009} Add failing seam-isolation tests for service/provider interaction in `/src/internal/service/weather_test.go` and `/src/internal/provider/openmeteo/client_test.go`
- [X] T013 [US3] {FR-006,FR-008} Refine the provider interface and service dependency boundary to keep Open-Meteo details internal in `/src/internal/provider/provider.go` and `/src/internal/service/weather.go`
- [X] T014 [US3] {FR-006,FR-007,FR-008} Constrain adapter-local request/response types and mapping helpers to the Open-Meteo package in `/src/internal/provider/openmeteo/request.go`, `/src/internal/provider/openmeteo/response.go`, and `/src/internal/provider/openmeteo/client.go`

## Phase 5: Polish & Cross-Cutting Concerns

- [X] T015 {FR-009} Add coverage-oriented regression cases and shared test fixtures for command, validation, service, and provider paths in `/src/cmd/weather/main_test.go`, `/src/internal/provider/openmeteo/client_test.go`, `/src/internal/service/weather_test.go`, and `/src/internal/validation/coordinates_test.go`
- [X] T016 {FR-009} Add notes for local verification commands (lint, vulnerability scan, and coverage checks) in `specs/00001-weather-retrieval-core/plan.md`

## Dependencies

- **Phase 1: Foundational** → blocks all delivery phases by establishing the module, shared internal models, and HTTP client policy.
- **Phase 2: US1** → depends on Phase 1 and establishes the runnable MVP command and provider-backed retrieval flow.
- **Phase 3: US2** → depends on Phase 1 and integrates fail-fast validation into the command flow created for US1.
- **Phase 4: US3** → depends on Phase 1 and hardens the service/provider seam after the basic retrieval path exists.
- **Phase 5: Polish** → depends on completion of Phases 2–4.
- **Parallel block**: `T002` and `T003` can run in parallel after `T001`.
