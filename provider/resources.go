// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package astronomer

import (
	_ "embed"
	"fmt"
	"path/filepath"

	shimprovider "github.com/astronomer/terraform-provider-astro/shim"
	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/ryan-pip/pulumi-astronomer/provider/pkg/version"
)

//go:embed cmd/pulumi-resource-astronomer/bridge-metadata.json
var bridgeMetadata []byte

// all of the token components used below.
const (
	mainPkg = "astronomer"
	mainMod = "index" // the astronomer module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pf.ShimProvider(shimprovider.NewProvider())

	delegateID := func(pulumiField string) tfbridge.ComputeID {
		return tfbridge.DelegateIDField(resource.PropertyKey(pulumiField),
			"astro", "https://github.com/ryan-pip/pulumi-astronomer")
	}

	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "astro",
		DisplayName:       "Astronomer",
		Publisher:         "ryan-pip",
		LogoURL:           "https://raw.githubusercontent.com/ryan-pip/pulumi-astronomer/main/docs/astronomer.svg",
		PluginDownloadURL: "github://api.github.com/ryan-pip/pulumi-astronomer",
		Description:       "A Pulumi package for creating and managing Astronomer Cloud resources",
		Keywords: []string{
			"pulumi",
			"astronomer",
			"category/infrastructure",
		},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/ryan-pip/pulumi-astronomer",
		Repository:        "https://github.com/ryan-pip/pulumi-astronomer",
		Version:           version.Version,
		GitHubOrg:         "astronomer",
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
		TFProviderVersion: "0.3.0",
		Config: map[string]*tfbridge.SchemaInfo{
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ASTRO_API_TOKEN"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"astro_hybrid_cluster_workspace_authorization": {
				ComputeID: delegateID("clusterId"),
			},
			"astro_team_roles": {ComputeID: delegateID("teamId")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@ryan-pip/pulumi_astronomer",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumi_astronomer",

			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/ryan-pip/pulumi-%[1]s/sdk/", "astronomer"),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				"astronomer",
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "RyanPip",

			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.ryan-pip",
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("astro_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.SetAutonaming(255, "-")

	return prov
}
