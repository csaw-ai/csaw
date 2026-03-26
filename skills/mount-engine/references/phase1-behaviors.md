# Phase 1 Behaviors

Target behaviors to preserve from the brief and `dotghost`:

- enumerate registry files and filter them through profiles, include globs, excludes, and ignore rules
- detect conflicts before writing symlinks
- stash overwritten originals into `.csaw-stash/`
- add mounted paths to `.git/info/exclude` under `# csaw-managed`
- detect broken or missing symlink targets
- repair broken state without silently deleting unrelated project files
