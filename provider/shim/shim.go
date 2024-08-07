package shim

import (
	p "github.com/astronomer/terraform-provider-astro/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/ryan-pip/pulumi-astronomer/provider/pkg/version"
)

func NewProvider() provider.Provider {
	return p.New(version.Version)()
}
