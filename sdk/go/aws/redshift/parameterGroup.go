// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Redshift Cluster parameter group resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := redshift.NewParameterGroup(ctx, "bar", &redshift.ParameterGroupArgs{
// 			Family: pulumi.String("redshift-1.0"),
// 			Parameters: redshift.ParameterGroupParameterArray{
// 				&redshift.ParameterGroupParameterArgs{
// 					Name:  pulumi.String("require_ssl"),
// 					Value: pulumi.String("true"),
// 				},
// 				&redshift.ParameterGroupParameterArgs{
// 					Name:  pulumi.String("query_group"),
// 					Value: pulumi.String("example"),
// 				},
// 				&redshift.ParameterGroupParameterArgs{
// 					Name:  pulumi.String("enable_user_activity_logging"),
// 					Value: pulumi.String("true"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ParameterGroup struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of parameter group
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the Redshift parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// The family of the Redshift parameter group.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the Redshift parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of Redshift parameters to apply.
	Parameters ParameterGroupParameterArrayOutput `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewParameterGroup(ctx *pulumi.Context,
	name string, args *ParameterGroupArgs, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	if args == nil || args.Family == nil {
		return nil, errors.New("missing required argument 'Family'")
	}
	if args == nil {
		args = &ParameterGroupArgs{}
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource ParameterGroup
	err := ctx.RegisterResource("aws:redshift/parameterGroup:ParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameterGroup gets an existing ParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterGroupState, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	var resource ParameterGroup
	err := ctx.ReadResource("aws:redshift/parameterGroup:ParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterGroup resources.
type parameterGroupState struct {
	// Amazon Resource Name (ARN) of parameter group
	Arn *string `pulumi:"arn"`
	// The description of the Redshift parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the Redshift parameter group.
	Family *string `pulumi:"family"`
	// The name of the Redshift parameter.
	Name *string `pulumi:"name"`
	// A list of Redshift parameters to apply.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ParameterGroupState struct {
	// Amazon Resource Name (ARN) of parameter group
	Arn pulumi.StringPtrInput
	// The description of the Redshift parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the Redshift parameter group.
	Family pulumi.StringPtrInput
	// The name of the Redshift parameter.
	Name pulumi.StringPtrInput
	// A list of Redshift parameters to apply.
	Parameters ParameterGroupParameterArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupState)(nil)).Elem()
}

type parameterGroupArgs struct {
	// The description of the Redshift parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the Redshift parameter group.
	Family string `pulumi:"family"`
	// The name of the Redshift parameter.
	Name *string `pulumi:"name"`
	// A list of Redshift parameters to apply.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ParameterGroup resource.
type ParameterGroupArgs struct {
	// The description of the Redshift parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the Redshift parameter group.
	Family pulumi.StringInput
	// The name of the Redshift parameter.
	Name pulumi.StringPtrInput
	// A list of Redshift parameters to apply.
	Parameters ParameterGroupParameterArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupArgs)(nil)).Elem()
}
