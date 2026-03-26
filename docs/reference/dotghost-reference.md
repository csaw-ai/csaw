# dotghost Reference

`dotghost` is the behavioral reference for `csaw`, not the implementation source.

Inspect the sibling `dotghost` checkout or the equivalent upstream source tree when matching behavior:

- `src/runtime.ts`
- `src/workspace.ts`
- `src/profiles.ts`
- `src/commands.ts`
- `src/matching.ts`

Carry forward these behaviors unless the product docs say otherwise:

- path normalization across platforms
- stash and restore semantics
- `.git/info/exclude` management
- profile inheritance and validation
- drift detection
- conflict-handling expectations

Do not port TypeScript line by line. Re-express the behavior idiomatically in Go and encode parity expectations in tests.
