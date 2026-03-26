# Project Management

## System Of Record

- GitHub Issues and one GitHub Project are the backlog and roadmap of record.
- `docs/exec-plans/active/` is the execution context for complex in-flight work.
- `docs/tech-debt-tracker.md` holds debt that should not block current delivery.

## GitHub Project Fields

Create these custom fields on the GitHub Project:

- `Status`
- `Type`
- `Area`
- `Phase`
- `Priority`
- `Size`
- `Agent Ready`
- `Target Release`

## Label Taxonomy

Use a tight and predictable label set:

- `type:feature`
- `type:bug`
- `type:research`
- `type:execution-plan`
- `area:cli`
- `area:sources`
- `area:profiles`
- `area:mount`
- `area:workspace`
- `area:docs`
- `area:ci`
- `phase:v0.1`
- `phase:v0.2`
- `phase:v0.3`
- `prio:p0`
- `prio:p1`
- `prio:p2`
- `blocked`
- `agent-ready`
- `needs-human`

## Agent-Ready Definition

An issue is agent-ready only when it:

- has explicit acceptance criteria
- identifies the affected subsystem or package
- names the commands or tests that prove completion
- points to the docs the implementer must read first
- is small enough to land in one PR unless an exec plan says otherwise
- is written so it can live in a public repository without internal names, secrets, or proprietary planning details

## Execution Plans

Create or update an exec plan when work:

- spans multiple subsystems
- changes public behavior or interfaces
- is likely to outlive a single coding session

Required headings for active plans:

- `## Summary`
- `## Success Criteria`
- `## Workstreams`
- `## Risks`
- `## Validation`

## Public Repo Guardrails

- Use generic examples and placeholders in docs, plans, and issue bodies.
- Keep private business strategy, customer-specific notes, and unpublished operational details out of versioned docs.
- Never commit secrets, API tokens, passwords, private keys, or workstation-specific paths.
