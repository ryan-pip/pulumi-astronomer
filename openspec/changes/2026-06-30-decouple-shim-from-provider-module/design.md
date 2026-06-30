## Context

`provider/shim` is a small Go module (`github.com/astronomer/terraform-provider-astro/shim`, wired in via `replace … => ./shim`) whose job is to expose the upstream Terraform provider's `provider.Provider` to the Pulumi bridge. Its only file, `shim.go`, calls the upstream constructor `p.New(version)()`.

Today the shim obtains that `version` by importing `github.com/ryan-pip/pulumi-astronomer/provider/pkg/version` — a package in the **parent** module. To satisfy that import, `provider/shim/go.mod` declares `require github.com/ryan-pip/pulumi-astronomer/provider v0.0.0-20240607020557-47813651c291`. This is a circular dependency (the shim is itself a sub-tree of the provider repo) pinned to a stale snapshot, and the snapshot's transitive closure is what reintroduces obsolete dependencies into the build graph.

## Goals / Non-Goals

- **Goal:** Remove the shim's dependency on the parent provider module so the build graph no longer inherits the parent's historical transitive dependencies.
- **Goal:** Keep version injection behavior identical — the value still originates from `provider/pkg/version.Version` (populated by the Go linker at build time).
- **Non-Goal:** Changing how `version.Version` itself is set (ldflags at build time) or any provider runtime behavior.
- **Non-Goal:** Altering the generated Pulumi schema or any SDK output.

## Decision

Invert the dependency direction: instead of the shim *reaching up* into the parent module for the version, the parent *passes the version down* as a function argument.

```go
// provider/shim/shim.go
func NewProvider(version string) provider.Provider {
    return p.New(version)()
}
```

```go
// provider/resources.go (call site)
p := pf.ShimProvider(shimprovider.NewProvider(version.Version))
```

`resources.go` already imports `…/provider/pkg/version`, so no new coupling is introduced at the provider layer; the version simply crosses the module boundary as data rather than as a package import. With the import gone, `provider/shim/go.mod` no longer needs to require the parent module, and `go mod tidy` removes the stale transitive deps — including `pulumi-java/pkg@v0.9.9`. The `exclude` band-aid in `provider/go.mod` then becomes unnecessary and is removed.

## Risks / Trade-offs

- **Schema drift:** A signature change in the shim could in principle ripple into generated artifacts. Mitigation: task 2.4 regenerates the schema and requires a zero diff before the change is accepted.
- **`upgrade-provider` regeneration:** The automated `upgrade-provider` tool regenerates `go.mod`/shim wiring on future upstream bumps. The pinned pseudo-version has been stable across runs (it was identical on `main` and PR #60), so the tool does not currently rewrite the shim's require block; this change should survive future runs. If a future `upgrade-provider` run reintroduces a parent-module require in the shim, that regression would be caught by the `provider-module-structure` spec check.

## Migration Plan

Single atomic change; no staged rollout. Apply the source edits, tidy the shim module, drop the `exclude`, and verify the three CI jobs plus a clean schema regeneration.
