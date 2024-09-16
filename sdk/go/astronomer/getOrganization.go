// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astronomer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

// Organization data source
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
//			exampleOrganization, err := astronomer.GetOrganization(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("organization", exampleOrganization)
//			return nil
//		})
//	}
//
// ```
func GetOrganization(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOrganizationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationResult
	err := ctx.Invoke("astronomer:index/getOrganization:getOrganization", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOrganization.
type GetOrganizationResult struct {
	// Organization billing email
	BillingEmail string `pulumi:"billingEmail"`
	// Organization creation timestamp
	CreatedAt string `pulumi:"createdAt"`
	// Organization creator
	CreatedBy GetOrganizationCreatedBy `pulumi:"createdBy"`
	// Organization identifier
	Id string `pulumi:"id"`
	// Whether SCIM is enabled for the organization
	IsScimEnabled bool `pulumi:"isScimEnabled"`
	// Organization name
	Name string `pulumi:"name"`
	// Organization payment method
	PaymentMethod string `pulumi:"paymentMethod"`
	// Organization product type
	Product string `pulumi:"product"`
	// Organization status
	Status string `pulumi:"status"`
	// Organization support plan
	SupportPlan string `pulumi:"supportPlan"`
	// Organization trial expiration timestamp
	TrialExpiresAt string `pulumi:"trialExpiresAt"`
	// Organization last updated timestamp
	UpdatedAt string `pulumi:"updatedAt"`
	// Organization updater
	UpdatedBy GetOrganizationUpdatedBy `pulumi:"updatedBy"`
}

func GetOrganizationOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetOrganizationResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetOrganizationResult, error) {
		r, err := GetOrganization(ctx, opts...)
		var s GetOrganizationResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetOrganizationResultOutput)
}

// A collection of values returned by getOrganization.
type GetOrganizationResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationResult)(nil)).Elem()
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutput() GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutputWithContext(ctx context.Context) GetOrganizationResultOutput {
	return o
}

// Organization billing email
func (o GetOrganizationResultOutput) BillingEmail() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.BillingEmail }).(pulumi.StringOutput)
}

// Organization creation timestamp
func (o GetOrganizationResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Organization creator
func (o GetOrganizationResultOutput) CreatedBy() GetOrganizationCreatedByOutput {
	return o.ApplyT(func(v GetOrganizationResult) GetOrganizationCreatedBy { return v.CreatedBy }).(GetOrganizationCreatedByOutput)
}

// Organization identifier
func (o GetOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether SCIM is enabled for the organization
func (o GetOrganizationResultOutput) IsScimEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.IsScimEnabled }).(pulumi.BoolOutput)
}

// Organization name
func (o GetOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Organization payment method
func (o GetOrganizationResultOutput) PaymentMethod() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.PaymentMethod }).(pulumi.StringOutput)
}

// Organization product type
func (o GetOrganizationResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Product }).(pulumi.StringOutput)
}

// Organization status
func (o GetOrganizationResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Status }).(pulumi.StringOutput)
}

// Organization support plan
func (o GetOrganizationResultOutput) SupportPlan() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.SupportPlan }).(pulumi.StringOutput)
}

// Organization trial expiration timestamp
func (o GetOrganizationResultOutput) TrialExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.TrialExpiresAt }).(pulumi.StringOutput)
}

// Organization last updated timestamp
func (o GetOrganizationResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Organization updater
func (o GetOrganizationResultOutput) UpdatedBy() GetOrganizationUpdatedByOutput {
	return o.ApplyT(func(v GetOrganizationResult) GetOrganizationUpdatedBy { return v.UpdatedBy }).(GetOrganizationUpdatedByOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationResultOutput{})
}
