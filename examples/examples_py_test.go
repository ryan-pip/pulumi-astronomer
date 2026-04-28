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

//go:build python || all
// +build python all

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	basePython := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePython
}

// requireLiveCreds skips the test unless workspace-scoped Astronomer credentials
// are present in the environment. Live tests run locally via `make test_python`
// and in CI via the `live` job (gated on the astronomer-live environment).
func requireLiveCreds(t *testing.T) string {
	if os.Getenv("ASTRO_API_TOKEN") == "" {
		t.Skip("ASTRO_API_TOKEN not set; skipping live test")
	}
	workspaceID := os.Getenv("ASTRO_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("ASTRO_WORKSPACE_ID not set; skipping live test")
	}
	return workspaceID
}

// TestAccApiTokenPy creates a workspace-scoped ApiToken and tears it down.
// Workspace-scoped, free on the developer tier.
func TestAccApiTokenPy(t *testing.T) {
	workspaceID := requireLiveCreds(t)
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "api_token", "python"),
			Config: map[string]string{
				"workspaceId": workspaceID,
			},
		})
	integration.ProgramTest(t, &test)
}

// TestAccGetWorkspacePy exercises the read-only get_workspace data source.
// Read-only, free; proves the bridge's invoke path works end-to-end.
func TestAccGetWorkspacePy(t *testing.T) {
	workspaceID := requireLiveCreds(t)
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "get_workspace", "python"),
			Config: map[string]string{
				"workspaceId": workspaceID,
			},
		})
	integration.ProgramTest(t, &test)
}
