// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astronomer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

// Astronomer Workspace Resource
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
//			_, err := astronomer.LookupWorkspace(ctx, &astronomer.LookupWorkspaceArgs{
//				Id: "cabcabcabcabcabcabcabcabcabc",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceResult
	err := ctx.Invoke("astronomer:index/getWorkspace:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkspace.
type LookupWorkspaceArgs struct {
	// The Workspace's identifier.
	Id string `pulumi:"id"`
}

// A collection of values returned by getWorkspace.
type LookupWorkspaceResult struct {
	// Whether new Deployments enforce CI/CD deploys by default.
	CicdEnforcedDefault bool `pulumi:"cicdEnforcedDefault"`
	// The Workspace's description
	Description string `pulumi:"description"`
	// The Workspace's identifier.
	Id string `pulumi:"id"`
	// The Workspace's name
	Name string `pulumi:"name"`
}

func LookupWorkspaceOutput(ctx *pulumi.Context, args LookupWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceResult, error) {
			args := v.(LookupWorkspaceArgs)
			r, err := LookupWorkspace(ctx, &args, opts...)
			var s LookupWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceResultOutput)
}

// A collection of arguments for invoking getWorkspace.
type LookupWorkspaceOutputArgs struct {
	// The Workspace's identifier.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}

// A collection of values returned by getWorkspace.
type LookupWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceResult)(nil)).Elem()
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutput() LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutputWithContext(ctx context.Context) LookupWorkspaceResultOutput {
	return o
}

// Whether new Deployments enforce CI/CD deploys by default.
func (o LookupWorkspaceResultOutput) CicdEnforcedDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) bool { return v.CicdEnforcedDefault }).(pulumi.BoolOutput)
}

// The Workspace's description
func (o LookupWorkspaceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Description }).(pulumi.StringOutput)
}

// The Workspace's identifier.
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Workspace's name
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
