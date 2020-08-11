// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Autoscaling Groups data source allows access to the list of AWS
// ASGs within a specific region. This will allow you to pass a list of AutoScaling Groups to other resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/autoscaling"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		groups, err := aws.GetAutoscalingGroups(ctx, &aws.GetAutoscalingGroupsArgs{
// 			Filters: []aws.GetAutoscalingGroupsFilter{
// 				aws.GetAutoscalingGroupsFilter{
// 					Name: "key",
// 					Values: []string{
// 						"Team",
// 					},
// 				},
// 				aws.GetAutoscalingGroupsFilter{
// 					Name: "value",
// 					Values: []string{
// 						"Pets",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = autoscaling.NewNotification(ctx, "slackNotifications", &autoscaling.NotificationArgs{
// 			GroupNames: toPulumiStringArray(groups.Names),
// 			Notifications: pulumi.StringArray{
// 				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH"),
// 				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE"),
// 				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH_ERROR"),
// 				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE_ERROR"),
// 			},
// 			TopicArn: pulumi.String("TOPIC ARN"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// func toPulumiStringArray(arr []string) pulumi.StringArray {
// 	var pulumiArr pulumi.StringArray
// 	for _, v := range arr {
// 		pulumiArr = append(pulumiArr, pulumi.String(v))
// 	}
// 	return pulumiArr
// }
// ```
func GetAutoscalingGroups(ctx *pulumi.Context, args *GetAutoscalingGroupsArgs, opts ...pulumi.InvokeOption) (*GetAutoscalingGroupsResult, error) {
	var rv GetAutoscalingGroupsResult
	err := ctx.Invoke("aws:index/getAutoscalingGroups:getAutoscalingGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutoscalingGroups.
type GetAutoscalingGroupsArgs struct {
	// A filter used to scope the list e.g. by tags. See [related docs](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Filter.html).
	Filters []GetAutoscalingGroupsFilter `pulumi:"filters"`
}

// A collection of values returned by getAutoscalingGroups.
type GetAutoscalingGroupsResult struct {
	// A list of the Autoscaling Groups Arns in the current region.
	Arns    []string                     `pulumi:"arns"`
	Filters []GetAutoscalingGroupsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the Autoscaling Groups in the current region.
	Names []string `pulumi:"names"`
}
