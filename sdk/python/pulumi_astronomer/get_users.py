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
    'GetUsersResult',
    'AwaitableGetUsersResult',
    'get_users',
    'get_users_output',
]

@pulumi.output_type
class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, deployment_id=None, id=None, users=None, workspace_id=None):
        if deployment_id and not isinstance(deployment_id, str):
            raise TypeError("Expected argument 'deployment_id' to be a str")
        pulumi.set(__self__, "deployment_id", deployment_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

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
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetUsersUserResult']:
        return pulumi.get(self, "users")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> Optional[str]:
        return pulumi.get(self, "workspace_id")


class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            deployment_id=self.deployment_id,
            id=self.id,
            users=self.users,
            workspace_id=self.workspace_id)


def get_users(deployment_id: Optional[str] = None,
              workspace_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsersResult:
    """
    Users data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_users_users = astronomer.get_users()
    example_users_filter_by_workspace_id = astronomer.get_users(workspace_id="clx42sxw501gl01o0gjenthnh")
    example_users_filter_by_deployment_id = astronomer.get_users(deployment_id="clx44jyu001m201m5dzsbexqr")
    pulumi.export("exampleUsers", example_users_users)
    ```
    """
    __args__ = dict()
    __args__['deploymentId'] = deployment_id
    __args__['workspaceId'] = workspace_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('astronomer:index/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult).value

    return AwaitableGetUsersResult(
        deployment_id=pulumi.get(__ret__, 'deployment_id'),
        id=pulumi.get(__ret__, 'id'),
        users=pulumi.get(__ret__, 'users'),
        workspace_id=pulumi.get(__ret__, 'workspace_id'))


@_utilities.lift_output_func(get_users)
def get_users_output(deployment_id: Optional[pulumi.Input[Optional[str]]] = None,
                     workspace_id: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUsersResult]:
    """
    Users data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_users_users = astronomer.get_users()
    example_users_filter_by_workspace_id = astronomer.get_users(workspace_id="clx42sxw501gl01o0gjenthnh")
    example_users_filter_by_deployment_id = astronomer.get_users(deployment_id="clx44jyu001m201m5dzsbexqr")
    pulumi.export("exampleUsers", example_users_users)
    ```
    """
    ...