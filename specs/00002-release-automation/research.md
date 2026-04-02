# Research: Release automation

## Topic: GitHub Actions release workflow
- Recommendation: keep CI and release concerns split into a validation workflow for push and pull request events plus a tag-driven release workflow for publishing artifacts.
- Why it matters: this isolates fast feedback from privileged release publishing and makes failure diagnosis clearer for maintainers.
- Source 1: GitHub Actions workflow trigger and reusable release patterns.
- Source 2: Existing project architecture decision ADR-005 in `specs/sad.md`.

## Topic: GoReleaser baseline
- Recommendation: use a single `.goreleaser.yaml` with explicit builds for Linux, macOS, and Windows, archive naming, checksum generation, and changelog disabling for deterministic local repository use.
- Why it matters: it centralizes packaging policy and reduces custom matrix scripting in CI.
- Source 1: GoReleaser configuration best practices for multi-platform CLI packaging.
- Source 2: ADR-005 in `specs/sad.md`.

## Topic: Validation strategy for release automation
- Recommendation: validate release automation through regular `go test ./...`, `go build ./...`, config-aware CI steps, and workflow structure checks rather than relying on live release publication during development.
- Why it matters: local development cannot safely simulate GitHub tag publication, but it can validate the commands and packaging intent that workflows invoke.
- Source 1: Project instructions quality policy.
- Source 2: Existing repository Go command baseline from `go.mod` and the implemented CLI packages.
