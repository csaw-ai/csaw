# Technical Debt Tracker

## Current Debt

- Expand `inspect` from source attribution to full layered provenance.
- Improve multi-source collision UX beyond the current fail-fast ambiguity error.
- Revisit `diff` semantics so it is useful for healthy mounts and drifted files, not just a command-level smoke path.
- Add Phase 2 structured context support.
- **Content-hash verification for protected files.** Today, protection is advisory — csaw's own commands honor it, but a user can manually replace a symlink with a modified file. Record SHA hashes of protected files at mount time and verify them in `csaw check --strict`. Would enable detection (not prevention) of tampering with team-mandated config.

## Usage

Add debt here only when it should not block the current change. Backlog-sized work still belongs in GitHub Issues and the project board.
