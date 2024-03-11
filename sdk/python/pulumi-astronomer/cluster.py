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
                 db_instance_type: Optional[pulumi.Input[str]] = None,
                 k8s_tags: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterK8sTagArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_pools: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]]] = None,
                 pod_subnet_range: Optional[pulumi.Input[str]] = None,
                 provider_account: Optional[pulumi.Input[str]] = None,
                 service_peering_range: Optional[pulumi.Input[str]] = None,
                 service_subnet_range: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[str] cloud_provider: The cluster's cloud provider.
        :param pulumi.Input[str] region: The cluster's region.
        :param pulumi.Input[str] type: The cluster's type.
        :param pulumi.Input[str] vpc_subnet_range: The VPC subnet range.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workspace_ids: The list of Workspaces that are authorized to the cluster.
        :param pulumi.Input[str] db_instance_type: The type of database instance that is used for the cluster. Required for Hybrid clusters.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterK8sTagArgs']]] k8s_tags: The Kubernetes tags in the cluster.
        :param pulumi.Input[str] name: The name of the node pool.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]] node_pools: The list of node pools to create in the cluster.
        :param pulumi.Input[str] pod_subnet_range: The subnet range for Pods. For GCP clusters only.
        :param pulumi.Input[str] provider_account: The provider account ID. Required for Hybrid clusters.
        :param pulumi.Input[str] service_peering_range: The service peering range. For GCP clusters only.
        :param pulumi.Input[str] service_subnet_range: The service subnet range. For GCP clusters only.
        :param pulumi.Input[str] tenant_id: The tenant ID. For Azure clusters only.
        """
        pulumi.set(__self__, "cloud_provider", cloud_provider)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "vpc_subnet_range", vpc_subnet_range)
        pulumi.set(__self__, "workspace_ids", workspace_ids)
        if db_instance_type is not None:
            pulumi.set(__self__, "db_instance_type", db_instance_type)
        if k8s_tags is not None:
            pulumi.set(__self__, "k8s_tags", k8s_tags)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_pools is not None:
            pulumi.set(__self__, "node_pools", node_pools)
        if pod_subnet_range is not None:
            pulumi.set(__self__, "pod_subnet_range", pod_subnet_range)
        if provider_account is not None:
            pulumi.set(__self__, "provider_account", provider_account)
        if service_peering_range is not None:
            pulumi.set(__self__, "service_peering_range", service_peering_range)
        if service_subnet_range is not None:
            pulumi.set(__self__, "service_subnet_range", service_subnet_range)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> pulumi.Input[str]:
        """
        The cluster's cloud provider.
        """
        return pulumi.get(self, "cloud_provider")

    @cloud_provider.setter
    def cloud_provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloud_provider", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The cluster's region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The cluster's type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcSubnetRange")
    def vpc_subnet_range(self) -> pulumi.Input[str]:
        """
        The VPC subnet range.
        """
        return pulumi.get(self, "vpc_subnet_range")

    @vpc_subnet_range.setter
    def vpc_subnet_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_subnet_range", value)

    @property
    @pulumi.getter(name="workspaceIds")
    def workspace_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The list of Workspaces that are authorized to the cluster.
        """
        return pulumi.get(self, "workspace_ids")

    @workspace_ids.setter
    def workspace_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "workspace_ids", value)

    @property
    @pulumi.getter(name="dbInstanceType")
    def db_instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of database instance that is used for the cluster. Required for Hybrid clusters.
        """
        return pulumi.get(self, "db_instance_type")

    @db_instance_type.setter
    def db_instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_type", value)

    @property
    @pulumi.getter(name="k8sTags")
    def k8s_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterK8sTagArgs']]]]:
        """
        The Kubernetes tags in the cluster.
        """
        return pulumi.get(self, "k8s_tags")

    @k8s_tags.setter
    def k8s_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterK8sTagArgs']]]]):
        pulumi.set(self, "k8s_tags", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the node pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]]]:
        """
        The list of node pools to create in the cluster.
        """
        return pulumi.get(self, "node_pools")

    @node_pools.setter
    def node_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]]]):
        pulumi.set(self, "node_pools", value)

    @property
    @pulumi.getter(name="podSubnetRange")
    def pod_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet range for Pods. For GCP clusters only.
        """
        return pulumi.get(self, "pod_subnet_range")

    @pod_subnet_range.setter
    def pod_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pod_subnet_range", value)

    @property
    @pulumi.getter(name="providerAccount")
    def provider_account(self) -> Optional[pulumi.Input[str]]:
        """
        The provider account ID. Required for Hybrid clusters.
        """
        return pulumi.get(self, "provider_account")

    @provider_account.setter
    def provider_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_account", value)

    @property
    @pulumi.getter(name="servicePeeringRange")
    def service_peering_range(self) -> Optional[pulumi.Input[str]]:
        """
        The service peering range. For GCP clusters only.
        """
        return pulumi.get(self, "service_peering_range")

    @service_peering_range.setter
    def service_peering_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_peering_range", value)

    @property
    @pulumi.getter(name="serviceSubnetRange")
    def service_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        The service subnet range. For GCP clusters only.
        """
        return pulumi.get(self, "service_subnet_range")

    @service_subnet_range.setter
    def service_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_subnet_range", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The tenant ID. For Azure clusters only.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class _ClusterState:
    def __init__(__self__, *,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 db_instance_type: Optional[pulumi.Input[str]] = None,
                 is_limited: Optional[pulumi.Input[bool]] = None,
                 k8s_tags: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterK8sTagArgs']]]] = None,
                 metadata: Optional[pulumi.Input['ClusterMetadataArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_pools: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 pod_subnet_range: Optional[pulumi.Input[str]] = None,
                 provider_account: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_peering_range: Optional[pulumi.Input[str]] = None,
                 service_subnet_range: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_subnet_range: Optional[pulumi.Input[str]] = None,
                 workspace_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[str] cloud_provider: The cluster's cloud provider.
        :param pulumi.Input[str] db_instance_type: The type of database instance that is used for the cluster. Required for Hybrid clusters.
        :param pulumi.Input[bool] is_limited: Whether the cluster is limited.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterK8sTagArgs']]] k8s_tags: The Kubernetes tags in the cluster.
        :param pulumi.Input['ClusterMetadataArgs'] metadata: The cluster's metadata.
        :param pulumi.Input[str] name: The name of the node pool.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]] node_pools: The list of node pools to create in the cluster.
        :param pulumi.Input[str] organization_id: The organization this cluster is associated with.
        :param pulumi.Input[str] pod_subnet_range: The subnet range for Pods. For GCP clusters only.
        :param pulumi.Input[str] provider_account: The provider account ID. Required for Hybrid clusters.
        :param pulumi.Input[str] region: The cluster's region.
        :param pulumi.Input[str] service_peering_range: The service peering range. For GCP clusters only.
        :param pulumi.Input[str] service_subnet_range: The service subnet range. For GCP clusters only.
        :param pulumi.Input[str] tenant_id: The tenant ID. For Azure clusters only.
        :param pulumi.Input[str] type: The cluster's type.
        :param pulumi.Input[str] vpc_subnet_range: The VPC subnet range.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workspace_ids: The list of Workspaces that are authorized to the cluster.
        """
        if cloud_provider is not None:
            pulumi.set(__self__, "cloud_provider", cloud_provider)
        if db_instance_type is not None:
            pulumi.set(__self__, "db_instance_type", db_instance_type)
        if is_limited is not None:
            pulumi.set(__self__, "is_limited", is_limited)
        if k8s_tags is not None:
            pulumi.set(__self__, "k8s_tags", k8s_tags)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_pools is not None:
            pulumi.set(__self__, "node_pools", node_pools)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
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
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vpc_subnet_range is not None:
            pulumi.set(__self__, "vpc_subnet_range", vpc_subnet_range)
        if workspace_ids is not None:
            pulumi.set(__self__, "workspace_ids", workspace_ids)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster's cloud provider.
        """
        return pulumi.get(self, "cloud_provider")

    @cloud_provider.setter
    def cloud_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_provider", value)

    @property
    @pulumi.getter(name="dbInstanceType")
    def db_instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of database instance that is used for the cluster. Required for Hybrid clusters.
        """
        return pulumi.get(self, "db_instance_type")

    @db_instance_type.setter
    def db_instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_type", value)

    @property
    @pulumi.getter(name="isLimited")
    def is_limited(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the cluster is limited.
        """
        return pulumi.get(self, "is_limited")

    @is_limited.setter
    def is_limited(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_limited", value)

    @property
    @pulumi.getter(name="k8sTags")
    def k8s_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterK8sTagArgs']]]]:
        """
        The Kubernetes tags in the cluster.
        """
        return pulumi.get(self, "k8s_tags")

    @k8s_tags.setter
    def k8s_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterK8sTagArgs']]]]):
        pulumi.set(self, "k8s_tags", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['ClusterMetadataArgs']]:
        """
        The cluster's metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['ClusterMetadataArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the node pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]]]:
        """
        The list of node pools to create in the cluster.
        """
        return pulumi.get(self, "node_pools")

    @node_pools.setter
    def node_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodePoolArgs']]]]):
        pulumi.set(self, "node_pools", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization this cluster is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="podSubnetRange")
    def pod_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet range for Pods. For GCP clusters only.
        """
        return pulumi.get(self, "pod_subnet_range")

    @pod_subnet_range.setter
    def pod_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pod_subnet_range", value)

    @property
    @pulumi.getter(name="providerAccount")
    def provider_account(self) -> Optional[pulumi.Input[str]]:
        """
        The provider account ID. Required for Hybrid clusters.
        """
        return pulumi.get(self, "provider_account")

    @provider_account.setter
    def provider_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_account", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster's region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="servicePeeringRange")
    def service_peering_range(self) -> Optional[pulumi.Input[str]]:
        """
        The service peering range. For GCP clusters only.
        """
        return pulumi.get(self, "service_peering_range")

    @service_peering_range.setter
    def service_peering_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_peering_range", value)

    @property
    @pulumi.getter(name="serviceSubnetRange")
    def service_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        The service subnet range. For GCP clusters only.
        """
        return pulumi.get(self, "service_subnet_range")

    @service_subnet_range.setter
    def service_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_subnet_range", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The tenant ID. For Azure clusters only.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster's type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcSubnetRange")
    def vpc_subnet_range(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC subnet range.
        """
        return pulumi.get(self, "vpc_subnet_range")

    @vpc_subnet_range.setter
    def vpc_subnet_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_subnet_range", value)

    @property
    @pulumi.getter(name="workspaceIds")
    def workspace_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of Workspaces that are authorized to the cluster.
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
                 db_instance_type: Optional[pulumi.Input[str]] = None,
                 k8s_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterK8sTagArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodePoolArgs']]]]] = None,
                 pod_subnet_range: Optional[pulumi.Input[str]] = None,
                 provider_account: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_peering_range: Optional[pulumi.Input[str]] = None,
                 service_subnet_range: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_subnet_range: Optional[pulumi.Input[str]] = None,
                 workspace_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        A cluster within an organization. An Astro cluster is a Kubernetes cluster that hosts the infrastructure required to run Deployments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi-astronomer as astronomer

        dedicated = astronomer.Workspace("dedicated",
            cicd_enforced_default=True,
            description="Workspace that demos a dedicated deployment set up")
        aws_dedicated = astronomer.Cluster("awsDedicated",
            cloud_provider="AWS",
            region="us-east-1",
            type="DEDICATED",
            vpc_subnet_range="172.20.0.0/20",
            k8s_tags=[],
            node_pools=[],
            workspace_ids=[dedicated.id])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_provider: The cluster's cloud provider.
        :param pulumi.Input[str] db_instance_type: The type of database instance that is used for the cluster. Required for Hybrid clusters.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterK8sTagArgs']]]] k8s_tags: The Kubernetes tags in the cluster.
        :param pulumi.Input[str] name: The name of the node pool.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodePoolArgs']]]] node_pools: The list of node pools to create in the cluster.
        :param pulumi.Input[str] pod_subnet_range: The subnet range for Pods. For GCP clusters only.
        :param pulumi.Input[str] provider_account: The provider account ID. Required for Hybrid clusters.
        :param pulumi.Input[str] region: The cluster's region.
        :param pulumi.Input[str] service_peering_range: The service peering range. For GCP clusters only.
        :param pulumi.Input[str] service_subnet_range: The service subnet range. For GCP clusters only.
        :param pulumi.Input[str] tenant_id: The tenant ID. For Azure clusters only.
        :param pulumi.Input[str] type: The cluster's type.
        :param pulumi.Input[str] vpc_subnet_range: The VPC subnet range.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workspace_ids: The list of Workspaces that are authorized to the cluster.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A cluster within an organization. An Astro cluster is a Kubernetes cluster that hosts the infrastructure required to run Deployments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi-astronomer as astronomer

        dedicated = astronomer.Workspace("dedicated",
            cicd_enforced_default=True,
            description="Workspace that demos a dedicated deployment set up")
        aws_dedicated = astronomer.Cluster("awsDedicated",
            cloud_provider="AWS",
            region="us-east-1",
            type="DEDICATED",
            vpc_subnet_range="172.20.0.0/20",
            k8s_tags=[],
            node_pools=[],
            workspace_ids=[dedicated.id])
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
                 db_instance_type: Optional[pulumi.Input[str]] = None,
                 k8s_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterK8sTagArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodePoolArgs']]]]] = None,
                 pod_subnet_range: Optional[pulumi.Input[str]] = None,
                 provider_account: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_peering_range: Optional[pulumi.Input[str]] = None,
                 service_subnet_range: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
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
            __props__.__dict__["db_instance_type"] = db_instance_type
            __props__.__dict__["k8s_tags"] = k8s_tags
            __props__.__dict__["name"] = name
            __props__.__dict__["node_pools"] = node_pools
            __props__.__dict__["pod_subnet_range"] = pod_subnet_range
            __props__.__dict__["provider_account"] = provider_account
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["service_peering_range"] = service_peering_range
            __props__.__dict__["service_subnet_range"] = service_subnet_range
            __props__.__dict__["tenant_id"] = tenant_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if vpc_subnet_range is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_subnet_range'")
            __props__.__dict__["vpc_subnet_range"] = vpc_subnet_range
            if workspace_ids is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_ids'")
            __props__.__dict__["workspace_ids"] = workspace_ids
            __props__.__dict__["is_limited"] = None
            __props__.__dict__["metadata"] = None
            __props__.__dict__["organization_id"] = None
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
            db_instance_type: Optional[pulumi.Input[str]] = None,
            is_limited: Optional[pulumi.Input[bool]] = None,
            k8s_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterK8sTagArgs']]]]] = None,
            metadata: Optional[pulumi.Input[pulumi.InputType['ClusterMetadataArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodePoolArgs']]]]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            pod_subnet_range: Optional[pulumi.Input[str]] = None,
            provider_account: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            service_peering_range: Optional[pulumi.Input[str]] = None,
            service_subnet_range: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vpc_subnet_range: Optional[pulumi.Input[str]] = None,
            workspace_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_provider: The cluster's cloud provider.
        :param pulumi.Input[str] db_instance_type: The type of database instance that is used for the cluster. Required for Hybrid clusters.
        :param pulumi.Input[bool] is_limited: Whether the cluster is limited.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterK8sTagArgs']]]] k8s_tags: The Kubernetes tags in the cluster.
        :param pulumi.Input[pulumi.InputType['ClusterMetadataArgs']] metadata: The cluster's metadata.
        :param pulumi.Input[str] name: The name of the node pool.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodePoolArgs']]]] node_pools: The list of node pools to create in the cluster.
        :param pulumi.Input[str] organization_id: The organization this cluster is associated with.
        :param pulumi.Input[str] pod_subnet_range: The subnet range for Pods. For GCP clusters only.
        :param pulumi.Input[str] provider_account: The provider account ID. Required for Hybrid clusters.
        :param pulumi.Input[str] region: The cluster's region.
        :param pulumi.Input[str] service_peering_range: The service peering range. For GCP clusters only.
        :param pulumi.Input[str] service_subnet_range: The service subnet range. For GCP clusters only.
        :param pulumi.Input[str] tenant_id: The tenant ID. For Azure clusters only.
        :param pulumi.Input[str] type: The cluster's type.
        :param pulumi.Input[str] vpc_subnet_range: The VPC subnet range.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workspace_ids: The list of Workspaces that are authorized to the cluster.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["cloud_provider"] = cloud_provider
        __props__.__dict__["db_instance_type"] = db_instance_type
        __props__.__dict__["is_limited"] = is_limited
        __props__.__dict__["k8s_tags"] = k8s_tags
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["name"] = name
        __props__.__dict__["node_pools"] = node_pools
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["pod_subnet_range"] = pod_subnet_range
        __props__.__dict__["provider_account"] = provider_account
        __props__.__dict__["region"] = region
        __props__.__dict__["service_peering_range"] = service_peering_range
        __props__.__dict__["service_subnet_range"] = service_subnet_range
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["type"] = type
        __props__.__dict__["vpc_subnet_range"] = vpc_subnet_range
        __props__.__dict__["workspace_ids"] = workspace_ids
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> pulumi.Output[str]:
        """
        The cluster's cloud provider.
        """
        return pulumi.get(self, "cloud_provider")

    @property
    @pulumi.getter(name="dbInstanceType")
    def db_instance_type(self) -> pulumi.Output[str]:
        """
        The type of database instance that is used for the cluster. Required for Hybrid clusters.
        """
        return pulumi.get(self, "db_instance_type")

    @property
    @pulumi.getter(name="isLimited")
    def is_limited(self) -> pulumi.Output[bool]:
        """
        Whether the cluster is limited.
        """
        return pulumi.get(self, "is_limited")

    @property
    @pulumi.getter(name="k8sTags")
    def k8s_tags(self) -> pulumi.Output[Optional[Sequence['outputs.ClusterK8sTag']]]:
        """
        The Kubernetes tags in the cluster.
        """
        return pulumi.get(self, "k8s_tags")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output['outputs.ClusterMetadata']:
        """
        The cluster's metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the node pool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> pulumi.Output[Optional[Sequence['outputs.ClusterNodePool']]]:
        """
        The list of node pools to create in the cluster.
        """
        return pulumi.get(self, "node_pools")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization this cluster is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="podSubnetRange")
    def pod_subnet_range(self) -> pulumi.Output[Optional[str]]:
        """
        The subnet range for Pods. For GCP clusters only.
        """
        return pulumi.get(self, "pod_subnet_range")

    @property
    @pulumi.getter(name="providerAccount")
    def provider_account(self) -> pulumi.Output[str]:
        """
        The provider account ID. Required for Hybrid clusters.
        """
        return pulumi.get(self, "provider_account")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The cluster's region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="servicePeeringRange")
    def service_peering_range(self) -> pulumi.Output[Optional[str]]:
        """
        The service peering range. For GCP clusters only.
        """
        return pulumi.get(self, "service_peering_range")

    @property
    @pulumi.getter(name="serviceSubnetRange")
    def service_subnet_range(self) -> pulumi.Output[Optional[str]]:
        """
        The service subnet range. For GCP clusters only.
        """
        return pulumi.get(self, "service_subnet_range")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The tenant ID. For Azure clusters only.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The cluster's type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcSubnetRange")
    def vpc_subnet_range(self) -> pulumi.Output[str]:
        """
        The VPC subnet range.
        """
        return pulumi.get(self, "vpc_subnet_range")

    @property
    @pulumi.getter(name="workspaceIds")
    def workspace_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of Workspaces that are authorized to the cluster.
        """
        return pulumi.get(self, "workspace_ids")

