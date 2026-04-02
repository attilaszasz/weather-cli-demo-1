# TESTING CHECKLIST: Weather retrieval core
**Created**: 2026-04-02 | **Feature**: `specs/00001-weather-retrieval-core/spec.md`

## Completeness

- [X] CHK001 Are unit test requirements explicitly defined for command parsing, validation, service orchestration, and provider parsing coverage? [Completeness, Spec §Requirements FR-009] <!-- Evaluator: Covered by spec.md §Requirements FR-009 and plan.md §Testing Strategy -->
- [X] CHK002 Are integration-oriented test requirements explicitly defined for provider request construction and provider response parsing? [Completeness, Spec §Requirements FR-009] <!-- Evaluator: Covered by spec.md §Requirements FR-009 and plan.md §Testing Strategy -->
- [X] CHK003 Are the required validation boundary cases identified for missing latitude, missing longitude, out-of-range latitude, and out-of-range longitude? [Completeness, Spec §Edge Cases & Boundaries] <!-- Evaluator: Covered by spec.md §Edge Cases & Boundaries and §User Story 2 Acceptance Scenarios -->
- [X] CHK004 Are malformed invocation-shape cases, including unexpected positional arguments, covered by stated testing expectations? [Completeness, Spec §Edge Cases & Boundaries] <!-- Evaluator: Covered by spec.md §Edge Cases & Boundaries and §User Story 2 Acceptance Scenarios -->
- [X] CHK005 Are successful-path testing expectations tied to the normalized internal weather data fields required by the clarified spec? [Completeness, Spec §Clarifications] <!-- Evaluator: Covered by spec.md §Clarifications, §Requirements FR-008, and §User Story 1 Acceptance Scenarios -->
- [X] CHK006 Is the outbound timeout and no-retry policy reflected in test planning for provider-facing behavior? [Completeness, Spec §Requirements FR-007] <!-- Evaluator: Covered by spec.md §Requirements FR-007 and plan.md §Testing Strategy -->

## Clarity

- [X] CHK007 Are the testing requirements written as artifact-quality expectations rather than implementation verification steps? [Clarity, Spec §User Scenarios & Testing] <!-- Evaluator: Covered by spec.md §User Scenarios & Testing and §Requirements -->
- [X] CHK008 Is the distinction between unit tests and integration-oriented tests unambiguous across the spec and plan? [Clarity, Spec §Requirements FR-009] <!-- Evaluator: Covered by spec.md §Requirements FR-009 and plan.md §Testing Strategy -->
- [X] CHK009 Are the normalized internal fields under test named consistently across spec, plan, and contract artifacts? [Clarity, Spec §Key Entities] <!-- Evaluator: Covered by spec.md §Key Entities, plan.md §Data Model Summary, and contracts/cli-command.md §Internal Success Payload -->
- [X] CHK010 Is the planned integration-test boundary clear about replacing the external provider with deterministic test doubles rather than live network dependence? [Clarity, Plan §Testing Strategy] <!-- Evaluator: Covered by plan.md §Testing Strategy and research.md §Go testing boundary for provider integration -->
- [X] CHK011 Is the testing target for help behavior explicit enough to avoid confusion with later version/output-contract work? [Clarity, Spec §Scope] <!-- Evaluator: Covered by spec.md §Scope Included/Excluded and §Requirements FR-002 -->

## Consistency

- [X] CHK012 Does the plan’s testing strategy align with the spec’s required test coverage areas without adding or omitting major test categories? [Consistency, Plan §Testing Strategy] <!-- Evaluator: Covered by spec.md §Requirements FR-009 and plan.md §Testing Strategy -->
- [X] CHK013 Do the requirement coverage map entries for FR-009 point to test files that match the stated testing strategy tiers? [Consistency, Plan §Requirement Coverage Map] <!-- Evaluator: Covered by plan.md §Requirement Coverage Map and §Testing Strategy -->
- [X] CHK014 Do the risk mitigations for provider payload variability and validation ambiguity align with the planned test approach? [Consistency, Plan §Risk Mitigation] <!-- Evaluator: Covered by plan.md §Risk Mitigation and §Testing Strategy -->
- [X] CHK015 Do the contract and data-model artifacts use the same testing-relevant terminology as the spec for command input, provider request, and normalized weather data? [Consistency, Spec §Key Entities] <!-- Evaluator: Covered by spec.md §Key Entities, data-model.md, and contracts/cli-command.md -->
- [X] CHK016 Is the quality-gate toolchain in the plan consistent with project instructions for linting, security scanning, and coverage? [Consistency, Plan §Testing Strategy] <!-- Evaluator: Covered by project-instructions.md §Testing & Quality Policy and plan.md §Testing Strategy -->

## Testability

- [X] CHK017 Can a reviewer determine from the artifacts what evidence would show FR-003 through FR-009 are adequately test-covered? [Testability, Plan §Requirement Coverage Map] <!-- Evaluator: Covered by plan.md §Requirement Coverage Map -->
- [X] CHK018 Are the command, validation, provider, and service seams sufficiently separated in the plan to support isolated tests and integration tests? [Testability, Plan §Architecture] <!-- Evaluator: Covered by plan.md §Architecture and §Project Structure -->
- [X] CHK019 Are the provider-facing tests designed around deterministic inputs and outputs rather than unspecified upstream behavior? [Testability, Plan §Testing Strategy] <!-- Evaluator: Covered by plan.md §Testing Strategy and research.md §Go testing boundary for provider integration -->
- [X] CHK020 Are the testing tools and install expectations specified well enough for a contributor to reproduce the planned quality checks? [Testability, Plan §Testing Strategy] <!-- Evaluator: Covered by plan.md §Testing Strategy -->
- [X] CHK021 Are the success criteria specific enough to support later verification of testing outcomes for the P1 and P2 stories? [Testability, Spec §Success Criteria] <!-- Evaluator: Covered by spec.md §Success Criteria and §User Scenarios & Testing -->
