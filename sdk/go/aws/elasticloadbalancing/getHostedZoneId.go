// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
// in a given region for the purpose of using in an AWS Route53 Alias.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elb"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := elb.GetHostedZoneId(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewRecord(ctx, "www", &route53.RecordArgs{
// 			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
// 			Name:   pulumi.String("example.com"),
// 			Type:   pulumi.String("A"),
// 			Aliases: route53.RecordAliasArray{
// 				&route53.RecordAliasArgs{
// 					Name:                 pulumi.Any(aws_elb.Main.Dns_name),
// 					ZoneId:               pulumi.String(main.Id),
// 					EvaluateTargetHealth: pulumi.Bool(true),
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
//
// Deprecated: aws.elasticloadbalancing.getHostedZoneId has been deprecated in favor of aws.elb.getHostedZoneId
func GetHostedZoneId(ctx *pulumi.Context, args *GetHostedZoneIdArgs, opts ...pulumi.InvokeOption) (*GetHostedZoneIdResult, error) {
	var rv GetHostedZoneIdResult
	err := ctx.Invoke("aws:elasticloadbalancing/getHostedZoneId:getHostedZoneId", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostedZoneId.
type GetHostedZoneIdArgs struct {
	// Name of the region whose AWS ELB HostedZoneId is desired.
	// Defaults to the region from the AWS provider configuration.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getHostedZoneId.
type GetHostedZoneIdResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Region *string `pulumi:"region"`
}
