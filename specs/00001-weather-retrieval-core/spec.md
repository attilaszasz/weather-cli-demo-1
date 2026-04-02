---
feature_branch: "00001-weather-retrieval-core"
created: "2026-04-02"
input: "Build the initial Go CLI runtime and provider integration for coordinate-based current weather retrieval."
spec_type: "product"
spec_maturity: "clarified"
epic_id: "E001"
epic_sources: "{PRD:CAP-001}{SAD:ADR-001}{SAD:ADR-002}{SAD:ADR-004}"
---

# Feature Specification: Weather retrieval core

**Feature Branch**: `00001-weather-retrieval-core`  
**Created**: 2026-04-02  
**Status**: Draft  
**Spec Type**: product  
**Spec Maturity**: clarified  
**Epic ID**: E001  
**Epic Sources**: {PRD:CAP-001}{SAD:ADR-001}{SAD:ADR-002}{SAD:ADR-004}  
**Product Document**: `specs/prd.md`

## Problem Statement

Weather CLI needs a runnable application foundation before it can offer stable public success and error contracts. Developers and automation users are affected because they need one explicit command path that accepts known coordinates, validates them deterministically, and reaches a weather provider through a controlled integration boundary. Without this foundation, later work on JSON contracts, error behavior, and release automation has no reliable execution path to build on.

## Scope

### Included

- A runnable Go CLI entrypoint under `/src` with explicit named latitude and longitude flags.
- Baseline `--help` behavior that makes the explicit coordinate invocation discoverable in the MVP foundation.
- Invocation-shape validation that rejects missing required coordinates before any outbound provider request.
- A provider abstraction with an initial Open-Meteo implementation using a single-request HTTPS timeout of 3 seconds and no retries.
- A service-layer flow that receives provider data and returns internal normalized weather data containing temperature, wind speed, weather code, and observation time for downstream formatting work.
- Baseline unit and integration-oriented test coverage for flag parsing, validation boundaries, provider request construction, and provider response parsing.

### Excluded

- Canonical stdout success JSON contract design — deferred to E002 so this epic can focus on runtime and provider foundation.
- Structured stderr error contract and exit-code policy — deferred to E003 because failure-surface stabilization is a separate product behavior.
- Release automation, packaging, and distribution workflows — deferred to E004 because they depend on the core runtime existing first.
- Baseline `--version` behavior — deferred to later work because it is not required to establish the core runtime and provider seam.
- Forecast, geocoding, unit-selection, or human-readable output modes — out of MVP scope to preserve a narrow current-weather baseline.

### Edge Cases & Boundaries

- Missing `--latitude` or `--longitude` values must fail before any network call.
- Invalid coordinate ranges must be rejected at the CLI boundary rather than being delegated to the provider.
- Unexpected extra positional arguments or malformed flag usage must not silently alter invocation meaning.
- Provider timeouts, transport failures, or malformed provider payloads are in scope only to the extent required for request construction and response parsing tests; stable public failure contracts remain out of scope for this epic.
- Internal normalized weather data must stay provider-agnostic enough for downstream contract work and must not expose raw provider payloads as the final public surface.

## User Scenarios & Testing

### User Story 1 - Retrieve weather with explicit coordinates (Priority: P1)

As a developer or automation user, I want to run a single Go CLI command with explicit latitude and longitude flags so I can initiate current-weather retrieval from a known location through a predictable runtime entrypoint. This story establishes the MVP command surface and the end-to-end path from flag parsing to service-layer weather retrieval.

**Why this priority**: Core value proposition — without a runnable coordinate-driven command path, the product has no usable MVP foundation.

**Independent Test**: Run the CLI with valid explicit latitude and longitude flags and verify that the request reaches the service flow and returns internal normalized weather data without exposing raw provider behavior as the public contract.

**Acceptance Scenarios**:

1. **Given** the runnable CLI is available and valid latitude and longitude flags are supplied, **When** the user invokes the command, **Then** the application parses the flags, calls the provider through the service layer, and produces internal normalized weather data containing temperature, wind speed, weather code, and observation time for downstream formatting.
2. **Given** the provider abstraction is configured with Open-Meteo, **When** the command performs a successful lookup, **Then** the outbound request uses HTTPS with a 3-second timeout, performs no retries, and only the provider adapter handles provider-specific response details.
3. **Given** a user needs invocation guidance, **When** the user requests `--help`, **Then** the CLI presents the explicit coordinate flag usage for the MVP command path.

### User Story 2 - Fail fast on invalid coordinate invocation (Priority: P1)

As a developer or automation user, I want invalid or incomplete coordinate invocation to be rejected before any outbound request so I can trust the CLI to behave deterministically in scripts and local terminal workflows. This story protects automation reliability and prevents wasted network calls on obviously invalid input.

**Why this priority**: MVP reliability depends on deterministic validation at the command boundary before any provider interaction occurs.

**Independent Test**: Invoke the CLI with missing, malformed, or out-of-range coordinates and verify that validation stops execution before any outbound provider request is attempted.

**Acceptance Scenarios**:

1. **Given** the user omits either required coordinate flag, **When** the command is invoked, **Then** the application rejects the invocation before any HTTP request is created.
2. **Given** the user supplies coordinates outside the accepted latitude or longitude ranges, **When** validation runs, **Then** the application stops before the provider layer is called.
3. **Given** the user supplies extra positional arguments alongside the named flags, **When** the command is parsed, **Then** the invocation shape is treated as invalid rather than interpreted loosely.

### User Story 3 - Maintain a swappable provider seam (Priority: P2)

As an engineering maintainer, I want Open-Meteo to sit behind a provider abstraction that returns normalized internal weather data so future provider or contract changes do not require the CLI entrypoint to depend on raw upstream payloads. This story reduces rework in later epics and keeps the MVP modular without broadening scope.

**Why this priority**: Important for maintainability and downstream epic isolation, but the MVP can still demonstrate user value once the basic runtime path exists.

**Independent Test**: Review the service and provider interaction boundary and verify tests cover provider request construction and response parsing independently of the CLI-owned public contract.

**Acceptance Scenarios**:

1. **Given** the weather service depends on a provider abstraction, **When** Open-Meteo is used as the initial implementation, **Then** provider-specific request and response handling remain isolated to the adapter package.
2. **Given** provider responses are parsed successfully, **When** the service returns weather data upstream, **Then** the returned structure is normalized for downstream formatting rather than mirroring the provider payload verbatim.

## Requirements

### Functional Requirements

- **FR-001**: System MUST provide a runnable Go CLI entrypoint under `/src` that accepts explicit named `--latitude` and `--longitude` flags for current-weather retrieval.
- **FR-002**: System MUST provide baseline `--help` behavior that documents the explicit coordinate-based invocation path.
- **FR-003**: System MUST validate the presence of required coordinate flags before performing any outbound provider request.
- **FR-004**: System MUST reject latitude values outside `-90` through `90` and longitude values outside `-180` through `180` before provider invocation.
- **FR-005**: System MUST validate invocation shape so unexpected positional arguments or malformed flag usage do not proceed as valid requests.
- **FR-006**: System MUST call weather retrieval through a provider abstraction with an initial Open-Meteo implementation rather than binding the command layer directly to provider-specific HTTP logic.
- **FR-007**: System MUST use HTTPS for Open-Meteo calls with a single-request timeout of 3 seconds and no automatic retries.
- **FR-008**: System MUST normalize successful provider responses into internal weather data containing temperature, wind speed, weather code, and observation time for downstream CLI contract formatting.
- **FR-009**: System MUST include unit and integration-oriented tests covering flag parsing, validation boundaries, provider request construction, and provider response parsing.

### Key Entities

- **CLI command input**: The explicit invocation data supplied by the user, including named latitude and longitude flags and overall invocation shape.
- **Provider request**: The provider-specific outbound HTTPS request built from validated coordinates, minimum current-weather query parameters, a 3-second timeout policy, and no-retry behavior.
- **Provider response envelope**: The JSON payload returned by Open-Meteo that must be parsed and mapped without becoming the public CLI contract.
- **Weather service**: The internal orchestration boundary that receives validated input, calls the provider abstraction, and returns normalized weather data with temperature, wind speed, weather code, and observation time.
- **HTTP client configuration**: The outbound request configuration that enforces HTTPS, a 3-second timeout, and no automatic retries for predictable CLI runtime behavior.

## Assumptions & Risks

### Assumptions

- Open-Meteo remains acceptable as the MVP provider baseline for coordinate-based current weather retrieval.
- The MVP foundation can rely on Go standard-library-first CLI and HTTP capabilities unless later implementation work shows a strong need for additional libraries.
- Internal normalized weather data can stay narrower than the eventual public success contract defined in E002.
- No persistent storage, caching, or background processing is needed for this epic.

### Risks

- **Provider payload variability** *(likelihood: medium, impact: medium)*: Open-Meteo response-shape or field-presence differences could complicate normalization unless tests lock down the expected parsing surface.
- **Validation ambiguity** *(likelihood: low, impact: high)*: If invocation-shape rules are left loose, downstream automation behavior may become inconsistent across shells and usage patterns.
- **Boundary drift** *(likelihood: medium, impact: medium)*: This epic could absorb public contract or error-policy work from E002/E003 unless the service and provider seam stays narrowly scoped.

## Implementation Signals

- `NEW-API` — Establish the CLI command entrypoint and explicit coordinate flag interface under `/src/cmd`.
- `EXTERNAL-SERVICE` — Integrate Open-Meteo through an HTTPS provider adapter with bounded timeout behavior.
- `NEW-ENTITY` — Define internal provider request, provider response envelope, and normalized weather service data structures.
- `NEW-CONFIG` — Introduce an explicit request-timeout policy for outbound provider calls.

## Success Criteria

### Measurable Outcomes

- **SC-001 [US1]**: A user can invoke the CLI with explicit latitude and longitude flags or request `--help` and reach a single documented command path for current-weather retrieval.
- **SC-002 [US2]**: Invalid or incomplete coordinate invocation is rejected before any outbound provider request in all covered validation test cases.
- **SC-003 [US3]**: Provider construction and response parsing tests demonstrate that Open-Meteo-specific behavior is isolated behind a provider abstraction and yields normalized internal weather data with temperature, wind speed, weather code, and observation time.

## Glossary

| Term | Definition |
|------|------------|
| Provider abstraction | The internal interface boundary that lets the CLI and service depend on weather retrieval capability without depending on one provider’s HTTP details. |
| Internal normalized weather data | The provider-agnostic weather structure returned by the service layer for downstream contract formatting work. |
| Invocation shape | The overall command structure, including required named flags and the absence of unsupported positional input. |

## Clarifications

### Session 2026-04-02

- Q: What minimum internal normalized weather fields should E001 guarantee reaches the service layer? -> A: temperature, wind speed, weather code, observation time
- Q: What timeout policy should this foundation epic require for outbound Open-Meteo requests? -> A: 3 seconds, no retries
- Q: Should E001 explicitly require baseline --help and --version behavior, or leave that to later work? -> A: require help now, defer version

### Compliance Check
**Target**: `spec.md`
**Status**: PASS

| Principle | Verdict | Notes |
|-----------|---------|-------|
| Simplicity-First Modularity | PASS | The spec keeps the scope limited to a small modular CLI foundation with clear command, provider, and service boundaries. |
| Stable CLI Contracts | PASS | The spec explicitly avoids exposing raw provider payloads as the public contract and defers final public contract work to later epics. |
| Test-Backed Delivery | PASS | `FR-009` and the included stories require unit and integration-oriented coverage for the touched behavior. |
| Automated Release Readiness | N/A | Release automation is explicitly excluded from this epic and assigned to E004. |
| ENFORCE_SRC_ROOT | PASS | Included scope and `FR-001` require the runnable entrypoint to live under `/src`. |

**Violations**:
None.
