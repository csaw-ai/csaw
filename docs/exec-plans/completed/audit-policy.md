# Audit Policy

## Summary

Add a first `csaw audit` command that verifies active AI workspace context against project policy. The initial scope is local assurance: mounted-link health, required sources, blocked sources, and required artifact kinds. It does not attempt content scanning, prompt-injection detection, or hard runtime enforcement.

## Success Criteria

- `csaw audit [path]` runs inside a git project and reports active context findings.
- `.csaw/policy.yml` supports `required_sources`, `blocked_sources`, and `required_kinds`.
- Default exit behavior fails on `error`; `--strict` fails on `warn` and `error`.
- `--json` emits a stable machine-readable report with the same exit semantics.
- New logic has unit tests and command-facing coverage.
- README, product overview, architecture, and cheat sheet document the command without overpromising enforcement.

## Workstreams

- Add an internal audit package for policy parsing, active mount inspection, findings, rendering, and exit semantics.
- Wire the Cobra command in `cmd/csaw`.
- Add tests for policy parsing, required/blocked checks, required kinds, strict behavior, and CLI happy path.
- Update docs and command references.

## Risks

- Policy language can sprawl if it tries to model enterprise enforcement too early.
- Local audit is advisory; docs must distinguish detection from prevention.
- Required source checks depend on mount state provenance, so manually discovered links cannot provide full assurance.

## Validation

```bash
mise exec -- gofmt -l .
git diff --check
mise exec -- go test ./...
mise exec -- go vet ./...
mise exec -- go build ./...
```
