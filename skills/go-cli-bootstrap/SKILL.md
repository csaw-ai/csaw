---
name: go-cli-bootstrap
description: Use when adding or reshaping csaw CLI commands, flags, help text, package seams, or Go tests for command-facing behavior.
---

# Go CLI Bootstrap

Keep command wiring in `cmd/csaw` and behavior in `internal/`.

## Workflow

1. Read [`ARCHITECTURE.md`](../../ARCHITECTURE.md) and [`AGENTS.md`](../../AGENTS.md).
2. Keep public commands aligned with the public product overview in [`../../docs/product/overview.md`](../../docs/product/overview.md), the CLI code, and the tests.
3. Prefer adding or extending behavior packages over stuffing logic into Cobra command handlers.
4. Add or update tests with the code change.
5. If command behavior, validation commands, or package boundaries change, update docs in the same patch.

## Read When Needed

- command and test patterns: [`references/commands-and-tests.md`](references/commands-and-tests.md)
- project workflow rules: [`../../docs/reference/project-management.md`](../../docs/reference/project-management.md)
