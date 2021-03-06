// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CloudWatch Event Target resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kinesis"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		console, err := cloudwatch.NewEventRule(ctx, "console", &cloudwatch.EventRuleArgs{
// 			Description:  pulumi.String("Capture all EC2 scaling events"),
// 			EventPattern: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"source\": [\n", "    \"aws.autoscaling\"\n", "  ],\n", "  \"detail-type\": [\n", "    \"EC2 Instance Launch Successful\",\n", "    \"EC2 Instance Terminate Successful\",\n", "    \"EC2 Instance Launch Unsuccessful\",\n", "    \"EC2 Instance Terminate Unsuccessful\"\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
// 			ShardCount: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewEventTarget(ctx, "yada", &cloudwatch.EventTargetArgs{
// 			Rule: console.Name,
// 			Arn:  testStream.Arn,
// 			RunCommandTargets: cloudwatch.EventTargetRunCommandTargetArray{
// 				&cloudwatch.EventTargetRunCommandTargetArgs{
// 					Key: pulumi.String("tag:Name"),
// 					Values: pulumi.StringArray{
// 						pulumi.String("FooBar"),
// 					},
// 				},
// 				&cloudwatch.EventTargetRunCommandTargetArgs{
// 					Key: pulumi.String("InstanceIds"),
// 					Values: pulumi.StringArray{
// 						pulumi.String("i-162058cd308bffec2"),
// 					},
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
// ## Example SSM Document Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ssmLifecycleTrust, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"sts:AssumeRole",
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "Service",
// 							Identifiers: []string{
// 								"events.amazonaws.com",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		stopInstance, err := ssm.NewDocument(ctx, "stopInstance", &ssm.DocumentArgs{
// 			DocumentType: pulumi.String("Command"),
// 			Content:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "  {\n", "    \"schemaVersion\": \"1.2\",\n", "    \"description\": \"Stop an instance\",\n", "    \"parameters\": {\n", "\n", "    },\n", "    \"runtimeConfig\": {\n", "      \"aws:runShellScript\": {\n", "        \"properties\": [\n", "          {\n", "            \"id\": \"0.aws:runShellScript\",\n", "            \"runCommand\": [\"halt\"]\n", "          }\n", "        ]\n", "      }\n", "    }\n", "  }\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ssmLifecycleRole, err := iam.NewRole(ctx, "ssmLifecycleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(ssmLifecycleTrust.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewPolicy(ctx, "ssmLifecyclePolicy", &iam.PolicyArgs{
// 			Policy: ssmLifecyclePolicyDocument.ApplyT(func(ssmLifecyclePolicyDocument iam.GetPolicyDocumentResult) (string, error) {
// 				return ssmLifecyclePolicyDocument.Json, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		stopInstancesEventRule, err := cloudwatch.NewEventRule(ctx, "stopInstancesEventRule", &cloudwatch.EventRuleArgs{
// 			Description:        pulumi.String("Stop instances nightly"),
// 			ScheduleExpression: pulumi.String("cron(0 0 * * ? *)"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewEventTarget(ctx, "stopInstancesEventTarget", &cloudwatch.EventTargetArgs{
// 			Arn:     stopInstance.Arn,
// 			Rule:    stopInstancesEventRule.Name,
// 			RoleArn: ssmLifecycleRole.Arn,
// 			RunCommandTargets: cloudwatch.EventTargetRunCommandTargetArray{
// 				&cloudwatch.EventTargetRunCommandTargetArgs{
// 					Key: pulumi.String("tag:Terminate"),
// 					Values: pulumi.StringArray{
// 						pulumi.String("midnight"),
// 					},
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
// ## Example RunCommand Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		stopInstancesEventRule, err := cloudwatch.NewEventRule(ctx, "stopInstancesEventRule", &cloudwatch.EventRuleArgs{
// 			Description:        pulumi.String("Stop instances nightly"),
// 			ScheduleExpression: pulumi.String("cron(0 0 * * ? *)"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewEventTarget(ctx, "stopInstancesEventTarget", &cloudwatch.EventTargetArgs{
// 			Arn:     pulumi.String(fmt.Sprintf("%v%v%v", "arn:aws:ssm:", _var.Aws_region, "::document/AWS-RunShellScript")),
// 			Input:   pulumi.String("{\"commands\":[\"halt\"]}"),
// 			Rule:    stopInstancesEventRule.Name,
// 			RoleArn: pulumi.Any(aws_iam_role.Ssm_lifecycle.Arn),
// 			RunCommandTargets: cloudwatch.EventTargetRunCommandTargetArray{
// 				&cloudwatch.EventTargetRunCommandTargetArgs{
// 					Key: pulumi.String("tag:Terminate"),
// 					Values: pulumi.StringArray{
// 						pulumi.String("midnight"),
// 					},
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
type EventTarget struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) associated of the target.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget EventTargetBatchTargetPtrOutput `pulumi:"batchTarget"`
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget EventTargetEcsTargetPtrOutput `pulumi:"ecsTarget"`
	// Valid JSON text passed to the target.
	Input pulumi.StringPtrOutput `pulumi:"input"`
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath pulumi.StringPtrOutput `pulumi:"inputPath"`
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer EventTargetInputTransformerPtrOutput `pulumi:"inputTransformer"`
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget EventTargetKinesisTargetPtrOutput `pulumi:"kinesisTarget"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// The name of the rule you want to add targets to.
	Rule pulumi.StringOutput `pulumi:"rule"`
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets EventTargetRunCommandTargetArrayOutput `pulumi:"runCommandTargets"`
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget EventTargetSqsTargetPtrOutput `pulumi:"sqsTarget"`
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
}

// NewEventTarget registers a new resource with the given unique name, arguments, and options.
func NewEventTarget(ctx *pulumi.Context,
	name string, args *EventTargetArgs, opts ...pulumi.ResourceOption) (*EventTarget, error) {
	if args == nil || args.Arn == nil {
		return nil, errors.New("missing required argument 'Arn'")
	}
	if args == nil || args.Rule == nil {
		return nil, errors.New("missing required argument 'Rule'")
	}
	if args == nil {
		args = &EventTargetArgs{}
	}
	var resource EventTarget
	err := ctx.RegisterResource("aws:cloudwatch/eventTarget:EventTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventTarget gets an existing EventTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventTargetState, opts ...pulumi.ResourceOption) (*EventTarget, error) {
	var resource EventTarget
	err := ctx.ReadResource("aws:cloudwatch/eventTarget:EventTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventTarget resources.
type eventTargetState struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn *string `pulumi:"arn"`
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget *EventTargetBatchTarget `pulumi:"batchTarget"`
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget *EventTargetEcsTarget `pulumi:"ecsTarget"`
	// Valid JSON text passed to the target.
	Input *string `pulumi:"input"`
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath *string `pulumi:"inputPath"`
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer *EventTargetInputTransformer `pulumi:"inputTransformer"`
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget *EventTargetKinesisTarget `pulumi:"kinesisTarget"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn *string `pulumi:"roleArn"`
	// The name of the rule you want to add targets to.
	Rule *string `pulumi:"rule"`
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets []EventTargetRunCommandTarget `pulumi:"runCommandTargets"`
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget *EventTargetSqsTarget `pulumi:"sqsTarget"`
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId *string `pulumi:"targetId"`
}

type EventTargetState struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn pulumi.StringPtrInput
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget EventTargetBatchTargetPtrInput
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget EventTargetEcsTargetPtrInput
	// Valid JSON text passed to the target.
	Input pulumi.StringPtrInput
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath pulumi.StringPtrInput
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer EventTargetInputTransformerPtrInput
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget EventTargetKinesisTargetPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn pulumi.StringPtrInput
	// The name of the rule you want to add targets to.
	Rule pulumi.StringPtrInput
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets EventTargetRunCommandTargetArrayInput
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget EventTargetSqsTargetPtrInput
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId pulumi.StringPtrInput
}

func (EventTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTargetState)(nil)).Elem()
}

type eventTargetArgs struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn string `pulumi:"arn"`
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget *EventTargetBatchTarget `pulumi:"batchTarget"`
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget *EventTargetEcsTarget `pulumi:"ecsTarget"`
	// Valid JSON text passed to the target.
	Input *string `pulumi:"input"`
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath *string `pulumi:"inputPath"`
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer *EventTargetInputTransformer `pulumi:"inputTransformer"`
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget *EventTargetKinesisTarget `pulumi:"kinesisTarget"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn *string `pulumi:"roleArn"`
	// The name of the rule you want to add targets to.
	Rule string `pulumi:"rule"`
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets []EventTargetRunCommandTarget `pulumi:"runCommandTargets"`
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget *EventTargetSqsTarget `pulumi:"sqsTarget"`
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId *string `pulumi:"targetId"`
}

// The set of arguments for constructing a EventTarget resource.
type EventTargetArgs struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn pulumi.StringInput
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget EventTargetBatchTargetPtrInput
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget EventTargetEcsTargetPtrInput
	// Valid JSON text passed to the target.
	Input pulumi.StringPtrInput
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath pulumi.StringPtrInput
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer EventTargetInputTransformerPtrInput
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget EventTargetKinesisTargetPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn pulumi.StringPtrInput
	// The name of the rule you want to add targets to.
	Rule pulumi.StringInput
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets EventTargetRunCommandTargetArrayInput
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget EventTargetSqsTargetPtrInput
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId pulumi.StringPtrInput
}

func (EventTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTargetArgs)(nil)).Elem()
}
