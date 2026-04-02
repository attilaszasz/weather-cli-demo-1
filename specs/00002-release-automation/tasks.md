# Task List: Release automation

**Feature**: `specs/00002-release-automation/spec.md`  
**Plan**: `specs/00002-release-automation/plan.md`

## Project Mode

- Brownfield

## Epic / Capability Map

- **OBJ1**: Establish CI validation for repository changes
- **OBJ2**: Publish tagged multi-platform releases
- **OBJ3**: Preserve maintainable release operations

## Phase 1: Foundational

- [X] T001 {OR-005} Add repository automation scaffolding in `.github/workflows/` and `.goreleaser.yaml`
- [X] T002 [P] {OR-001} Add CI validation workflow for `go test ./...` and `go build ./...` in `.github/workflows/ci.yml`
- [X] T003 [P] {OR-002,OR-003,OR-004} Add GoReleaser configuration for Linux, macOS, and Windows archives plus checksums in `.goreleaser.yaml`

## Phase 2: 🎯 P1 Objective 2 - Publish tagged multi-platform releases

- [X] T004 [OBJ2] {OR-002,OR-003,OR-004} Add tag-triggered GitHub Actions release workflow invoking GoReleaser with full git history in `.github/workflows/release.yml`
- [X] T005 [OBJ2] {OR-003,OR-004,OR-005} Align GoReleaser build targets, binary path, archives, and checksum naming with the existing CLI entrypoint in `.goreleaser.yaml`

## Phase 3: 🎯 P1 Objective 1 - Establish CI validation for repository changes

- [X] T006 [OBJ1] {OR-001,OR-005} Finalize push and pull request CI triggers and Go toolchain setup in `.github/workflows/ci.yml`

## Phase 4: Objective 3 - Preserve maintainable release operations

- [X] T007 [OBJ3] {RR-001,RR-002} Keep workflow naming, step structure, and artifact naming rules clear enough for maintainer handoff in `.github/workflows/release.yml` and `.goreleaser.yaml`

## Phase 5: Polish & Cross-Cutting Concerns

- [X] T008 {OR-001,OR-002,OR-003,OR-004,OR-005,RR-001,RR-002} Validate repository commands locally and update QC-facing evidence files for the release automation feature workspace

## Dependencies

- **Phase 1: Foundational** → establishes the repository config files required by all later objectives.
- **Phase 2: OBJ2** → depends on Phase 1 and wires tagged publication through GoReleaser.
- **Phase 3: OBJ1** → depends on Phase 1 and finalizes CI validation behavior.
- **Phase 4: OBJ3** → depends on Phases 2–3 so maintainer-facing clarity reflects the final workflow shape.
- **Phase 5: Polish** → depends on completion of Phases 2–4.
- **Parallel block**: `T002` and `T003` can run in parallel after `T001`.
