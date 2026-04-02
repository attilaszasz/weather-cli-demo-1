---
feature_branch: "00002-release-automation"
created: "2026-04-02"
input: "E004 Release automation"
spec_type: "operational"
spec_maturity: "draft"
epic_id: "E004"
epic_sources: "{PRD:CAP-005}{SAD:ADR-005}"
---

# Feature Specification: Release automation

**Feature Branch**: `00002-release-automation`  
**Created**: 2026-04-02  
**Status**: Draft  
**Spec Type**: operational  
**Spec Maturity**: draft  
**Epic ID**: E004  
**Epic Sources**: {PRD:CAP-005}{SAD:ADR-005}  
**Product Document**: `specs/prd.md`

## Problem Statement

Engineering maintainers need a repeatable way to validate and publish the Weather CLI for Linux, macOS, and Windows without relying on manual local packaging steps. Without automated CI and tagged release orchestration, binary delivery is error-prone, checksum generation is inconsistent, and the project cannot reliably satisfy the standalone executable distribution goal in the product and architecture documents.

## Scope

### Included

- GitHub Actions workflows that run the repository validation commands on repository changes and on release-related paths.
- A GoReleaser configuration that packages the CLI for Linux, macOS, and Windows with deterministic archive and checksum outputs.
- Maintainer-oriented release automation guidance embedded in workflow naming, step structure, and repository configuration.

### Excluded

- Package-manager publishing such as Homebrew, Scoop, or Winget — deferred because ADR-005 only requires GitHub Releases plus GoReleaser.
- Code signing, notarization, and SBOM generation — reserved for later hardening after baseline release automation is stable.

### Edge Cases & Boundaries

- Pull request and push validation must not attempt to publish a release.
- Tagged release automation must use repository history and tags required by GoReleaser.
- Archive naming must remain deterministic across Linux, macOS, and Windows targets.
- Release automation must operate from the existing CLI entrypoint and test baseline without moving source files outside `/src`.

## Operational Objectives

### Objective 1 - Establish CI validation for repository changes (Priority: P1)

Add repository automation that runs the project validation commands whenever maintainers change code or prepare release-related updates.

**Why this priority**: Release automation is unsafe without a green automated validation path.

**Rationale**: Maintainers need rapid feedback that the repository still builds and tests before tagging a release.

**Deliverables**:
- `.github/workflows/ci.yml`
- CI steps that execute the project Go test and build commands

**Verification Criteria**:
1. **Given** a repository change on the default branch or in a pull request, **When** the CI workflow runs, **Then** it executes the documented Go validation commands without attempting to publish artifacts.
2. **Given** release automation files change, **When** the CI workflow runs, **Then** it validates the same repository baseline that release workflows depend on.

### Objective 2 - Publish tagged multi-platform releases (Priority: P1)

Add a tag-triggered GitHub Actions workflow that invokes GoReleaser to package and publish the Weather CLI for Linux, macOS, and Windows.

**Why this priority**: Cross-platform artifact publication is the primary outcome of the epic.

**Rationale**: The product requires standalone executable distribution through GitHub releases, and GoReleaser is the accepted packaging strategy.

**Deliverables**:
- `.github/workflows/release.yml`
- `.goreleaser.yaml`
- deterministic archive and checksum naming rules

**Verification Criteria**:
1. **Given** a semantic version tag, **When** the release workflow runs, **Then** it checks out full git history and invokes GoReleaser to publish Linux, macOS, and Windows artifacts.
2. **Given** a release build completes, **When** artifacts are produced, **Then** archives and checksums follow deterministic names suitable for GitHub release distribution.

### Objective 3 - Preserve maintainable release operations (Priority: P2)

Keep the release automation understandable and easy for maintainers to extend for future signing, SBOM, or package-manager work.

**Why this priority**: Baseline automation should not block later release hardening.

**Rationale**: Operational handoff quality matters because release automation is maintained over time rather than executed once.

**Deliverables**:
- workflow structure and naming that communicate purpose clearly
- GoReleaser settings that leave room for future hardening

**Verification Criteria**:
1. **Given** a maintainer needs to cut a release, **When** they inspect repository automation files, **Then** the tag trigger, validation steps, and publishing path are clear enough to follow without reverse engineering.

### Operational Constraints

- GitHub Actions and GoReleaser are mandatory implementation choices for this epic.
- Release automation must target Linux, macOS, and Windows.
- The existing Go CLI module and `/src/cmd/weather` entrypoint remain the packaging target.
- Workflow validation must preserve the repository quality baseline from `project-instructions.md`.

## Integration Points

- **IP-001**: E004 depends on the E001 command runtime via the existing Go module and `/src/cmd/weather` entrypoint.
- **IP-002**: Release workflows depend on GitHub Actions hosted runners for CI execution and on GoReleaser for packaging and GitHub Release publication.
- **IP-003**: Maintainers depend on repository-level workflow files and `.goreleaser.yaml` as the operational interface for cutting releases.

## Requirements

### Operational Requirements

- **OR-001**: The repository MUST provide a GitHub Actions CI workflow that runs `go test ./...` and `go build ./...` on pushes and pull requests.
- **OR-002**: The repository MUST provide a tag-triggered GitHub Actions release workflow that uses GoReleaser to publish release artifacts to GitHub Releases.
- **OR-003**: The release workflow MUST build archives for Linux, macOS, and Windows.
- **OR-004**: Release outputs MUST include checksums and deterministic artifact naming.
- **OR-005**: Release automation MUST preserve the current repository source layout and Go command entrypoint under `/src`.

### Runbook Requirements

- **RR-001**: A runbook MUST exist for maintainers to understand how tagged releases are triggered and what validation occurs before publication.
- **RR-002**: A runbook MUST exist for maintainers to understand where archive and checksum naming are defined and how to extend them safely.

## Assumptions & Risks

### Assumptions

- GitHub-hosted runners are acceptable for the initial CI and release workflows.
- Publishing through GitHub Releases is sufficient for the current distribution scope.
- The module path and command entrypoint created in E001 remain stable during this epic.

### Risks

- **Workflow drift** *(likelihood: medium, impact: medium)*: CI commands may diverge from local verification practices unless both are kept aligned in one place.
- **Packaging misconfiguration** *(likelihood: medium, impact: high)*: incorrect GoReleaser settings could produce incomplete or misnamed platform artifacts.
- **Future hardening pressure** *(likelihood: low, impact: medium)*: later signing or SBOM work could require follow-up refactoring if the baseline configuration is too rigid.

## Implementation Signals

- `NEW-CONFIG` — add GitHub Actions workflow files and `.goreleaser.yaml` at repository root.
- `EXTERNAL-SERVICE` — integrate repository automation with GitHub Actions runners and GitHub Releases.

## Success Criteria

### Measurable Outcomes

- **SC-001 [OBJ1]**: Repository automation runs `go test ./...` and `go build ./...` on push and pull request events.
- **SC-002 [OBJ2]**: A tag-triggered release workflow invokes GoReleaser with full git history to publish Linux, macOS, and Windows artifacts.
- **SC-003 [OBJ2]**: Release outputs include deterministic archives and a checksum file.
- **SC-004 [OBJ3]**: Maintainers can identify the validation and publishing path directly from repository automation files without separate undocumented steps.

## Glossary

| Term | Definition |
|------|------------|
| GoReleaser | A release automation tool used to build, archive, checksum, and publish CLI binaries. |
| GitHub Actions | The CI/CD service that runs repository workflows for validation and release publication. |
| Tagged release | A repository release flow triggered from a git tag, used here to publish versioned binary artifacts. |
