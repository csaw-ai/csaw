# Product Overview

csaw (like "see-saw") is a CLI tool that **mounts, not installs** your AI workspace configuration. It manages agent instructions (AGENTS.md), skills (SKILL.md), and other AI config files from one or more git-backed registries — symlinked into your projects, excluded from git history, and fully inspectable at all times.

## How it works

Your AI config files live in registries (personal, team, or community). csaw symlinks them into your project so your AI tools can see them, but your repo stays clean. Mount when you need them. Unmount when you don't. Switch between configurations. Detect drift. Inspect everything.

## Key ideas

- **Mount, not install.** Symlinks from a registry, not copies in your repo. Reversible, temporary, clean.
- **No hidden defaults.** `csaw inspect` shows the full state — what's mounted, where it came from, and whether it's healthy.
- **Files, not formats.** csaw manages standard files (AGENTS.md, SKILL.md, plain markdown). Every file in a registry is usable without csaw.
- **Multi-source composition.** Layer team registries, personal overrides, and community skills. Profiles with inheritance control what mounts where.

## Where to learn more

- [README.md](../../README.md) — quick start and command overview
- [ARCHITECTURE.md](../../ARCHITECTURE.md) — package structure and interfaces
- [CONTRIBUTING.md](../../CONTRIBUTING.md) — contributor workflow
- [Distribution strategy](../reference/distribution.md) — how csaw is packaged and released
