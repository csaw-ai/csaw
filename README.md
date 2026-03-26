# csaw

`csaw` is a Go CLI for mounting AI workspace configuration from git-backed registries into a project with symlinks instead of copied files.

The repository is bootstrapped as an agent-first codebase:

- [`AGENTS.md`](AGENTS.md) is the top-level operating manual for coding agents.
- [`docs/`](docs/index.md) holds durable public product, architecture, implementation, and decision context.
- [`skills/`](skills/) contains repo-local skills for repeatable engineering workflows.

## Status

The repo now contains the foundational scaffold for Phase 1:

- Go module and CLI command surface
- mount and unmount lifecycle with stash, restore, `--force`, `--skip-conflicts`, and `--restore`
- source catalogs, cross-source profile resolution, per-source `.csawignore`, drift repair, and personal-registry push
- internal package boundaries for sources, profiles, mount planning, workspace state, drift, inspect, and output
- agent-facing docs, skills, issue templates, PR template, and CI
- repository validation tests for docs and skills

Remaining gaps are mostly follow-up hardening work: richer layered provenance in `inspect`, clearer multi-source collision UX, and Phase 2 context features. The active implementation plan lives in [`docs/exec-plans/active/bootstrap-foundation.md`](docs/exec-plans/active/bootstrap-foundation.md).

## Quick Start

```bash
make fmt
make test
make vet
go run ./cmd/csaw --help
```

## Current Command Surface

The v0.1 CLI surface is present now:

- `csaw source add|remove|list`
- `csaw mount`
- `csaw unmount`
- `csaw inspect`
- `csaw check`
- `csaw update`
- `csaw diff`
- `csaw pull`
- `csaw push`
- `csaw status`
- `csaw version`

Some commands are still scaffolded rather than feature-complete. See [`ARCHITECTURE.md`](ARCHITECTURE.md) and the active exec plan for current implementation status.

## Documentation

- [`AGENTS.md`](AGENTS.md)
- [`ARCHITECTURE.md`](ARCHITECTURE.md)
- [`CONTRIBUTING.md`](CONTRIBUTING.md)
- [`SECURITY.md`](SECURITY.md)
- [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md)
- [`docs/index.md`](docs/index.md)
