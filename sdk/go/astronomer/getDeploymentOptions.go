// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astronomer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

// Deployment options data source
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleDeploymentOptions, err := astronomer.GetDeploymentOptions(ctx, &astronomer.GetDeploymentOptionsArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = astronomer.GetDeploymentOptions(ctx, &astronomer.GetDeploymentOptionsArgs{
//				DeploymentId: pulumi.StringRef("clozc036j01to01jrlgvueo8t"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = astronomer.GetDeploymentOptions(ctx, &astronomer.GetDeploymentOptionsArgs{
//				DeploymentType: pulumi.StringRef("DEDICATED"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = astronomer.GetDeploymentOptions(ctx, &astronomer.GetDeploymentOptionsArgs{
//				Executor: pulumi.StringRef("CELERY"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = astronomer.GetDeploymentOptions(ctx, &astronomer.GetDeploymentOptionsArgs{
//				CloudProvider: pulumi.StringRef("AWS"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("deploymentOptions", exampleDeploymentOptions)
//			return nil
//		})
//	}
//
// ```
func GetDeploymentOptions(ctx *pulumi.Context, args *GetDeploymentOptionsArgs, opts ...pulumi.InvokeOption) (*GetDeploymentOptionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDeploymentOptionsResult
	err := ctx.Invoke("astronomer:index/getDeploymentOptions:getDeploymentOptions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDeploymentOptions.
type GetDeploymentOptionsArgs struct {
	// Cloud provider
	CloudProvider *string `pulumi:"cloudProvider"`
	// Deployment ID
	DeploymentId *string `pulumi:"deploymentId"`
	// Deployment type
	DeploymentType *string `pulumi:"deploymentType"`
	// Executor
	Executor *string `pulumi:"executor"`
}

// A collection of values returned by getDeploymentOptions.
type GetDeploymentOptionsResult struct {
	// Cloud provider
	CloudProvider *string `pulumi:"cloudProvider"`
	// Deployment ID
	DeploymentId *string `pulumi:"deploymentId"`
	// Deployment type
	DeploymentType *string `pulumi:"deploymentType"`
	// Executor
	Executor *string `pulumi:"executor"`
	// Available executors
	Executors []string `pulumi:"executors"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Resource quota options
	ResourceQuotas GetDeploymentOptionsResourceQuotas `pulumi:"resourceQuotas"`
	// Available Astro Runtime versions
	RuntimeReleases []GetDeploymentOptionsRuntimeRelease `pulumi:"runtimeReleases"`
	// Available scheduler sizes
	SchedulerMachines []GetDeploymentOptionsSchedulerMachine `pulumi:"schedulerMachines"`
	// Available worker machine types
	WorkerMachines []GetDeploymentOptionsWorkerMachine `pulumi:"workerMachines"`
	// Available worker queue options
	WorkerQueues GetDeploymentOptionsWorkerQueues `pulumi:"workerQueues"`
	// Available workload identity options
	WorkloadIdentityOptions []GetDeploymentOptionsWorkloadIdentityOption `pulumi:"workloadIdentityOptions"`
}

func GetDeploymentOptionsOutput(ctx *pulumi.Context, args GetDeploymentOptionsOutputArgs, opts ...pulumi.InvokeOption) GetDeploymentOptionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeploymentOptionsResultOutput, error) {
			args := v.(GetDeploymentOptionsArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetDeploymentOptionsResult
			secret, err := ctx.InvokePackageRaw("astronomer:index/getDeploymentOptions:getDeploymentOptions", args, &rv, "", opts...)
			if err != nil {
				return GetDeploymentOptionsResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetDeploymentOptionsResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetDeploymentOptionsResultOutput), nil
			}
			return output, nil
		}).(GetDeploymentOptionsResultOutput)
}

// A collection of arguments for invoking getDeploymentOptions.
type GetDeploymentOptionsOutputArgs struct {
	// Cloud provider
	CloudProvider pulumi.StringPtrInput `pulumi:"cloudProvider"`
	// Deployment ID
	DeploymentId pulumi.StringPtrInput `pulumi:"deploymentId"`
	// Deployment type
	DeploymentType pulumi.StringPtrInput `pulumi:"deploymentType"`
	// Executor
	Executor pulumi.StringPtrInput `pulumi:"executor"`
}

func (GetDeploymentOptionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentOptionsArgs)(nil)).Elem()
}

// A collection of values returned by getDeploymentOptions.
type GetDeploymentOptionsResultOutput struct{ *pulumi.OutputState }

func (GetDeploymentOptionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentOptionsResult)(nil)).Elem()
}

func (o GetDeploymentOptionsResultOutput) ToGetDeploymentOptionsResultOutput() GetDeploymentOptionsResultOutput {
	return o
}

func (o GetDeploymentOptionsResultOutput) ToGetDeploymentOptionsResultOutputWithContext(ctx context.Context) GetDeploymentOptionsResultOutput {
	return o
}

// Cloud provider
func (o GetDeploymentOptionsResultOutput) CloudProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) *string { return v.CloudProvider }).(pulumi.StringPtrOutput)
}

// Deployment ID
func (o GetDeploymentOptionsResultOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

// Deployment type
func (o GetDeploymentOptionsResultOutput) DeploymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) *string { return v.DeploymentType }).(pulumi.StringPtrOutput)
}

// Executor
func (o GetDeploymentOptionsResultOutput) Executor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) *string { return v.Executor }).(pulumi.StringPtrOutput)
}

// Available executors
func (o GetDeploymentOptionsResultOutput) Executors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) []string { return v.Executors }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDeploymentOptionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource quota options
func (o GetDeploymentOptionsResultOutput) ResourceQuotas() GetDeploymentOptionsResourceQuotasOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) GetDeploymentOptionsResourceQuotas { return v.ResourceQuotas }).(GetDeploymentOptionsResourceQuotasOutput)
}

// Available Astro Runtime versions
func (o GetDeploymentOptionsResultOutput) RuntimeReleases() GetDeploymentOptionsRuntimeReleaseArrayOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) []GetDeploymentOptionsRuntimeRelease { return v.RuntimeReleases }).(GetDeploymentOptionsRuntimeReleaseArrayOutput)
}

// Available scheduler sizes
func (o GetDeploymentOptionsResultOutput) SchedulerMachines() GetDeploymentOptionsSchedulerMachineArrayOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) []GetDeploymentOptionsSchedulerMachine { return v.SchedulerMachines }).(GetDeploymentOptionsSchedulerMachineArrayOutput)
}

// Available worker machine types
func (o GetDeploymentOptionsResultOutput) WorkerMachines() GetDeploymentOptionsWorkerMachineArrayOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) []GetDeploymentOptionsWorkerMachine { return v.WorkerMachines }).(GetDeploymentOptionsWorkerMachineArrayOutput)
}

// Available worker queue options
func (o GetDeploymentOptionsResultOutput) WorkerQueues() GetDeploymentOptionsWorkerQueuesOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) GetDeploymentOptionsWorkerQueues { return v.WorkerQueues }).(GetDeploymentOptionsWorkerQueuesOutput)
}

// Available workload identity options
func (o GetDeploymentOptionsResultOutput) WorkloadIdentityOptions() GetDeploymentOptionsWorkloadIdentityOptionArrayOutput {
	return o.ApplyT(func(v GetDeploymentOptionsResult) []GetDeploymentOptionsWorkloadIdentityOption {
		return v.WorkloadIdentityOptions
	}).(GetDeploymentOptionsWorkloadIdentityOptionArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeploymentOptionsResultOutput{})
}
