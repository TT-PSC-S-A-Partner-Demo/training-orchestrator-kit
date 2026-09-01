# Feature request: invoices service

Add a new `invoices` service alongside `orders`, `billing`, and `shipping`.

It needs a `normalize(x)` function that cleans up user-entered text the same way the
other services do, so invoice references are stored consistently.

## Acceptance criteria

- `normalize("  INV-100  ")` returns `"INV-100"` (surrounding whitespace trimmed).
- Lives in `services/invoices.py`, same shape as the sibling services.

## Out of scope

- Persistence, formatting, numbering schemes - just the `normalize` step for now.

---

> This is the request as it arrived. Build to it faithfully. If the tests then disagree
> with the code, do not just patch the code - ask whether this spec is missing something
> the other services already assume.
