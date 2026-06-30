## ADDED Requirements

### Requirement: The shim module follows the canonical bridged-PF pattern

The Terraform shim module (`provider/shim`, module path `github.com/astronomer/terraform-provider-astro/shim`) SHALL follow the canonical Pulumi bridged Plugin-Framework shim pattern used by `pulumi-random/provider/shim` (the bridge's `new-pf-provider.md` reference example) and by this repo's sibling `pulumi-fivetran`. Specifically, the shim's `go.mod` SHALL require only the upstream Terraform provider (`github.com/astronomer/terraform-provider-astro`) and `github.com/hashicorp/terraform-plugin-framework` (plus their indirect dependencies), and SHALL NOT declare a requirement on — nor import any package from — the parent Pulumi provider module `github.com/ryan-pip/pulumi-astronomer/provider`. Where the upstream constructor takes a version argument, the shim SHALL pass a constant literal (e.g. `"dev"`) rather than importing the parent's `version` package. As a consequence, `provider/go.mod` SHALL NOT need an `exclude` directive for the obsolete `github.com/pulumi/pulumi-java/pkg` submodule in order to build.

#### Scenario: shim go.mod has no parent-module requirement

- **WHEN** `provider/shim/go.mod` is inspected
- **THEN** it contains no `require` (direct or indirect) on `github.com/ryan-pip/pulumi-astronomer/provider`, and `provider/shim/shim.go` imports nothing from that module

#### Scenario: shim dependency surface matches sibling providers

- **WHEN** the shim module's direct requirements are listed
- **THEN** they are exactly the upstream `github.com/astronomer/terraform-provider-astro` and `github.com/hashicorp/terraform-plugin-framework`, the same shape as the `pulumi-random` and `pulumi-fivetran` shims

#### Scenario: version literal is passed to the upstream constructor

- **WHEN** `shim.NewProvider()` constructs the upstream provider
- **THEN** it calls `p.New("dev")()` with a constant literal, and the bridged provider's reported version is still set independently in `resources.go` via `info.Version = version.Version`

#### Scenario: build resolves pulumi-java unambiguously without an exclude

- **WHEN** the provider module is built with no `exclude github.com/pulumi/pulumi-java/pkg` directive in `provider/go.mod`
- **THEN** `go build ./cmd/pulumi-tfgen-astronomer` and `go build ./pkg/...` succeed and `go mod graph` contains no `github.com/pulumi/pulumi-java/pkg@v0.9.9`
