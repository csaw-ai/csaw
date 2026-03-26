---
name: mount-engine
description: Use when implementing or revising csaw mount, unmount, stash, exclude, drift, or symlink behavior across platforms.
---

# Mount Engine

Treat `dotghost` as the behavioral reference and the public code plus docs as the repository contract.

## Workflow

1. Read [`../../docs/product/overview.md`](../../docs/product/overview.md).
2. Read [`../../docs/reference/dotghost-reference.md`](../../docs/reference/dotghost-reference.md).
3. Inspect the current Go packages in `internal/mount`, `internal/workspace`, `internal/drift`, and `internal/runtime`.
4. Encode behavior in tests before broad refactors when parity risk is high.
5. Preserve naming from the brief: `.csawignore`, `.csaw-stash`, `manifest.json`, and `# csaw-managed`.

## Watchouts

- Windows path normalization and link behavior
- CRLF and BOM handling
- `.git/info/exclude` idempotence
- restoring originals without damaging unrelated files

## Read When Needed

- phase 1 behavior checklist: [`references/phase1-behaviors.md`](references/phase1-behaviors.md)
