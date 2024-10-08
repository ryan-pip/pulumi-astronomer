// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

var _ = internal.GetEnvOrDefault

// API host to use for the provider. Default is `https://api.astronomer.io`
func GetHost(ctx *pulumi.Context) string {
	return config.Get(ctx, "astronomer:host")
}

// Organization ID this provider will operate on.
func GetOrganizationId(ctx *pulumi.Context) string {
	return config.Get(ctx, "astronomer:organizationId")
}

// Astro API Token. Can be set with an `ASTRO_API_TOKEN` env var.
func GetToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "astronomer:token")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "ASTRO_API_TOKEN"); d != nil {
		value = d.(string)
	}
	return value
}
