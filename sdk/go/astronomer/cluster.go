// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astronomer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

// A cluster within an organization. An Astro cluster is a Kubernetes cluster that hosts the infrastructure required to run Deployments.
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
//			dedicated, err := astronomer.NewWorkspace(ctx, "dedicated", &astronomer.WorkspaceArgs{
//				CicdEnforcedDefault: pulumi.Bool(true),
//				Description:         pulumi.String("Workspace that demos a dedicated deployment set up"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = astronomer.NewCluster(ctx, "awsDedicated", &astronomer.ClusterArgs{
//				CloudProvider:  pulumi.String("AWS"),
//				Region:         pulumi.String("us-east-1"),
//				Type:           pulumi.String("DEDICATED"),
//				VpcSubnetRange: pulumi.String("172.20.0.0/20"),
//				K8sTags:        astronomer.ClusterK8sTagArray{},
//				NodePools:      astronomer.ClusterNodePoolArray{},
//				WorkspaceIds: pulumi.StringArray{
//					dedicated.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The cluster's cloud provider.
	CloudProvider pulumi.StringOutput `pulumi:"cloudProvider"`
	// The type of database instance that is used for the cluster. Required for Hybrid clusters.
	DbInstanceType pulumi.StringOutput `pulumi:"dbInstanceType"`
	// Whether the cluster is limited.
	IsLimited pulumi.BoolOutput `pulumi:"isLimited"`
	// The Kubernetes tags in the cluster.
	K8sTags ClusterK8sTagArrayOutput `pulumi:"k8sTags"`
	// The cluster's metadata.
	Metadata ClusterMetadataOutput `pulumi:"metadata"`
	// The name of the node pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of node pools to create in the cluster.
	NodePools ClusterNodePoolArrayOutput `pulumi:"nodePools"`
	// The organization this cluster is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The subnet range for Pods. For GCP clusters only.
	PodSubnetRange pulumi.StringPtrOutput `pulumi:"podSubnetRange"`
	// The provider account ID. Required for Hybrid clusters.
	ProviderAccount pulumi.StringOutput `pulumi:"providerAccount"`
	// The cluster's region.
	Region pulumi.StringOutput `pulumi:"region"`
	// The service peering range. For GCP clusters only.
	ServicePeeringRange pulumi.StringPtrOutput `pulumi:"servicePeeringRange"`
	// The service subnet range. For GCP clusters only.
	ServiceSubnetRange pulumi.StringPtrOutput `pulumi:"serviceSubnetRange"`
	// The tenant ID. For Azure clusters only.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The cluster's type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VPC subnet range.
	VpcSubnetRange pulumi.StringOutput `pulumi:"vpcSubnetRange"`
	// The list of Workspaces that are authorized to the cluster.
	WorkspaceIds pulumi.StringArrayOutput `pulumi:"workspaceIds"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudProvider == nil {
		return nil, errors.New("invalid value for required argument 'CloudProvider'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.VpcSubnetRange == nil {
		return nil, errors.New("invalid value for required argument 'VpcSubnetRange'")
	}
	if args.WorkspaceIds == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("astronomer:index/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("astronomer:index/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The cluster's cloud provider.
	CloudProvider *string `pulumi:"cloudProvider"`
	// The type of database instance that is used for the cluster. Required for Hybrid clusters.
	DbInstanceType *string `pulumi:"dbInstanceType"`
	// Whether the cluster is limited.
	IsLimited *bool `pulumi:"isLimited"`
	// The Kubernetes tags in the cluster.
	K8sTags []ClusterK8sTag `pulumi:"k8sTags"`
	// The cluster's metadata.
	Metadata *ClusterMetadata `pulumi:"metadata"`
	// The name of the node pool.
	Name *string `pulumi:"name"`
	// The list of node pools to create in the cluster.
	NodePools []ClusterNodePool `pulumi:"nodePools"`
	// The organization this cluster is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The subnet range for Pods. For GCP clusters only.
	PodSubnetRange *string `pulumi:"podSubnetRange"`
	// The provider account ID. Required for Hybrid clusters.
	ProviderAccount *string `pulumi:"providerAccount"`
	// The cluster's region.
	Region *string `pulumi:"region"`
	// The service peering range. For GCP clusters only.
	ServicePeeringRange *string `pulumi:"servicePeeringRange"`
	// The service subnet range. For GCP clusters only.
	ServiceSubnetRange *string `pulumi:"serviceSubnetRange"`
	// The tenant ID. For Azure clusters only.
	TenantId *string `pulumi:"tenantId"`
	// The cluster's type.
	Type *string `pulumi:"type"`
	// The VPC subnet range.
	VpcSubnetRange *string `pulumi:"vpcSubnetRange"`
	// The list of Workspaces that are authorized to the cluster.
	WorkspaceIds []string `pulumi:"workspaceIds"`
}

type ClusterState struct {
	// The cluster's cloud provider.
	CloudProvider pulumi.StringPtrInput
	// The type of database instance that is used for the cluster. Required for Hybrid clusters.
	DbInstanceType pulumi.StringPtrInput
	// Whether the cluster is limited.
	IsLimited pulumi.BoolPtrInput
	// The Kubernetes tags in the cluster.
	K8sTags ClusterK8sTagArrayInput
	// The cluster's metadata.
	Metadata ClusterMetadataPtrInput
	// The name of the node pool.
	Name pulumi.StringPtrInput
	// The list of node pools to create in the cluster.
	NodePools ClusterNodePoolArrayInput
	// The organization this cluster is associated with.
	OrganizationId pulumi.StringPtrInput
	// The subnet range for Pods. For GCP clusters only.
	PodSubnetRange pulumi.StringPtrInput
	// The provider account ID. Required for Hybrid clusters.
	ProviderAccount pulumi.StringPtrInput
	// The cluster's region.
	Region pulumi.StringPtrInput
	// The service peering range. For GCP clusters only.
	ServicePeeringRange pulumi.StringPtrInput
	// The service subnet range. For GCP clusters only.
	ServiceSubnetRange pulumi.StringPtrInput
	// The tenant ID. For Azure clusters only.
	TenantId pulumi.StringPtrInput
	// The cluster's type.
	Type pulumi.StringPtrInput
	// The VPC subnet range.
	VpcSubnetRange pulumi.StringPtrInput
	// The list of Workspaces that are authorized to the cluster.
	WorkspaceIds pulumi.StringArrayInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The cluster's cloud provider.
	CloudProvider string `pulumi:"cloudProvider"`
	// The type of database instance that is used for the cluster. Required for Hybrid clusters.
	DbInstanceType *string `pulumi:"dbInstanceType"`
	// The Kubernetes tags in the cluster.
	K8sTags []ClusterK8sTag `pulumi:"k8sTags"`
	// The name of the node pool.
	Name *string `pulumi:"name"`
	// The list of node pools to create in the cluster.
	NodePools []ClusterNodePool `pulumi:"nodePools"`
	// The subnet range for Pods. For GCP clusters only.
	PodSubnetRange *string `pulumi:"podSubnetRange"`
	// The provider account ID. Required for Hybrid clusters.
	ProviderAccount *string `pulumi:"providerAccount"`
	// The cluster's region.
	Region string `pulumi:"region"`
	// The service peering range. For GCP clusters only.
	ServicePeeringRange *string `pulumi:"servicePeeringRange"`
	// The service subnet range. For GCP clusters only.
	ServiceSubnetRange *string `pulumi:"serviceSubnetRange"`
	// The tenant ID. For Azure clusters only.
	TenantId *string `pulumi:"tenantId"`
	// The cluster's type.
	Type string `pulumi:"type"`
	// The VPC subnet range.
	VpcSubnetRange string `pulumi:"vpcSubnetRange"`
	// The list of Workspaces that are authorized to the cluster.
	WorkspaceIds []string `pulumi:"workspaceIds"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The cluster's cloud provider.
	CloudProvider pulumi.StringInput
	// The type of database instance that is used for the cluster. Required for Hybrid clusters.
	DbInstanceType pulumi.StringPtrInput
	// The Kubernetes tags in the cluster.
	K8sTags ClusterK8sTagArrayInput
	// The name of the node pool.
	Name pulumi.StringPtrInput
	// The list of node pools to create in the cluster.
	NodePools ClusterNodePoolArrayInput
	// The subnet range for Pods. For GCP clusters only.
	PodSubnetRange pulumi.StringPtrInput
	// The provider account ID. Required for Hybrid clusters.
	ProviderAccount pulumi.StringPtrInput
	// The cluster's region.
	Region pulumi.StringInput
	// The service peering range. For GCP clusters only.
	ServicePeeringRange pulumi.StringPtrInput
	// The service subnet range. For GCP clusters only.
	ServiceSubnetRange pulumi.StringPtrInput
	// The tenant ID. For Azure clusters only.
	TenantId pulumi.StringPtrInput
	// The cluster's type.
	Type pulumi.StringInput
	// The VPC subnet range.
	VpcSubnetRange pulumi.StringInput
	// The list of Workspaces that are authorized to the cluster.
	WorkspaceIds pulumi.StringArrayInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//	ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//	ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// The cluster's cloud provider.
func (o ClusterOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CloudProvider }).(pulumi.StringOutput)
}

// The type of database instance that is used for the cluster. Required for Hybrid clusters.
func (o ClusterOutput) DbInstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DbInstanceType }).(pulumi.StringOutput)
}

// Whether the cluster is limited.
func (o ClusterOutput) IsLimited() pulumi.BoolOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolOutput { return v.IsLimited }).(pulumi.BoolOutput)
}

// The Kubernetes tags in the cluster.
func (o ClusterOutput) K8sTags() ClusterK8sTagArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterK8sTagArrayOutput { return v.K8sTags }).(ClusterK8sTagArrayOutput)
}

// The cluster's metadata.
func (o ClusterOutput) Metadata() ClusterMetadataOutput {
	return o.ApplyT(func(v *Cluster) ClusterMetadataOutput { return v.Metadata }).(ClusterMetadataOutput)
}

// The name of the node pool.
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of node pools to create in the cluster.
func (o ClusterOutput) NodePools() ClusterNodePoolArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterNodePoolArrayOutput { return v.NodePools }).(ClusterNodePoolArrayOutput)
}

// The organization this cluster is associated with.
func (o ClusterOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The subnet range for Pods. For GCP clusters only.
func (o ClusterOutput) PodSubnetRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.PodSubnetRange }).(pulumi.StringPtrOutput)
}

// The provider account ID. Required for Hybrid clusters.
func (o ClusterOutput) ProviderAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProviderAccount }).(pulumi.StringOutput)
}

// The cluster's region.
func (o ClusterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The service peering range. For GCP clusters only.
func (o ClusterOutput) ServicePeeringRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ServicePeeringRange }).(pulumi.StringPtrOutput)
}

// The service subnet range. For GCP clusters only.
func (o ClusterOutput) ServiceSubnetRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ServiceSubnetRange }).(pulumi.StringPtrOutput)
}

// The tenant ID. For Azure clusters only.
func (o ClusterOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The cluster's type.
func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The VPC subnet range.
func (o ClusterOutput) VpcSubnetRange() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.VpcSubnetRange }).(pulumi.StringOutput)
}

// The list of Workspaces that are authorized to the cluster.
func (o ClusterOutput) WorkspaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.WorkspaceIds }).(pulumi.StringArrayOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
