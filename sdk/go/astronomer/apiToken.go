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

// API Token resource
type ApiToken struct {
	pulumi.CustomResourceState

	// API Token creation timestamp
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// API Token creator
	CreatedBy ApiTokenCreatedByOutput `pulumi:"createdBy"`
	// API Token description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// time when the API token will expire in UTC
	EndAt pulumi.StringOutput `pulumi:"endAt"`
	// API Token expiry period in days
	ExpiryPeriodInDays pulumi.IntPtrOutput `pulumi:"expiryPeriodInDays"`
	// API Token last used timestamp
	LastUsedAt pulumi.StringOutput `pulumi:"lastUsedAt"`
	// API Token name
	Name pulumi.StringOutput `pulumi:"name"`
	// The roles assigned to the API Token
	Roles ApiTokenRoleArrayOutput `pulumi:"roles"`
	// API Token short token
	ShortToken pulumi.StringOutput `pulumi:"shortToken"`
	// time when the API token will become valid in UTC
	StartAt pulumi.StringOutput `pulumi:"startAt"`
	Token   pulumi.StringOutput `pulumi:"token"`
	// API Token type - if changing this value, the API Token will be recreated with the new type
	Type pulumi.StringOutput `pulumi:"type"`
	// API Token last updated timestamp
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// API Token updater
	UpdatedBy ApiTokenUpdatedByOutput `pulumi:"updatedBy"`
}

// NewApiToken registers a new resource with the given unique name, arguments, and options.
func NewApiToken(ctx *pulumi.Context,
	name string, args *ApiTokenArgs, opts ...pulumi.ResourceOption) (*ApiToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiToken
	err := ctx.RegisterResource("astronomer:index/apiToken:ApiToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiToken gets an existing ApiToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiTokenState, opts ...pulumi.ResourceOption) (*ApiToken, error) {
	var resource ApiToken
	err := ctx.ReadResource("astronomer:index/apiToken:ApiToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiToken resources.
type apiTokenState struct {
	// API Token creation timestamp
	CreatedAt *string `pulumi:"createdAt"`
	// API Token creator
	CreatedBy *ApiTokenCreatedBy `pulumi:"createdBy"`
	// API Token description
	Description *string `pulumi:"description"`
	// time when the API token will expire in UTC
	EndAt *string `pulumi:"endAt"`
	// API Token expiry period in days
	ExpiryPeriodInDays *int `pulumi:"expiryPeriodInDays"`
	// API Token last used timestamp
	LastUsedAt *string `pulumi:"lastUsedAt"`
	// API Token name
	Name *string `pulumi:"name"`
	// The roles assigned to the API Token
	Roles []ApiTokenRole `pulumi:"roles"`
	// API Token short token
	ShortToken *string `pulumi:"shortToken"`
	// time when the API token will become valid in UTC
	StartAt *string `pulumi:"startAt"`
	Token   *string `pulumi:"token"`
	// API Token type - if changing this value, the API Token will be recreated with the new type
	Type *string `pulumi:"type"`
	// API Token last updated timestamp
	UpdatedAt *string `pulumi:"updatedAt"`
	// API Token updater
	UpdatedBy *ApiTokenUpdatedBy `pulumi:"updatedBy"`
}

type ApiTokenState struct {
	// API Token creation timestamp
	CreatedAt pulumi.StringPtrInput
	// API Token creator
	CreatedBy ApiTokenCreatedByPtrInput
	// API Token description
	Description pulumi.StringPtrInput
	// time when the API token will expire in UTC
	EndAt pulumi.StringPtrInput
	// API Token expiry period in days
	ExpiryPeriodInDays pulumi.IntPtrInput
	// API Token last used timestamp
	LastUsedAt pulumi.StringPtrInput
	// API Token name
	Name pulumi.StringPtrInput
	// The roles assigned to the API Token
	Roles ApiTokenRoleArrayInput
	// API Token short token
	ShortToken pulumi.StringPtrInput
	// time when the API token will become valid in UTC
	StartAt pulumi.StringPtrInput
	Token   pulumi.StringPtrInput
	// API Token type - if changing this value, the API Token will be recreated with the new type
	Type pulumi.StringPtrInput
	// API Token last updated timestamp
	UpdatedAt pulumi.StringPtrInput
	// API Token updater
	UpdatedBy ApiTokenUpdatedByPtrInput
}

func (ApiTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiTokenState)(nil)).Elem()
}

type apiTokenArgs struct {
	// API Token description
	Description *string `pulumi:"description"`
	// API Token expiry period in days
	ExpiryPeriodInDays *int `pulumi:"expiryPeriodInDays"`
	// API Token name
	Name *string `pulumi:"name"`
	// The roles assigned to the API Token
	Roles []ApiTokenRole `pulumi:"roles"`
	// API Token type - if changing this value, the API Token will be recreated with the new type
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ApiToken resource.
type ApiTokenArgs struct {
	// API Token description
	Description pulumi.StringPtrInput
	// API Token expiry period in days
	ExpiryPeriodInDays pulumi.IntPtrInput
	// API Token name
	Name pulumi.StringPtrInput
	// The roles assigned to the API Token
	Roles ApiTokenRoleArrayInput
	// API Token type - if changing this value, the API Token will be recreated with the new type
	Type pulumi.StringInput
}

func (ApiTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiTokenArgs)(nil)).Elem()
}

type ApiTokenInput interface {
	pulumi.Input

	ToApiTokenOutput() ApiTokenOutput
	ToApiTokenOutputWithContext(ctx context.Context) ApiTokenOutput
}

func (*ApiToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiToken)(nil)).Elem()
}

func (i *ApiToken) ToApiTokenOutput() ApiTokenOutput {
	return i.ToApiTokenOutputWithContext(context.Background())
}

func (i *ApiToken) ToApiTokenOutputWithContext(ctx context.Context) ApiTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiTokenOutput)
}

// ApiTokenArrayInput is an input type that accepts ApiTokenArray and ApiTokenArrayOutput values.
// You can construct a concrete instance of `ApiTokenArrayInput` via:
//
//	ApiTokenArray{ ApiTokenArgs{...} }
type ApiTokenArrayInput interface {
	pulumi.Input

	ToApiTokenArrayOutput() ApiTokenArrayOutput
	ToApiTokenArrayOutputWithContext(context.Context) ApiTokenArrayOutput
}

type ApiTokenArray []ApiTokenInput

func (ApiTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiToken)(nil)).Elem()
}

func (i ApiTokenArray) ToApiTokenArrayOutput() ApiTokenArrayOutput {
	return i.ToApiTokenArrayOutputWithContext(context.Background())
}

func (i ApiTokenArray) ToApiTokenArrayOutputWithContext(ctx context.Context) ApiTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiTokenArrayOutput)
}

// ApiTokenMapInput is an input type that accepts ApiTokenMap and ApiTokenMapOutput values.
// You can construct a concrete instance of `ApiTokenMapInput` via:
//
//	ApiTokenMap{ "key": ApiTokenArgs{...} }
type ApiTokenMapInput interface {
	pulumi.Input

	ToApiTokenMapOutput() ApiTokenMapOutput
	ToApiTokenMapOutputWithContext(context.Context) ApiTokenMapOutput
}

type ApiTokenMap map[string]ApiTokenInput

func (ApiTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiToken)(nil)).Elem()
}

func (i ApiTokenMap) ToApiTokenMapOutput() ApiTokenMapOutput {
	return i.ToApiTokenMapOutputWithContext(context.Background())
}

func (i ApiTokenMap) ToApiTokenMapOutputWithContext(ctx context.Context) ApiTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiTokenMapOutput)
}

type ApiTokenOutput struct{ *pulumi.OutputState }

func (ApiTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiToken)(nil)).Elem()
}

func (o ApiTokenOutput) ToApiTokenOutput() ApiTokenOutput {
	return o
}

func (o ApiTokenOutput) ToApiTokenOutputWithContext(ctx context.Context) ApiTokenOutput {
	return o
}

// API Token creation timestamp
func (o ApiTokenOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// API Token creator
func (o ApiTokenOutput) CreatedBy() ApiTokenCreatedByOutput {
	return o.ApplyT(func(v *ApiToken) ApiTokenCreatedByOutput { return v.CreatedBy }).(ApiTokenCreatedByOutput)
}

// API Token description
func (o ApiTokenOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// time when the API token will expire in UTC
func (o ApiTokenOutput) EndAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringOutput { return v.EndAt }).(pulumi.StringOutput)
}

// API Token expiry period in days
func (o ApiTokenOutput) ExpiryPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.IntPtrOutput { return v.ExpiryPeriodInDays }).(pulumi.IntPtrOutput)
}

// API Token last used timestamp
func (o ApiTokenOutput) LastUsedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringOutput { return v.LastUsedAt }).(pulumi.StringOutput)
}

// API Token name
func (o ApiTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The roles assigned to the API Token
func (o ApiTokenOutput) Roles() ApiTokenRoleArrayOutput {
	return o.ApplyT(func(v *ApiToken) ApiTokenRoleArrayOutput { return v.Roles }).(ApiTokenRoleArrayOutput)
}

// API Token short token
func (o ApiTokenOutput) ShortToken() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringOutput { return v.ShortToken }).(pulumi.StringOutput)
}

// time when the API token will become valid in UTC
func (o ApiTokenOutput) StartAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringOutput { return v.StartAt }).(pulumi.StringOutput)
}

func (o ApiTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// API Token type - if changing this value, the API Token will be recreated with the new type
func (o ApiTokenOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// API Token last updated timestamp
func (o ApiTokenOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiToken) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// API Token updater
func (o ApiTokenOutput) UpdatedBy() ApiTokenUpdatedByOutput {
	return o.ApplyT(func(v *ApiToken) ApiTokenUpdatedByOutput { return v.UpdatedBy }).(ApiTokenUpdatedByOutput)
}

type ApiTokenArrayOutput struct{ *pulumi.OutputState }

func (ApiTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiToken)(nil)).Elem()
}

func (o ApiTokenArrayOutput) ToApiTokenArrayOutput() ApiTokenArrayOutput {
	return o
}

func (o ApiTokenArrayOutput) ToApiTokenArrayOutputWithContext(ctx context.Context) ApiTokenArrayOutput {
	return o
}

func (o ApiTokenArrayOutput) Index(i pulumi.IntInput) ApiTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiToken {
		return vs[0].([]*ApiToken)[vs[1].(int)]
	}).(ApiTokenOutput)
}

type ApiTokenMapOutput struct{ *pulumi.OutputState }

func (ApiTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiToken)(nil)).Elem()
}

func (o ApiTokenMapOutput) ToApiTokenMapOutput() ApiTokenMapOutput {
	return o
}

func (o ApiTokenMapOutput) ToApiTokenMapOutputWithContext(ctx context.Context) ApiTokenMapOutput {
	return o
}

func (o ApiTokenMapOutput) MapIndex(k pulumi.StringInput) ApiTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiToken {
		return vs[0].(map[string]*ApiToken)[vs[1].(string)]
	}).(ApiTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiTokenInput)(nil)).Elem(), &ApiToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiTokenArrayInput)(nil)).Elem(), ApiTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiTokenMapInput)(nil)).Elem(), ApiTokenMap{})
	pulumi.RegisterOutputType(ApiTokenOutput{})
	pulumi.RegisterOutputType(ApiTokenArrayOutput{})
	pulumi.RegisterOutputType(ApiTokenMapOutput{})
}
