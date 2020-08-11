// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package workspaces

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Retrieve information about an AWS WorkSpaces bundle.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/workspaces"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "Value with Windows 10 and Office 2016"
// 		opt1 := "AMAZON"
// 		_, err := workspaces.GetBundle(ctx, &workspaces.GetBundleArgs{
// 			Name:  &opt0,
// 			Owner: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetBundle(ctx *pulumi.Context, args *GetBundleArgs, opts ...pulumi.InvokeOption) (*GetBundleResult, error) {
	var rv GetBundleResult
	err := ctx.Invoke("aws:workspaces/getBundle:getBundle", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBundle.
type GetBundleArgs struct {
	// The ID of the bundle.
	BundleId *string `pulumi:"bundleId"`
	// The name of the bundle. You cannot combine this parameter with `bundleId`.
	Name *string `pulumi:"name"`
	// The owner of the bundles. You have to leave it blank for own bundles. You cannot combine this parameter with `bundleId`.
	Owner *string `pulumi:"owner"`
}

// A collection of values returned by getBundle.
type GetBundleResult struct {
	// The ID of the bundle.
	BundleId *string `pulumi:"bundleId"`
	// The compute type. See supported fields below.
	ComputeTypes []GetBundleComputeType `pulumi:"computeTypes"`
	// The description of the bundle.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the compute type.
	Name *string `pulumi:"name"`
	// The owner of the bundle.
	Owner *string `pulumi:"owner"`
	// The root volume. See supported fields below.
	RootStorages []GetBundleRootStorage `pulumi:"rootStorages"`
	// The user storage. See supported fields below.
	UserStorages []GetBundleUserStorage `pulumi:"userStorages"`
}
