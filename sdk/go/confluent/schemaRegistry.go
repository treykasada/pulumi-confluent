// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package confluent

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SchemaRegistry struct {
	pulumi.CustomResourceState

	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Environment ID
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// where
	Region pulumi.StringOutput `pulumi:"region"`
	// Cloud provider
	ServiceProvider pulumi.StringOutput `pulumi:"serviceProvider"`
}

// NewSchemaRegistry registers a new resource with the given unique name, arguments, and options.
func NewSchemaRegistry(ctx *pulumi.Context,
	name string, args *SchemaRegistryArgs, opts ...pulumi.ResourceOption) (*SchemaRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.ServiceProvider == nil {
		return nil, errors.New("invalid value for required argument 'ServiceProvider'")
	}
	var resource SchemaRegistry
	err := ctx.RegisterResource("confluent:index/schemaRegistry:SchemaRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchemaRegistry gets an existing SchemaRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchemaRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaRegistryState, opts ...pulumi.ResourceOption) (*SchemaRegistry, error) {
	var resource SchemaRegistry
	err := ctx.ReadResource("confluent:index/schemaRegistry:SchemaRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SchemaRegistry resources.
type schemaRegistryState struct {
	Endpoint *string `pulumi:"endpoint"`
	// Environment ID
	EnvironmentId *string `pulumi:"environmentId"`
	// where
	Region *string `pulumi:"region"`
	// Cloud provider
	ServiceProvider *string `pulumi:"serviceProvider"`
}

type SchemaRegistryState struct {
	Endpoint pulumi.StringPtrInput
	// Environment ID
	EnvironmentId pulumi.StringPtrInput
	// where
	Region pulumi.StringPtrInput
	// Cloud provider
	ServiceProvider pulumi.StringPtrInput
}

func (SchemaRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaRegistryState)(nil)).Elem()
}

type schemaRegistryArgs struct {
	// Environment ID
	EnvironmentId string `pulumi:"environmentId"`
	// where
	Region string `pulumi:"region"`
	// Cloud provider
	ServiceProvider string `pulumi:"serviceProvider"`
}

// The set of arguments for constructing a SchemaRegistry resource.
type SchemaRegistryArgs struct {
	// Environment ID
	EnvironmentId pulumi.StringInput
	// where
	Region pulumi.StringInput
	// Cloud provider
	ServiceProvider pulumi.StringInput
}

func (SchemaRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaRegistryArgs)(nil)).Elem()
}

type SchemaRegistryInput interface {
	pulumi.Input

	ToSchemaRegistryOutput() SchemaRegistryOutput
	ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput
}

func (*SchemaRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaRegistry)(nil))
}

func (i *SchemaRegistry) ToSchemaRegistryOutput() SchemaRegistryOutput {
	return i.ToSchemaRegistryOutputWithContext(context.Background())
}

func (i *SchemaRegistry) ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaRegistryOutput)
}

func (i *SchemaRegistry) ToSchemaRegistryPtrOutput() SchemaRegistryPtrOutput {
	return i.ToSchemaRegistryPtrOutputWithContext(context.Background())
}

func (i *SchemaRegistry) ToSchemaRegistryPtrOutputWithContext(ctx context.Context) SchemaRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaRegistryPtrOutput)
}

type SchemaRegistryPtrInput interface {
	pulumi.Input

	ToSchemaRegistryPtrOutput() SchemaRegistryPtrOutput
	ToSchemaRegistryPtrOutputWithContext(ctx context.Context) SchemaRegistryPtrOutput
}

type schemaRegistryPtrType SchemaRegistryArgs

func (*schemaRegistryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaRegistry)(nil))
}

func (i *schemaRegistryPtrType) ToSchemaRegistryPtrOutput() SchemaRegistryPtrOutput {
	return i.ToSchemaRegistryPtrOutputWithContext(context.Background())
}

func (i *schemaRegistryPtrType) ToSchemaRegistryPtrOutputWithContext(ctx context.Context) SchemaRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaRegistryPtrOutput)
}

// SchemaRegistryArrayInput is an input type that accepts SchemaRegistryArray and SchemaRegistryArrayOutput values.
// You can construct a concrete instance of `SchemaRegistryArrayInput` via:
//
//          SchemaRegistryArray{ SchemaRegistryArgs{...} }
type SchemaRegistryArrayInput interface {
	pulumi.Input

	ToSchemaRegistryArrayOutput() SchemaRegistryArrayOutput
	ToSchemaRegistryArrayOutputWithContext(context.Context) SchemaRegistryArrayOutput
}

type SchemaRegistryArray []SchemaRegistryInput

func (SchemaRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SchemaRegistry)(nil))
}

func (i SchemaRegistryArray) ToSchemaRegistryArrayOutput() SchemaRegistryArrayOutput {
	return i.ToSchemaRegistryArrayOutputWithContext(context.Background())
}

func (i SchemaRegistryArray) ToSchemaRegistryArrayOutputWithContext(ctx context.Context) SchemaRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaRegistryArrayOutput)
}

// SchemaRegistryMapInput is an input type that accepts SchemaRegistryMap and SchemaRegistryMapOutput values.
// You can construct a concrete instance of `SchemaRegistryMapInput` via:
//
//          SchemaRegistryMap{ "key": SchemaRegistryArgs{...} }
type SchemaRegistryMapInput interface {
	pulumi.Input

	ToSchemaRegistryMapOutput() SchemaRegistryMapOutput
	ToSchemaRegistryMapOutputWithContext(context.Context) SchemaRegistryMapOutput
}

type SchemaRegistryMap map[string]SchemaRegistryInput

func (SchemaRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SchemaRegistry)(nil))
}

func (i SchemaRegistryMap) ToSchemaRegistryMapOutput() SchemaRegistryMapOutput {
	return i.ToSchemaRegistryMapOutputWithContext(context.Background())
}

func (i SchemaRegistryMap) ToSchemaRegistryMapOutputWithContext(ctx context.Context) SchemaRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaRegistryMapOutput)
}

type SchemaRegistryOutput struct {
	*pulumi.OutputState
}

func (SchemaRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaRegistry)(nil))
}

func (o SchemaRegistryOutput) ToSchemaRegistryOutput() SchemaRegistryOutput {
	return o
}

func (o SchemaRegistryOutput) ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput {
	return o
}

func (o SchemaRegistryOutput) ToSchemaRegistryPtrOutput() SchemaRegistryPtrOutput {
	return o.ToSchemaRegistryPtrOutputWithContext(context.Background())
}

func (o SchemaRegistryOutput) ToSchemaRegistryPtrOutputWithContext(ctx context.Context) SchemaRegistryPtrOutput {
	return o.ApplyT(func(v SchemaRegistry) *SchemaRegistry {
		return &v
	}).(SchemaRegistryPtrOutput)
}

type SchemaRegistryPtrOutput struct {
	*pulumi.OutputState
}

func (SchemaRegistryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaRegistry)(nil))
}

func (o SchemaRegistryPtrOutput) ToSchemaRegistryPtrOutput() SchemaRegistryPtrOutput {
	return o
}

func (o SchemaRegistryPtrOutput) ToSchemaRegistryPtrOutputWithContext(ctx context.Context) SchemaRegistryPtrOutput {
	return o
}

type SchemaRegistryArrayOutput struct{ *pulumi.OutputState }

func (SchemaRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SchemaRegistry)(nil))
}

func (o SchemaRegistryArrayOutput) ToSchemaRegistryArrayOutput() SchemaRegistryArrayOutput {
	return o
}

func (o SchemaRegistryArrayOutput) ToSchemaRegistryArrayOutputWithContext(ctx context.Context) SchemaRegistryArrayOutput {
	return o
}

func (o SchemaRegistryArrayOutput) Index(i pulumi.IntInput) SchemaRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SchemaRegistry {
		return vs[0].([]SchemaRegistry)[vs[1].(int)]
	}).(SchemaRegistryOutput)
}

type SchemaRegistryMapOutput struct{ *pulumi.OutputState }

func (SchemaRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SchemaRegistry)(nil))
}

func (o SchemaRegistryMapOutput) ToSchemaRegistryMapOutput() SchemaRegistryMapOutput {
	return o
}

func (o SchemaRegistryMapOutput) ToSchemaRegistryMapOutputWithContext(ctx context.Context) SchemaRegistryMapOutput {
	return o
}

func (o SchemaRegistryMapOutput) MapIndex(k pulumi.StringInput) SchemaRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SchemaRegistry {
		return vs[0].(map[string]SchemaRegistry)[vs[1].(string)]
	}).(SchemaRegistryOutput)
}

func init() {
	pulumi.RegisterOutputType(SchemaRegistryOutput{})
	pulumi.RegisterOutputType(SchemaRegistryPtrOutput{})
	pulumi.RegisterOutputType(SchemaRegistryArrayOutput{})
	pulumi.RegisterOutputType(SchemaRegistryMapOutput{})
}
