## Context

`provider/shim` is a small Go module (`github.com/astronomer/terraform-provider-astro/shim`, wired in via `replace … => ./shim`) whose job is to expose the upstream Terraform provider's `provider.Provider` to the Pulumi bridge, because Plugin-Framework providers define their constructor in an `internal/` package that cannot be imported directly.

Today the shim obtains the version it forwards to the upstream constructor by importing `github.com/ryan-pip/pulumi-astronomer/provider/pkg/version` — a package in the **parent** module:

```go
// provider/shim/shim.go (current)
import (
    p "github.com/astronomer/terraform-provider-astro/internal/provider"
    "github.com/hashicorp/terraform-plugin-framework/provider"
    "github.com/ryan-pip/pulumi-astronomer/provider/pkg/version"
)
func NewProvider() provider.Provider {
    return p.New(version.Version)()
}
```

To satisfy that import, `provider/shim/go.mod` declares `require github.com/ryan-pip/pulumi-astronomer/provider v0.0.0-20240607020557-47813651c291`. This is a circular dependency (the shim is a sub-tree of the provider repo) pinned to a stale snapshot, and the snapshot's transitive closure is what reintroduces obsolete dependencies (notably `pulumi-java/pkg@v0.9.9`) into the build graph.

## Research: what real bridged-PF providers do

The fix is not invented — it matches the established pattern verified across three independent sources:

- **`pulumi-random/provider/shim`** — the example the bridge's own `new-pf-provider.md` guide points at. `shim.go`: `func NewProvider() tfpf.Provider { return provider.New()() }`. `go.mod` requires only `github.com/terraform-providers/terraform-provider-random` and `github.com/hashicorp/terraform-plugin-framework`. No parent Pulumi module, no `replace`.
- **`ryan-pip/pulumi-fivetran/provider/shim`** — this repo's sibling, same author and generator. `shim.go`: `func NewProvider() provider.Provider { return p.FivetranProvider() }` (no parent import). `go.mod` requires only `github.com/fivetran/terraform-provider-fivetran` and `github.com/hashicorp/terraform-plugin-framework`. No parent module, no `replace`.
- **airbyte shim (bridge blog example)** — where the upstream constructor takes a version argument, it passes a literal: `return provider.New("dev")()`.

In none of these does the shim import the parent Pulumi provider's `version` package or require the parent module. `astronomer` is the only one that does.

## Goals / Non-Goals

- **Goal:** Make the shim follow the canonical pattern — depend only on the upstream Terraform provider + `terraform-plugin-framework`, with no dependency on the parent Pulumi provider module — so the build graph no longer inherits the parent's historical transitive dependencies.
- **Non-Goal:** Changing how the operative Pulumi provider version is set. That is `info.Version = version.Version` in `resources.go` and stays exactly as-is.
- **Non-Goal:** Altering the generated Pulumi schema, the reported provider version, or any SDK output.

## Decision

Match the canonical pattern: drop the parent import from the shim and pass the conventional `"dev"` literal to the upstream constructor (astronomer's `p.New` takes a version argument, like airbyte's).

```go
// provider/shim/shim.go (after)
import (
    p "github.com/astronomer/terraform-provider-astro/internal/provider"
    "github.com/hashicorp/terraform-plugin-framework/provider"
)
func NewProvider() provider.Provider {
    return p.New("dev")()
}
```

`NewProvider()` keeps its no-argument signature (identical to `pulumi-random`/`pulumi-fivetran`), so `resources.go`'s `shimprovider.NewProvider()` call site is untouched. With the import gone, `provider/shim/go.mod` no longer needs the parent module, and `go mod tidy` removes the stale transitive deps — including `pulumi-java/pkg@v0.9.9`. The `exclude` band-aid in `provider/go.mod` then becomes unnecessary and is removed.

### Why not keep the real version by adding a parameter?

An alternative is `NewProvider(version string)` called with `version.Version` from `resources.go`. That also removes the circular dependency, but it diverges from every real bridged-PF shim surveyed (all use either no argument or a literal) and is not what the user asked for ("don't make it up"). The version handed to the **upstream** provider only affects that provider's internal version string (user-agent/telemetry); the bridged provider's reported version comes from `resources.go`. So the literal-`"dev"` convention loses nothing operationally and keeps the shim identical in shape to its siblings.

## Risks / Trade-offs

- **Upstream internal version becomes `"dev"`:** The upstream Terraform provider will report `"dev"` as its internal version (e.g. in its own user-agent). This is the same trade-off `pulumi-random` and `pulumi-fivetran` already accept, and it does not affect the Pulumi provider's reported version or schema.
- **Schema drift:** A change in the shim could in principle ripple into generated artifacts. Mitigation: task 2.5 regenerates the schema and requires a zero diff before acceptance.
- **`upgrade-provider` regeneration:** The pinned pseudo-version has been stable across runs (identical on `main` and PR #60), so `upgrade-provider` does not currently rewrite the shim's require block; this change should survive future runs. If a future run reintroduces a parent-module require in the shim, the `provider-module-structure` spec check catches the regression.

## Migration Plan

Single atomic change; no staged rollout. Edit `shim.go`, tidy the shim module, drop the `exclude`, and verify the three CI jobs plus a clean schema regeneration.
