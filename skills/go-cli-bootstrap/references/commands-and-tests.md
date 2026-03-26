# Commands And Tests

- Keep command handlers thin. Parse flags, build services, and delegate.
- Prefer package-level tests for domain behavior over brittle command-output tests.
- Use table-driven tests for profiles, glob filtering, path normalization, and config parsing.
- Use `t.TempDir()` for filesystem behavior instead of mocks where symlinks or path handling matter.
