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
	"context"
	"os"
	"strings"
	"testing"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
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

// resourcesWithoutStableID have no upstream "id" and nothing scalar to delegate
// to, so tfbridge.MissingIDPlaceholder stands. Value is the reason.
var resourcesWithoutStableID = map[string]string{
	"astro_alerts": "upstream exposes only a required `alerts` map",
}

func TestProviderInfo(t *testing.T) {
	info := Provider()

	if info.Name != "astro" {
		t.Errorf("Name = %q, want %q", info.Name, "astro")
	}
	if info.Version == "" {
		t.Errorf("Version is empty; expected a non-empty version")
	}

}

// TestTokensAreWellFormed guards the surface user code is written against. The
// maps in ProviderInfo are keyed by Terraform name, so a bridge upgrade that
// changed tokenization would rename every Pulumi type without disturbing a
// single key.
func TestTokensAreWellFormed(t *testing.T) {
	info := Provider()
	prefix := mainPkg + ":" + mainMod + "/"

	check := func(kind, tfName, tok string) {
		t.Helper()
		switch {
		case tok == "":
			t.Errorf("%s %s: empty token, so MustComputeTokens left it unmapped", kind, tfName)
		case !strings.HasPrefix(tok, prefix):
			t.Errorf("%s %s: token %q, want prefix %q", kind, tfName, tok, prefix)
		}
	}

	if len(info.Resources) == 0 {
		t.Fatal("no resources mapped")
	}
	for name, r := range info.Resources {
		check("resource", name, string(r.Tok))
	}

	if len(info.DataSources) == 0 {
		t.Fatal("no data sources mapped")
	}
	for name, d := range info.DataSources {
		check("data source", name, string(d.Tok))
	}
}

// idOf reports what the bridge would hand Pulumi as this resource's ID.
// MustComputeTokens installs a ComputeID on every resource lacking an upstream
// "id", so a non-nil ComputeID proves nothing — only calling it does.
func idOf(t *testing.T, res *tfbridge.ResourceInfo) (string, bool) {
	t.Helper()
	if res.ComputeID == nil {
		return "", false // upstream's computed string "id" fills the ID slot
	}
	id, err := res.ComputeID(context.Background(), resource.PropertyMap{})
	if err != nil {
		return "", false // a real delegate, asking for state we didn't supply
	}
	return string(id), true
}

// TestEveryResourceHasAUsableID catches the one ID defect tfgen stays quiet
// about: with no upstream "id" and no delegate, tokens.fixMissingID installs
// MissingIDComputeID and every instance of the resource answers to "missing ID".
func TestEveryResourceHasAUsableID(t *testing.T) {
	for name, res := range Provider().Resources {
		id, ok := idOf(t, res)
		if !ok || id != tfbridge.MissingIDPlaceholder {
			continue
		}
		if _, allowed := resourcesWithoutStableID[name]; allowed {
			continue
		}
		t.Errorf("%s: no upstream \"id\" and no delegate, so every instance gets %q. "+
			"Add a delegateID entry in resources.go, or list it in "+
			"resourcesWithoutStableID with a reason.", name, tfbridge.MissingIDPlaceholder)
	}
}

// TestResourcesWithoutStableIDAreStillBroken keeps the allowlist from rotting
// once upstream gives one of these resources a real ID.
func TestResourcesWithoutStableIDAreStillBroken(t *testing.T) {
	resources := Provider().Resources
	for name, reason := range resourcesWithoutStableID {
		res, ok := resources[name]
		if !ok {
			t.Errorf("%s: allowlisted but no longer mapped; drop it from resourcesWithoutStableID", name)
			continue
		}
		if id, got := idOf(t, res); !got || id != tfbridge.MissingIDPlaceholder {
			t.Errorf("%s: now resolves a real ID (%q) - drop it from resourcesWithoutStableID (was: %s)",
				name, id, reason)
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
