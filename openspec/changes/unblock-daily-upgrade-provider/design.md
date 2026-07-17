# Design — unblock-daily-upgrade-provider

## Context

The daily `upgrade-provider` workflow automates bridge/upstream bumps end-to-end (upgrade → regen → PR → auto-merge). It has been red since 2026-07-09:

- 07-09/07-10: upgrade to `terraform-provider-astro` v1.4.0 reached `make generate_sdks` and failed at `.make/build_dotnet` (CI showed only make's stderr; local repro surfaced `Alerts.cs(165,74): error CS0542`).
- 07-11 onward (and pulumi-fivetran from the same date): bridge v3.134.0 declares `go >= 1.25.11`; the `[env] GOTOOLCHAIN = "go1.25.9"` pin in `mise.toml` makes the workflow's `go get` fail before anything else happens, masking the dotnet failure.

Key mechanics established during diagnosis:

- mise shims re-export `[env]` values over the caller's environment (verified: `GOTOOLCHAIN=auto go get …` still ran under `go1.25.9`). The pin therefore cannot be worked around per-step; it must change in `mise.toml`.
- `mise.lock` already pins go exactly (1.25.11 in both repos) — the env pin is a redundant second copy that has drifted twice before (`go1.24.13` phase-0, `go1.25.9` phase-2, now again).
- Verified with the fix applied in a scratch worktree: `go get pulumi-terraform-bridge/v3@v3.134.0` under `auto` bumps `provider/go.mod` to `go 1.25.11`, Go drops the stale `toolchain` line on the next workspace command, and `go build ./provider/...` passes. The workflow's existing `Sync go.work go directive` step commits the equivalent skew fix on upgrade branches.

## Goals / Non-Goals

**Goals:**

- Daily upgrade automation goes green in both repos and stays green across future bridge Go-requirement bumps without manual toolchain bumps.
- Exactly one pinned source of the installed Go version (`mise.lock`).
- The astro v1.2.6 → v1.4.0 upgrade PR can regenerate and build all four SDKs.

**Non-Goals:**

- No workflow YAML changes (the sync + auto-merge machinery is sound).
- No change to how other tools are pinned (lockfile model per `developer-tooling` spec stands).
- Not upgrading bridge/upstream by hand — the bot does that once unblocked (surgical-change principle).

## Decisions

### D1: `GOTOOLCHAIN = "auto"` (not `local`, not a fresher exact pin)

- **Fresher exact pin (`go1.25.11`)** — rejected: recreates the exact incident on the next drift; this is the "we have done this a few times" loop.
- **`local`** — rejected: single-sources the version to `mise.lock`, but the daily job still hard-fails whenever the bridge requires a Go newer than the lock (bridge adopts Go patch releases within weeks; nothing automated bumps `mise.lock`). Same recurrence, different error text.
- **`auto` (chosen)** — the go command uses the mise-installed toolchain until a module's `go` directive requires newer, then downloads that exact toolchain (sumdb-checksummed). The requirement lives in committed `go.mod`/`go.work` directives, so builds stay reproducible at the git level. This is also Go's upstream default; the repo stops fighting it.
- **Templating the env pin from the lockfile** — rejected: mise env templates have no supported "resolved tool version" variable; an `exec()` hack adds fragility for no benefit over `auto`.

### D2: Fix CS0542 with a `CSharpName` field override in `resources.go`

The bridge's supported mechanism for language-level name collisions. `"alerts": {CSharpName: "AlertsCollection"}` renames only the C# member; nodejs/python/go SDKs and the Pulumi schema property name are unchanged. Alternatives — patching generated code (forbidden; regenerated every upgrade) or renaming the Pulumi-level property (breaks all four SDKs' surface) — rejected. Name choice `AlertsCollection` is arbitrary but descriptive (any identifier ≠ `Alerts` compiles); bike-shedding welcome at review.

### D3: Ship the override in the same PR as the astro v1.4.0 bump

The override cannot land ahead of the upstream bump: verified that against v1.2.6 (which has no `astro_alerts` resource) tfgen panics in the bridge's autonaming pass on the unknown-resource entry. Therefore this change's PR does the v1.2.6 → v1.4.0 bump + override + regen together (phases 3–4 of the bridge-upgrade-flow rule), after which the daily bot resumes handling future bumps. This is the one place this change touches the dependency itself, and it follows the documented manual upgrade flow.

### D4: Mirror in pulumi-fivetran as a sibling OpenSpec change

Same `GOTOOLCHAIN` decision, no codegen fix needed (no pending upstream tag). Kept as its own change in that repo per the established pattern (migrate-mise-env-layering precedent).

## Risks / Trade-offs

- [Go may run a toolchain newer than `mise.lock`'s in CI] → Downloads are exact-version, checksummed against the Go toolchain sumdb, and pinned by committed `go` directives; `mise up` periodically re-syncs the installed baseline. A stale lock degrades to a download, not a failure.
- [`golangci-lint` can lag a new Go language version and fail the lint job on a bot PR] → Pre-existing risk under any GOTOOLCHAIN setting (under the exact pin it failed earlier and harder). Resolved when it occurs by `mise up` bumping golangci-lint; the failure mode is a red PR check, not a silent merge.
- [`AlertsCollection` naming is permanent API surface for the dotnet SDK] → It's a brand-new resource (v1.4.0); no existing consumers. Rename cost is zero until the upgrade PR merges.
- [Auto-downloaded toolchains add a network dependency to CI] → Only on the window between a bridge Go-bump and the next `mise up`; mise itself already downloads tools from the network in the same jobs.

## Migration Plan

1. Apply this change (mise env + resources.go override + astro v1.4.0 bump + regen) as a normal PR; CI proves all SDKs build.
2. Apply the fivetran sibling change (one-line mise env) as a PR there.
3. Next scheduled runs (05:00 UTC) should go green: fivetran bot opens the bridge v3.134.0 PR; astronomer bot finds bridge/upstream current (or opens the next available bump).
4. Rollback: revert the mise env line (restores the exact-pin behavior); the resources.go override only matters while upstream ≥ v1.4.0 is in go.mod, and reverting the whole PR restores v1.2.6.

## Open Questions

- None blocking. Optional follow-up: teach the `Sync go.work go directive` step to also drop a stale `toolchain` line from `go.work` (observed to be handled by Go itself on the next workspace command; cosmetic).
