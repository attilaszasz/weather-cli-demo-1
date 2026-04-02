---
feature_branch: "00003-canonical-weather-json"
created: "2026-04-02"
input: "E002 Canonical weather JSON"
spec_type: "product"
spec_maturity: "draft"
epic_id: "E002"
epic_sources: "{PRD:CAP-002}{SAD:ADR-003}"
---

# Feature Specification: Canonical weather JSON

**Feature Branch**: `00003-canonical-weather-json`  
**Created**: 2026-04-02  
**Status**: Draft  
**Spec Type**: product  
**Spec Maturity**: draft  
**Epic ID**: E002  
**Epic Sources**: {PRD:CAP-002}{SAD:ADR-003}  
**Product Document**: `specs/prd.md`

## Problem Statement

Automation users need the Weather CLI to emit a stable machine-readable success payload that is owned by the CLI rather than by the upstream provider. Without a canonical stdout contract, downstream scripts are coupled to internal or provider-specific shapes and successful command output can drift in ways that break automation unexpectedly.

## Scope

### Included

- Define a CLI-owned canonical success JSON contract for successful weather lookups.
- Map internal normalized weather data into the canonical success payload written to stdout.
- Add contract-focused tests for valid JSON, stable field mapping, and success-path stdout behavior.

### Excluded

- Structured error JSON and exit-code policy — deferred to `E003`.
- Human-readable output modes or alternate formats — out of scope for the MVP default JSON mode.

### Edge Cases & Boundaries

- Successful command output must write valid JSON to stdout only.
- Help and validation failures are not part of the success contract and remain outside this epic’s output mapping scope.
- The public contract must not mirror provider-specific raw payload structure directly.
- Future schema evolution should prefer additive changes over breaking renames or removals.

## User Scenarios & Testing

### User Story 1 - Parse successful CLI output in automation (Priority: P1)

As an automation user, I want successful weather lookups to emit valid canonical JSON so my scripts can parse the CLI output reliably.

**Why this priority**: This is the primary public value of the epic and the main automation-facing contract.

**Independent Test**: Run the CLI with valid coordinates and verify stdout parses as JSON matching the canonical schema while stderr stays empty.

**Acceptance Scenarios**:

1. **Given** valid coordinates and a successful weather lookup, **When** the command completes, **Then** stdout contains valid canonical JSON and stderr remains empty.
2. **Given** the provider returns normalized weather data, **When** the command maps it to the CLI response, **Then** the output uses CLI-owned field names rather than provider-specific envelope fields.

### User Story 2 - Rely on stable success fields across provider changes (Priority: P1)

As an engineering maintainer, I want the CLI success schema to be separated from provider structs so provider evolution does not silently change the public contract.

**Why this priority**: Contract stability is required to keep downstream consumers insulated from provider churn.

**Independent Test**: Validate command and contract tests against a stubbed weather service and confirm the expected JSON keys and values are preserved.

**Acceptance Scenarios**:

1. **Given** a successful internal weather response, **When** it is encoded for stdout, **Then** only canonical contract fields are exposed.
2. **Given** future internal or provider refactoring, **When** success-path tests run, **Then** they fail if canonical field mapping changes unexpectedly.

## Requirements

### Functional Requirements

- **FR-001**: System MUST define a canonical success JSON contract independent of the raw provider payload.
- **FR-002**: System MUST emit successful command output as valid parseable JSON on stdout only.
- **FR-003**: System MUST map provider-derived weather values consistently into the documented success contract.
- **FR-004**: System MUST verify success-path JSON parsing and stable field mapping with automated tests.

### Key Entities

- **Canonical Success Contract**: The CLI-owned JSON payload exposed on stdout for successful weather lookups.
- **Normalized Weather Payload**: The internal weather data returned from the service layer before command-level contract formatting.

## Assumptions & Risks

### Assumptions

- The normalized internal weather data from `E001` is a suitable source for the first public success contract.
- Default CLI mode remains JSON-first for successful responses.
- Existing command tests can be extended to verify success-path contract behavior.

### Risks

- **Contract leakage** *(likelihood: medium, impact: high)*: provider or internal structs may leak into the public JSON contract unless command output uses a separate contract type.
- **Silent breaking change** *(likelihood: medium, impact: high)*: field renames or shape drift could break automation consumers without explicit contract tests.
- **Output contamination** *(likelihood: low, impact: medium)*: success-path diagnostics could reach stdout and make JSON unparsable.

## Implementation Signals

- `NEW-CONFIG` — none
- `NEW-API` — add a CLI-owned success contract package and stdout JSON writer behavior.

## Success Criteria

### Measurable Outcomes

- **SC-001 [US1]**: Successful CLI invocations emit valid JSON on stdout that parses without transformation.
- **SC-002 [US2]**: Automated tests fail if canonical JSON field names or value mapping change unexpectedly.
