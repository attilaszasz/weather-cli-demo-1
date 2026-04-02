# Task List: Canonical weather JSON

**Feature**: `specs/00003-canonical-weather-json/spec.md`  
**Plan**: `specs/00003-canonical-weather-json/plan.md`

## Project Mode

- Brownfield

## User Story Map

- **US1**: Parse successful CLI output in automation
- **US2**: Rely on stable success fields across provider changes

## Phase 1: Foundational

- [X] T001 {FR-001} Add a CLI-owned canonical success contract in `src/internal/contract/success.go`
- [X] T002 {FR-002,FR-003} Replace ad hoc success formatting with canonical JSON stdout writing in `src/cmd/weather/run.go`

## Phase 2: 🎯 P1 User Story 1 - Parse successful CLI output in automation

- [X] T003 [US1] {FR-002,FR-004} Extend command success tests to assert valid JSON stdout and empty stderr in `src/cmd/weather/main_test.go`

## Phase 3: 🎯 P1 User Story 2 - Rely on stable success fields across provider changes

- [X] T004 [US2] {FR-001,FR-003,FR-004} Add stable field mapping assertions for the canonical contract in `src/cmd/weather/main_test.go`

## Phase 4: Polish & Cross-Cutting Concerns

- [X] T005 {FR-001,FR-002,FR-003,FR-004} Run repository test/build QC and update the E002 feature evidence files

## Dependencies

- **Phase 1: Foundational** → establishes the canonical contract type and command output path.
- **Phase 2: US1** → depends on Phase 1 and verifies parseable JSON stdout behavior.
- **Phase 3: US2** → depends on Phase 1 and verifies stable canonical field mapping.
- **Phase 4: Polish** → depends on Phases 2–3.
