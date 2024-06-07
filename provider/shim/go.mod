module github.com/GK-Consulting/terraform-provider-astronomer/shim

go 1.21.5

toolchain go1.21.6

replace (
	github.com/GK-Consulting/terraform-provider-astronomer => ../../upstream
	github.com/openglshaders/astronomer-api/v2 => ../../upstream/internal/api
	terraform-provider-astronomer => github.com/ryan-pip/terraform-provider-astronomer v0.0.0-00010101000000-000000000000
)

require (
	github.com/GK-Consulting/terraform-provider-astronomer v0.0.0-00010101000000-000000000000
	github.com/hashicorp/terraform-plugin-framework v1.4.2
	github.com/ryan-pip/pulumi-astronomer/provider v0.0.0-20240311084552-d5ad20b3423b
)

require (
	github.com/fatih/color v1.15.0 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.19.1 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/openglshaders/astronomer-api/v2 v2.0.0-00010101000000-000000000000 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
)
