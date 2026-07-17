# Tasks — unblock-daily-upgrade-provider

## 1. Go toolchain fix

- [x] 1.1 In `mise.toml` `[env]`, change `GOTOOLCHAIN = "go1.25.9"` to `GOTOOLCHAIN = "auto"` with a one-line comment stating why an exact pin is forbidden (drift vs `mise.lock` broke the daily upgrade three times)
- [x] 1.2 Verify locally: from the repo root, `cd provider && go get github.com/pulumi/pulumi-terraform-bridge/v3@v3.134.0 && go mod tidy` succeeds and bumps `provider/go.mod` to `go 1.25.11` (do NOT commit the bridge bump itself — revert it after verifying; the bot owns bridge bumps)

## 2. astro_alerts C# collision fix + v1.4.0 upgrade

- [x] 2.1 Add to the `Resources` map in `provider/resources.go`: `"astro_alerts": {Fields: map[string]*tfbridge.SchemaInfo{"alerts": {CSharpName: "AlertsCollection"}}}` with a comment naming CS0542
- [x] 2.2 Bump upstream per bridge-upgrade-flow phase 3: `cd provider && go get github.com/astronomer/terraform-provider-astro@v1.4.0 && go mod tidy`, same in `provider/shim`
- [x] 2.3 Regenerate: `make schema build_sdks`; confirm the diff under `sdk/` is generated-code only and `schema.json` gains the v1.4.0 entities (alerts, notification channels, etc.)
- [x] 2.4 Verify the collision fix: `cd sdk/dotnet && dotnet build` → 0 errors; `grep AlertsCollection sdk/dotnet/Alerts.cs` hits
- [x] 2.5 Run provider unit tests: `cd provider && go test -v -short -parallel 10 . ./pkg/...`

## 3. PR + automation verification

- [ ] 3.1 Open the PR (branch + Jira ref per repo conventions); confirm `pull-request.yml` is green across all SDK shards under the new toolchain setting
- [ ] 3.2 After merge, confirm the next scheduled `upgrade-provider` run (05:00 UTC) completes: `gh run list --workflow upgrade-provider.yml --limit 1` shows success, and any PR it opens auto-merges per the auto-merge-safety rule
- [ ] 3.3 Confirm the run's `go get pulumi-terraform-bridge` step no longer mentions `GOTOOLCHAIN=go1.25.9`

## 4. Mirror in pulumi-fivetran

- [ ] 4.1 Apply the sibling change `unblock-daily-upgrade-provider` in pulumi-fivetran (mise env line only) and open its PR
- [ ] 4.2 After merge, confirm fivetran's next scheduled run opens the bridge v3.134.0 upgrade PR and it auto-merges
