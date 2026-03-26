# Public Repo Policy

## Goal

Standard open-source hygiene — keep secrets, credentials, and local environment details out of the repository.

## Do not commit

- Secrets, credentials, API keys, or private keys
- `.env` files or local environment configuration
- Workstation-specific absolute paths (e.g., `/Users/yourname/...`)
- Files containing internal URLs, hostnames, or infrastructure details

## Guardrails

- `.gitignore` blocks common sensitive file patterns (`.env`, `.private/`, `*.local.*`)
- CI runs validation on every push and pull request
- PR authors should verify no sensitive material was introduced

## Local-only paths

Use `.private/` or other gitignored paths for any local notes, scratch files, or working material you don't want committed. These paths are blocked by `.gitignore`.
