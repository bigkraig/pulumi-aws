// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IAM instance profile.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
// 			Path:             pulumi.String("/"),
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Action\": \"sts:AssumeRole\",\n", "            \"Principal\": {\n", "               \"Service\": \"ec2.amazonaws.com\"\n", "            },\n", "            \"Effect\": \"Allow\",\n", "            \"Sid\": \"\"\n", "        }\n", "    ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewInstanceProfile(ctx, "testProfile", &iam.InstanceProfileArgs{
// 			Role: role.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type InstanceProfile struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS to the instance profile.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation timestamp of the instance profile.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`
	// The profile's name. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// Path in which to create the profile.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The role name to include in the profile.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// The [unique ID][1] assigned by AWS.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewInstanceProfile registers a new resource with the given unique name, arguments, and options.
func NewInstanceProfile(ctx *pulumi.Context,
	name string, args *InstanceProfileArgs, opts ...pulumi.ResourceOption) (*InstanceProfile, error) {
	if args == nil {
		args = &InstanceProfileArgs{}
	}
	var resource InstanceProfile
	err := ctx.RegisterResource("aws:iam/instanceProfile:InstanceProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceProfile gets an existing InstanceProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceProfileState, opts ...pulumi.ResourceOption) (*InstanceProfile, error) {
	var resource InstanceProfile
	err := ctx.ReadResource("aws:iam/instanceProfile:InstanceProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceProfile resources.
type instanceProfileState struct {
	// The ARN assigned by AWS to the instance profile.
	Arn *string `pulumi:"arn"`
	// The creation timestamp of the instance profile.
	CreateDate *string `pulumi:"createDate"`
	// The profile's name. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Path in which to create the profile.
	Path *string `pulumi:"path"`
	// The role name to include in the profile.
	Role *string `pulumi:"role"`
	// The [unique ID][1] assigned by AWS.
	UniqueId *string `pulumi:"uniqueId"`
}

type InstanceProfileState struct {
	// The ARN assigned by AWS to the instance profile.
	Arn pulumi.StringPtrInput
	// The creation timestamp of the instance profile.
	CreateDate pulumi.StringPtrInput
	// The profile's name. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Path in which to create the profile.
	Path pulumi.StringPtrInput
	// The role name to include in the profile.
	Role pulumi.StringPtrInput
	// The [unique ID][1] assigned by AWS.
	UniqueId pulumi.StringPtrInput
}

func (InstanceProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceProfileState)(nil)).Elem()
}

type instanceProfileArgs struct {
	// The profile's name. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Path in which to create the profile.
	Path *string `pulumi:"path"`
	// The role name to include in the profile.
	Role interface{} `pulumi:"role"`
}

// The set of arguments for constructing a InstanceProfile resource.
type InstanceProfileArgs struct {
	// The profile's name. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Path in which to create the profile.
	Path pulumi.StringPtrInput
	// The role name to include in the profile.
	Role pulumi.Input
}

func (InstanceProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceProfileArgs)(nil)).Elem()
}
