// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astronomer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan.pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

// An Astro Deployment is an Airflow environment that is powered by all core Airflow components.
type Deployment struct {
	pulumi.CustomResourceState

	// Deployment's Astro Runtime version.
	AstroRuntimeVersion pulumi.StringPtrOutput `pulumi:"astroRuntimeVersion"`
	// The cloud provider for the Deployment's cluster. Optional if `ClusterId` is specified.
	CloudProvider pulumi.StringOutput `pulumi:"cloudProvider"`
	// The ID of the cluster where the Deployment will be created.
	ClusterId pulumi.StringPtrOutput `pulumi:"clusterId"`
	// The default CPU resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in number of CPU cores.
	DefaultTaskPodCpu pulumi.StringOutput `pulumi:"defaultTaskPodCpu"`
	// The default memory resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in `Gi`. This value must always be twice the value of `DefaultTaskPodCpu`.
	DefaultTaskPodMemory pulumi.StringOutput `pulumi:"defaultTaskPodMemory"`
	// The Deployment's description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of environment variables to add to the Deployment.
	EnvironmentVariables DeploymentEnvironmentVariableArrayOutput `pulumi:"environmentVariables"`
	// The Deployment's executor type.
	Executor pulumi.StringOutput `pulumi:"executor"`
	// Whether the Deployment requires that all deploys are made through CI/CD.
	IsCicdEnforced pulumi.BoolOutput `pulumi:"isCicdEnforced"`
	// Whether the Deployment has DAG deploys enabled.
	IsDagDeployEnabled pulumi.BoolOutput `pulumi:"isDagDeployEnabled"`
	// Whether the Deployment is configured for high availability. If `true`, multiple scheduler pods will be online.
	IsHighAvailability pulumi.BoolOutput `pulumi:"isHighAvailability"`
	// The Deployment's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region to host the Deployment in. Optional if `ClusterId` is specified.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The CPU quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current CPU usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in number of CPU cores.
	ResourceQuotaCpu pulumi.StringOutput `pulumi:"resourceQuotaCpu"`
	// The memory quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current memory usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in `Gi`. This value must always be twice the value of `ResourceQuotaCpu`.
	ResourceQuotaMemory pulumi.StringOutput `pulumi:"resourceQuotaMemory"`
	// The size of the scheduler pod.
	SchedulerSize pulumi.StringOutput `pulumi:"schedulerSize"`
	// The node pool ID for the task pods. For KUBERNETES executor only.
	TaskPodNodePoolId pulumi.StringPtrOutput `pulumi:"taskPodNodePoolId"`
	// The type of the Deployment.
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of worker queues configured for the Deployment. Applies only when `Executor` is `CELERY`. At least 1 worker queue is needed. All Deployments need at least 1 worker queue called `default`.
	WorkerQueues DeploymentWorkerQueueArrayOutput `pulumi:"workerQueues"`
	// The Deployment's workload identity.
	WorkloadIdentity pulumi.StringOutput `pulumi:"workloadIdentity"`
	// The ID of the workspace to which the Deployment belongs.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultTaskPodCpu == nil {
		return nil, errors.New("invalid value for required argument 'DefaultTaskPodCpu'")
	}
	if args.DefaultTaskPodMemory == nil {
		return nil, errors.New("invalid value for required argument 'DefaultTaskPodMemory'")
	}
	if args.Executor == nil {
		return nil, errors.New("invalid value for required argument 'Executor'")
	}
	if args.IsCicdEnforced == nil {
		return nil, errors.New("invalid value for required argument 'IsCicdEnforced'")
	}
	if args.IsDagDeployEnabled == nil {
		return nil, errors.New("invalid value for required argument 'IsDagDeployEnabled'")
	}
	if args.IsHighAvailability == nil {
		return nil, errors.New("invalid value for required argument 'IsHighAvailability'")
	}
	if args.ResourceQuotaCpu == nil {
		return nil, errors.New("invalid value for required argument 'ResourceQuotaCpu'")
	}
	if args.ResourceQuotaMemory == nil {
		return nil, errors.New("invalid value for required argument 'ResourceQuotaMemory'")
	}
	if args.SchedulerSize == nil {
		return nil, errors.New("invalid value for required argument 'SchedulerSize'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Deployment
	err := ctx.RegisterResource("astronomer:index/deployment:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("astronomer:index/deployment:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
	// Deployment's Astro Runtime version.
	AstroRuntimeVersion *string `pulumi:"astroRuntimeVersion"`
	// The cloud provider for the Deployment's cluster. Optional if `ClusterId` is specified.
	CloudProvider *string `pulumi:"cloudProvider"`
	// The ID of the cluster where the Deployment will be created.
	ClusterId *string `pulumi:"clusterId"`
	// The default CPU resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in number of CPU cores.
	DefaultTaskPodCpu *string `pulumi:"defaultTaskPodCpu"`
	// The default memory resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in `Gi`. This value must always be twice the value of `DefaultTaskPodCpu`.
	DefaultTaskPodMemory *string `pulumi:"defaultTaskPodMemory"`
	// The Deployment's description.
	Description *string `pulumi:"description"`
	// List of environment variables to add to the Deployment.
	EnvironmentVariables []DeploymentEnvironmentVariable `pulumi:"environmentVariables"`
	// The Deployment's executor type.
	Executor *string `pulumi:"executor"`
	// Whether the Deployment requires that all deploys are made through CI/CD.
	IsCicdEnforced *bool `pulumi:"isCicdEnforced"`
	// Whether the Deployment has DAG deploys enabled.
	IsDagDeployEnabled *bool `pulumi:"isDagDeployEnabled"`
	// Whether the Deployment is configured for high availability. If `true`, multiple scheduler pods will be online.
	IsHighAvailability *bool `pulumi:"isHighAvailability"`
	// The Deployment's name.
	Name *string `pulumi:"name"`
	// The region to host the Deployment in. Optional if `ClusterId` is specified.
	Region *string `pulumi:"region"`
	// The CPU quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current CPU usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in number of CPU cores.
	ResourceQuotaCpu *string `pulumi:"resourceQuotaCpu"`
	// The memory quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current memory usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in `Gi`. This value must always be twice the value of `ResourceQuotaCpu`.
	ResourceQuotaMemory *string `pulumi:"resourceQuotaMemory"`
	// The size of the scheduler pod.
	SchedulerSize *string `pulumi:"schedulerSize"`
	// The node pool ID for the task pods. For KUBERNETES executor only.
	TaskPodNodePoolId *string `pulumi:"taskPodNodePoolId"`
	// The type of the Deployment.
	Type *string `pulumi:"type"`
	// The list of worker queues configured for the Deployment. Applies only when `Executor` is `CELERY`. At least 1 worker queue is needed. All Deployments need at least 1 worker queue called `default`.
	WorkerQueues []DeploymentWorkerQueue `pulumi:"workerQueues"`
	// The Deployment's workload identity.
	WorkloadIdentity *string `pulumi:"workloadIdentity"`
	// The ID of the workspace to which the Deployment belongs.
	WorkspaceId *string `pulumi:"workspaceId"`
}

type DeploymentState struct {
	// Deployment's Astro Runtime version.
	AstroRuntimeVersion pulumi.StringPtrInput
	// The cloud provider for the Deployment's cluster. Optional if `ClusterId` is specified.
	CloudProvider pulumi.StringPtrInput
	// The ID of the cluster where the Deployment will be created.
	ClusterId pulumi.StringPtrInput
	// The default CPU resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in number of CPU cores.
	DefaultTaskPodCpu pulumi.StringPtrInput
	// The default memory resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in `Gi`. This value must always be twice the value of `DefaultTaskPodCpu`.
	DefaultTaskPodMemory pulumi.StringPtrInput
	// The Deployment's description.
	Description pulumi.StringPtrInput
	// List of environment variables to add to the Deployment.
	EnvironmentVariables DeploymentEnvironmentVariableArrayInput
	// The Deployment's executor type.
	Executor pulumi.StringPtrInput
	// Whether the Deployment requires that all deploys are made through CI/CD.
	IsCicdEnforced pulumi.BoolPtrInput
	// Whether the Deployment has DAG deploys enabled.
	IsDagDeployEnabled pulumi.BoolPtrInput
	// Whether the Deployment is configured for high availability. If `true`, multiple scheduler pods will be online.
	IsHighAvailability pulumi.BoolPtrInput
	// The Deployment's name.
	Name pulumi.StringPtrInput
	// The region to host the Deployment in. Optional if `ClusterId` is specified.
	Region pulumi.StringPtrInput
	// The CPU quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current CPU usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in number of CPU cores.
	ResourceQuotaCpu pulumi.StringPtrInput
	// The memory quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current memory usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in `Gi`. This value must always be twice the value of `ResourceQuotaCpu`.
	ResourceQuotaMemory pulumi.StringPtrInput
	// The size of the scheduler pod.
	SchedulerSize pulumi.StringPtrInput
	// The node pool ID for the task pods. For KUBERNETES executor only.
	TaskPodNodePoolId pulumi.StringPtrInput
	// The type of the Deployment.
	Type pulumi.StringPtrInput
	// The list of worker queues configured for the Deployment. Applies only when `Executor` is `CELERY`. At least 1 worker queue is needed. All Deployments need at least 1 worker queue called `default`.
	WorkerQueues DeploymentWorkerQueueArrayInput
	// The Deployment's workload identity.
	WorkloadIdentity pulumi.StringPtrInput
	// The ID of the workspace to which the Deployment belongs.
	WorkspaceId pulumi.StringPtrInput
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// Deployment's Astro Runtime version.
	AstroRuntimeVersion *string `pulumi:"astroRuntimeVersion"`
	// The cloud provider for the Deployment's cluster. Optional if `ClusterId` is specified.
	CloudProvider *string `pulumi:"cloudProvider"`
	// The ID of the cluster where the Deployment will be created.
	ClusterId *string `pulumi:"clusterId"`
	// The default CPU resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in number of CPU cores.
	DefaultTaskPodCpu string `pulumi:"defaultTaskPodCpu"`
	// The default memory resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in `Gi`. This value must always be twice the value of `DefaultTaskPodCpu`.
	DefaultTaskPodMemory string `pulumi:"defaultTaskPodMemory"`
	// The Deployment's description.
	Description *string `pulumi:"description"`
	// List of environment variables to add to the Deployment.
	EnvironmentVariables []DeploymentEnvironmentVariable `pulumi:"environmentVariables"`
	// The Deployment's executor type.
	Executor string `pulumi:"executor"`
	// Whether the Deployment requires that all deploys are made through CI/CD.
	IsCicdEnforced bool `pulumi:"isCicdEnforced"`
	// Whether the Deployment has DAG deploys enabled.
	IsDagDeployEnabled bool `pulumi:"isDagDeployEnabled"`
	// Whether the Deployment is configured for high availability. If `true`, multiple scheduler pods will be online.
	IsHighAvailability bool `pulumi:"isHighAvailability"`
	// The Deployment's name.
	Name *string `pulumi:"name"`
	// The region to host the Deployment in. Optional if `ClusterId` is specified.
	Region *string `pulumi:"region"`
	// The CPU quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current CPU usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in number of CPU cores.
	ResourceQuotaCpu string `pulumi:"resourceQuotaCpu"`
	// The memory quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current memory usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in `Gi`. This value must always be twice the value of `ResourceQuotaCpu`.
	ResourceQuotaMemory string `pulumi:"resourceQuotaMemory"`
	// The size of the scheduler pod.
	SchedulerSize string `pulumi:"schedulerSize"`
	// The node pool ID for the task pods. For KUBERNETES executor only.
	TaskPodNodePoolId *string `pulumi:"taskPodNodePoolId"`
	// The type of the Deployment.
	Type string `pulumi:"type"`
	// The list of worker queues configured for the Deployment. Applies only when `Executor` is `CELERY`. At least 1 worker queue is needed. All Deployments need at least 1 worker queue called `default`.
	WorkerQueues []DeploymentWorkerQueue `pulumi:"workerQueues"`
	// The ID of the workspace to which the Deployment belongs.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// Deployment's Astro Runtime version.
	AstroRuntimeVersion pulumi.StringPtrInput
	// The cloud provider for the Deployment's cluster. Optional if `ClusterId` is specified.
	CloudProvider pulumi.StringPtrInput
	// The ID of the cluster where the Deployment will be created.
	ClusterId pulumi.StringPtrInput
	// The default CPU resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in number of CPU cores.
	DefaultTaskPodCpu pulumi.StringInput
	// The default memory resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in `Gi`. This value must always be twice the value of `DefaultTaskPodCpu`.
	DefaultTaskPodMemory pulumi.StringInput
	// The Deployment's description.
	Description pulumi.StringPtrInput
	// List of environment variables to add to the Deployment.
	EnvironmentVariables DeploymentEnvironmentVariableArrayInput
	// The Deployment's executor type.
	Executor pulumi.StringInput
	// Whether the Deployment requires that all deploys are made through CI/CD.
	IsCicdEnforced pulumi.BoolInput
	// Whether the Deployment has DAG deploys enabled.
	IsDagDeployEnabled pulumi.BoolInput
	// Whether the Deployment is configured for high availability. If `true`, multiple scheduler pods will be online.
	IsHighAvailability pulumi.BoolInput
	// The Deployment's name.
	Name pulumi.StringPtrInput
	// The region to host the Deployment in. Optional if `ClusterId` is specified.
	Region pulumi.StringPtrInput
	// The CPU quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current CPU usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in number of CPU cores.
	ResourceQuotaCpu pulumi.StringInput
	// The memory quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current memory usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in `Gi`. This value must always be twice the value of `ResourceQuotaCpu`.
	ResourceQuotaMemory pulumi.StringInput
	// The size of the scheduler pod.
	SchedulerSize pulumi.StringInput
	// The node pool ID for the task pods. For KUBERNETES executor only.
	TaskPodNodePoolId pulumi.StringPtrInput
	// The type of the Deployment.
	Type pulumi.StringInput
	// The list of worker queues configured for the Deployment. Applies only when `Executor` is `CELERY`. At least 1 worker queue is needed. All Deployments need at least 1 worker queue called `default`.
	WorkerQueues DeploymentWorkerQueueArrayInput
	// The ID of the workspace to which the Deployment belongs.
	WorkspaceId pulumi.StringInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

// DeploymentArrayInput is an input type that accepts DeploymentArray and DeploymentArrayOutput values.
// You can construct a concrete instance of `DeploymentArrayInput` via:
//
//	DeploymentArray{ DeploymentArgs{...} }
type DeploymentArrayInput interface {
	pulumi.Input

	ToDeploymentArrayOutput() DeploymentArrayOutput
	ToDeploymentArrayOutputWithContext(context.Context) DeploymentArrayOutput
}

type DeploymentArray []DeploymentInput

func (DeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deployment)(nil)).Elem()
}

func (i DeploymentArray) ToDeploymentArrayOutput() DeploymentArrayOutput {
	return i.ToDeploymentArrayOutputWithContext(context.Background())
}

func (i DeploymentArray) ToDeploymentArrayOutputWithContext(ctx context.Context) DeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentArrayOutput)
}

// DeploymentMapInput is an input type that accepts DeploymentMap and DeploymentMapOutput values.
// You can construct a concrete instance of `DeploymentMapInput` via:
//
//	DeploymentMap{ "key": DeploymentArgs{...} }
type DeploymentMapInput interface {
	pulumi.Input

	ToDeploymentMapOutput() DeploymentMapOutput
	ToDeploymentMapOutputWithContext(context.Context) DeploymentMapOutput
}

type DeploymentMap map[string]DeploymentInput

func (DeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deployment)(nil)).Elem()
}

func (i DeploymentMap) ToDeploymentMapOutput() DeploymentMapOutput {
	return i.ToDeploymentMapOutputWithContext(context.Background())
}

func (i DeploymentMap) ToDeploymentMapOutputWithContext(ctx context.Context) DeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentMapOutput)
}

type DeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

// Deployment's Astro Runtime version.
func (o DeploymentOutput) AstroRuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.AstroRuntimeVersion }).(pulumi.StringPtrOutput)
}

// The cloud provider for the Deployment's cluster. Optional if `ClusterId` is specified.
func (o DeploymentOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.CloudProvider }).(pulumi.StringOutput)
}

// The ID of the cluster where the Deployment will be created.
func (o DeploymentOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The default CPU resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in number of CPU cores.
func (o DeploymentOutput) DefaultTaskPodCpu() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.DefaultTaskPodCpu }).(pulumi.StringOutput)
}

// The default memory resource usage for a worker Pod when running the Kubernetes executor or KubernetesPodOperator. Units are in `Gi`. This value must always be twice the value of `DefaultTaskPodCpu`.
func (o DeploymentOutput) DefaultTaskPodMemory() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.DefaultTaskPodMemory }).(pulumi.StringOutput)
}

// The Deployment's description.
func (o DeploymentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of environment variables to add to the Deployment.
func (o DeploymentOutput) EnvironmentVariables() DeploymentEnvironmentVariableArrayOutput {
	return o.ApplyT(func(v *Deployment) DeploymentEnvironmentVariableArrayOutput { return v.EnvironmentVariables }).(DeploymentEnvironmentVariableArrayOutput)
}

// The Deployment's executor type.
func (o DeploymentOutput) Executor() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Executor }).(pulumi.StringOutput)
}

// Whether the Deployment requires that all deploys are made through CI/CD.
func (o DeploymentOutput) IsCicdEnforced() pulumi.BoolOutput {
	return o.ApplyT(func(v *Deployment) pulumi.BoolOutput { return v.IsCicdEnforced }).(pulumi.BoolOutput)
}

// Whether the Deployment has DAG deploys enabled.
func (o DeploymentOutput) IsDagDeployEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Deployment) pulumi.BoolOutput { return v.IsDagDeployEnabled }).(pulumi.BoolOutput)
}

// Whether the Deployment is configured for high availability. If `true`, multiple scheduler pods will be online.
func (o DeploymentOutput) IsHighAvailability() pulumi.BoolOutput {
	return o.ApplyT(func(v *Deployment) pulumi.BoolOutput { return v.IsHighAvailability }).(pulumi.BoolOutput)
}

// The Deployment's name.
func (o DeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region to host the Deployment in. Optional if `ClusterId` is specified.
func (o DeploymentOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The CPU quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current CPU usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in number of CPU cores.
func (o DeploymentOutput) ResourceQuotaCpu() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.ResourceQuotaCpu }).(pulumi.StringOutput)
}

// The memory quota for worker Pods when running the Kubernetes executor or KubernetesPodOperator. If current memory usage across all workers exceeds the quota, no new worker Pods can be scheduled. Units are in `Gi`. This value must always be twice the value of `ResourceQuotaCpu`.
func (o DeploymentOutput) ResourceQuotaMemory() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.ResourceQuotaMemory }).(pulumi.StringOutput)
}

// The size of the scheduler pod.
func (o DeploymentOutput) SchedulerSize() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.SchedulerSize }).(pulumi.StringOutput)
}

// The node pool ID for the task pods. For KUBERNETES executor only.
func (o DeploymentOutput) TaskPodNodePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.TaskPodNodePoolId }).(pulumi.StringPtrOutput)
}

// The type of the Deployment.
func (o DeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of worker queues configured for the Deployment. Applies only when `Executor` is `CELERY`. At least 1 worker queue is needed. All Deployments need at least 1 worker queue called `default`.
func (o DeploymentOutput) WorkerQueues() DeploymentWorkerQueueArrayOutput {
	return o.ApplyT(func(v *Deployment) DeploymentWorkerQueueArrayOutput { return v.WorkerQueues }).(DeploymentWorkerQueueArrayOutput)
}

// The Deployment's workload identity.
func (o DeploymentOutput) WorkloadIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.WorkloadIdentity }).(pulumi.StringOutput)
}

// The ID of the workspace to which the Deployment belongs.
func (o DeploymentOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type DeploymentArrayOutput struct{ *pulumi.OutputState }

func (DeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deployment)(nil)).Elem()
}

func (o DeploymentArrayOutput) ToDeploymentArrayOutput() DeploymentArrayOutput {
	return o
}

func (o DeploymentArrayOutput) ToDeploymentArrayOutputWithContext(ctx context.Context) DeploymentArrayOutput {
	return o
}

func (o DeploymentArrayOutput) Index(i pulumi.IntInput) DeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Deployment {
		return vs[0].([]*Deployment)[vs[1].(int)]
	}).(DeploymentOutput)
}

type DeploymentMapOutput struct{ *pulumi.OutputState }

func (DeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deployment)(nil)).Elem()
}

func (o DeploymentMapOutput) ToDeploymentMapOutput() DeploymentMapOutput {
	return o
}

func (o DeploymentMapOutput) ToDeploymentMapOutputWithContext(ctx context.Context) DeploymentMapOutput {
	return o
}

func (o DeploymentMapOutput) MapIndex(k pulumi.StringInput) DeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Deployment {
		return vs[0].(map[string]*Deployment)[vs[1].(string)]
	}).(DeploymentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentInput)(nil)).Elem(), &Deployment{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentArrayInput)(nil)).Elem(), DeploymentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentMapInput)(nil)).Elem(), DeploymentMap{})
	pulumi.RegisterOutputType(DeploymentOutput{})
	pulumi.RegisterOutputType(DeploymentArrayOutput{})
	pulumi.RegisterOutputType(DeploymentMapOutput{})
}
