// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astronomer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

// Cluster data source
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
//			_, err := astronomer.LookupCluster(ctx, &astronomer.LookupClusterArgs{
//				Id: "clozc036j01to01jrlgvueo8t",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("astronomer:index/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// Cluster identifier
	Id string `pulumi:"id"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// Cluster cloud provider
	CloudProvider string `pulumi:"cloudProvider"`
	// Cluster creation timestamp
	CreatedAt string `pulumi:"createdAt"`
	// Cluster database instance type
	DbInstanceType string `pulumi:"dbInstanceType"`
	// Cluster identifier
	Id string `pulumi:"id"`
	// Whether the cluster is limited
	IsLimited bool `pulumi:"isLimited"`
	// Cluster metadata
	Metadata GetClusterMetadata `pulumi:"metadata"`
	// Cluster name
	Name string `pulumi:"name"`
	// Cluster node pools
	NodePools []GetClusterNodePool `pulumi:"nodePools"`
	// Cluster pod subnet range
	PodSubnetRange string `pulumi:"podSubnetRange"`
	// Cluster provider account
	ProviderAccount string `pulumi:"providerAccount"`
	// Cluster region
	Region string `pulumi:"region"`
	// Cluster service peering range
	ServicePeeringRange string `pulumi:"servicePeeringRange"`
	// Cluster service subnet range
	ServiceSubnetRange string `pulumi:"serviceSubnetRange"`
	// Cluster status
	Status string `pulumi:"status"`
	// Cluster tags
	Tags []GetClusterTag `pulumi:"tags"`
	// Cluster tenant ID
	TenantId string `pulumi:"tenantId"`
	// Cluster type
	Type string `pulumi:"type"`
	// Cluster last updated timestamp
	UpdatedAt string `pulumi:"updatedAt"`
	// Cluster VPC subnet range
	VpcSubnetRange string `pulumi:"vpcSubnetRange"`
	// Cluster workspace IDs
	WorkspaceIds []string `pulumi:"workspaceIds"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

// A collection of arguments for invoking getCluster.
type LookupClusterOutputArgs struct {
	// Cluster identifier
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

// Cluster cloud provider
func (o LookupClusterResultOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CloudProvider }).(pulumi.StringOutput)
}

// Cluster creation timestamp
func (o LookupClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Cluster database instance type
func (o LookupClusterResultOutput) DbInstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DbInstanceType }).(pulumi.StringOutput)
}

// Cluster identifier
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether the cluster is limited
func (o LookupClusterResultOutput) IsLimited() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClusterResult) bool { return v.IsLimited }).(pulumi.BoolOutput)
}

// Cluster metadata
func (o LookupClusterResultOutput) Metadata() GetClusterMetadataOutput {
	return o.ApplyT(func(v LookupClusterResult) GetClusterMetadata { return v.Metadata }).(GetClusterMetadataOutput)
}

// Cluster name
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Cluster node pools
func (o LookupClusterResultOutput) NodePools() GetClusterNodePoolArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterNodePool { return v.NodePools }).(GetClusterNodePoolArrayOutput)
}

// Cluster pod subnet range
func (o LookupClusterResultOutput) PodSubnetRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.PodSubnetRange }).(pulumi.StringOutput)
}

// Cluster provider account
func (o LookupClusterResultOutput) ProviderAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProviderAccount }).(pulumi.StringOutput)
}

// Cluster region
func (o LookupClusterResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Region }).(pulumi.StringOutput)
}

// Cluster service peering range
func (o LookupClusterResultOutput) ServicePeeringRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ServicePeeringRange }).(pulumi.StringOutput)
}

// Cluster service subnet range
func (o LookupClusterResultOutput) ServiceSubnetRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ServiceSubnetRange }).(pulumi.StringOutput)
}

// Cluster status
func (o LookupClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// Cluster tags
func (o LookupClusterResultOutput) Tags() GetClusterTagArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterTag { return v.Tags }).(GetClusterTagArrayOutput)
}

// Cluster tenant ID
func (o LookupClusterResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Cluster type
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// Cluster last updated timestamp
func (o LookupClusterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Cluster VPC subnet range
func (o LookupClusterResultOutput) VpcSubnetRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.VpcSubnetRange }).(pulumi.StringOutput)
}

// Cluster workspace IDs
func (o LookupClusterResultOutput) WorkspaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.WorkspaceIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
