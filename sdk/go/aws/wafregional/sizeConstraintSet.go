// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Regional Size Constraint Set Resource for use with Application Load Balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafregional.NewSizeConstraintSet(ctx, "sizeConstraintSet", &wafregional.SizeConstraintSetArgs{
// 			SizeConstraints: wafregional.SizeConstraintSetSizeConstraintArray{
// 				&wafregional.SizeConstraintSetSizeConstraintArgs{
// 					ComparisonOperator: pulumi.String("EQ"),
// 					FieldToMatch: &wafregional.SizeConstraintSetSizeConstraintFieldToMatchArgs{
// 						Type: pulumi.String("BODY"),
// 					},
// 					Size:               pulumi.Int(4096),
// 					TextTransformation: pulumi.String("NONE"),
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
type SizeConstraintSet struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description of the Size Constraint Set.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints SizeConstraintSetSizeConstraintArrayOutput `pulumi:"sizeConstraints"`
}

// NewSizeConstraintSet registers a new resource with the given unique name, arguments, and options.
func NewSizeConstraintSet(ctx *pulumi.Context,
	name string, args *SizeConstraintSetArgs, opts ...pulumi.ResourceOption) (*SizeConstraintSet, error) {
	if args == nil {
		args = &SizeConstraintSetArgs{}
	}
	var resource SizeConstraintSet
	err := ctx.RegisterResource("aws:wafregional/sizeConstraintSet:SizeConstraintSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSizeConstraintSet gets an existing SizeConstraintSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSizeConstraintSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SizeConstraintSetState, opts ...pulumi.ResourceOption) (*SizeConstraintSet, error) {
	var resource SizeConstraintSet
	err := ctx.ReadResource("aws:wafregional/sizeConstraintSet:SizeConstraintSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SizeConstraintSet resources.
type sizeConstraintSetState struct {
	Arn *string `pulumi:"arn"`
	// The name or description of the Size Constraint Set.
	Name *string `pulumi:"name"`
	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints []SizeConstraintSetSizeConstraint `pulumi:"sizeConstraints"`
}

type SizeConstraintSetState struct {
	Arn pulumi.StringPtrInput
	// The name or description of the Size Constraint Set.
	Name pulumi.StringPtrInput
	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints SizeConstraintSetSizeConstraintArrayInput
}

func (SizeConstraintSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*sizeConstraintSetState)(nil)).Elem()
}

type sizeConstraintSetArgs struct {
	// The name or description of the Size Constraint Set.
	Name *string `pulumi:"name"`
	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints []SizeConstraintSetSizeConstraint `pulumi:"sizeConstraints"`
}

// The set of arguments for constructing a SizeConstraintSet resource.
type SizeConstraintSetArgs struct {
	// The name or description of the Size Constraint Set.
	Name pulumi.StringPtrInput
	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints SizeConstraintSetSizeConstraintArrayInput
}

func (SizeConstraintSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sizeConstraintSetArgs)(nil)).Elem()
}
