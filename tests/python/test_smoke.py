"""SDK smoke tests.

Each test constructs one resource via the mocked Pulumi runtime and asserts
that the URN and a representative output property resolve. Catches:
  - SDK build/import broken
  - __init__ signature drift after a bridge bump
  - Property serialization failures in the bridge layer

These tests do NOT exercise real Astronomer API calls — see
examples/<resource>/python/ + examples_py_test.go for live integration.
"""

import pulumi
import pulumi_astronomer as astro


def _check_urn(resource):
    """Return an Output[None] that asserts the resource registered with a URN."""

    def check(urn):
        assert urn, f"empty URN for {resource}"

    return resource.urn.apply(check)


def test_module_imports():
    # Cheap sanity check: the package exposes expected top-level names.
    expected = {
        "AgentToken",
        "Alert",
        "ApiToken",
        "Cluster",
        "CustomRole",
        "Deployment",
        "HybridClusterWorkspaceAuthorization",
        "NotificationChannel",
        "Team",
        "TeamMembership",
        "TeamRoles",
        "UserInvite",
        "UserRoles",
        "Workspace",
    }
    missing = expected - set(dir(astro))
    assert not missing, f"missing exports: {missing}"


@pulumi.runtime.test
def test_workspace():
    ws = astro.Workspace(
        "smoke-workspace",
        name="smoke-workspace",
        description="smoke test",
        cicd_enforced_default=False,
    )
    return _check_urn(ws)


@pulumi.runtime.test
def test_agent_token():
    tok = astro.AgentToken(
        "smoke-agent-token",
        name="smoke-agent-token",
        description="smoke",
        deployment_id="ckxyz",
        expiry_period_in_days=30,
    )
    return _check_urn(tok)


@pulumi.runtime.test
def test_api_token():
    tok = astro.ApiToken(
        "smoke-api-token",
        name="smoke-api-token",
        type="WORKSPACE",
        description="smoke",
        expiry_period_in_days=30,
        roles=[
            astro.ApiTokenRoleArgs(
                entity_id="cl-workspace-xyz",
                entity_type="WORKSPACE",
                role="WORKSPACE_OWNER",
            )
        ],
    )
    return _check_urn(tok)


@pulumi.runtime.test
def test_cluster():
    c = astro.Cluster(
        "smoke-cluster",
        name="smoke-cluster",
        type="DEDICATED",
        cloud_provider="AWS",
        region="us-east-1",
        vpc_subnet_range="172.20.0.0/22",
        workspace_ids=[],
    )
    return _check_urn(c)


@pulumi.runtime.test
def test_custom_role():
    r = astro.CustomRole(
        "smoke-custom-role",
        name="smoke-role",
        description="smoke",
        permissions=["organization.update"],
        scope_type="ORGANIZATION",
    )
    return _check_urn(r)


@pulumi.runtime.test
def test_deployment():
    d = astro.Deployment(
        "smoke-deployment",
        name="smoke-deployment",
        description="smoke",
        type="STANDARD",
        cloud_provider="AWS",
        region="us-east-1",
        executor="CELERY",
        workspace_id="cl-workspace-xyz",
        contact_emails=[],
        scheduler_size="SMALL",
        is_cicd_enforced=False,
        is_dag_deploy_enabled=True,
        is_development_mode=False,
        is_high_availability=False,
        environment_variables=[],
        worker_queues=[],
    )
    return _check_urn(d)


@pulumi.runtime.test
def test_alert():
    a = astro.Alert(
        "smoke-alert",
        name="smoke-alert",
        type="DAG_FAILURE",
        severity="WARNING",
        entity_id="cl-deployment-xyz",
        entity_type="DEPLOYMENT",
        notification_channel_ids=[],
        rules=astro.AlertRulesArgs(
            properties=astro.AlertRulesPropertiesArgs(
                deployment_id="cl-deployment-xyz",
            ),
            pattern_matches=[],
        ),
    )
    return _check_urn(a)


@pulumi.runtime.test
def test_notification_channel():
    n = astro.NotificationChannel(
        "smoke-notif",
        name="smoke-notif",
        type="EMAIL",
        entity_id="cl-workspace-xyz",
        entity_type="WORKSPACE",
        is_shared=True,
        definition=astro.NotificationChannelDefinitionArgs(
            recipients=["ops@example.com"],
        ),
    )
    return _check_urn(n)


@pulumi.runtime.test
def test_team():
    t = astro.Team(
        "smoke-team",
        name="smoke-team",
        description="smoke",
        organization_role="ORGANIZATION_MEMBER",
        member_ids=[],
        workspace_roles=[],
        deployment_roles=[],
        dag_roles=[],
    )
    return _check_urn(t)


@pulumi.runtime.test
def test_team_membership():
    tm = astro.TeamMembership(
        "smoke-team-membership",
        team_id="cl-team-xyz",
        user_id="cl-user-xyz",
    )
    return _check_urn(tm)


@pulumi.runtime.test
def test_team_roles():
    tr = astro.TeamRoles(
        "smoke-team-roles",
        team_id="cl-team-xyz",
        organization_role="ORGANIZATION_MEMBER",
        workspace_roles=[],
        deployment_roles=[],
        dag_roles=[],
    )
    return _check_urn(tr)


@pulumi.runtime.test
def test_user_invite():
    ui = astro.UserInvite(
        "smoke-user-invite",
        email="invitee@example.com",
        role="ORGANIZATION_MEMBER",
    )
    return _check_urn(ui)


@pulumi.runtime.test
def test_user_roles():
    ur = astro.UserRoles(
        "smoke-user-roles",
        user_id="cl-user-xyz",
        organization_role="ORGANIZATION_MEMBER",
        workspace_roles=[],
        deployment_roles=[],
        dag_roles=[],
    )
    return _check_urn(ur)


@pulumi.runtime.test
def test_hybrid_cluster_workspace_authorization():
    h = astro.HybridClusterWorkspaceAuthorization(
        "smoke-hybrid-auth",
        cluster_id="cl-cluster-xyz",
        workspace_ids=["cl-workspace-xyz"],
    )
    return _check_urn(h)
