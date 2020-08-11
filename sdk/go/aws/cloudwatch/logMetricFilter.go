// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CloudWatch Log Metric Filter resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		dada, err := cloudwatch.NewLogGroup(ctx, "dada", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewLogMetricFilter(ctx, "yada", &cloudwatch.LogMetricFilterArgs{
// 			Pattern:      pulumi.String(""),
// 			LogGroupName: dada.Name,
// 			MetricTransformation: &cloudwatch.LogMetricFilterMetricTransformationArgs{
// 				Name:      pulumi.String("EventCount"),
// 				Namespace: pulumi.String("YourNamespace"),
// 				Value:     pulumi.String("1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LogMetricFilter struct {
	pulumi.CustomResourceState

	// The name of the log group to associate the metric filter with.
	LogGroupName pulumi.StringOutput `pulumi:"logGroupName"`
	// A block defining collection of information
	// needed to define how metric data gets emitted. See below.
	MetricTransformation LogMetricFilterMetricTransformationOutput `pulumi:"metricTransformation"`
	// A name for the metric filter.
	Name pulumi.StringOutput `pulumi:"name"`
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern pulumi.StringOutput `pulumi:"pattern"`
}

// NewLogMetricFilter registers a new resource with the given unique name, arguments, and options.
func NewLogMetricFilter(ctx *pulumi.Context,
	name string, args *LogMetricFilterArgs, opts ...pulumi.ResourceOption) (*LogMetricFilter, error) {
	if args == nil || args.LogGroupName == nil {
		return nil, errors.New("missing required argument 'LogGroupName'")
	}
	if args == nil || args.MetricTransformation == nil {
		return nil, errors.New("missing required argument 'MetricTransformation'")
	}
	if args == nil || args.Pattern == nil {
		return nil, errors.New("missing required argument 'Pattern'")
	}
	if args == nil {
		args = &LogMetricFilterArgs{}
	}
	var resource LogMetricFilter
	err := ctx.RegisterResource("aws:cloudwatch/logMetricFilter:LogMetricFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogMetricFilter gets an existing LogMetricFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogMetricFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogMetricFilterState, opts ...pulumi.ResourceOption) (*LogMetricFilter, error) {
	var resource LogMetricFilter
	err := ctx.ReadResource("aws:cloudwatch/logMetricFilter:LogMetricFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogMetricFilter resources.
type logMetricFilterState struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName *string `pulumi:"logGroupName"`
	// A block defining collection of information
	// needed to define how metric data gets emitted. See below.
	MetricTransformation *LogMetricFilterMetricTransformation `pulumi:"metricTransformation"`
	// A name for the metric filter.
	Name *string `pulumi:"name"`
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern *string `pulumi:"pattern"`
}

type LogMetricFilterState struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName pulumi.StringPtrInput
	// A block defining collection of information
	// needed to define how metric data gets emitted. See below.
	MetricTransformation LogMetricFilterMetricTransformationPtrInput
	// A name for the metric filter.
	Name pulumi.StringPtrInput
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern pulumi.StringPtrInput
}

func (LogMetricFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logMetricFilterState)(nil)).Elem()
}

type logMetricFilterArgs struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName string `pulumi:"logGroupName"`
	// A block defining collection of information
	// needed to define how metric data gets emitted. See below.
	MetricTransformation LogMetricFilterMetricTransformation `pulumi:"metricTransformation"`
	// A name for the metric filter.
	Name *string `pulumi:"name"`
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern string `pulumi:"pattern"`
}

// The set of arguments for constructing a LogMetricFilter resource.
type LogMetricFilterArgs struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName pulumi.StringInput
	// A block defining collection of information
	// needed to define how metric data gets emitted. See below.
	MetricTransformation LogMetricFilterMetricTransformationInput
	// A name for the metric filter.
	Name pulumi.StringPtrInput
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern pulumi.StringInput
}

func (LogMetricFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logMetricFilterArgs)(nil)).Elem()
}
