// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SSM Maintenance Window Task resource
//
// ## Example Usage
// ### Run Command Tasks
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ssm.NewMaintenanceWindowTask(ctx, "example", &ssm.MaintenanceWindowTaskArgs{
// 			MaxConcurrency: pulumi.String("2"),
// 			MaxErrors:      pulumi.String("1"),
// 			Priority:       pulumi.Int(1),
// 			ServiceRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
// 			TaskArn:        pulumi.String("AWS-RunShellScript"),
// 			TaskType:       pulumi.String("RUN_COMMAND"),
// 			WindowId:       pulumi.Any(aws_ssm_maintenance_window.Example.Id),
// 			Targets: ssm.MaintenanceWindowTaskTargetArray{
// 				&ssm.MaintenanceWindowTaskTargetArgs{
// 					Key: pulumi.String("InstanceIds"),
// 					Values: pulumi.StringArray{
// 						pulumi.Any(aws_instance.Example.Id),
// 					},
// 				},
// 			},
// 			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
// 				RunCommandParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersArgs{
// 					OutputS3Bucket:    pulumi.Any(aws_s3_bucket.Example.Bucket),
// 					OutputS3KeyPrefix: pulumi.String("output"),
// 					ServiceRoleArn:    pulumi.Any(aws_iam_role.Example.Arn),
// 					TimeoutSeconds:    pulumi.Int(600),
// 					NotificationConfig: &ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigArgs{
// 						NotificationArn: pulumi.Any(aws_sns_topic.Example.Arn),
// 						NotificationEvents: pulumi.StringArray{
// 							pulumi.String("All"),
// 						},
// 						NotificationType: pulumi.String("Command"),
// 					},
// 					Parameters: ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterArray{
// 						&ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterArgs{
// 							Name: pulumi.String("commands"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("date"),
// 							},
// 						},
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
// ### Step Function Tasks
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ssm.NewMaintenanceWindowTask(ctx, "example", &ssm.MaintenanceWindowTaskArgs{
// 			MaxConcurrency: pulumi.String("2"),
// 			MaxErrors:      pulumi.String("1"),
// 			Priority:       pulumi.Int(1),
// 			ServiceRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
// 			TaskArn:        pulumi.Any(aws_sfn_activity.Example.Id),
// 			TaskType:       pulumi.String("STEP_FUNCTIONS"),
// 			WindowId:       pulumi.Any(aws_ssm_maintenance_window.Example.Id),
// 			Targets: ssm.MaintenanceWindowTaskTargetArray{
// 				&ssm.MaintenanceWindowTaskTargetArgs{
// 					Key: pulumi.String("InstanceIds"),
// 					Values: pulumi.StringArray{
// 						pulumi.Any(aws_instance.Example.Id),
// 					},
// 				},
// 			},
// 			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
// 				StepFunctionsParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersArgs{
// 					Input: pulumi.String("{\"key1\":\"value1\"}"),
// 					Name:  pulumi.String("example"),
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
type MaintenanceWindowTask struct {
	pulumi.CustomResourceState

	// The description of the maintenance window task.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency pulumi.StringOutput `pulumi:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors pulumi.StringOutput `pulumi:"maxErrors"`
	// The name of the maintenance window task.
	Name pulumi.StringOutput `pulumi:"name"`
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The role that should be assumed when executing the task.
	ServiceRoleArn pulumi.StringOutput `pulumi:"serviceRoleArn"`
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets MaintenanceWindowTaskTargetArrayOutput `pulumi:"targets"`
	// The ARN of the task to execute.
	TaskArn pulumi.StringOutput `pulumi:"taskArn"`
	// Configuration block with parameters for task execution.
	TaskInvocationParameters MaintenanceWindowTaskTaskInvocationParametersPtrOutput `pulumi:"taskInvocationParameters"`
	// The type of task being registered. The only allowed value is `RUN_COMMAND`.
	TaskType pulumi.StringOutput `pulumi:"taskType"`
	// The Id of the maintenance window to register the task with.
	WindowId pulumi.StringOutput `pulumi:"windowId"`
}

// NewMaintenanceWindowTask registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindowTask(ctx *pulumi.Context,
	name string, args *MaintenanceWindowTaskArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindowTask, error) {
	if args == nil || args.MaxConcurrency == nil {
		return nil, errors.New("missing required argument 'MaxConcurrency'")
	}
	if args == nil || args.MaxErrors == nil {
		return nil, errors.New("missing required argument 'MaxErrors'")
	}
	if args == nil || args.ServiceRoleArn == nil {
		return nil, errors.New("missing required argument 'ServiceRoleArn'")
	}
	if args == nil || args.Targets == nil {
		return nil, errors.New("missing required argument 'Targets'")
	}
	if args == nil || args.TaskArn == nil {
		return nil, errors.New("missing required argument 'TaskArn'")
	}
	if args == nil || args.TaskType == nil {
		return nil, errors.New("missing required argument 'TaskType'")
	}
	if args == nil || args.WindowId == nil {
		return nil, errors.New("missing required argument 'WindowId'")
	}
	if args == nil {
		args = &MaintenanceWindowTaskArgs{}
	}
	var resource MaintenanceWindowTask
	err := ctx.RegisterResource("aws:ssm/maintenanceWindowTask:MaintenanceWindowTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindowTask gets an existing MaintenanceWindowTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindowTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowTaskState, opts ...pulumi.ResourceOption) (*MaintenanceWindowTask, error) {
	var resource MaintenanceWindowTask
	err := ctx.ReadResource("aws:ssm/maintenanceWindowTask:MaintenanceWindowTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindowTask resources.
type maintenanceWindowTaskState struct {
	// The description of the maintenance window task.
	Description *string `pulumi:"description"`
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency *string `pulumi:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors *string `pulumi:"maxErrors"`
	// The name of the maintenance window task.
	Name *string `pulumi:"name"`
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority *int `pulumi:"priority"`
	// The role that should be assumed when executing the task.
	ServiceRoleArn *string `pulumi:"serviceRoleArn"`
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets []MaintenanceWindowTaskTarget `pulumi:"targets"`
	// The ARN of the task to execute.
	TaskArn *string `pulumi:"taskArn"`
	// Configuration block with parameters for task execution.
	TaskInvocationParameters *MaintenanceWindowTaskTaskInvocationParameters `pulumi:"taskInvocationParameters"`
	// The type of task being registered. The only allowed value is `RUN_COMMAND`.
	TaskType *string `pulumi:"taskType"`
	// The Id of the maintenance window to register the task with.
	WindowId *string `pulumi:"windowId"`
}

type MaintenanceWindowTaskState struct {
	// The description of the maintenance window task.
	Description pulumi.StringPtrInput
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency pulumi.StringPtrInput
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors pulumi.StringPtrInput
	// The name of the maintenance window task.
	Name pulumi.StringPtrInput
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority pulumi.IntPtrInput
	// The role that should be assumed when executing the task.
	ServiceRoleArn pulumi.StringPtrInput
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets MaintenanceWindowTaskTargetArrayInput
	// The ARN of the task to execute.
	TaskArn pulumi.StringPtrInput
	// Configuration block with parameters for task execution.
	TaskInvocationParameters MaintenanceWindowTaskTaskInvocationParametersPtrInput
	// The type of task being registered. The only allowed value is `RUN_COMMAND`.
	TaskType pulumi.StringPtrInput
	// The Id of the maintenance window to register the task with.
	WindowId pulumi.StringPtrInput
}

func (MaintenanceWindowTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowTaskState)(nil)).Elem()
}

type maintenanceWindowTaskArgs struct {
	// The description of the maintenance window task.
	Description *string `pulumi:"description"`
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency string `pulumi:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors string `pulumi:"maxErrors"`
	// The name of the maintenance window task.
	Name *string `pulumi:"name"`
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority *int `pulumi:"priority"`
	// The role that should be assumed when executing the task.
	ServiceRoleArn string `pulumi:"serviceRoleArn"`
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets []MaintenanceWindowTaskTarget `pulumi:"targets"`
	// The ARN of the task to execute.
	TaskArn string `pulumi:"taskArn"`
	// Configuration block with parameters for task execution.
	TaskInvocationParameters *MaintenanceWindowTaskTaskInvocationParameters `pulumi:"taskInvocationParameters"`
	// The type of task being registered. The only allowed value is `RUN_COMMAND`.
	TaskType string `pulumi:"taskType"`
	// The Id of the maintenance window to register the task with.
	WindowId string `pulumi:"windowId"`
}

// The set of arguments for constructing a MaintenanceWindowTask resource.
type MaintenanceWindowTaskArgs struct {
	// The description of the maintenance window task.
	Description pulumi.StringPtrInput
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency pulumi.StringInput
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors pulumi.StringInput
	// The name of the maintenance window task.
	Name pulumi.StringPtrInput
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority pulumi.IntPtrInput
	// The role that should be assumed when executing the task.
	ServiceRoleArn pulumi.StringInput
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets MaintenanceWindowTaskTargetArrayInput
	// The ARN of the task to execute.
	TaskArn pulumi.StringInput
	// Configuration block with parameters for task execution.
	TaskInvocationParameters MaintenanceWindowTaskTaskInvocationParametersPtrInput
	// The type of task being registered. The only allowed value is `RUN_COMMAND`.
	TaskType pulumi.StringInput
	// The Id of the maintenance window to register the task with.
	WindowId pulumi.StringInput
}

func (MaintenanceWindowTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowTaskArgs)(nil)).Elem()
}
