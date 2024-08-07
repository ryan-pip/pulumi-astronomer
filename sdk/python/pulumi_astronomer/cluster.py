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
from ._inputs import *

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 cloud_provider: pulumi.Input[str],
                 region: pulumi.Input[str],
                 type: pulumi.Input[str],
                 vpc_subnet_range: pulumi.Input[str],
                 workspace_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: Optional[pulumi.Input[str]] = None,
                 pod_subnet_range: Optional[pulumi.Input[str]] = None,
                 service_peering_range: Optional[pulumi.Input[str]] = None,
                 service_subnet_range: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input['ClusterTimeoutsArgs']] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[str] cloud_provider: Cluster cloud provider - if changed, the cluster will be recreated.
        :param pulumi.Input[str] region: Cluster region - if changed, the cluster will be recreated.
        :param pulumi.Input[str] type: Cluster type
        :param pulumi.Input[str] vpc_subnet_range: Cluster VPC subnet range. If changed, the cluster will be recreated.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workspace_ids: Cluster workspace IDs
        :param pulumi.Input[str] name: Cluster name
        :param pulumi.Input[str] pod_subnet_range: Cluster pod subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] service_peering_range: Cluster service peering range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] service_subnet_range: Cluster service subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        pulumi.set(__self__, "cloud_provider", cloud_provider)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "vpc_subnet_range", vpc_subnet_range)
        pulumi.set(__self__, "workspace_ids", workspace_ids)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pod_subnet_range is not None:
            pulumi.set(__self__, "pod_subnet_range", pod_subnet_range)
        if service_peering_range is not None:
            pulumi.set(__self__, "service_peering_range", service_peering_range)
        if service_subnet_range is not None:
            pulumi.set(__self__, "service_subnet_range", service_subnet_range)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> pulumi.Input[str]:
        """
        Cluster cloud provider - if changed, the cluster will be recreated.
        """
        return pulumi.get(self, "cloud_provider")

    @cloud_provider.setter
    def cloud_provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloud_provider", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Cluster region - if changed, the cluster will be recreated.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Cluster type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcSubnetRange")
    def vpc_subnet_range(self) -> pulumi.Input[str]:
        """
        Cluster VPC subnet range. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "vpc_subnet_range")

    @vpc_subnet_range.setter
    def vpc_subnet_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_subnet_range", value)

    @property
    @pulumi.getter(name="workspaceIds")
    def workspace_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Cluster workspace IDs
        """
        return pulumi.get(self, "workspace_ids")

    @workspace_ids.setter
    def workspace_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "workspace_ids", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="podSubnetRange")
    def pod_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster pod subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "pod_subnet_range")

    @pod_subnet_range.setter
    def pod_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pod_subnet_range", value)

    @property
    @pulumi.getter(name="servicePeeringRange")
    def service_peering_range(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster service peering range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "service_peering_range")

    @service_peering_range.setter
    def service_peering_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_peering_range", value)

    @property
    @pulumi.getter(name="serviceSubnetRange")
    def service_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster service subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "service_subnet_range")

    @service_subnet_range.setter
    def service_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_subnet_range", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['ClusterTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['ClusterTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)


@pulumi.input_type
class _ClusterState:
    def __init__(__self__, *,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 db_instance_type: Optional[pulumi.Input[str]] = None,
                 is_limited: Optional[pulumi.Input[bool]] = None,
                 metadata: Optional[pulumi.Input['ClusterMetadataArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_pools: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]]] = None,
                 pod_subnet_range: Optional[pulumi.Input[str]] = None,
                 provider_account: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_peering_range: Optional[pulumi.Input[str]] = None,
                 service_subnet_range: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input['ClusterTimeoutsArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 vpc_subnet_range: Optional[pulumi.Input[str]] = None,
                 workspace_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[str] cloud_provider: Cluster cloud provider - if changed, the cluster will be recreated.
        :param pulumi.Input[str] created_at: Cluster creation timestamp
        :param pulumi.Input[str] db_instance_type: Cluster database instance type
        :param pulumi.Input[bool] is_limited: Whether the cluster is limited
        :param pulumi.Input['ClusterMetadataArgs'] metadata: Cluster metadata
        :param pulumi.Input[str] name: Cluster name
        :param pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]] node_pools: Cluster node pools
        :param pulumi.Input[str] pod_subnet_range: Cluster pod subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] provider_account: Cluster provider account
        :param pulumi.Input[str] region: Cluster region - if changed, the cluster will be recreated.
        :param pulumi.Input[str] service_peering_range: Cluster service peering range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] service_subnet_range: Cluster service subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] status: Cluster status
        :param pulumi.Input[str] tenant_id: Cluster tenant ID
        :param pulumi.Input[str] type: Cluster type
        :param pulumi.Input[str] updated_at: Cluster last updated timestamp
        :param pulumi.Input[str] vpc_subnet_range: Cluster VPC subnet range. If changed, the cluster will be recreated.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workspace_ids: Cluster workspace IDs
        """
        if cloud_provider is not None:
            pulumi.set(__self__, "cloud_provider", cloud_provider)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if db_instance_type is not None:
            pulumi.set(__self__, "db_instance_type", db_instance_type)
        if is_limited is not None:
            pulumi.set(__self__, "is_limited", is_limited)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_pools is not None:
            pulumi.set(__self__, "node_pools", node_pools)
        if pod_subnet_range is not None:
            pulumi.set(__self__, "pod_subnet_range", pod_subnet_range)
        if provider_account is not None:
            pulumi.set(__self__, "provider_account", provider_account)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if service_peering_range is not None:
            pulumi.set(__self__, "service_peering_range", service_peering_range)
        if service_subnet_range is not None:
            pulumi.set(__self__, "service_subnet_range", service_subnet_range)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if vpc_subnet_range is not None:
            pulumi.set(__self__, "vpc_subnet_range", vpc_subnet_range)
        if workspace_ids is not None:
            pulumi.set(__self__, "workspace_ids", workspace_ids)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster cloud provider - if changed, the cluster will be recreated.
        """
        return pulumi.get(self, "cloud_provider")

    @cloud_provider.setter
    def cloud_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_provider", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster creation timestamp
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dbInstanceType")
    def db_instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster database instance type
        """
        return pulumi.get(self, "db_instance_type")

    @db_instance_type.setter
    def db_instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_type", value)

    @property
    @pulumi.getter(name="isLimited")
    def is_limited(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the cluster is limited
        """
        return pulumi.get(self, "is_limited")

    @is_limited.setter
    def is_limited(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_limited", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['ClusterMetadataArgs']]:
        """
        Cluster metadata
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['ClusterMetadataArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]]]:
        """
        Cluster node pools
        """
        return pulumi.get(self, "node_pools")

    @node_pools.setter
    def node_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]]]):
        pulumi.set(self, "node_pools", value)

    @property
    @pulumi.getter(name="podSubnetRange")
    def pod_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster pod subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "pod_subnet_range")

    @pod_subnet_range.setter
    def pod_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pod_subnet_range", value)

    @property
    @pulumi.getter(name="providerAccount")
    def provider_account(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster provider account
        """
        return pulumi.get(self, "provider_account")

    @provider_account.setter
    def provider_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_account", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster region - if changed, the cluster will be recreated.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="servicePeeringRange")
    def service_peering_range(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster service peering range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "service_peering_range")

    @service_peering_range.setter
    def service_peering_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_peering_range", value)

    @property
    @pulumi.getter(name="serviceSubnetRange")
    def service_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster service subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "service_subnet_range")

    @service_subnet_range.setter
    def service_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_subnet_range", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster tenant ID
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['ClusterTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['ClusterTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster last updated timestamp
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="vpcSubnetRange")
    def vpc_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster VPC subnet range. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "vpc_subnet_range")

    @vpc_subnet_range.setter
    def vpc_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_subnet_range", value)

    @property
    @pulumi.getter(name="workspaceIds")
    def workspace_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Cluster workspace IDs
        """
        return pulumi.get(self, "workspace_ids")

    @workspace_ids.setter
    def workspace_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "workspace_ids", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pod_subnet_range: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_peering_range: Optional[pulumi.Input[str]] = None,
                 service_subnet_range: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input[pulumi.InputType['ClusterTimeoutsArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_subnet_range: Optional[pulumi.Input[str]] = None,
                 workspace_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Cluster resource. If creating multiple clusters, add a delay between each cluster creation to avoid cluster creation limiting errors.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_astronomer as astronomer

        aws_example = astronomer.Cluster("awsExample",
            type="DEDICATED",
            region="us-east-1",
            cloud_provider="AWS",
            vpc_subnet_range="172.20.0.0/20",
            workspace_ids=[],
            timeouts=astronomer.ClusterTimeoutsArgs(
                create="3h",
                update="2h",
                delete="1h",
            ))
        azure_example = astronomer.Cluster("azureExample",
            type="DEDICATED",
            region="westus2",
            cloud_provider="AZURE",
            vpc_subnet_range="172.20.0.0/19",
            workspace_ids=["clv4wcf6f003u01m3zp7gsvzg"])
        gcp_example = astronomer.Cluster("gcpExample",
            type="DEDICATED",
            region="us-central1",
            cloud_provider="GCP",
            pod_subnet_range="172.21.0.0/19",
            service_peering_range="172.23.0.0/20",
            service_subnet_range="172.22.0.0/22",
            vpc_subnet_range="172.20.0.0/22",
            workspace_ids=[])
        imported_cluster = astronomer.Cluster("importedCluster",
            type="DEDICATED",
            region="us-central1",
            cloud_provider="GCP",
            pod_subnet_range="172.21.0.0/19",
            service_peering_range="172.23.0.0/20",
            service_subnet_range="172.22.0.0/22",
            vpc_subnet_range="172.20.0.0/22",
            workspace_ids=[])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_provider: Cluster cloud provider - if changed, the cluster will be recreated.
        :param pulumi.Input[str] name: Cluster name
        :param pulumi.Input[str] pod_subnet_range: Cluster pod subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] region: Cluster region - if changed, the cluster will be recreated.
        :param pulumi.Input[str] service_peering_range: Cluster service peering range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] service_subnet_range: Cluster service subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] type: Cluster type
        :param pulumi.Input[str] vpc_subnet_range: Cluster VPC subnet range. If changed, the cluster will be recreated.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workspace_ids: Cluster workspace IDs
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Cluster resource. If creating multiple clusters, add a delay between each cluster creation to avoid cluster creation limiting errors.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_astronomer as astronomer

        aws_example = astronomer.Cluster("awsExample",
            type="DEDICATED",
            region="us-east-1",
            cloud_provider="AWS",
            vpc_subnet_range="172.20.0.0/20",
            workspace_ids=[],
            timeouts=astronomer.ClusterTimeoutsArgs(
                create="3h",
                update="2h",
                delete="1h",
            ))
        azure_example = astronomer.Cluster("azureExample",
            type="DEDICATED",
            region="westus2",
            cloud_provider="AZURE",
            vpc_subnet_range="172.20.0.0/19",
            workspace_ids=["clv4wcf6f003u01m3zp7gsvzg"])
        gcp_example = astronomer.Cluster("gcpExample",
            type="DEDICATED",
            region="us-central1",
            cloud_provider="GCP",
            pod_subnet_range="172.21.0.0/19",
            service_peering_range="172.23.0.0/20",
            service_subnet_range="172.22.0.0/22",
            vpc_subnet_range="172.20.0.0/22",
            workspace_ids=[])
        imported_cluster = astronomer.Cluster("importedCluster",
            type="DEDICATED",
            region="us-central1",
            cloud_provider="GCP",
            pod_subnet_range="172.21.0.0/19",
            service_peering_range="172.23.0.0/20",
            service_subnet_range="172.22.0.0/22",
            vpc_subnet_range="172.20.0.0/22",
            workspace_ids=[])
        ```

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pod_subnet_range: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_peering_range: Optional[pulumi.Input[str]] = None,
                 service_subnet_range: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input[pulumi.InputType['ClusterTimeoutsArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_subnet_range: Optional[pulumi.Input[str]] = None,
                 workspace_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            if cloud_provider is None and not opts.urn:
                raise TypeError("Missing required property 'cloud_provider'")
            __props__.__dict__["cloud_provider"] = cloud_provider
            __props__.__dict__["name"] = name
            __props__.__dict__["pod_subnet_range"] = pod_subnet_range
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["service_peering_range"] = service_peering_range
            __props__.__dict__["service_subnet_range"] = service_subnet_range
            __props__.__dict__["timeouts"] = timeouts
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if vpc_subnet_range is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_subnet_range'")
            __props__.__dict__["vpc_subnet_range"] = vpc_subnet_range
            if workspace_ids is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_ids'")
            __props__.__dict__["workspace_ids"] = workspace_ids
            __props__.__dict__["created_at"] = None
            __props__.__dict__["db_instance_type"] = None
            __props__.__dict__["is_limited"] = None
            __props__.__dict__["metadata"] = None
            __props__.__dict__["node_pools"] = None
            __props__.__dict__["provider_account"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["tenant_id"] = None
            __props__.__dict__["updated_at"] = None
        super(Cluster, __self__).__init__(
            'astronomer:index/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cloud_provider: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            db_instance_type: Optional[pulumi.Input[str]] = None,
            is_limited: Optional[pulumi.Input[bool]] = None,
            metadata: Optional[pulumi.Input[pulumi.InputType['ClusterMetadataArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodePoolArgs']]]]] = None,
            pod_subnet_range: Optional[pulumi.Input[str]] = None,
            provider_account: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            service_peering_range: Optional[pulumi.Input[str]] = None,
            service_subnet_range: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            timeouts: Optional[pulumi.Input[pulumi.InputType['ClusterTimeoutsArgs']]] = None,
            type: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            vpc_subnet_range: Optional[pulumi.Input[str]] = None,
            workspace_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_provider: Cluster cloud provider - if changed, the cluster will be recreated.
        :param pulumi.Input[str] created_at: Cluster creation timestamp
        :param pulumi.Input[str] db_instance_type: Cluster database instance type
        :param pulumi.Input[bool] is_limited: Whether the cluster is limited
        :param pulumi.Input[pulumi.InputType['ClusterMetadataArgs']] metadata: Cluster metadata
        :param pulumi.Input[str] name: Cluster name
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodePoolArgs']]]] node_pools: Cluster node pools
        :param pulumi.Input[str] pod_subnet_range: Cluster pod subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] provider_account: Cluster provider account
        :param pulumi.Input[str] region: Cluster region - if changed, the cluster will be recreated.
        :param pulumi.Input[str] service_peering_range: Cluster service peering range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] service_subnet_range: Cluster service subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        :param pulumi.Input[str] status: Cluster status
        :param pulumi.Input[str] tenant_id: Cluster tenant ID
        :param pulumi.Input[str] type: Cluster type
        :param pulumi.Input[str] updated_at: Cluster last updated timestamp
        :param pulumi.Input[str] vpc_subnet_range: Cluster VPC subnet range. If changed, the cluster will be recreated.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workspace_ids: Cluster workspace IDs
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["cloud_provider"] = cloud_provider
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["db_instance_type"] = db_instance_type
        __props__.__dict__["is_limited"] = is_limited
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["name"] = name
        __props__.__dict__["node_pools"] = node_pools
        __props__.__dict__["pod_subnet_range"] = pod_subnet_range
        __props__.__dict__["provider_account"] = provider_account
        __props__.__dict__["region"] = region
        __props__.__dict__["service_peering_range"] = service_peering_range
        __props__.__dict__["service_subnet_range"] = service_subnet_range
        __props__.__dict__["status"] = status
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["timeouts"] = timeouts
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["vpc_subnet_range"] = vpc_subnet_range
        __props__.__dict__["workspace_ids"] = workspace_ids
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> pulumi.Output[str]:
        """
        Cluster cloud provider - if changed, the cluster will be recreated.
        """
        return pulumi.get(self, "cloud_provider")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Cluster creation timestamp
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dbInstanceType")
    def db_instance_type(self) -> pulumi.Output[str]:
        """
        Cluster database instance type
        """
        return pulumi.get(self, "db_instance_type")

    @property
    @pulumi.getter(name="isLimited")
    def is_limited(self) -> pulumi.Output[bool]:
        """
        Whether the cluster is limited
        """
        return pulumi.get(self, "is_limited")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output['outputs.ClusterMetadata']:
        """
        Cluster metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Cluster name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> pulumi.Output[Sequence['outputs.ClusterNodePool']]:
        """
        Cluster node pools
        """
        return pulumi.get(self, "node_pools")

    @property
    @pulumi.getter(name="podSubnetRange")
    def pod_subnet_range(self) -> pulumi.Output[Optional[str]]:
        """
        Cluster pod subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "pod_subnet_range")

    @property
    @pulumi.getter(name="providerAccount")
    def provider_account(self) -> pulumi.Output[str]:
        """
        Cluster provider account
        """
        return pulumi.get(self, "provider_account")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Cluster region - if changed, the cluster will be recreated.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="servicePeeringRange")
    def service_peering_range(self) -> pulumi.Output[Optional[str]]:
        """
        Cluster service peering range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "service_peering_range")

    @property
    @pulumi.getter(name="serviceSubnetRange")
    def service_subnet_range(self) -> pulumi.Output[Optional[str]]:
        """
        Cluster service subnet range - required for 'GCP' clusters. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "service_subnet_range")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Cluster status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        Cluster tenant ID
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def timeouts(self) -> pulumi.Output[Optional['outputs.ClusterTimeouts']]:
        return pulumi.get(self, "timeouts")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Cluster type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Cluster last updated timestamp
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="vpcSubnetRange")
    def vpc_subnet_range(self) -> pulumi.Output[str]:
        """
        Cluster VPC subnet range. If changed, the cluster will be recreated.
        """
        return pulumi.get(self, "vpc_subnet_range")

    @property
    @pulumi.getter(name="workspaceIds")
    def workspace_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Cluster workspace IDs
        """
        return pulumi.get(self, "workspace_ids")

