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
    'GetDeploymentOptionsResult',
    'AwaitableGetDeploymentOptionsResult',
    'get_deployment_options',
    'get_deployment_options_output',
]

@pulumi.output_type
class GetDeploymentOptionsResult:
    """
    A collection of values returned by getDeploymentOptions.
    """
    def __init__(__self__, cloud_provider=None, deployment_id=None, deployment_type=None, executor=None, executors=None, id=None, resource_quotas=None, runtime_releases=None, scheduler_machines=None, worker_machines=None, worker_queues=None, workload_identity_options=None):
        if cloud_provider and not isinstance(cloud_provider, str):
            raise TypeError("Expected argument 'cloud_provider' to be a str")
        pulumi.set(__self__, "cloud_provider", cloud_provider)
        if deployment_id and not isinstance(deployment_id, str):
            raise TypeError("Expected argument 'deployment_id' to be a str")
        pulumi.set(__self__, "deployment_id", deployment_id)
        if deployment_type and not isinstance(deployment_type, str):
            raise TypeError("Expected argument 'deployment_type' to be a str")
        pulumi.set(__self__, "deployment_type", deployment_type)
        if executor and not isinstance(executor, str):
            raise TypeError("Expected argument 'executor' to be a str")
        pulumi.set(__self__, "executor", executor)
        if executors and not isinstance(executors, list):
            raise TypeError("Expected argument 'executors' to be a list")
        pulumi.set(__self__, "executors", executors)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if resource_quotas and not isinstance(resource_quotas, dict):
            raise TypeError("Expected argument 'resource_quotas' to be a dict")
        pulumi.set(__self__, "resource_quotas", resource_quotas)
        if runtime_releases and not isinstance(runtime_releases, list):
            raise TypeError("Expected argument 'runtime_releases' to be a list")
        pulumi.set(__self__, "runtime_releases", runtime_releases)
        if scheduler_machines and not isinstance(scheduler_machines, list):
            raise TypeError("Expected argument 'scheduler_machines' to be a list")
        pulumi.set(__self__, "scheduler_machines", scheduler_machines)
        if worker_machines and not isinstance(worker_machines, list):
            raise TypeError("Expected argument 'worker_machines' to be a list")
        pulumi.set(__self__, "worker_machines", worker_machines)
        if worker_queues and not isinstance(worker_queues, dict):
            raise TypeError("Expected argument 'worker_queues' to be a dict")
        pulumi.set(__self__, "worker_queues", worker_queues)
        if workload_identity_options and not isinstance(workload_identity_options, list):
            raise TypeError("Expected argument 'workload_identity_options' to be a list")
        pulumi.set(__self__, "workload_identity_options", workload_identity_options)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> Optional[str]:
        """
        Cloud provider
        """
        return pulumi.get(self, "cloud_provider")

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[str]:
        """
        Deployment ID
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter(name="deploymentType")
    def deployment_type(self) -> Optional[str]:
        """
        Deployment type
        """
        return pulumi.get(self, "deployment_type")

    @property
    @pulumi.getter
    def executor(self) -> Optional[str]:
        """
        Executor
        """
        return pulumi.get(self, "executor")

    @property
    @pulumi.getter
    def executors(self) -> Sequence[str]:
        """
        Available executors
        """
        return pulumi.get(self, "executors")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceQuotas")
    def resource_quotas(self) -> 'outputs.GetDeploymentOptionsResourceQuotasResult':
        """
        Resource quota options
        """
        return pulumi.get(self, "resource_quotas")

    @property
    @pulumi.getter(name="runtimeReleases")
    def runtime_releases(self) -> Sequence['outputs.GetDeploymentOptionsRuntimeReleaseResult']:
        """
        Available Astro Runtime versions
        """
        return pulumi.get(self, "runtime_releases")

    @property
    @pulumi.getter(name="schedulerMachines")
    def scheduler_machines(self) -> Sequence['outputs.GetDeploymentOptionsSchedulerMachineResult']:
        """
        Available scheduler sizes
        """
        return pulumi.get(self, "scheduler_machines")

    @property
    @pulumi.getter(name="workerMachines")
    def worker_machines(self) -> Sequence['outputs.GetDeploymentOptionsWorkerMachineResult']:
        """
        Available worker machine types
        """
        return pulumi.get(self, "worker_machines")

    @property
    @pulumi.getter(name="workerQueues")
    def worker_queues(self) -> 'outputs.GetDeploymentOptionsWorkerQueuesResult':
        """
        Available worker queue options
        """
        return pulumi.get(self, "worker_queues")

    @property
    @pulumi.getter(name="workloadIdentityOptions")
    def workload_identity_options(self) -> Sequence['outputs.GetDeploymentOptionsWorkloadIdentityOptionResult']:
        """
        Available workload identity options
        """
        return pulumi.get(self, "workload_identity_options")


class AwaitableGetDeploymentOptionsResult(GetDeploymentOptionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeploymentOptionsResult(
            cloud_provider=self.cloud_provider,
            deployment_id=self.deployment_id,
            deployment_type=self.deployment_type,
            executor=self.executor,
            executors=self.executors,
            id=self.id,
            resource_quotas=self.resource_quotas,
            runtime_releases=self.runtime_releases,
            scheduler_machines=self.scheduler_machines,
            worker_machines=self.worker_machines,
            worker_queues=self.worker_queues,
            workload_identity_options=self.workload_identity_options)


def get_deployment_options(cloud_provider: Optional[str] = None,
                           deployment_id: Optional[str] = None,
                           deployment_type: Optional[str] = None,
                           executor: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeploymentOptionsResult:
    """
    Deployment options data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example = astronomer.get_deployment_options()
    example_with_deployment_id_query_param = astronomer.get_deployment_options(deployment_id="clozc036j01to01jrlgvueo8t")
    example_with_deployment_type_query_param = astronomer.get_deployment_options(deployment_type="DEDICATED")
    example_with_executor_query_param = astronomer.get_deployment_options(executor="CELERY")
    example_with_cloud_provider_query_param = astronomer.get_deployment_options(cloud_provider="AWS")
    ```


    :param str cloud_provider: Cloud provider
    :param str deployment_id: Deployment ID
    :param str deployment_type: Deployment type
    :param str executor: Executor
    """
    __args__ = dict()
    __args__['cloudProvider'] = cloud_provider
    __args__['deploymentId'] = deployment_id
    __args__['deploymentType'] = deployment_type
    __args__['executor'] = executor
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('astronomer:index/getDeploymentOptions:getDeploymentOptions', __args__, opts=opts, typ=GetDeploymentOptionsResult).value

    return AwaitableGetDeploymentOptionsResult(
        cloud_provider=pulumi.get(__ret__, 'cloud_provider'),
        deployment_id=pulumi.get(__ret__, 'deployment_id'),
        deployment_type=pulumi.get(__ret__, 'deployment_type'),
        executor=pulumi.get(__ret__, 'executor'),
        executors=pulumi.get(__ret__, 'executors'),
        id=pulumi.get(__ret__, 'id'),
        resource_quotas=pulumi.get(__ret__, 'resource_quotas'),
        runtime_releases=pulumi.get(__ret__, 'runtime_releases'),
        scheduler_machines=pulumi.get(__ret__, 'scheduler_machines'),
        worker_machines=pulumi.get(__ret__, 'worker_machines'),
        worker_queues=pulumi.get(__ret__, 'worker_queues'),
        workload_identity_options=pulumi.get(__ret__, 'workload_identity_options'))


@_utilities.lift_output_func(get_deployment_options)
def get_deployment_options_output(cloud_provider: Optional[pulumi.Input[Optional[str]]] = None,
                                  deployment_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  deployment_type: Optional[pulumi.Input[Optional[str]]] = None,
                                  executor: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeploymentOptionsResult]:
    """
    Deployment options data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astronomer as astronomer

    example = astronomer.get_deployment_options()
    example_with_deployment_id_query_param = astronomer.get_deployment_options(deployment_id="clozc036j01to01jrlgvueo8t")
    example_with_deployment_type_query_param = astronomer.get_deployment_options(deployment_type="DEDICATED")
    example_with_executor_query_param = astronomer.get_deployment_options(executor="CELERY")
    example_with_cloud_provider_query_param = astronomer.get_deployment_options(cloud_provider="AWS")
    ```


    :param str cloud_provider: Cloud provider
    :param str deployment_id: Deployment ID
    :param str deployment_type: Deployment type
    :param str executor: Executor
    """
    ...