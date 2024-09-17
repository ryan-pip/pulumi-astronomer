# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetApiTokensResult',
    'AwaitableGetApiTokensResult',
    'get_api_tokens',
    'get_api_tokens_output',
]

@pulumi.output_type
class GetApiTokensResult:
    """
    A collection of values returned by getApiTokens.
    """
    def __init__(__self__, api_tokens=None, deployment_id=None, id=None, include_only_organization_tokens=None, workspace_id=None):
        if api_tokens and not isinstance(api_tokens, list):
            raise TypeError("Expected argument 'api_tokens' to be a list")
        pulumi.set(__self__, "api_tokens", api_tokens)
        if deployment_id and not isinstance(deployment_id, str):
            raise TypeError("Expected argument 'deployment_id' to be a str")
        pulumi.set(__self__, "deployment_id", deployment_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_only_organization_tokens and not isinstance(include_only_organization_tokens, bool):
            raise TypeError("Expected argument 'include_only_organization_tokens' to be a bool")
        pulumi.set(__self__, "include_only_organization_tokens", include_only_organization_tokens)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter(name="apiTokens")
    def api_tokens(self) -> Sequence['outputs.GetApiTokensApiTokenResult']:
        return pulumi.get(self, "api_tokens")

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[str]:
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeOnlyOrganizationTokens")
    def include_only_organization_tokens(self) -> Optional[bool]:
        return pulumi.get(self, "include_only_organization_tokens")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> Optional[str]:
        return pulumi.get(self, "workspace_id")


class AwaitableGetApiTokensResult(GetApiTokensResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiTokensResult(
            api_tokens=self.api_tokens,
            deployment_id=self.deployment_id,
            id=self.id,
            include_only_organization_tokens=self.include_only_organization_tokens,
            workspace_id=self.workspace_id)


def get_api_tokens(deployment_id: Optional[str] = None,
                   include_only_organization_tokens: Optional[bool] = None,
                   workspace_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiTokensResult:
    """
    API Tokens data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_api_tokens = astronomer.get_api_tokens()
    organization_only_example = astronomer.get_api_tokens(include_only_organization_tokens=True)
    workspace_example = astronomer.get_api_tokens(workspace_id="clx42sxw501gl01o0gjenthnh")
    deployment_example = astronomer.get_api_tokens(deployment_id="clx44jyu001m201m5dzsbexqr")
    pulumi.export("apiTokens", example_api_tokens)
    ```
    """
    __args__ = dict()
    __args__['deploymentId'] = deployment_id
    __args__['includeOnlyOrganizationTokens'] = include_only_organization_tokens
    __args__['workspaceId'] = workspace_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('astronomer:index/getApiTokens:getApiTokens', __args__, opts=opts, typ=GetApiTokensResult).value

    return AwaitableGetApiTokensResult(
        api_tokens=pulumi.get(__ret__, 'api_tokens'),
        deployment_id=pulumi.get(__ret__, 'deployment_id'),
        id=pulumi.get(__ret__, 'id'),
        include_only_organization_tokens=pulumi.get(__ret__, 'include_only_organization_tokens'),
        workspace_id=pulumi.get(__ret__, 'workspace_id'))


@_utilities.lift_output_func(get_api_tokens)
def get_api_tokens_output(deployment_id: Optional[pulumi.Input[Optional[str]]] = None,
                          include_only_organization_tokens: Optional[pulumi.Input[Optional[bool]]] = None,
                          workspace_id: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApiTokensResult]:
    """
    API Tokens data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_api_tokens = astronomer.get_api_tokens()
    organization_only_example = astronomer.get_api_tokens(include_only_organization_tokens=True)
    workspace_example = astronomer.get_api_tokens(workspace_id="clx42sxw501gl01o0gjenthnh")
    deployment_example = astronomer.get_api_tokens(deployment_id="clx44jyu001m201m5dzsbexqr")
    pulumi.export("apiTokens", example_api_tokens)
    ```
    """
    ...