# Roadmap

## Product Thesis

csaw is the git-native context control plane for AI-assisted development.

Repos remain the source of truth for project-owned context. csaw is for context that crosses repo, tool, team, client, privacy, or provenance boundaries: shared rules, client policy, personal skills, MCP configuration, reusable agents, and local assurance that the right context is active.

## Current State

`v0.4.0` establishes the core governance direction:

- multi-source mounting from personal, team, client, and community sources
- priority-based composition and protected files
- first-class artifact kinds: instructions, rules, agents, skills, and MCP
- per-tool projection for Claude Code, Cursor, Codex, OpenCode, Windsurf, and shared fallback directories
- per-project source pinning
- fork and promote workflows
- `csaw inspect`, `check`, `status`, and `diff`
- `csaw audit` with `.csaw/policy.yml`, required sources, blocked sources, required kinds, strict mode, and JSON output
- distribution through GitHub Releases, Homebrew, Scoop, and PyPI

## Roadmap Principles

- Prefer repo-local files when context belongs to one repo.
- Build csaw where context must be composed, switched, kept private, projected across tools, or audited.
- Keep security language precise: csaw provides local assurance and detection, not hard endpoint enforcement.
- Make provenance obvious before adding more automation.
- Prefer stable files and JSON schemas over hidden state.
- Keep every new public behavior documented and tested.

## Immediate Follow-Up: v0.4.x

Goal: make the new audit surface practical and trustworthy without expanding the product too far.

| Work | Why | Acceptance |
|---|---|---|
| Document audit JSON schema | Teams need stable CI/reporting integration | `docs/reference/audit-json.md` defines report fields, severities, finding IDs, and exit behavior |
| Add `csaw audit --init` | Users need an easy way to adopt project policy | Command writes a minimal `.csaw/policy.yml` or refuses to overwrite without `--force` |
| Verify required source URL/ref | Policy currently checks active source names only | `required_sources` with `url` and `ref` emits findings when active context does not match |
| Add policy examples | Make team/client use cases concrete | README and cheat sheet include client isolation and team policy examples |
| Clean up local toolchain guidance | Avoid ad hoc local setup files | Decide whether to commit a pinned `mise.toml` or keep it out of repo |

## v0.5: Context Assurance

Goal: make `csaw audit --strict` the reliable answer to “am I in the right AI workspace?”

| Work | Why | Acceptance |
|---|---|---|
| Protected-file hash verification | Protected files are currently advisory within csaw | Mount records hashes for protected entries; audit/check detects replacement or content drift |
| Required pins | Clients may require an exact source ref | Policy can require a source branch/tag/SHA; audit reports mismatch |
| Blocked path/kind checks | Some projects may forbid MCP or personal agents entirely | Policy supports blocked kinds and blocked project paths |
| Layered provenance in inspect | Users need to know not just what won, but why | Inspect shows winning source, losing candidates, priority/protection reason, and pin state |
| Better collision UX | Fail-fast conflicts are correct but rough | Conflicts explain candidate sources, priorities, protection, and suggested fixes |

## v0.6: Context UX

Goal: make the user-facing vocabulary match the product: contexts, not just sources and profiles.

| Work | Why | Acceptance |
|---|---|---|
| `csaw context status` | Users need one command for “what mode am I in?” | Shows active sources, profile, policy, pins, audit summary, mounted kinds |
| `csaw context use <name>` | Profiles are implementation detail | Context command resolves to source/profile policy without breaking existing mount behavior |
| `csaw context leave` | Context switching needs a clean exit | Equivalent to safe unmount plus policy-aware summary |
| Context aliases | Client/team names should be ergonomic | Config supports named aliases that point to source/profile combinations |
| Explain mode | Make magic debuggable | `--explain` shows why each file was mounted or skipped |

## v0.7: Continuity And Handoff

Goal: preserve enough state that users can pause and resume work without polluting repos.

| Work | Why | Acceptance |
|---|---|---|
| `csaw pause` | Context switching is expensive | Captures active mount state, git branch, dirty summary, pins, audit result, and notes |
| `csaw resume` | Returning should be explicit and verifiable | Restores/remounts the recorded context and reports drift since pause |
| Local context ledger | Audit and handoff need history | Append-only local records of context switches and audit summaries |
| Handoff bundle | Contractors and teams need transfer artifacts | Generates a sanitized markdown/json handoff with active context and next steps |

## v0.8: Source And Agent Ecosystem

Goal: make reusable AI workspace artifacts easier to distribute without inventing a central registry too early.

| Work | Why | Acceptance |
|---|---|---|
| `csaw install <git-url>` alias | Installing sources should feel natural | Adds source, pulls, lists profiles, and suggests mount/context commands |
| Source metadata | Users need to evaluate trust and purpose | Optional metadata file supports description, owner, license, supported tools, and kinds |
| Source validation | Shared sources need quality gates | `csaw source validate` checks profile syntax, paths, kinds, protected entries, and metadata |
| Publish workflow polish | Teams need contribution loops | Existing push/fork/promote flows gain clearer status and docs |

## v0.9: Risk Scanning And Integrations

Goal: extend audit without turning csaw into a bespoke security scanner.

| Work | Why | Acceptance |
|---|---|---|
| Scanner protocol | Content/risk checks should be delegated | External scanners can emit normalized findings consumed by `csaw audit` |
| MCP risk summary | MCP is the highest-risk context surface | Audit reports command/env/path risk from known MCP config files |
| CI examples | Teams need enforcement points | Docs include GitHub Actions examples for `csaw audit --strict --json` |
| Optional npm distribution | Meet frontend users where they are | npm wrapper installs platform-specific binaries without postinstall scripts |

## v1.0 Criteria

`v1.0` should mean csaw is boringly reliable for real team/client use:

- stable `csaw.yml` profile behavior
- stable `.csaw/policy.yml` schema
- stable audit JSON schema and finding IDs
- cross-platform mount/unmount/restore confidence on Linux, macOS, and Windows
- clear context/provenance UX
- release channels working consistently
- no known data-loss bugs in stash/restore or unmount flows
- docs explain when to use repo-local context instead of csaw

## Research Tracks

These are intentionally not the next implementation steps.

### Personal Operating System For Work Modes

The broad idea is compelling: `client-acme`, `deep-work`, `incident-response`, and `writing` modes that switch tools, browser profiles, AI context, notes, tasks, and reminders.

csaw should not chase that directly yet. The credible path is to first win at AI-assisted technical work modes: source composition, audit, context status, pause/resume, and handoff. Generic app/browser/calendar automation can come later or remain integrations.

### Context Firewall

Hard prevention would require shell wrappers, IDE hooks, agent runtime integration, OS users, containers, or endpoint control. csaw should keep building local assurance and policy drift detection before claiming enforcement.

### Agent Package Manager

Agent/skill distribution needs trust, metadata, provenance, and network effects. csaw can grow toward this through git-backed source install and validation before attempting a registry ecosystem.

## Not Now

- SaaS control plane
- central hosted registry
- hard sandboxing
- custom prompt-injection scanner built into core
- generic dev environment management that competes with `direnv`, `mise`, Nix, or devcontainers
- broad consumer productivity mode switching

## Next Best Issue

The highest-leverage next implementation issue is:

**Add `csaw audit --init` and document the audit JSON schema.**

This improves adoption of the just-released governance surface, keeps the scope small, and gives future CI/reporting integrations a stable contract.
