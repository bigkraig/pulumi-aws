// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AutoScaling Group with Notification support, via SNS Topics. Each of
// the `notifications` map to a [Notification Configuration](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeNotificationConfigurations.html) inside Amazon Web
// Services, and are applied to each AutoScaling Group you supply.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/autoscaling"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := sns.NewTopic(ctx, "example", nil)
// 		if err != nil {
// 			return err
// 		}
// 		bar, err := autoscaling.NewGroup(ctx, "bar", nil)
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := autoscaling.NewGroup(ctx, "foo", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = autoscaling.NewNotification(ctx, "exampleNotifications", &autoscaling.NotificationArgs{
// 			GroupNames: pulumi.StringArray{
// 				bar.Name,
// 				foo.Name,
// 			},
// 			Notifications: pulumi.StringArray{
// 				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH"),
// 				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE"),
// 				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH_ERROR"),
// 				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE_ERROR"),
// 			},
// 			TopicArn: example.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Notification struct {
	pulumi.CustomResourceState

	// A list of AutoScaling Group Names
	GroupNames pulumi.StringArrayOutput `pulumi:"groupNames"`
	// A list of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications pulumi.StringArrayOutput `pulumi:"notifications"`
	// The Topic ARN for notifications to be sent through
	TopicArn pulumi.StringOutput `pulumi:"topicArn"`
}

// NewNotification registers a new resource with the given unique name, arguments, and options.
func NewNotification(ctx *pulumi.Context,
	name string, args *NotificationArgs, opts ...pulumi.ResourceOption) (*Notification, error) {
	if args == nil || args.GroupNames == nil {
		return nil, errors.New("missing required argument 'GroupNames'")
	}
	if args == nil || args.Notifications == nil {
		return nil, errors.New("missing required argument 'Notifications'")
	}
	if args == nil || args.TopicArn == nil {
		return nil, errors.New("missing required argument 'TopicArn'")
	}
	if args == nil {
		args = &NotificationArgs{}
	}
	var resource Notification
	err := ctx.RegisterResource("aws:autoscaling/notification:Notification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotification gets an existing Notification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationState, opts ...pulumi.ResourceOption) (*Notification, error) {
	var resource Notification
	err := ctx.ReadResource("aws:autoscaling/notification:Notification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Notification resources.
type notificationState struct {
	// A list of AutoScaling Group Names
	GroupNames []string `pulumi:"groupNames"`
	// A list of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications []string `pulumi:"notifications"`
	// The Topic ARN for notifications to be sent through
	TopicArn *string `pulumi:"topicArn"`
}

type NotificationState struct {
	// A list of AutoScaling Group Names
	GroupNames pulumi.StringArrayInput
	// A list of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications pulumi.StringArrayInput
	// The Topic ARN for notifications to be sent through
	TopicArn pulumi.StringPtrInput
}

func (NotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationState)(nil)).Elem()
}

type notificationArgs struct {
	// A list of AutoScaling Group Names
	GroupNames []string `pulumi:"groupNames"`
	// A list of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications []string `pulumi:"notifications"`
	// The Topic ARN for notifications to be sent through
	TopicArn string `pulumi:"topicArn"`
}

// The set of arguments for constructing a Notification resource.
type NotificationArgs struct {
	// A list of AutoScaling Group Names
	GroupNames pulumi.StringArrayInput
	// A list of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications pulumi.StringArrayInput
	// The Topic ARN for notifications to be sent through
	TopicArn pulumi.StringInput
}

func (NotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationArgs)(nil)).Elem()
}
