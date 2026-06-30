## 1. Align the shim with the canonical bridged-PF pattern

- [ ] 1.1 In `provider/shim/shim.go`, remove the import of `github.com/ryan-pip/pulumi-astronomer/provider/pkg/version` and change `return p.New(version.Version)()` to `return p.New("dev")()`. Keep the existing `func NewProvider() provider.Provider` signature (no parameter), matching `pulumi-random`/`pulumi-fivetran`. Do **not** change `provider/resources.go` — `shimprovider.NewProvider()` already takes no arguments.
- [ ] 1.2 In `provider/shim/go.mod`, remove the `github.com/ryan-pip/pulumi-astronomer/provider v0.0.0-20240607020557-47813651c291` requirement, then run `go mod tidy` in `provider/shim`. The result should require only `github.com/astronomer/terraform-provider-astro` and `github.com/hashicorp/terraform-plugin-framework` (plus their indirect deps), mirroring the sibling `pulumi-fivetran` shim.
- [ ] 1.3 Remove the `exclude github.com/pulumi/pulumi-java/pkg v0.9.9` directive (and its explanatory comment) from `provider/go.mod`.

## 2. Verify

- [ ] 2.1 In `provider/shim`, confirm `go mod why github.com/ryan-pip/pulumi-astronomer/provider` reports the module is not needed, and the parent module no longer appears in `provider/shim/go.mod`.
- [ ] 2.2 In `provider/`, confirm `go mod graph | grep pulumi-java/pkg@v0.9.9` returns nothing **without** the `exclude` directive.
- [ ] 2.3 Build the generator and provider: `go build ./cmd/pulumi-tfgen-astronomer` and `go build ./pkg/...` both succeed.
- [ ] 2.4 Run the `prerequisites`, `go-test`, and `lint` jobs (locally or via CI) and confirm all three pass.
- [ ] 2.5 Regenerate the schema and confirm there is **no** diff — the upstream version literal change must not alter the generated Pulumi schema or the reported provider version (which comes from `resources.go`'s `version.Version`), per the `regenerate-sdks` workflow.
