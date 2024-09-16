// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astronomer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

// API Token data source
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
//			exampleApiToken, err := astronomer.LookupApiToken(ctx, &astronomer.LookupApiTokenArgs{
//				Id: "clxm4836f00ql01me3nigmcr6",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("apiToken", exampleApiToken)
//			return nil
//		})
//	}
//
// ```
func LookupApiToken(ctx *pulumi.Context, args *LookupApiTokenArgs, opts ...pulumi.InvokeOption) (*LookupApiTokenResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiTokenResult
	err := ctx.Invoke("astronomer:index/getApiToken:getApiToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApiToken.
type LookupApiTokenArgs struct {
	// API Token identifier
	Id string `pulumi:"id"`
}

// A collection of values returned by getApiToken.
type LookupApiTokenResult struct {
	// API Token creation timestamp
	CreatedAt string `pulumi:"createdAt"`
	// API Token creator
	CreatedBy GetApiTokenCreatedBy `pulumi:"createdBy"`
	// API Token description
	Description string `pulumi:"description"`
	// time when the API token will expire in UTC
	EndAt string `pulumi:"endAt"`
	// API Token expiry period in days
	ExpiryPeriodInDays int `pulumi:"expiryPeriodInDays"`
	// API Token identifier
	Id string `pulumi:"id"`
	// API Token last used timestamp
	LastUsedAt string `pulumi:"lastUsedAt"`
	// API Token name
	Name string `pulumi:"name"`
	// The roles assigned to the API Token
	Roles []GetApiTokenRole `pulumi:"roles"`
	// API Token short token
	ShortToken string `pulumi:"shortToken"`
	// time when the API token will become valid in UTC
	StartAt string `pulumi:"startAt"`
	// API Token type
	Type string `pulumi:"type"`
	// API Token last updated timestamp
	UpdatedAt string `pulumi:"updatedAt"`
	// API Token updater
	UpdatedBy GetApiTokenUpdatedBy `pulumi:"updatedBy"`
}

func LookupApiTokenOutput(ctx *pulumi.Context, args LookupApiTokenOutputArgs, opts ...pulumi.InvokeOption) LookupApiTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiTokenResult, error) {
			args := v.(LookupApiTokenArgs)
			r, err := LookupApiToken(ctx, &args, opts...)
			var s LookupApiTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiTokenResultOutput)
}

// A collection of arguments for invoking getApiToken.
type LookupApiTokenOutputArgs struct {
	// API Token identifier
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupApiTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiTokenArgs)(nil)).Elem()
}

// A collection of values returned by getApiToken.
type LookupApiTokenResultOutput struct{ *pulumi.OutputState }

func (LookupApiTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiTokenResult)(nil)).Elem()
}

func (o LookupApiTokenResultOutput) ToLookupApiTokenResultOutput() LookupApiTokenResultOutput {
	return o
}

func (o LookupApiTokenResultOutput) ToLookupApiTokenResultOutputWithContext(ctx context.Context) LookupApiTokenResultOutput {
	return o
}

// API Token creation timestamp
func (o LookupApiTokenResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// API Token creator
func (o LookupApiTokenResultOutput) CreatedBy() GetApiTokenCreatedByOutput {
	return o.ApplyT(func(v LookupApiTokenResult) GetApiTokenCreatedBy { return v.CreatedBy }).(GetApiTokenCreatedByOutput)
}

// API Token description
func (o LookupApiTokenResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.Description }).(pulumi.StringOutput)
}

// time when the API token will expire in UTC
func (o LookupApiTokenResultOutput) EndAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.EndAt }).(pulumi.StringOutput)
}

// API Token expiry period in days
func (o LookupApiTokenResultOutput) ExpiryPeriodInDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupApiTokenResult) int { return v.ExpiryPeriodInDays }).(pulumi.IntOutput)
}

// API Token identifier
func (o LookupApiTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

// API Token last used timestamp
func (o LookupApiTokenResultOutput) LastUsedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.LastUsedAt }).(pulumi.StringOutput)
}

// API Token name
func (o LookupApiTokenResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.Name }).(pulumi.StringOutput)
}

// The roles assigned to the API Token
func (o LookupApiTokenResultOutput) Roles() GetApiTokenRoleArrayOutput {
	return o.ApplyT(func(v LookupApiTokenResult) []GetApiTokenRole { return v.Roles }).(GetApiTokenRoleArrayOutput)
}

// API Token short token
func (o LookupApiTokenResultOutput) ShortToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.ShortToken }).(pulumi.StringOutput)
}

// time when the API token will become valid in UTC
func (o LookupApiTokenResultOutput) StartAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.StartAt }).(pulumi.StringOutput)
}

// API Token type
func (o LookupApiTokenResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.Type }).(pulumi.StringOutput)
}

// API Token last updated timestamp
func (o LookupApiTokenResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTokenResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// API Token updater
func (o LookupApiTokenResultOutput) UpdatedBy() GetApiTokenUpdatedByOutput {
	return o.ApplyT(func(v LookupApiTokenResult) GetApiTokenUpdatedBy { return v.UpdatedBy }).(GetApiTokenUpdatedByOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiTokenResultOutput{})
}
