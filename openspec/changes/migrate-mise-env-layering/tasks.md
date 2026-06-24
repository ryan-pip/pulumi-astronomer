# Tasks

## 1. Restructure mise config
- [x] 1.1 Rewrite `mise.toml` to the lean base set (today's `mise.ci.toml` contents, minus `github-cli` — `gh` ships on the GitHub-hosted runner image, so CI doesn't need it; see design.md).
- [x] 1.2 Create `mise.dev.toml` with the dev-only delta (github-cli, java, gopls, upgrade-provider, hk, pkl, yamllint, taplo, actionlint; `[settings] experimental = true`; `[env] HK_MISE = 1`).
- [x] 1.3 Create `.miserc.toml` with `env = ["dev", "local"]`.
- [x] 1.4 Delete `mise.ci.toml`.
- [x] 1.5 Run `mise lock`; commit the regenerated `mise.lock` (base) and new `mise.dev.lock` (dev).
- [x] 1.6 Pin `node` to an exact version (`24.17.0`, ≥24.1 per pulumi/ci-mgmt) so upgrade-provider's `mise upgrade --raw` finds nothing newer and never reinstalls node via its gpg-verify path (which fails on CI runners). `pulumi` intentionally left unpinned. Regenerate the lockfile.

## 2. Drop the composite wrapper
- [x] 2.1 Delete `.github/actions/setup-mise/`.

## 3. Rewrite workflows to install mise directly
- [x] 3.1 `pull-request.yml`: add `env: MISE_ENV: ci` at workflow level; replace each `uses: ./.github/actions/setup-mise` with a direct `jdx/mise-action` step (no `version:` input; `cache_key: "mise-{{platform}}-{{file_hash}}"`); keep `cache_save: true` on exactly the one prerequisites job, restore-only elsewhere.
- [x] 3.2 `release.yml`: same treatment (`MISE_ENV: ci`, direct `jdx/mise-action`, same cache key, `cache_save: true` on the one prereq/build job).
- [x] 3.3 `upgrade-provider.yml`: add `env: MISE_ENV: ci`; replace the direct full-set `jdx/mise-action` step (drop `version: 2026.3.7`); set `cache_save: true`; keep cache key `mise-{{platform}}-{{file_hash}}` (drop the `mise-full-` prefix so it shares the warm cache). Leave the `Install upgrade-provider` (`go install ...@main`) step intact.
- [x] 3.4 `upgrade-provider.yml`: de-hardcode the `pulumi-release-bot` identity in both places it appears, using the token step's `app-slug` output for this repo's own app:
  - `Set up git identity` step — set `user.name = "${{ steps.app-token.outputs.app-slug }}[bot]"` and `user.email = "<bot-user-id>+<name>@users.noreply.github.com"` (fetch the bot user id via `gh api /users/<name> --jq .id`, with `GH_TOKEN` from `steps.app-token.outputs.token`). Keep the step — git needs a committer identity to make the upgrade commit(s) and push the PR branch.
  - `Enable auto-merge on upgrade PR` step — change the author safety filter from `app/pulumi-release-bot` to `app/${{ steps.app-token.outputs.app-slug }}` so auto-merge isn't silently skipped on this repo's bot PRs.
- [x] 3.5 `upgrade-provider.yml`: port `pulumi-fivetran`'s `Sync go.work go directive` step (verbatim — identical module layout `provider/go.mod`, `provider/shim/go.mod`, `sdk/go.mod`), placed after `Run upgrade-provider` and before `Enable auto-merge`, gated on the same upgrade-branch prefixes. It bumps `go.work`'s `go` directive to the max of the module directives so a bridge bump doesn't leave the workspace behind.

## 4. Verify
- [x] 4.1 `MISE_ENV=ci mise install` locally resolves only the base set; a plain `mise install` resolves base+dev.
- [x] 4.2 `actionlint`/`yamllint` clean on all three workflows.
- [ ] 4.3 Trigger `upgrade-provider.yml` via `workflow_dispatch` — `Setup mise` passes, `upgrade-provider` runs to completion (or opens its PR), and the ported `Sync go.work go directive` step runs without error.
- [ ] 4.4 A pull-request run stays green (base set still covers golangci-lint, build, SDKs).
- [x] 4.5 Confirm the pattern matches `pulumi-fivetran`'s equivalent change.
