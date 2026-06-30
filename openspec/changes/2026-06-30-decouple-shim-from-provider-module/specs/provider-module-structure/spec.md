## ADDED Requirements

### Requirement: The shim module does not depend on the parent provider module

The Terraform shim module (`provider/shim`, module path `github.com/astronomer/terraform-provider-astro/shim`) SHALL NOT declare a requirement on the parent provider module `github.com/ryan-pip/pulumi-astronomer/provider`, and SHALL NOT import any package from it. The provider version that the shim forwards to the upstream provider constructor SHALL be supplied by the caller as a function argument (`NewProvider(version string)`), originating from `provider/pkg/version.Version` at the provider layer. As a consequence, `provider/go.mod` SHALL NOT need an `exclude` directive for the obsolete `github.com/pulumi/pulumi-java/pkg` submodule to build.

#### Scenario: shim go.mod has no parent-module requirement

- **WHEN** `provider/shim/go.mod` is inspected
- **THEN** it contains no `require` (direct or indirect) on `github.com/ryan-pip/pulumi-astronomer/provider`, and `provider/shim/shim.go` imports nothing from that module

#### Scenario: version is injected through the call boundary

- **WHEN** the provider constructs the shimmed upstream provider in `resources.go`
- **THEN** it calls `shimprovider.NewProvider(version.Version)`, passing the linker-populated version string as an argument rather than the shim importing it

#### Scenario: build resolves pulumi-java unambiguously without an exclude

- **WHEN** the provider module is built with no `exclude github.com/pulumi/pulumi-java/pkg` directive in `provider/go.mod`
- **THEN** `go build ./cmd/pulumi-tfgen-astronomer` and `go build ./pkg/...` succeed and `go mod graph` contains no `github.com/pulumi/pulumi-java/pkg@v0.9.9`
