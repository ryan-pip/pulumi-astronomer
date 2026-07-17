# Unblock daily upgrade-provider: Go toolchain drift + astro_alerts C# codegen

## Why

The scheduled `upgrade-provider` workflow has failed every day since 2026-07-09, so no upstream/bridge upgrade PRs are being opened. Two independent failures are stacked (verified by reproducing both locally against `origin/main`):

1. **Go toolchain pin drift (fails first, both this repo and pulumi-fivetran).** `mise.toml` `[env]` hard-pins `GOTOOLCHAIN = "go1.25.9"` — a second, hand-maintained copy of the Go version that `mise.lock` already pins (currently go 1.25.11). The pin drifted behind the lockfile (PR #60 bumped the lock, not the env pin), and bridge v3.134.0 now declares `go >= 1.25.11`, so the workflow's `go get pulumi-terraform-bridge/v3` hard-fails: `requires go >= 1.25.11 (running go 1.25.9; GOTOOLCHAIN=go1.25.9)`. Because mise shims re-export `[env]` over the caller's environment, the pin is absolute in CI. This is the third time this pin has needed a hand bump (`go1.24.13` → `go1.25.9` → …); the class of failure recurs every time the bridge adopts a newer Go patch release.
2. **astro_alerts C# member/type collision (masked by 1; failed 07-09/07-10 before bridge v3.134.0 existed).** The pending upstream upgrade `terraform-provider-astro` v1.2.6 → v1.4.0 adds an `astro_alerts` resource. tfgen generates C# class `Alerts` with an output property `Alerts`, which C# forbids (CS0542: member names cannot be the same as their enclosing type), so `make generate_sdks` dies at `build_dotnet` and no upgrade PR is opened.

## What Changes

- Replace `GOTOOLCHAIN = "go1.25.9"` with `GOTOOLCHAIN = "auto"` in `mise.toml` `[env]`, removing the duplicated exact Go version. `mise.lock` stays the single pinned source of the installed toolchain; module `go` directives (managed by `go get`/upgrade-provider and the workflow's go.work sync step) declare the minimum, and Go resolves/downloads a newer checksummed toolchain only when a module requires it. Verified: with `auto`, `go get pulumi-terraform-bridge/v3@v3.134.0` succeeds, bumps `provider/go.mod` to go 1.25.11, and the workspace builds.
- Add an `astro_alerts` entry to the `Resources` map in `provider/resources.go` with `Fields: {"alerts": {CSharpName: "AlertsCollection"}}`, then regenerate schema + SDKs. Verified: dotnet SDK builds with 0 errors after the override; other language SDKs are unaffected by `CSharpName`.
- Mirror the `GOTOOLCHAIN` change in `pulumi-fivetran` via its own OpenSpec change of the same name (fivetran has no pending upstream tag and no codegen collision; it only needs the toolchain fix).

## Capabilities

### New Capabilities

- `sdk-generation`: generated SDKs must compile in all four languages; upstream schema additions that collide with language keywords/rules (here: C# CS0542) are resolved with bridge naming overrides (`CSharpName`) in `resources.go`, not by patching generated code.

### Modified Capabilities

- `developer-tooling`: new requirement — the mise env SHALL NOT hard-pin an exact Go toolchain (`GOTOOLCHAIN` = exact version). Toolchain selection SHALL be `auto` so the exact installed version is pinned in one place (`mise.lock`) and module `go` directives can pull a newer toolchain without manual intervention.

## Impact

- `mise.toml` (`[env]` block) — one-line change; affects every mise-run Go invocation (CI and dev) in this repo, and the same in pulumi-fivetran.
- `provider/resources.go` — one `Resources` map entry; `make schema build_sdks` regen follows (schema.json + all four SDKs pick up astro v1.4.0 when the bot PR lands, or the override merges first and the bot regen lands cleanly on top).
- No workflow YAML changes. The existing `Sync go.work go directive` step already handles the directive skew the bridge bump creates.
- Risk note: with `auto`, CI may download a Go toolchain newer than `mise.lock`'s (checksummed via the Go toolchain sumdb, recorded in committed `go` directives). `mise.lock` and `golangci-lint` should still be bumped periodically (`mise up`), but a stale lockfile no longer breaks the daily automation.
