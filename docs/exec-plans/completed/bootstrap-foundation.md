# Bootstrap Foundation

## Summary

Bootstrap `csaw` into an agent-first Go CLI repository with the Phase 1 command surface, behavior-oriented package seams, repo-local skills, and project-management scaffolding needed for iterative mount-engine delivery.

## Success Criteria

- Root docs and `docs/` structure are present and internally linked.
- `AGENTS.md` points agents to the right product, architecture, and workflow docs.
- The Go module compiles and exposes the v0.1 command surface.
- Core packages exist for sources, profiles, mount planning, workspace state, drift, inspect, and output.
- `mount`, `unmount`, `check`, `update`, `pull`, `push`, and `status` are wired to real Phase 1 behavior.
- Cross-source profiles and per-source `.csawignore` are implemented.
- Repo-local skills exist for bootstrap, mount-engine work, `dotghost` parity, exec-plan maintenance, and doc gardening.
- CI runs `gofmt`, `go test`, and `go vet` across macOS, Linux, and Windows.

## Workstreams

### 1. Repo bootstrap

- establish root docs, license, ignore rules, Makefile, and module metadata
- publish a public-safe product overview and keep the detailed brief outside the repo
- define `docs/` as the durable knowledge base

### 2. CLI and package seams

- add Cobra entrypoint and Phase 1 command surface
- add multi-source source catalogs, cross-source profile resolution, and per-source ignore handling
- implement mount and unmount orchestration, restore snapshots, drift repair, and personal-registry push
- leave layered provenance and Phase 2 context work for follow-up

### 3. Agent workflow

- add repo-local skills with concise descriptions and linked references
- add docs and tests that validate `AGENTS.md`, active exec plans, and skill frontmatter
- add GitHub issue and PR templates aligned with the hybrid planning model

## Risks

- Multi-source collisions currently fail fast instead of offering richer source-selection UX.
- Inspect shows source attribution, but not the full layered provenance model described for later phases.
- External Go dependencies may require network access the first time the module is built.

## Validation

```bash
make fmt
make test
make vet
make docs-check
go run ./cmd/csaw --help
```
