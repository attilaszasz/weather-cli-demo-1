# Task List: Structured CLI errors

**Feature**: `specs/00004-structured-cli-errors/spec.md`  
**Plan**: `specs/00004-structured-cli-errors/plan.md`

## Project Mode

- Brownfield

## User Story Map

- **US1**: Distinguish validation failures from downstream failures
- **US2**: Keep failure behavior stable for maintainers and scripts

## Phase 1: Foundational

- [X] T001 {FR-001} Add a CLI-owned canonical error contract and exit code policy in `src/internal/contract/error.go`
- [X] T002 {FR-002,FR-003} Update command failure handling to emit structured stderr JSON in `src/cmd/weather/run.go`
- [X] T003 {FR-002,FR-003} Update process exit handling to use deterministic exit codes in `src/cmd/weather/main.go`

## Phase 2: 🎯 P1 User Story 1 - Distinguish validation failures from downstream failures

- [X] T004 [US1] {FR-001,FR-002,FR-003,FR-004} Add validation and downstream failure-path command tests in `src/cmd/weather/main_test.go`

## Phase 3: 🎯 P1 User Story 2 - Keep failure behavior stable for maintainers and scripts

- [X] T005 [US2] {FR-001,FR-003,FR-004} Add internal-failure and exit-code mapping assertions in `src/cmd/weather/main_test.go`

## Phase 4: Polish & Cross-Cutting Concerns

- [X] T006 {FR-001,FR-002,FR-003,FR-004} Run repository QC and update the E003 feature evidence files

## Dependencies

- **Phase 1: Foundational** → establishes the error contract, stderr writer, and exit policy.
- **Phase 2: US1** → depends on Phase 1 and verifies canonical validation/downstream behavior.
- **Phase 3: US2** → depends on Phase 1 and verifies stable internal and exit-code behavior.
- **Phase 4: Polish** → depends on Phases 2–3.
