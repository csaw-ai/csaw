---
name: doc-gardening
description: Use when a code change affects csaw documentation, AGENTS guidance, architecture notes, execution plans, or debt tracking.
---

# Doc Gardening

Keep the repo docs synchronized with the codebase.

## Workflow

1. Update `AGENTS.md` when the agent entry path, validation commands, or repo map changes.
2. Update `ARCHITECTURE.md` when package seams or implementation state changes.
3. Update the active exec plan if the work is still in flight.
4. Add debt to `docs/tech-debt-tracker.md` only when it should not block the current change.
5. Prefer editing existing durable docs over creating new one-off notes.
6. Keep public docs sanitized: no local workstation paths, secrets, or internal-only planning details.

## Read When Needed

- update checklist: [`references/update-checklist.md`](references/update-checklist.md)
