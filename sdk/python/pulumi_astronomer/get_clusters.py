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
    'GetClustersResult',
    'AwaitableGetClustersResult',
    'get_clusters',
    'get_clusters_output',
]

@pulumi.output_type
class GetClustersResult:
    """
    A collection of values returned by getClusters.
    """
    def __init__(__self__, cloud_provider=None, clusters=None, id=None, names=None):
        if cloud_provider and not isinstance(cloud_provider, str):
            raise TypeError("Expected argument 'cloud_provider' to be a str")
        pulumi.set(__self__, "cloud_provider", cloud_provider)
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> Optional[str]:
        return pulumi.get(self, "cloud_provider")

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetClustersClusterResult']:
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "names")


class AwaitableGetClustersResult(GetClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClustersResult(
            cloud_provider=self.cloud_provider,
            clusters=self.clusters,
            id=self.id,
            names=self.names)


def get_clusters(cloud_provider: Optional[str] = None,
                 names: Optional[Sequence[str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClustersResult:
    """
    Clusters data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_clusters = astronomer.get_clusters()
    example_clusters_filter_by_names = astronomer.get_clusters(names=["my cluster"])
    example_clusters_filter_by_cloud_provider = astronomer.get_clusters(cloud_provider="AWS")
    pulumi.export("clusters", example_clusters)
    ```
    """
    __args__ = dict()
    __args__['cloudProvider'] = cloud_provider
    __args__['names'] = names
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('astronomer:index/getClusters:getClusters', __args__, opts=opts, typ=GetClustersResult).value

    return AwaitableGetClustersResult(
        cloud_provider=pulumi.get(__ret__, 'cloud_provider'),
        clusters=pulumi.get(__ret__, 'clusters'),
        id=pulumi.get(__ret__, 'id'),
        names=pulumi.get(__ret__, 'names'))


@_utilities.lift_output_func(get_clusters)
def get_clusters_output(cloud_provider: Optional[pulumi.Input[Optional[str]]] = None,
                        names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClustersResult]:
    """
    Clusters data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example_clusters = astronomer.get_clusters()
    example_clusters_filter_by_names = astronomer.get_clusters(names=["my cluster"])
    example_clusters_filter_by_cloud_provider = astronomer.get_clusters(cloud_provider="AWS")
    pulumi.export("clusters", example_clusters)
    ```
    """
    ...
