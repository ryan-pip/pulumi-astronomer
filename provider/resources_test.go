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
	"os"
	"sort"
	"testing"

	"github.com/ryan-pip/pulumi-astronomer/provider/pkg/version"
)

// TestMain sets a non-empty Version so Provider() can call
// tfbridge.GetModuleMajorVersion without panicking. In production builds the
// Makefile injects this via -ldflags; in `go test` we have to do it ourselves.
func TestMain(m *testing.M) {
	if version.Version == "" {
		version.Version = "1.0.0"
	}
	os.Exit(m.Run())
}

// expectedResources is the full list of upstream Terraform resources we map.
// Adding/removing a resource here is a deliberate change that should be
// reviewed alongside the corresponding resources.go edit.
var expectedResources = []string{
	"astro_agent_token",
	"astro_alert",
	"astro_api_token",
	"astro_cluster",
	"astro_custom_role",
	"astro_deployment",
	"astro_hybrid_cluster_workspace_authorization",
	"astro_notification_channel",
	"astro_team",
	"astro_team_membership",
	"astro_team_roles",
	"astro_user_invite",
	"astro_user_roles",
	"astro_workspace",
}

// resourcesWithCustomComputeID are the resources whose Pulumi ID we delegate
// to a non-default field (see resources.go). If any of these regress to a nil
// ComputeID we'll silently start producing wrong URN/ID pairs.
var resourcesWithCustomComputeID = []string{
	"astro_hybrid_cluster_workspace_authorization",
	"astro_team_roles",
	"astro_user_invite",
	"astro_user_roles",
}

func TestProviderInfo(t *testing.T) {
	info := Provider()

	if info.Name != "astro" {
		t.Errorf("Name = %q, want %q", info.Name, "astro")
	}
	if info.Version == "" {
		t.Errorf("Version is empty; expected a non-empty version")
	}

	for _, k := range expectedResources {
		if _, ok := info.Resources[k]; !ok {
			t.Errorf("missing expected resource mapping: %s", k)
		}
	}

	got := make([]string, 0, len(info.Resources))
	for k := range info.Resources {
		got = append(got, k)
	}
	sort.Strings(got)
	if len(got) != len(expectedResources) {
		t.Errorf("resource count = %d (%v), want %d (%v)",
			len(got), got, len(expectedResources), expectedResources)
	}

	for _, k := range resourcesWithCustomComputeID {
		r, ok := info.Resources[k]
		if !ok {
			continue // already reported above
		}
		if r.ComputeID == nil {
			t.Errorf("%s: ComputeID is nil; expected delegate ID set in resources.go", k)
		}
	}
}

func TestProviderConfigTokenEnvVar(t *testing.T) {
	info := Provider()
	tok, ok := info.Config["token"]
	if !ok {
		t.Fatalf("Config.token is missing; provider auth wiring lost")
	}
	if tok.Default == nil {
		t.Fatalf("Config.token has no Default; expected ASTRO_API_TOKEN env var binding")
	}
	found := false
	for _, env := range tok.Default.EnvVars {
		if env == "ASTRO_API_TOKEN" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Config.token EnvVars = %v, want to include ASTRO_API_TOKEN", tok.Default.EnvVars)
	}
}
