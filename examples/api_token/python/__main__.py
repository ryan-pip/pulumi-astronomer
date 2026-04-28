import pulumi
import pulumi_astronomer as astro

cfg = pulumi.Config()
workspace_id = cfg.require("workspaceId")

token = astro.ApiToken(
    "pulumi-test-api-token",
    name="pulumi-test-api-token",
    type="WORKSPACE",
    description="Created by pulumi-astronomer integration tests",
    expiry_period_in_days=1,
    roles=[
        astro.ApiTokenRoleArgs(
            entity_id=workspace_id,
            entity_type="WORKSPACE",
            role="WORKSPACE_OWNER",
        )
    ],
)

pulumi.export("token_id", token.id)
pulumi.export("token_name", token.name)
