package shim

import (
	"github.com/GK-Consulting/terraform-provider-astronomer/internal/provider"
	tf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/ryan-pip/pulumi-astronomer/provider/pkg/version"
)

func NewProvider() tf.Provider {
	return provider.New(version.Version)()
}
