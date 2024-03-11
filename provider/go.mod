module github.com/ryan.pip/pulumi-astronomer/provider

go 1.21

replace (
	terraform-provider-astronomer/shim => ./shim
)

require (
	github.com/ettle/strcase v0.1.1
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.73.0
	github.com/pulumi/pulumi-terraform-bridge/pf v0.26.0
	github.com/pulumi/pulumi/sdk/v3 v3.104.2
)
