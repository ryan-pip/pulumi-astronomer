# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetWorkspaceResult',
    'AwaitableGetWorkspaceResult',
    'get_workspace',
    'get_workspace_output',
]

@pulumi.output_type
class GetWorkspaceResult:
    """
    A collection of values returned by getWorkspace.
    """
    def __init__(__self__, cicd_enforced_default=None, created_at=None, created_by=None, description=None, id=None, name=None, updated_at=None, updated_by=None):
        if cicd_enforced_default and not isinstance(cicd_enforced_default, bool):
            raise TypeError("Expected argument 'cicd_enforced_default' to be a bool")
        pulumi.set(__self__, "cicd_enforced_default", cicd_enforced_default)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if created_by and not isinstance(created_by, dict):
            raise TypeError("Expected argument 'created_by' to be a dict")
        pulumi.set(__self__, "created_by", created_by)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if updated_by and not isinstance(updated_by, dict):
            raise TypeError("Expected argument 'updated_by' to be a dict")
        pulumi.set(__self__, "updated_by", updated_by)

    @property
    @pulumi.getter(name="cicdEnforcedDefault")
    def cicd_enforced_default(self) -> bool:
        """
        Whether new Deployments enforce CI/CD deploys by default
        """
        return pulumi.get(self, "cicd_enforced_default")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Workspace creation timestamp
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> 'outputs.GetWorkspaceCreatedByResult':
        """
        Workspace creator
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Workspace description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Workspace identifier
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Workspace name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Workspace last updated timestamp
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> 'outputs.GetWorkspaceUpdatedByResult':
        """
        Workspace updater
        """
        return pulumi.get(self, "updated_by")


class AwaitableGetWorkspaceResult(GetWorkspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkspaceResult(
            cicd_enforced_default=self.cicd_enforced_default,
            created_at=self.created_at,
            created_by=self.created_by,
            description=self.description,
            id=self.id,
            name=self.name,
            updated_at=self.updated_at,
            updated_by=self.updated_by)


def get_workspace(id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkspaceResult:
    """
    Workspace data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_workspace = astronomer.get_workspace(id="clozc036j01to01jrlgvueo8t")
    pulumi.export("workspace", example_workspace)
    ```


    :param str id: Workspace identifier
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('astronomer:index/getWorkspace:getWorkspace', __args__, opts=opts, typ=GetWorkspaceResult).value

    return AwaitableGetWorkspaceResult(
        cicd_enforced_default=pulumi.get(__ret__, 'cicd_enforced_default'),
        created_at=pulumi.get(__ret__, 'created_at'),
        created_by=pulumi.get(__ret__, 'created_by'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        updated_by=pulumi.get(__ret__, 'updated_by'))
def get_workspace_output(id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWorkspaceResult]:
    """
    Workspace data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_workspace = astronomer.get_workspace(id="clozc036j01to01jrlgvueo8t")
    pulumi.export("workspace", example_workspace)
    ```


    :param str id: Workspace identifier
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('astronomer:index/getWorkspace:getWorkspace', __args__, opts=opts, typ=GetWorkspaceResult)
    return __ret__.apply(lambda __response__: GetWorkspaceResult(
        cicd_enforced_default=pulumi.get(__response__, 'cicd_enforced_default'),
        created_at=pulumi.get(__response__, 'created_at'),
        created_by=pulumi.get(__response__, 'created_by'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        updated_at=pulumi.get(__response__, 'updated_at'),
        updated_by=pulumi.get(__response__, 'updated_by')))
