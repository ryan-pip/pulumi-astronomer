import pulumi
import pulumi_astronomer as astro

cfg = pulumi.Config()
workspace_id = cfg.require("workspaceId")

ws = astro.get_workspace(id=workspace_id)

pulumi.export("workspace_id", ws.id)
pulumi.export("workspace_name", ws.name)
pulumi.export("cicd_enforced_default", ws.cicd_enforced_default)
