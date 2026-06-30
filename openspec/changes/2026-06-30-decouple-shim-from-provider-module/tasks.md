## 1. Align the shim with the canonical bridged-PF pattern

- [x] 1.1 In `provider/shim/shim.go`, remove the import of `github.com/ryan-pip/pulumi-astronomer/provider/pkg/version` and change `return p.New(version.Version)()` to `return p.New("dev")()`. Keep the existing `func NewProvider() provider.Provider` signature (no parameter), matching `pulumi-random`/`pulumi-fivetran`. Do **not** change `provider/resources.go` — `shimprovider.NewProvider()` already takes no arguments.
- [x] 1.2 In `provider/shim/go.mod`, remove the `github.com/ryan-pip/pulumi-astronomer/provider v0.0.0-20240607020557-47813651c291` requirement, then run `go mod tidy` in `provider/shim`. The result should require only `github.com/astronomer/terraform-provider-astro` and `github.com/hashicorp/terraform-plugin-framework` (plus their indirect deps), mirroring the sibling `pulumi-fivetran` shim.
- [x] 1.3 Remove the `exclude github.com/pulumi/pulumi-java/pkg v0.9.9` directive (and its explanatory comment) from `provider/go.mod`.

## 2. Verify

- [x] 2.1 In `provider/shim`, confirm the parent module no longer appears in `provider/shim/go.mod`. (Note: `go mod why` runs in workspace mode here, where the parent is a workspace member, so the `go.mod` require list is the authoritative check — and it no longer lists the parent.)
- [x] 2.2 In `provider/`, confirm `go mod graph | grep pulumi-java/pkg@v0.9.9` returns nothing **without** the `exclude` directive. (Verified: count 0.)
- [x] 2.3 Build the generator and provider: `go build ./cmd/pulumi-tfgen-astronomer` and `go build ./pkg/...` both succeed. (Both build; provider root package also builds.)
- [x] 2.4 Run the `prerequisites`, `go-test`, and `lint` jobs and confirm all three pass. VERIFIED locally with the locked toolchain: `lint` (exact CI command) reports `0 issues`; `go test -short ./...` passes (`ok` for the provider package, no failures); and `make provider` builds `bin/pulumi-resource-astronomer` cleanly (prerequisites).
- [x] 2.5 Regenerate the schema and confirm there is **no** diff. VERIFIED: `make schema` with the locked toolchain (pulumi 3.248.0) produces a byte-identical `schema.json`, `bridge-metadata.json`, and `schema-embed.json` — zero drift. The change only swaps the version string handed to the upstream provider, which is not part of the resource schema; since the SDKs derive from the schema, they are unaffected as well.
