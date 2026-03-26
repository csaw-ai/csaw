---
name: dotghost-parity
description: Use when csaw behavior should be compared to the existing dotghost implementation without porting TypeScript source directly.
---

# dotghost Parity

Use `dotghost` to answer behavior questions, not to dictate Go structure.

## Workflow

1. Start with [`../../docs/reference/dotghost-reference.md`](../../docs/reference/dotghost-reference.md).
2. Inspect the minimal `dotghost` files needed to answer the question.
3. Write down the behavior in tests or docs before mirroring it in Go.
4. Prefer black-box parity: inputs, outputs, side effects, and edge cases.
5. If csaw intentionally diverges, document the divergence in the exec plan or decision log.

## Read When Needed

- suggested files and parity notes: [`references/reference-files.md`](references/reference-files.md)
