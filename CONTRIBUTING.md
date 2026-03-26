# Contributing

## Workflow

1. Start from a GitHub issue or an active execution plan.
2. If the work spans multiple subsystems, public behavior, or more than one session, create or update an exec plan in `docs/exec-plans/active/`.
3. Keep docs, tests, and code together in the same change.
4. Move completed plans to `docs/exec-plans/completed/` when the work is done.

## Validation

Baseline validation before opening a PR:

```bash
make fmt
make test
make vet
make docs-check
```

If you touch command behavior, also run:

```bash
go run ./cmd/csaw --help
```

## Backlog And Tracking

- GitHub Issues and the project board are the backlog of record.
- `docs/exec-plans/active/` is for in-flight implementation context.
- `docs/tech-debt-tracker.md` is for debt that should not block the current change.
Do not create a long-lived root `TODO.md`.

## Repo Standards

- Preserve the command surface defined in the architecture and product docs.
- Prefer small behavior-oriented packages.
- Treat `dotghost` as behavioral reference, not source code to translate.
- Keep `AGENTS.md` concise and push detail into linked docs.
- Do not commit secrets, credentials, or workstation-specific paths.
