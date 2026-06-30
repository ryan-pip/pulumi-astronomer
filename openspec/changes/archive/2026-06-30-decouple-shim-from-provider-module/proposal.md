## Why

The Terraform shim module (`provider/shim`) has a **circular self-reference**: its `go.mod` requires the parent provider module `github.com/ryan-pip/pulumi-astronomer/provider` at a frozen 2024 pseudo-version (`v0.0.0-20240607020557-47813651c291`), purely so `shim.go` can import `…/provider/pkg/version` and pass `version.Version` into the upstream provider constructor (`p.New(version.Version)()`).

That stale snapshot's `go.mod` pins the obsolete `github.com/pulumi/pulumi-java/pkg v0.9.9` submodule. As long as the bridge resolved `pulumi-java/pkg/codegen/java` through the `/pkg` submodule (bridge ≤ v3.128.0, which used `pulumi-java/pkg v1.21.3`), MVS upgraded `v0.9.9` to `v1.21.3` and a single module served the package — no conflict. But pulumi-java reorganized its modules: in v1.x the `codegen/java` package moved into the **root** `github.com/pulumi/pulumi-java` module at path `pkg/codegen/java`. When `pulumi-terraform-bridge` v3.133.0 switched to importing it from the root module (`pulumi-java v1.30.0`), the still-present `/pkg@v0.9.9` submodule began serving the *same* import path — producing an `ambiguous import` build failure across `prerequisites`, `go-test`, and `lint` (PR #60).

A targeted `exclude github.com/pulumi/pulumi-java/pkg v0.9.9` in `provider/go.mod` resolves the immediate failure, but it is a band-aid: the circular shim self-reference remains, leaving the build fragile to future bridge/dependency reshuffles.

**This deviates from how bridged Plugin-Framework providers are actually built.** The canonical shim pattern — the example the bridge's own PF guide points at (`pulumi-random/provider/shim`), and the pattern this repo's sibling `ryan-pip/pulumi-fivetran` already follows — has the shim depend **only** on the upstream Terraform provider and `terraform-plugin-framework`, with **no dependency on the parent Pulumi provider module**:

- `pulumi-random/provider/shim/go.mod`: requires only `github.com/terraform-providers/terraform-provider-random` and `github.com/hashicorp/terraform-plugin-framework`; `shim.go` is `func NewProvider() tfpf.Provider { return provider.New()() }`.
- `ryan-pip/pulumi-fivetran/provider/shim/go.mod`: requires only `github.com/fivetran/terraform-provider-fivetran` and `github.com/hashicorp/terraform-plugin-framework`; `shim.go` is `func NewProvider() provider.Provider { return p.FivetranProvider() }` — no parent import.
- Where the upstream constructor takes a version argument (as astronomer's `p.New(version)` does), the established convention is to pass a literal such as `"dev"` (e.g. the airbyte shim's `provider.New("dev")()`), not to reach into the parent for the build version.

`astronomer` is the outlier: it is the only one importing the parent's `version` package into the shim, and that is the sole reason `provider/shim/go.mod` requires the parent module. Aligning with the canonical pattern removes the root cause so the `exclude` is no longer needed.

## What Changes

- In `provider/shim/shim.go`, drop the import of `github.com/ryan-pip/pulumi-astronomer/provider/pkg/version` and pass a constant version literal to the upstream constructor — `return p.New("dev")()` — matching the canonical shim pattern (`pulumi-random`, the airbyte example) and this repo's sibling `pulumi-fivetran`. `NewProvider()` keeps its existing no-argument signature, so the call site in `resources.go` (`shimprovider.NewProvider()`) is unchanged.
- Remove the `github.com/ryan-pip/pulumi-astronomer/provider` requirement from `provider/shim/go.mod` and re-tidy the shim module, leaving only the upstream `terraform-provider-astro` and `terraform-plugin-framework` (plus their indirect deps) — the same dependency surface `pulumi-random` and `pulumi-fivetran` shims have. This drops the stale transitive dependencies (including `pulumi-java/pkg@v0.9.9`) the self-reference dragged in.
- Remove the `exclude github.com/pulumi/pulumi-java/pkg v0.9.9` directive from `provider/go.mod`, since the obsolete submodule no longer enters the build graph once the circular reference is gone.

## Capabilities

### New Capabilities
- `provider-module-structure`: the Terraform shim module must follow the canonical bridged-PF pattern and not depend on the parent provider module.

### Modified Capabilities
<!-- none -->

## Impact

- Modified: `provider/shim/shim.go`, `provider/shim/go.mod`, `provider/shim/go.sum`, `provider/go.mod`. `provider/resources.go` is **not** changed.
- Build/module hygiene only — no change to the generated Pulumi schema or to any SDK. The operative Pulumi provider version is still set by `resources.go` (`info.Version = version.Version`, unchanged). The only behavioral difference is that the **upstream** Terraform provider's internal version string becomes the conventional `"dev"` literal instead of the build version — the same trade-off pulumi-random and pulumi-fivetran already make, and it does not affect the bridged provider's reported version, schema, or resources.
- Removes the dependency-resolution fragility that broke the bridge bump in PR #60, and brings `pulumi-astronomer`'s shim in line with `pulumi-fivetran` and the upstream canonical example.
