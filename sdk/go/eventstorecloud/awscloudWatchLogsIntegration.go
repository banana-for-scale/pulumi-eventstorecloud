// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventstorecloud

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AWSCloudWatchLogsIntegration struct {
	pulumi.CustomResourceState

	// The access key ID of IAM credentials which have permissions to create and write to the log group
	AccessKeyId pulumi.StringPtrOutput `pulumi:"accessKeyId"`
	// Clusters to be used with this integration
	ClusterIds pulumi.StringArrayOutput `pulumi:"clusterIds"`
	// Human readable description of the integration
	Description pulumi.StringOutput `pulumi:"description"`
	// Name of the CloudWatch group
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// ID of the project to which the integration applies
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// AWS region for group
	Region pulumi.StringOutput `pulumi:"region"`
	// The secret access key of IAM credentials which will be used to write to the log groups
	SecretAccessKey pulumi.StringPtrOutput `pulumi:"secretAccessKey"`
}

// NewAWSCloudWatchLogsIntegration registers a new resource with the given unique name, arguments, and options.
func NewAWSCloudWatchLogsIntegration(ctx *pulumi.Context,
	name string, args *AWSCloudWatchLogsIntegrationArgs, opts ...pulumi.ResourceOption) (*AWSCloudWatchLogsIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterIds == nil {
		return nil, errors.New("invalid value for required argument 'ClusterIds'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AWSCloudWatchLogsIntegration
	err := ctx.RegisterResource("eventstorecloud:index/aWSCloudWatchLogsIntegration:AWSCloudWatchLogsIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAWSCloudWatchLogsIntegration gets an existing AWSCloudWatchLogsIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAWSCloudWatchLogsIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AWSCloudWatchLogsIntegrationState, opts ...pulumi.ResourceOption) (*AWSCloudWatchLogsIntegration, error) {
	var resource AWSCloudWatchLogsIntegration
	err := ctx.ReadResource("eventstorecloud:index/aWSCloudWatchLogsIntegration:AWSCloudWatchLogsIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AWSCloudWatchLogsIntegration resources.
type awscloudWatchLogsIntegrationState struct {
	// The access key ID of IAM credentials which have permissions to create and write to the log group
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Clusters to be used with this integration
	ClusterIds []string `pulumi:"clusterIds"`
	// Human readable description of the integration
	Description *string `pulumi:"description"`
	// Name of the CloudWatch group
	GroupName *string `pulumi:"groupName"`
	// ID of the project to which the integration applies
	ProjectId *string `pulumi:"projectId"`
	// AWS region for group
	Region *string `pulumi:"region"`
	// The secret access key of IAM credentials which will be used to write to the log groups
	SecretAccessKey *string `pulumi:"secretAccessKey"`
}

type AWSCloudWatchLogsIntegrationState struct {
	// The access key ID of IAM credentials which have permissions to create and write to the log group
	AccessKeyId pulumi.StringPtrInput
	// Clusters to be used with this integration
	ClusterIds pulumi.StringArrayInput
	// Human readable description of the integration
	Description pulumi.StringPtrInput
	// Name of the CloudWatch group
	GroupName pulumi.StringPtrInput
	// ID of the project to which the integration applies
	ProjectId pulumi.StringPtrInput
	// AWS region for group
	Region pulumi.StringPtrInput
	// The secret access key of IAM credentials which will be used to write to the log groups
	SecretAccessKey pulumi.StringPtrInput
}

func (AWSCloudWatchLogsIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*awscloudWatchLogsIntegrationState)(nil)).Elem()
}

type awscloudWatchLogsIntegrationArgs struct {
	// The access key ID of IAM credentials which have permissions to create and write to the log group
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Clusters to be used with this integration
	ClusterIds []string `pulumi:"clusterIds"`
	// Human readable description of the integration
	Description string `pulumi:"description"`
	// Name of the CloudWatch group
	GroupName string `pulumi:"groupName"`
	// ID of the project to which the integration applies
	ProjectId string `pulumi:"projectId"`
	// AWS region for group
	Region string `pulumi:"region"`
	// The secret access key of IAM credentials which will be used to write to the log groups
	SecretAccessKey *string `pulumi:"secretAccessKey"`
}

// The set of arguments for constructing a AWSCloudWatchLogsIntegration resource.
type AWSCloudWatchLogsIntegrationArgs struct {
	// The access key ID of IAM credentials which have permissions to create and write to the log group
	AccessKeyId pulumi.StringPtrInput
	// Clusters to be used with this integration
	ClusterIds pulumi.StringArrayInput
	// Human readable description of the integration
	Description pulumi.StringInput
	// Name of the CloudWatch group
	GroupName pulumi.StringInput
	// ID of the project to which the integration applies
	ProjectId pulumi.StringInput
	// AWS region for group
	Region pulumi.StringInput
	// The secret access key of IAM credentials which will be used to write to the log groups
	SecretAccessKey pulumi.StringPtrInput
}

func (AWSCloudWatchLogsIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awscloudWatchLogsIntegrationArgs)(nil)).Elem()
}

type AWSCloudWatchLogsIntegrationInput interface {
	pulumi.Input

	ToAWSCloudWatchLogsIntegrationOutput() AWSCloudWatchLogsIntegrationOutput
	ToAWSCloudWatchLogsIntegrationOutputWithContext(ctx context.Context) AWSCloudWatchLogsIntegrationOutput
}

func (*AWSCloudWatchLogsIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**AWSCloudWatchLogsIntegration)(nil)).Elem()
}

func (i *AWSCloudWatchLogsIntegration) ToAWSCloudWatchLogsIntegrationOutput() AWSCloudWatchLogsIntegrationOutput {
	return i.ToAWSCloudWatchLogsIntegrationOutputWithContext(context.Background())
}

func (i *AWSCloudWatchLogsIntegration) ToAWSCloudWatchLogsIntegrationOutputWithContext(ctx context.Context) AWSCloudWatchLogsIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AWSCloudWatchLogsIntegrationOutput)
}

// AWSCloudWatchLogsIntegrationArrayInput is an input type that accepts AWSCloudWatchLogsIntegrationArray and AWSCloudWatchLogsIntegrationArrayOutput values.
// You can construct a concrete instance of `AWSCloudWatchLogsIntegrationArrayInput` via:
//
//          AWSCloudWatchLogsIntegrationArray{ AWSCloudWatchLogsIntegrationArgs{...} }
type AWSCloudWatchLogsIntegrationArrayInput interface {
	pulumi.Input

	ToAWSCloudWatchLogsIntegrationArrayOutput() AWSCloudWatchLogsIntegrationArrayOutput
	ToAWSCloudWatchLogsIntegrationArrayOutputWithContext(context.Context) AWSCloudWatchLogsIntegrationArrayOutput
}

type AWSCloudWatchLogsIntegrationArray []AWSCloudWatchLogsIntegrationInput

func (AWSCloudWatchLogsIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AWSCloudWatchLogsIntegration)(nil)).Elem()
}

func (i AWSCloudWatchLogsIntegrationArray) ToAWSCloudWatchLogsIntegrationArrayOutput() AWSCloudWatchLogsIntegrationArrayOutput {
	return i.ToAWSCloudWatchLogsIntegrationArrayOutputWithContext(context.Background())
}

func (i AWSCloudWatchLogsIntegrationArray) ToAWSCloudWatchLogsIntegrationArrayOutputWithContext(ctx context.Context) AWSCloudWatchLogsIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AWSCloudWatchLogsIntegrationArrayOutput)
}

// AWSCloudWatchLogsIntegrationMapInput is an input type that accepts AWSCloudWatchLogsIntegrationMap and AWSCloudWatchLogsIntegrationMapOutput values.
// You can construct a concrete instance of `AWSCloudWatchLogsIntegrationMapInput` via:
//
//          AWSCloudWatchLogsIntegrationMap{ "key": AWSCloudWatchLogsIntegrationArgs{...} }
type AWSCloudWatchLogsIntegrationMapInput interface {
	pulumi.Input

	ToAWSCloudWatchLogsIntegrationMapOutput() AWSCloudWatchLogsIntegrationMapOutput
	ToAWSCloudWatchLogsIntegrationMapOutputWithContext(context.Context) AWSCloudWatchLogsIntegrationMapOutput
}

type AWSCloudWatchLogsIntegrationMap map[string]AWSCloudWatchLogsIntegrationInput

func (AWSCloudWatchLogsIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AWSCloudWatchLogsIntegration)(nil)).Elem()
}

func (i AWSCloudWatchLogsIntegrationMap) ToAWSCloudWatchLogsIntegrationMapOutput() AWSCloudWatchLogsIntegrationMapOutput {
	return i.ToAWSCloudWatchLogsIntegrationMapOutputWithContext(context.Background())
}

func (i AWSCloudWatchLogsIntegrationMap) ToAWSCloudWatchLogsIntegrationMapOutputWithContext(ctx context.Context) AWSCloudWatchLogsIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AWSCloudWatchLogsIntegrationMapOutput)
}

type AWSCloudWatchLogsIntegrationOutput struct{ *pulumi.OutputState }

func (AWSCloudWatchLogsIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AWSCloudWatchLogsIntegration)(nil)).Elem()
}

func (o AWSCloudWatchLogsIntegrationOutput) ToAWSCloudWatchLogsIntegrationOutput() AWSCloudWatchLogsIntegrationOutput {
	return o
}

func (o AWSCloudWatchLogsIntegrationOutput) ToAWSCloudWatchLogsIntegrationOutputWithContext(ctx context.Context) AWSCloudWatchLogsIntegrationOutput {
	return o
}

type AWSCloudWatchLogsIntegrationArrayOutput struct{ *pulumi.OutputState }

func (AWSCloudWatchLogsIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AWSCloudWatchLogsIntegration)(nil)).Elem()
}

func (o AWSCloudWatchLogsIntegrationArrayOutput) ToAWSCloudWatchLogsIntegrationArrayOutput() AWSCloudWatchLogsIntegrationArrayOutput {
	return o
}

func (o AWSCloudWatchLogsIntegrationArrayOutput) ToAWSCloudWatchLogsIntegrationArrayOutputWithContext(ctx context.Context) AWSCloudWatchLogsIntegrationArrayOutput {
	return o
}

func (o AWSCloudWatchLogsIntegrationArrayOutput) Index(i pulumi.IntInput) AWSCloudWatchLogsIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AWSCloudWatchLogsIntegration {
		return vs[0].([]*AWSCloudWatchLogsIntegration)[vs[1].(int)]
	}).(AWSCloudWatchLogsIntegrationOutput)
}

type AWSCloudWatchLogsIntegrationMapOutput struct{ *pulumi.OutputState }

func (AWSCloudWatchLogsIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AWSCloudWatchLogsIntegration)(nil)).Elem()
}

func (o AWSCloudWatchLogsIntegrationMapOutput) ToAWSCloudWatchLogsIntegrationMapOutput() AWSCloudWatchLogsIntegrationMapOutput {
	return o
}

func (o AWSCloudWatchLogsIntegrationMapOutput) ToAWSCloudWatchLogsIntegrationMapOutputWithContext(ctx context.Context) AWSCloudWatchLogsIntegrationMapOutput {
	return o
}

func (o AWSCloudWatchLogsIntegrationMapOutput) MapIndex(k pulumi.StringInput) AWSCloudWatchLogsIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AWSCloudWatchLogsIntegration {
		return vs[0].(map[string]*AWSCloudWatchLogsIntegration)[vs[1].(string)]
	}).(AWSCloudWatchLogsIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AWSCloudWatchLogsIntegrationInput)(nil)).Elem(), &AWSCloudWatchLogsIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*AWSCloudWatchLogsIntegrationArrayInput)(nil)).Elem(), AWSCloudWatchLogsIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AWSCloudWatchLogsIntegrationMapInput)(nil)).Elem(), AWSCloudWatchLogsIntegrationMap{})
	pulumi.RegisterOutputType(AWSCloudWatchLogsIntegrationOutput{})
	pulumi.RegisterOutputType(AWSCloudWatchLogsIntegrationArrayOutput{})
	pulumi.RegisterOutputType(AWSCloudWatchLogsIntegrationMapOutput{})
}
