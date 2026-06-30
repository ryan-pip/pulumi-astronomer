## 1. Decouple the shim from the parent provider module

- [ ] 1.1 In `provider/shim/shim.go`, change `func NewProvider() provider.Provider` to `func NewProvider(version string) provider.Provider` and pass `version` through to `p.New(version)()`. Remove the `github.com/ryan-pip/pulumi-astronomer/provider/pkg/version` import.
- [ ] 1.2 In `provider/resources.go`, update the call site `shimprovider.NewProvider()` to `shimprovider.NewProvider(version.Version)` (the `version` package is already imported there).
- [ ] 1.3 In `provider/shim/go.mod`, remove the `github.com/ryan-pip/pulumi-astronomer/provider v0.0.0-20240607020557-47813651c291` requirement, then run `go mod tidy` in `provider/shim` to drop the now-unused transitive dependencies.
- [ ] 1.4 Remove the `exclude github.com/pulumi/pulumi-java/pkg v0.9.9` directive (and its explanatory comment) from `provider/go.mod`.

## 2. Verify

- [ ] 2.1 In `provider/`, confirm `go mod graph | grep pulumi-java/pkg@v0.9.9` returns nothing without the `exclude` directive.
- [ ] 2.2 Build the generator and provider: `go build ./cmd/pulumi-tfgen-astronomer` and `go build ./pkg/...` both succeed.
- [ ] 2.3 Run `make` / the `prerequisites`, `go-test`, and `lint` targets locally (or confirm via CI) that all three pass.
- [ ] 2.4 Regenerate the schema and confirm there is **no** diff (the shim signature change must not alter the generated Pulumi schema), per the `regenerate-sdks` workflow.
