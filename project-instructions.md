<!-- template-version: 2 -->
# Weather CLI Project Instructions

## Core Principles

### I. Simplicity-First Modularity

Implementation MUST keep the CLI as a small modular Go application with clear package boundaries for command handling, provider integration, contracts, and output behavior — this preserves maintainability and keeps a single-purpose tool from drifting into an over-engineered system.

### II. Stable CLI Contracts

The public CLI success and error outputs MUST remain CLI-owned, machine-readable, and backward-compatible by default — this protects automation users from upstream provider churn and makes the tool reliable in scripts.

### III. Test-Backed Delivery

Changes MUST ship with `go test` coverage for the touched behavior, including unit and integration tests for provider mapping, contract behavior, and failure paths — this reduces regression risk for a contract-driven CLI with external service dependencies.

### IV. Automated Release Readiness

Build, lint, test, and release steps SHOULD be automatable through GitHub Actions and GoReleaser-compatible workflows — this keeps Linux, macOS, and Windows artifacts reproducible and lowers release risk.

### V. Agent Output Style

All agent output MUST be concise and outcome-oriented. This principle supersedes any verbose defaults.

- **Progress reports**: Facts and outcomes only — no narration, no restating the task.
- **Artifacts**: Emit required sections only — no preamble paragraphs, no summary epilogues.
- **Reasoning**: Omit unless the user asks "why" or the decision is non-obvious.
- **Errors / blockers**: State the problem, the attempted fix, and the result — nothing else.
- **Phase-boundary reports**: ≤ 5 bullet points.
- **Preserve without compressing**: Artifact template structure and required sections; explicit decision / registration / validation guidance in shared skills; delegation constraints and sub-agent role definitions; existing size limits (spec ≤ 10 KB, research ≤ 4 KB, stories ≤ 200 words).

## Technology Stack

<!-- Downstream phases (Plan, QC, Autopilot) read this section as the authoritative tech-stack reference. -->

- **Language/Runtime**: Go 1.24
- **Frameworks**: Go standard library first, optional Cobra-style CLI layer when justified, GoReleaser for packaging
- **Storage**: none
- **Infrastructure**: local CLI runtime, GitHub Actions for CI/CD, GitHub Releases for artifact distribution

## Testing & Quality Policy

<!-- QC extracts enforcement rules from this section. Use the keywords below so automated checks activate correctly. -->
<!-- Keywords recognised by QC: lint, static analysis, code quality, coverage, security, vulnerability, OWASP, WCAG, accessibility, benchmark, performance -->

- **Coverage Target**: 80%
- **Required QC Categories**: linting, security scanning, coverage
- **Test Strategy**: Test-after implementation with unit + integration testing for critical CLI, contract, and provider flows
- **Linting / Formatting**: golangci-lint, gofmt

## Source Code Layout

- **Policy**: ENFORCE_SRC_ROOT
- **Convention**: Source code under `/src`; internal packages under `/src/internal`; command entrypoints under `/src/cmd`; repository-level config at root

## Development Workflow

- **Branching**: Feature branches from main with squash merge
- **Commit Convention**: Conventional Commits
- **CI Requirements**: All tests pass, lint clean, release workflow remains green, no contract-breaking CLI changes without explicit documentation

<!-- Optional: add additional sections below (Security Requirements, Performance Standards, Compliance, etc.) -->

## Governance

- Project instructions supersede all other documentation and practices.
- Amendments require a version bump with ISO-dated changelog entry.
- All implementations MUST pass the Instructions Check gate during planning.
- Complexity beyond these principles MUST be justified and documented.

- Bootstrap document registrations in `.github/sddp-config.md` MUST be preserved unless the user explicitly replaces them.
- P1 work MUST preserve a demonstrable CLI MVP that accepts coordinates and returns machine-readable weather output.
- Security-sensitive values such as future API keys MUST be supplied through environment variables or CI secrets and MUST NOT be committed.

**Version**: 1.0.0 | **Last Amended**: 2026-04-02
