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
from ._inputs import *

__all__ = ['UserRolesArgs', 'UserRoles']

@pulumi.input_type
class UserRolesArgs:
    def __init__(__self__, *,
                 organization_role: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 deployment_roles: Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesDeploymentRoleArgs']]]] = None,
                 workspace_roles: Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesWorkspaceRoleArgs']]]] = None):
        """
        The set of arguments for constructing a UserRoles resource.
        :param pulumi.Input[str] organization_role: The role to assign to the organization
        :param pulumi.Input[str] user_id: The ID of the user to assign the roles to
        :param pulumi.Input[Sequence[pulumi.Input['UserRolesDeploymentRoleArgs']]] deployment_roles: The roles to assign to the deployments
        :param pulumi.Input[Sequence[pulumi.Input['UserRolesWorkspaceRoleArgs']]] workspace_roles: The roles to assign to the workspaces
        """
        pulumi.set(__self__, "organization_role", organization_role)
        pulumi.set(__self__, "user_id", user_id)
        if deployment_roles is not None:
            pulumi.set(__self__, "deployment_roles", deployment_roles)
        if workspace_roles is not None:
            pulumi.set(__self__, "workspace_roles", workspace_roles)

    @property
    @pulumi.getter(name="organizationRole")
    def organization_role(self) -> pulumi.Input[str]:
        """
        The role to assign to the organization
        """
        return pulumi.get(self, "organization_role")

    @organization_role.setter
    def organization_role(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_role", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        The ID of the user to assign the roles to
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="deploymentRoles")
    def deployment_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesDeploymentRoleArgs']]]]:
        """
        The roles to assign to the deployments
        """
        return pulumi.get(self, "deployment_roles")

    @deployment_roles.setter
    def deployment_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesDeploymentRoleArgs']]]]):
        pulumi.set(self, "deployment_roles", value)

    @property
    @pulumi.getter(name="workspaceRoles")
    def workspace_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesWorkspaceRoleArgs']]]]:
        """
        The roles to assign to the workspaces
        """
        return pulumi.get(self, "workspace_roles")

    @workspace_roles.setter
    def workspace_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesWorkspaceRoleArgs']]]]):
        pulumi.set(self, "workspace_roles", value)


@pulumi.input_type
class _UserRolesState:
    def __init__(__self__, *,
                 deployment_roles: Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesDeploymentRoleArgs']]]] = None,
                 organization_role: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 workspace_roles: Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesWorkspaceRoleArgs']]]] = None):
        """
        Input properties used for looking up and filtering UserRoles resources.
        :param pulumi.Input[Sequence[pulumi.Input['UserRolesDeploymentRoleArgs']]] deployment_roles: The roles to assign to the deployments
        :param pulumi.Input[str] organization_role: The role to assign to the organization
        :param pulumi.Input[str] user_id: The ID of the user to assign the roles to
        :param pulumi.Input[Sequence[pulumi.Input['UserRolesWorkspaceRoleArgs']]] workspace_roles: The roles to assign to the workspaces
        """
        if deployment_roles is not None:
            pulumi.set(__self__, "deployment_roles", deployment_roles)
        if organization_role is not None:
            pulumi.set(__self__, "organization_role", organization_role)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)
        if workspace_roles is not None:
            pulumi.set(__self__, "workspace_roles", workspace_roles)

    @property
    @pulumi.getter(name="deploymentRoles")
    def deployment_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesDeploymentRoleArgs']]]]:
        """
        The roles to assign to the deployments
        """
        return pulumi.get(self, "deployment_roles")

    @deployment_roles.setter
    def deployment_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesDeploymentRoleArgs']]]]):
        pulumi.set(self, "deployment_roles", value)

    @property
    @pulumi.getter(name="organizationRole")
    def organization_role(self) -> Optional[pulumi.Input[str]]:
        """
        The role to assign to the organization
        """
        return pulumi.get(self, "organization_role")

    @organization_role.setter
    def organization_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_role", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the user to assign the roles to
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="workspaceRoles")
    def workspace_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesWorkspaceRoleArgs']]]]:
        """
        The roles to assign to the workspaces
        """
        return pulumi.get(self, "workspace_roles")

    @workspace_roles.setter
    def workspace_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserRolesWorkspaceRoleArgs']]]]):
        pulumi.set(self, "workspace_roles", value)


class UserRoles(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_roles: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserRolesDeploymentRoleArgs', 'UserRolesDeploymentRoleArgsDict']]]]] = None,
                 organization_role: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 workspace_roles: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserRolesWorkspaceRoleArgs', 'UserRolesWorkspaceRoleArgsDict']]]]] = None,
                 __props__=None):
        """
        User Roles resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['UserRolesDeploymentRoleArgs', 'UserRolesDeploymentRoleArgsDict']]]] deployment_roles: The roles to assign to the deployments
        :param pulumi.Input[str] organization_role: The role to assign to the organization
        :param pulumi.Input[str] user_id: The ID of the user to assign the roles to
        :param pulumi.Input[Sequence[pulumi.Input[Union['UserRolesWorkspaceRoleArgs', 'UserRolesWorkspaceRoleArgsDict']]]] workspace_roles: The roles to assign to the workspaces
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserRolesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        User Roles resource

        :param str resource_name: The name of the resource.
        :param UserRolesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserRolesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_roles: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserRolesDeploymentRoleArgs', 'UserRolesDeploymentRoleArgsDict']]]]] = None,
                 organization_role: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 workspace_roles: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserRolesWorkspaceRoleArgs', 'UserRolesWorkspaceRoleArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserRolesArgs.__new__(UserRolesArgs)

            __props__.__dict__["deployment_roles"] = deployment_roles
            if organization_role is None and not opts.urn:
                raise TypeError("Missing required property 'organization_role'")
            __props__.__dict__["organization_role"] = organization_role
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["workspace_roles"] = workspace_roles
        super(UserRoles, __self__).__init__(
            'astronomer:index/userRoles:UserRoles',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deployment_roles: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserRolesDeploymentRoleArgs', 'UserRolesDeploymentRoleArgsDict']]]]] = None,
            organization_role: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None,
            workspace_roles: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserRolesWorkspaceRoleArgs', 'UserRolesWorkspaceRoleArgsDict']]]]] = None) -> 'UserRoles':
        """
        Get an existing UserRoles resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['UserRolesDeploymentRoleArgs', 'UserRolesDeploymentRoleArgsDict']]]] deployment_roles: The roles to assign to the deployments
        :param pulumi.Input[str] organization_role: The role to assign to the organization
        :param pulumi.Input[str] user_id: The ID of the user to assign the roles to
        :param pulumi.Input[Sequence[pulumi.Input[Union['UserRolesWorkspaceRoleArgs', 'UserRolesWorkspaceRoleArgsDict']]]] workspace_roles: The roles to assign to the workspaces
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserRolesState.__new__(_UserRolesState)

        __props__.__dict__["deployment_roles"] = deployment_roles
        __props__.__dict__["organization_role"] = organization_role
        __props__.__dict__["user_id"] = user_id
        __props__.__dict__["workspace_roles"] = workspace_roles
        return UserRoles(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deploymentRoles")
    def deployment_roles(self) -> pulumi.Output[Optional[Sequence['outputs.UserRolesDeploymentRole']]]:
        """
        The roles to assign to the deployments
        """
        return pulumi.get(self, "deployment_roles")

    @property
    @pulumi.getter(name="organizationRole")
    def organization_role(self) -> pulumi.Output[str]:
        """
        The role to assign to the organization
        """
        return pulumi.get(self, "organization_role")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        The ID of the user to assign the roles to
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="workspaceRoles")
    def workspace_roles(self) -> pulumi.Output[Optional[Sequence['outputs.UserRolesWorkspaceRole']]]:
        """
        The roles to assign to the workspaces
        """
        return pulumi.get(self, "workspace_roles")

