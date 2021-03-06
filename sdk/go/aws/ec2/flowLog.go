// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VPC/Subnet/ENI Flow Log to capture IP traffic for a specific network
// interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group or a S3 Bucket.
//
// ## Example Usage
// ### CloudWatch Logging
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"vpc-flow-logs.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewFlowLog(ctx, "exampleFlowLog", &ec2.FlowLogArgs{
// 			IamRoleArn:     exampleRole.Arn,
// 			LogDestination: exampleLogGroup.Arn,
// 			TrafficType:    pulumi.String("ALL"),
// 			VpcId:          pulumi.Any(aws_vpc.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "exampleRolePolicy", &iam.RolePolicyArgs{
// 			Role:   exampleRole.ID(),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"logs:CreateLogGroup\",\n", "        \"logs:CreateLogStream\",\n", "        \"logs:PutLogEvents\",\n", "        \"logs:DescribeLogGroups\",\n", "        \"logs:DescribeLogStreams\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### S3 Logging
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleBucket, err := s3.NewBucket(ctx, "exampleBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewFlowLog(ctx, "exampleFlowLog", &ec2.FlowLogArgs{
// 			LogDestination:     exampleBucket.Arn,
// 			LogDestinationType: pulumi.String("s3"),
// 			TrafficType:        pulumi.String("ALL"),
// 			VpcId:              pulumi.Any(aws_vpc.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FlowLog struct {
	pulumi.CustomResourceState

	// The ARN of the Flow Log.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Elastic Network Interface ID to attach to
	EniId pulumi.StringPtrOutput `pulumi:"eniId"`
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn pulumi.StringPtrOutput `pulumi:"iamRoleArn"`
	// The ARN of the logging destination.
	LogDestination pulumi.StringOutput `pulumi:"logDestination"`
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
	LogDestinationType pulumi.StringPtrOutput `pulumi:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat pulumi.StringOutput `pulumi:"logFormat"`
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName pulumi.StringOutput `pulumi:"logGroupName"`
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`.
	MaxAggregationInterval pulumi.IntPtrOutput `pulumi:"maxAggregationInterval"`
	// Subnet ID to attach to
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType pulumi.StringOutput `pulumi:"trafficType"`
	// VPC ID to attach to
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewFlowLog registers a new resource with the given unique name, arguments, and options.
func NewFlowLog(ctx *pulumi.Context,
	name string, args *FlowLogArgs, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	if args == nil || args.TrafficType == nil {
		return nil, errors.New("missing required argument 'TrafficType'")
	}
	if args == nil {
		args = &FlowLogArgs{}
	}
	var resource FlowLog
	err := ctx.RegisterResource("aws:ec2/flowLog:FlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowLog gets an existing FlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowLogState, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	var resource FlowLog
	err := ctx.ReadResource("aws:ec2/flowLog:FlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowLog resources.
type flowLogState struct {
	// The ARN of the Flow Log.
	Arn *string `pulumi:"arn"`
	// Elastic Network Interface ID to attach to
	EniId *string `pulumi:"eniId"`
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// The ARN of the logging destination.
	LogDestination *string `pulumi:"logDestination"`
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
	LogDestinationType *string `pulumi:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat *string `pulumi:"logFormat"`
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName *string `pulumi:"logGroupName"`
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`.
	MaxAggregationInterval *int `pulumi:"maxAggregationInterval"`
	// Subnet ID to attach to
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType *string `pulumi:"trafficType"`
	// VPC ID to attach to
	VpcId *string `pulumi:"vpcId"`
}

type FlowLogState struct {
	// The ARN of the Flow Log.
	Arn pulumi.StringPtrInput
	// Elastic Network Interface ID to attach to
	EniId pulumi.StringPtrInput
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn pulumi.StringPtrInput
	// The ARN of the logging destination.
	LogDestination pulumi.StringPtrInput
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
	LogDestinationType pulumi.StringPtrInput
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat pulumi.StringPtrInput
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName pulumi.StringPtrInput
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`.
	MaxAggregationInterval pulumi.IntPtrInput
	// Subnet ID to attach to
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType pulumi.StringPtrInput
	// VPC ID to attach to
	VpcId pulumi.StringPtrInput
}

func (FlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogState)(nil)).Elem()
}

type flowLogArgs struct {
	// Elastic Network Interface ID to attach to
	EniId *string `pulumi:"eniId"`
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// The ARN of the logging destination.
	LogDestination *string `pulumi:"logDestination"`
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
	LogDestinationType *string `pulumi:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat *string `pulumi:"logFormat"`
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName *string `pulumi:"logGroupName"`
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`.
	MaxAggregationInterval *int `pulumi:"maxAggregationInterval"`
	// Subnet ID to attach to
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType string `pulumi:"trafficType"`
	// VPC ID to attach to
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a FlowLog resource.
type FlowLogArgs struct {
	// Elastic Network Interface ID to attach to
	EniId pulumi.StringPtrInput
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn pulumi.StringPtrInput
	// The ARN of the logging destination.
	LogDestination pulumi.StringPtrInput
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
	LogDestinationType pulumi.StringPtrInput
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat pulumi.StringPtrInput
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName pulumi.StringPtrInput
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`.
	MaxAggregationInterval pulumi.IntPtrInput
	// Subnet ID to attach to
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType pulumi.StringInput
	// VPC ID to attach to
	VpcId pulumi.StringPtrInput
}

func (FlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogArgs)(nil)).Elem()
}
