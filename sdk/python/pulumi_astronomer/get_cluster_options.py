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
    'GetClusterOptionsResult',
    'AwaitableGetClusterOptionsResult',
    'get_cluster_options',
    'get_cluster_options_output',
]

@pulumi.output_type
class GetClusterOptionsResult:
    """
    A collection of values returned by getClusterOptions.
    """
    def __init__(__self__, cloud_provider=None, cluster_options=None, id=None, type=None):
        if cloud_provider and not isinstance(cloud_provider, str):
            raise TypeError("Expected argument 'cloud_provider' to be a str")
        pulumi.set(__self__, "cloud_provider", cloud_provider)
        if cluster_options and not isinstance(cluster_options, list):
            raise TypeError("Expected argument 'cluster_options' to be a list")
        pulumi.set(__self__, "cluster_options", cluster_options)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> Optional[str]:
        return pulumi.get(self, "cloud_provider")

    @property
    @pulumi.getter(name="clusterOptions")
    def cluster_options(self) -> Sequence['outputs.GetClusterOptionsClusterOptionResult']:
        return pulumi.get(self, "cluster_options")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetClusterOptionsResult(GetClusterOptionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterOptionsResult(
            cloud_provider=self.cloud_provider,
            cluster_options=self.cluster_options,
            id=self.id,
            type=self.type)


def get_cluster_options(cloud_provider: Optional[str] = None,
                        type: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterOptionsResult:
    """
    ClusterOptions data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_cluster_options = astronomer.get_cluster_options(type="HYBRID")
    example_cluster_options_filter_by_provider = astronomer.get_cluster_options(type="HYBRID",
        cloud_provider="AWS")
    pulumi.export("clusterOptions", example_cluster_options)
    ```
    """
    __args__ = dict()
    __args__['cloudProvider'] = cloud_provider
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('astronomer:index/getClusterOptions:getClusterOptions', __args__, opts=opts, typ=GetClusterOptionsResult).value

    return AwaitableGetClusterOptionsResult(
        cloud_provider=pulumi.get(__ret__, 'cloud_provider'),
        cluster_options=pulumi.get(__ret__, 'cluster_options'),
        id=pulumi.get(__ret__, 'id'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_cluster_options)
def get_cluster_options_output(cloud_provider: Optional[pulumi.Input[Optional[str]]] = None,
                               type: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterOptionsResult]:
    """
    ClusterOptions data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_cluster_options = astronomer.get_cluster_options(type="HYBRID")
    example_cluster_options_filter_by_provider = astronomer.get_cluster_options(type="HYBRID",
        cloud_provider="AWS")
    pulumi.export("clusterOptions", example_cluster_options)
    ```
    """
    ...
