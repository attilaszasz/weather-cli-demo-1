---
feature_branch: "00004-structured-cli-errors"
created: "2026-04-02"
input: "E003 Structured CLI errors"
spec_type: "product"
spec_maturity: "draft"
epic_id: "E003"
epic_sources: "{PRD:CAP-004}{SAD:ADR-003}{SAD:ADR-004}"
---

# Feature Specification: Structured CLI errors

**Feature Branch**: `00004-structured-cli-errors`  
**Created**: 2026-04-02  
**Status**: Draft  
**Spec Type**: product  
**Spec Maturity**: draft  
**Epic ID**: E003  
**Epic Sources**: {PRD:CAP-004}{SAD:ADR-003}{SAD:ADR-004}  
**Product Document**: `specs/prd.md`

## Problem Statement

Automation users need deterministic machine-readable failure output so scripts can distinguish invalid input from downstream and internal faults without parsing ad hoc text. Without a canonical stderr error contract and stable non-zero exit codes, failure handling is brittle, diagnostics may leak unstable details, and success/failure streams are not safely separated.

## Scope

### Included

- Define a CLI-owned canonical error JSON contract for validation, downstream transport, downstream provider, and internal failures.
- Map current command and provider failure paths into deterministic error categories and exit codes.
- Add tests that verify stderr JSON, clean stdout on failure, and stable failure behavior for representative error classes.

### Excluded

- Additional output formats or debug logging modes — out of scope for the MVP error contract.
- Retry logic, backoff, or resilience policies — deferred beyond the MVP failure-surface contract.

### Edge Cases & Boundaries

- Validation failures must fail before any outbound provider call.
- Successful command output must remain on stdout, while error payloads must be written to stderr.
- Provider or transport failures must avoid leaking raw internal details unsafely.
- Exit code semantics must remain deterministic and testable.

## User Scenarios & Testing

### User Story 1 - Distinguish validation failures from downstream failures (Priority: P1)

As an automation user, I want stderr to contain structured error JSON with deterministic exit semantics so my scripts can branch safely on failure type.

**Why this priority**: Failure handling is a core automation contract and cannot depend on free-form text parsing.

**Independent Test**: Run the command with invalid input and downstream failure stubs, then verify stderr parses as canonical JSON and stdout stays empty.

**Acceptance Scenarios**:

1. **Given** invalid coordinates or unexpected arguments, **When** the command fails validation, **Then** stderr contains canonical validation error JSON and the exit semantics are deterministic.
2. **Given** a downstream transport or provider failure, **When** the command fails after invoking the service layer, **Then** stderr contains canonical downstream error JSON without corrupting stdout.

### User Story 2 - Keep failure behavior stable for maintainers and scripts (Priority: P1)

As an engineering maintainer, I want error classification and exit-code mapping centralized in CLI-owned code so failure behavior stays stable as implementation details evolve.

**Why this priority**: Centralizing failure mapping prevents accidental contract drift during future refactors.

**Independent Test**: Run contract-focused tests against representative validation, provider, and internal failure cases and verify the expected category and exit code mapping.

**Acceptance Scenarios**:

1. **Given** a known failure class, **When** the CLI maps it to the public error contract, **Then** the category and exit code match the documented policy.
2. **Given** future internal implementation changes, **When** failure-path tests run, **Then** they fail if canonical error fields or exit mappings drift unexpectedly.

## Requirements

### Functional Requirements

- **FR-001**: System MUST define a canonical error JSON contract for validation, downstream transport, downstream provider, and internal failures.
- **FR-002**: System MUST fail invalid input before network activity and return a deterministic non-zero exit code.
- **FR-003**: System MUST emit downstream and internal failure payloads to stderr as structured JSON without corrupting stdout.
- **FR-004**: System MUST verify validation, transport/provider, and internal failure-path behavior with automated tests.

### Key Entities

- **Canonical Error Contract**: The CLI-owned JSON payload emitted on stderr for failed command execution.
- **Exit Code Policy**: The deterministic mapping between canonical failure categories and process exit codes.

## Assumptions & Risks

### Assumptions

- The existing command flow and provider seam from `E001` are sufficient to classify current failure paths.
- `main()` can be updated to honor a CLI-owned exit code policy without changing the success path.
- Existing command tests can be extended to cover failure contracts.

### Risks

- **Category drift** *(likelihood: medium, impact: high)*: unstructured errors could bypass the canonical mapping unless command failure handling is centralized.
- **Information leakage** *(likelihood: medium, impact: high)*: raw provider or internal errors could expose unstable details if written directly to stderr.
- **Stream contamination** *(likelihood: low, impact: medium)*: mixed stdout/stderr behavior could break automation users expecting clean separation.

## Implementation Signals

- `NEW-API` — add a CLI-owned error contract type, classifier, and stderr writer behavior.

## Success Criteria

### Measurable Outcomes

- **SC-001 [US1]**: Failure-path command invocations emit parseable canonical error JSON on stderr with empty stdout.
- **SC-002 [US2]**: Automated tests fail if canonical error categories, fields, or exit-code mappings change unexpectedly.
