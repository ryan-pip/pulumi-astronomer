// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astronomer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

// ClusterOptions data source
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
//			exampleClusterOptions, err := astronomer.GetClusterOptions(ctx, &astronomer.GetClusterOptionsArgs{
//				Type: "HYBRID",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = astronomer.GetClusterOptions(ctx, &astronomer.GetClusterOptionsArgs{
//				Type:          "HYBRID",
//				CloudProvider: pulumi.StringRef("AWS"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("clusterOptions", exampleClusterOptions)
//			return nil
//		})
//	}
//
// ```
func GetClusterOptions(ctx *pulumi.Context, args *GetClusterOptionsArgs, opts ...pulumi.InvokeOption) (*GetClusterOptionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetClusterOptionsResult
	err := ctx.Invoke("astronomer:index/getClusterOptions:getClusterOptions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterOptions.
type GetClusterOptionsArgs struct {
	CloudProvider *string `pulumi:"cloudProvider"`
	Type          string  `pulumi:"type"`
}

// A collection of values returned by getClusterOptions.
type GetClusterOptionsResult struct {
	CloudProvider  *string                          `pulumi:"cloudProvider"`
	ClusterOptions []GetClusterOptionsClusterOption `pulumi:"clusterOptions"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Type string `pulumi:"type"`
}

func GetClusterOptionsOutput(ctx *pulumi.Context, args GetClusterOptionsOutputArgs, opts ...pulumi.InvokeOption) GetClusterOptionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetClusterOptionsResultOutput, error) {
			args := v.(GetClusterOptionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("astronomer:index/getClusterOptions:getClusterOptions", args, GetClusterOptionsResultOutput{}, options).(GetClusterOptionsResultOutput), nil
		}).(GetClusterOptionsResultOutput)
}

// A collection of arguments for invoking getClusterOptions.
type GetClusterOptionsOutputArgs struct {
	CloudProvider pulumi.StringPtrInput `pulumi:"cloudProvider"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (GetClusterOptionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterOptionsArgs)(nil)).Elem()
}

// A collection of values returned by getClusterOptions.
type GetClusterOptionsResultOutput struct{ *pulumi.OutputState }

func (GetClusterOptionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterOptionsResult)(nil)).Elem()
}

func (o GetClusterOptionsResultOutput) ToGetClusterOptionsResultOutput() GetClusterOptionsResultOutput {
	return o
}

func (o GetClusterOptionsResultOutput) ToGetClusterOptionsResultOutputWithContext(ctx context.Context) GetClusterOptionsResultOutput {
	return o
}

func (o GetClusterOptionsResultOutput) CloudProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterOptionsResult) *string { return v.CloudProvider }).(pulumi.StringPtrOutput)
}

func (o GetClusterOptionsResultOutput) ClusterOptions() GetClusterOptionsClusterOptionArrayOutput {
	return o.ApplyT(func(v GetClusterOptionsResult) []GetClusterOptionsClusterOption { return v.ClusterOptions }).(GetClusterOptionsClusterOptionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClusterOptionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterOptionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetClusterOptionsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterOptionsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterOptionsResultOutput{})
}
