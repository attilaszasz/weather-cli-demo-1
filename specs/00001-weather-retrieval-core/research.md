# Research: Weather retrieval core
 > Feature: E001 Weather retrieval core | Date: 2026-04-02 | Purpose: inform technical planning and QC choices
 
 ## Open-Meteo provider boundary
 - **Decision**: Use Open-Meteo `/v1/forecast` behind a provider adapter that requests only the minimum current-weather fields needed by E001.
 - **Rationale**: The documented coordinate-based HTTPS API and stable required-parameter stance support a narrow, swappable adapter boundary.
 - **Rejected**: Directly exposing provider payloads, because it would couple the CLI contract to upstream schema details.
 - **Pitfalls**: Do not assume optional provider fields are always present or let raw payload shapes escape the adapter.
 - **Sources**: https://open-meteo.com/en/docs
 
 ## Go CLI parsing baseline
 - **Decision**: Use Go standard-library flag parsing as the MVP baseline with explicit named flags and strict invocation-shape validation.
 - **Rationale**: The standard `flag` package already supports the required named-flag workflow without introducing extra dependency surface.
 - **Rejected**: Positional arguments or a heavier CLI framework up front, because the MVP interface is narrow and does not need subcommand complexity yet.
 - **Pitfalls**: Do not allow mixed positional/named invocation patterns that weaken automation predictability.
 - **Sources**: https://pkg.go.dev/flag
 
 ## Go testing boundary for provider integration
 - **Decision**: Use `go test` with table-driven tests and `net/http/httptest` servers for provider request/response coverage.
 - **Rationale**: `httptest` provides standard request, recorder, and in-process server utilities that fit adapter and command-path integration testing.
 - **Rejected**: Live-provider tests as the primary baseline, because they add nondeterminism and external dependency noise.
 - **Pitfalls**: Do not let integration tests depend on public network availability for routine validation.
 - **Sources**: https://pkg.go.dev/net/http/httptest, https://pkg.go.dev/net/http
 
 ## Go QC tooling baseline
 - **Decision**: Use `golangci-lint` for linting, `govulncheck` for vulnerability scanning, and `go test -coverprofile` for coverage measurement.
 - **Rationale**: These tools align with project instructions and official Go ecosystem guidance for linting, security, and coverage.
 - **Rejected**: Ad hoc single-linter selection or dependency-only scanning, because the project requires broad lint and security coverage.
 - **Pitfalls**: Do not leave tool versions implicit in CI planning or skip code-reachability-aware vulnerability checks.
 - **Sources**: https://golangci-lint.run/docs/welcome/install/, https://go.dev/doc/security/vuln/
 
 ## Summary
 | Topic | Decision | Rationale |
 |-------|----------|-----------|
 | Open-Meteo provider boundary | Adapter over `/v1/forecast` with minimum fields | Preserves a swappable provider seam |
 | Go CLI parsing baseline | Standard `flag` with explicit named flags | Meets MVP needs with minimal dependency surface |
 | Go testing boundary | `go test` plus `httptest` | Keeps provider-path tests deterministic |
 | Go QC tooling baseline | `golangci-lint`, `govulncheck`, coverage via `go test` | Matches project quality gates |
 
 ## Sources Index
 | URL | Topic | Fetched |
 |-----|-------|---------|
 | https://open-meteo.com/en/docs | Open-Meteo provider boundary | 2026-04-02 |
 | https://pkg.go.dev/flag | Go CLI parsing baseline | 2026-04-02 |
 | https://pkg.go.dev/net/http/httptest | Go testing boundary for provider integration | 2026-04-02 |
 | https://pkg.go.dev/net/http | Go testing boundary for provider integration | 2026-04-02 |
 | https://golangci-lint.run/docs/welcome/install/ | Go QC tooling baseline | 2026-04-02 |
 | https://go.dev/doc/security/vuln/ | Go QC tooling baseline | 2026-04-02 |
