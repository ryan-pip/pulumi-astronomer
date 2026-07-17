# sdk-generation Specification

## ADDED Requirements

### Requirement: Generated SDKs compile in all four languages

`make generate_sdks` output SHALL compile for every published SDK (nodejs, python, go, dotnet). When an upstream schema addition produces an identifier that a target language rejects (e.g. C# CS0542: a member named the same as its enclosing type), the collision SHALL be resolved with the bridge's per-language naming override in `provider/resources.go` (`SchemaInfo.CSharpName` for C#), never by hand-editing generated files under `sdk/`.

#### Scenario: astro_alerts C# member/type collision is resolved by override

- **WHEN** the schema is generated with upstream `terraform-provider-astro` >= v1.4.0, whose `astro_alerts` resource has an `alerts` property
- **THEN** `provider/resources.go` maps `astro_alerts.alerts` with `CSharpName: "AlertsCollection"` and `cd sdk/dotnet && dotnet build` succeeds with zero errors

#### Scenario: naming overrides do not leak into other SDKs

- **WHEN** a `CSharpName` override is applied to a property
- **THEN** the Pulumi schema property name and the nodejs/python/go SDK member names are unchanged (only the dotnet member is renamed)

### Requirement: Naming overrides reference only existing upstream entities

A `Resources`/`DataSources` override entry in `provider/resources.go` SHALL only be introduced together with (or after) the upstream version that defines the referenced entity, because tfgen fails against an upstream version that lacks it.

#### Scenario: override is coupled to the upstream bump

- **WHEN** the `astro_alerts` override lands
- **THEN** the same commit series bumps `provider/go.mod` (and shim) to `terraform-provider-astro` >= v1.4.0 and regenerates schema + SDKs, so tfgen never runs with the override against v1.2.6
