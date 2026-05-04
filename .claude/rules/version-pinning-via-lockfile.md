---
description: Pin exact tool versions in mise.lock, not in mise.toml. Declarations stay loose (e.g. `pulumi = "3"`) so `mise up` works ergonomically; the lockfile is the single source of truth for "what version is everyone on right now".
paths:
  - "mise.toml"
  - "mise.ci.toml"
  - "mise.lock"
  - ".github/actions/setup-mise/action.yml"
---

# Version Pinning via mise.lock

`mise.toml` and `mise.ci.toml` declare **ranges** (`pulumi = "3"`, `golangci-lint = "2"`). `mise.lock` pins **exact** versions (`pulumi 3.233.0`, `golangci-lint 2.11.4`). Same model as `package.json` + `package-lock.json`.

`[settings] lockfile = true` is set in both mise files; `mise install` reads from the lockfile when present.

## Bumping

- `mise up` — upgrade within the declared range, update `mise.lock`. (e.g. `pulumi = "3"` lets it move from 3.233.0 to 3.234.0.)
- `mise up --bump` — widen the range in `mise.toml` and update `mise.lock`. (e.g. `pulumi 3` → `pulumi 4`.)

Either way, commit the resulting `mise.lock` change as part of the bump.

## CI plumbing

`.github/actions/setup-mise/action.yml` stages both `mise.ci.toml` (renamed to `mise.toml` in `$RUNNER_TEMP/mise-ci/`) **and** `mise.lock` so mise-action installs from the lockfile. If you forget the `mise.lock` copy, CI silently falls back to range resolution and drift returns.

## Don't

- Don't pin exact versions in `mise.toml` (e.g. `pulumi = "3.233.0"`) to fix a drift bug — that's the wrong layer. Lockfile is the right place; the declaration should describe intent ("we want pulumi 3.x"), not the snapshot.
- Don't leave a tool unpinned (e.g. `golangci-lint = "2"` with no lockfile entry) — silent minor-version drift breaks CI when new lint rules ship.
- Don't `mise use --pin <tool>@<version>` to pin via the lockfile — that flag writes the exact version into `mise.toml`, the opposite of what we want. Use `mise install <tool>@<version>` then `mise lock`.
