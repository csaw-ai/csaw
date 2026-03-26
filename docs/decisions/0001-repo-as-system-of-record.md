# 0001: Repo As System Of Record

## Status

Accepted

## Decision

Treat the repository itself as the durable system of record for agent execution context:

- `AGENTS.md` is the concise map for agents entering the repo.
- `docs/` contains durable product, architecture, reference, and implementation context.
- GitHub Issues and the GitHub Project are the backlog of record.
- Active execution context for multi-step work lives in versioned exec plans under `docs/exec-plans/active/`.

## Rationale

Agent work degrades when context is split across ephemeral chats, undocumented conventions, and ad hoc TODO files. A versioned repo-local knowledge base keeps the operational map close to the code while allowing the backlog itself to live in GitHub where triage and prioritization belong.

## Consequences

- Contributors must update docs when behavior or workflows change.
- The repo should prefer a few durable documents over many overlapping notes.
- Tool-specific instruction files should stay minimal and point back to `AGENTS.md` when possible.
