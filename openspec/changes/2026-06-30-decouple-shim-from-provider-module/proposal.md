## Why

The Terraform shim module (`provider/shim`) has a **circular self-reference**: its `go.mod` requires the parent provider module `github.com/ryan-pip/pulumi-astronomer/provider` at a frozen 2026 pseudo-version (`v0.0.0-20240607020557-47813651c291`), purely so `shim.go` can import `…/provider/pkg/version` and pass `version.Version` into the upstream provider constructor.

That stale snapshot's `go.mod` pins the obsolete `github.com/pulumi/pulumi-java/pkg v0.9.9` submodule. As long as the bridge resolved `pulumi-java/pkg/codegen/java` through the `/pkg` submodule (bridge ≤ v3.128.0, which used `pulumi-java/pkg v1.21.3`), MVS upgraded `v0.9.9` to `v1.21.3` and a single module served the package — no conflict. But pulumi-java reorganized its modules: in v1.x the `codegen/java` package moved into the **root** `github.com/pulumi/pulumi-java` module at path `pkg/codegen/java`. When `pulumi-terraform-bridge` v3.133.0 switched to importing it from the root module (`pulumi-java v1.30.0`), the still-present `/pkg@v0.9.9` submodule began serving the *same* import path — producing an `ambiguous import` build failure across `prerequisites`, `go-test`, and `lint` (PR #60).

A targeted `exclude github.com/pulumi/pulumi-java/pkg v0.9.9` in `provider/go.mod` resolves the immediate failure, but it is a band-aid: the circular shim self-reference remains, leaving the build fragile to future bridge/dependency reshuffles. This change removes the root cause so the `exclude` is no longer needed.

## What Changes

- Change `shim.NewProvider()` in `provider/shim/shim.go` to accept the version as a parameter — `NewProvider(version string)` — and drop the import of `github.com/ryan-pip/pulumi-astronomer/provider/pkg/version`. Version injection stays at the provider layer (`resources.go`), where the ldflags-populated `version.Version` already lives, instead of reaching across the module boundary.
- Update the call site in `provider/resources.go`: `shimprovider.NewProvider()` → `shimprovider.NewProvider(version.Version)`.
- Remove the `github.com/ryan-pip/pulumi-astronomer/provider` requirement from `provider/shim/go.mod` and re-tidy the shim module, dropping the stale transitive dependencies (including `pulumi-java/pkg@v0.9.9`) the self-reference dragged in.
- Remove the `exclude github.com/pulumi/pulumi-java/pkg v0.9.9` directive from `provider/go.mod`, since the obsolete submodule no longer enters the build graph once the circular reference is gone.

## Capabilities

### New Capabilities
- `provider-module-structure`: the Terraform shim module must not depend on the parent provider module.

### Modified Capabilities
<!-- none -->

## Impact

- Modified: `provider/shim/shim.go`, `provider/shim/go.mod`, `provider/shim/go.sum`, `provider/resources.go`, `provider/go.mod`.
- Build/module hygiene only — no change to the generated Pulumi schema or to any SDK. The version string flows to the upstream provider exactly as before; only the path it takes (a function argument instead of a cross-module import) changes.
- Removes the dependency-resolution fragility that broke the bridge bump in PR #60. Future bridge/pulumi-java upgrades no longer risk reintroducing the obsolete `pulumi-java/pkg` submodule via the shim.
