// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a way to check whether default EBS encryption is enabled for your AWS account in the current AWS region.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ebs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ebs.LookupEncryptionByDefault(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupEncryptionByDefault(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupEncryptionByDefaultResult, error) {
	var rv LookupEncryptionByDefaultResult
	err := ctx.Invoke("aws:ebs/getEncryptionByDefault:getEncryptionByDefault", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getEncryptionByDefault.
type LookupEncryptionByDefaultResult struct {
	// Whether or not default EBS encryption is enabled. Returns as `true` or `false`.
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
