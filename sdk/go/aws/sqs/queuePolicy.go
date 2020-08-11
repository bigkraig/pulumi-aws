// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows you to set a policy of an SQS Queue
// while referencing ARN of the queue within the policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sqs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		queue, err := sqs.NewQueue(ctx, "queue", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sqs.NewQueuePolicy(ctx, "test", &sqs.QueuePolicyArgs{
// 			QueueUrl: queue.ID(),
// 			Policy: queue.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Id\": \"sqspolicy\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"First\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": \"*\",\n", "      \"Action\": \"sqs:SendMessage\",\n", "      \"Resource\": \"", arn, "\",\n", "      \"Condition\": {\n", "        \"ArnEquals\": {\n", "          \"aws:SourceArn\": \"", aws_sns_topic.Example.Arn, "\"\n", "        }\n", "      }\n", "    }\n", "  ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type QueuePolicy struct {
	pulumi.CustomResourceState

	// The JSON policy for the SQS queue.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl pulumi.StringOutput `pulumi:"queueUrl"`
}

// NewQueuePolicy registers a new resource with the given unique name, arguments, and options.
func NewQueuePolicy(ctx *pulumi.Context,
	name string, args *QueuePolicyArgs, opts ...pulumi.ResourceOption) (*QueuePolicy, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil || args.QueueUrl == nil {
		return nil, errors.New("missing required argument 'QueueUrl'")
	}
	if args == nil {
		args = &QueuePolicyArgs{}
	}
	var resource QueuePolicy
	err := ctx.RegisterResource("aws:sqs/queuePolicy:QueuePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueuePolicy gets an existing QueuePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueuePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueuePolicyState, opts ...pulumi.ResourceOption) (*QueuePolicy, error) {
	var resource QueuePolicy
	err := ctx.ReadResource("aws:sqs/queuePolicy:QueuePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueuePolicy resources.
type queuePolicyState struct {
	// The JSON policy for the SQS queue.
	Policy *string `pulumi:"policy"`
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl *string `pulumi:"queueUrl"`
}

type QueuePolicyState struct {
	// The JSON policy for the SQS queue.
	Policy pulumi.StringPtrInput
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl pulumi.StringPtrInput
}

func (QueuePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*queuePolicyState)(nil)).Elem()
}

type queuePolicyArgs struct {
	// The JSON policy for the SQS queue.
	Policy interface{} `pulumi:"policy"`
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl string `pulumi:"queueUrl"`
}

// The set of arguments for constructing a QueuePolicy resource.
type QueuePolicyArgs struct {
	// The JSON policy for the SQS queue.
	Policy pulumi.Input
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl pulumi.StringInput
}

func (QueuePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queuePolicyArgs)(nil)).Elem()
}
