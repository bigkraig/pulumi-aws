// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **Note:** To prevent a race condition during service deletion, make sure to set `dependsOn` to the related `iam.RolePolicy`; otherwise, the policy may be destroyed too soon and the ECS service will then get stuck in the `DRAINING` state.
//
// Provides an ECS service - effectively a task that is expected to run until an error occurs or a user terminates it (typically a webserver or a database).
//
// See [ECS Services section in AWS developer guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewService(ctx, "mongo", &ecs.ServiceArgs{
// 			Cluster:        pulumi.Any(aws_ecs_cluster.Foo.Id),
// 			TaskDefinition: pulumi.Any(aws_ecs_task_definition.Mongo.Arn),
// 			DesiredCount:   pulumi.Int(3),
// 			IamRole:        pulumi.Any(aws_iam_role.Foo.Arn),
// 			OrderedPlacementStrategies: ecs.ServiceOrderedPlacementStrategyArray{
// 				&ecs.ServiceOrderedPlacementStrategyArgs{
// 					Type:  pulumi.String("binpack"),
// 					Field: pulumi.String("cpu"),
// 				},
// 			},
// 			LoadBalancers: ecs.ServiceLoadBalancerArray{
// 				&ecs.ServiceLoadBalancerArgs{
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Foo.Arn),
// 					ContainerName:  pulumi.String("mongo"),
// 					ContainerPort:  pulumi.Int(8080),
// 				},
// 			},
// 			PlacementConstraints: ecs.ServicePlacementConstraintArray{
// 				&ecs.ServicePlacementConstraintArgs{
// 					Type:       pulumi.String("memberOf"),
// 					Expression: pulumi.String("attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"),
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			aws_iam_role_policy.Foo,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Ignoring Changes to Desired Count
//
// You can use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to create an ECS service with an initial count of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewService(ctx, "example", &ecs.ServiceArgs{
// 			DesiredCount: pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Daemon Scheduling Strategy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewService(ctx, "bar", &ecs.ServiceArgs{
// 			Cluster:            pulumi.Any(aws_ecs_cluster.Foo.Id),
// 			TaskDefinition:     pulumi.Any(aws_ecs_task_definition.Bar.Arn),
// 			SchedulingStrategy: pulumi.String("DAEMON"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### External Deployment Controller
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewService(ctx, "example", &ecs.ServiceArgs{
// 			Cluster: pulumi.Any(aws_ecs_cluster.Example.Id),
// 			DeploymentController: &ecs.ServiceDeploymentControllerArgs{
// 				Type: pulumi.String("EXTERNAL"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Service struct {
	pulumi.CustomResourceState

	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies ServiceCapacityProviderStrategyArrayOutput `pulumi:"capacityProviderStrategies"`
	// ARN of an ECS cluster
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// Configuration block containing deployment controller configuration. Defined below.
	DeploymentController ServiceDeploymentControllerPtrOutput `pulumi:"deploymentController"`
	// The upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
	DeploymentMaximumPercent pulumi.IntPtrOutput `pulumi:"deploymentMaximumPercent"`
	// The lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
	DeploymentMinimumHealthyPercent pulumi.IntPtrOutput `pulumi:"deploymentMinimumHealthyPercent"`
	// The number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
	DesiredCount pulumi.IntPtrOutput `pulumi:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	EnableEcsManagedTags pulumi.BoolPtrOutput `pulumi:"enableEcsManagedTags"`
	// Enable to force a new task deployment of the service. This can be used to update tasks to use a newer Docker image with same image/tag combination (e.g. `myimage:latest`), roll Fargate tasks onto a newer platform version, or immediately deploy `orderedPlacementStrategy` and `placementConstraints` updates.
	ForceNewDeployment pulumi.BoolPtrOutput `pulumi:"forceNewDeployment"`
	// Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
	HealthCheckGracePeriodSeconds pulumi.IntPtrOutput `pulumi:"healthCheckGracePeriodSeconds"`
	// ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
	IamRole pulumi.StringOutput `pulumi:"iamRole"`
	// The launch type on which to run your service. The valid values are `EC2` and `FARGATE`. Defaults to `EC2`.
	LaunchType pulumi.StringOutput `pulumi:"launchType"`
	// A load balancer block. Load balancers documented below.
	LoadBalancers ServiceLoadBalancerArrayOutput `pulumi:"loadBalancers"`
	// The name of the service (up to 255 letters, numbers, hyphens, and underscores)
	Name pulumi.StringOutput `pulumi:"name"`
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes.
	NetworkConfiguration ServiceNetworkConfigurationPtrOutput `pulumi:"networkConfiguration"`
	// Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. The maximum number of `orderedPlacementStrategy` blocks is `5`. Defined below.
	OrderedPlacementStrategies ServiceOrderedPlacementStrategyArrayOutput `pulumi:"orderedPlacementStrategies"`
	// rules that are taken into consideration during task placement. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. Maximum number of `placementConstraints` is `10`. Defined below.
	PlacementConstraints ServicePlacementConstraintArrayOutput `pulumi:"placementConstraints"`
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion pulumi.StringOutput `pulumi:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
	PropagateTags pulumi.StringPtrOutput `pulumi:"propagateTags"`
	// The scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [*Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy*](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html).
	SchedulingStrategy pulumi.StringPtrOutput `pulumi:"schedulingStrategy"`
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`.
	ServiceRegistries ServiceServiceRegistriesPtrOutput `pulumi:"serviceRegistries"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service. Required unless using the `EXTERNAL` deployment controller. If a revision is not specified, the latest `ACTIVE` revision is used.
	TaskDefinition     pulumi.StringPtrOutput `pulumi:"taskDefinition"`
	WaitForSteadyState pulumi.BoolPtrOutput   `pulumi:"waitForSteadyState"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		args = &ServiceArgs{}
	}
	var resource Service
	err := ctx.RegisterResource("aws:ecs/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("aws:ecs/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies []ServiceCapacityProviderStrategy `pulumi:"capacityProviderStrategies"`
	// ARN of an ECS cluster
	Cluster *string `pulumi:"cluster"`
	// Configuration block containing deployment controller configuration. Defined below.
	DeploymentController *ServiceDeploymentController `pulumi:"deploymentController"`
	// The upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
	DeploymentMaximumPercent *int `pulumi:"deploymentMaximumPercent"`
	// The lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
	DeploymentMinimumHealthyPercent *int `pulumi:"deploymentMinimumHealthyPercent"`
	// The number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
	DesiredCount *int `pulumi:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	EnableEcsManagedTags *bool `pulumi:"enableEcsManagedTags"`
	// Enable to force a new task deployment of the service. This can be used to update tasks to use a newer Docker image with same image/tag combination (e.g. `myimage:latest`), roll Fargate tasks onto a newer platform version, or immediately deploy `orderedPlacementStrategy` and `placementConstraints` updates.
	ForceNewDeployment *bool `pulumi:"forceNewDeployment"`
	// Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
	HealthCheckGracePeriodSeconds *int `pulumi:"healthCheckGracePeriodSeconds"`
	// ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
	IamRole *string `pulumi:"iamRole"`
	// The launch type on which to run your service. The valid values are `EC2` and `FARGATE`. Defaults to `EC2`.
	LaunchType *string `pulumi:"launchType"`
	// A load balancer block. Load balancers documented below.
	LoadBalancers []ServiceLoadBalancer `pulumi:"loadBalancers"`
	// The name of the service (up to 255 letters, numbers, hyphens, and underscores)
	Name *string `pulumi:"name"`
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes.
	NetworkConfiguration *ServiceNetworkConfiguration `pulumi:"networkConfiguration"`
	// Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. The maximum number of `orderedPlacementStrategy` blocks is `5`. Defined below.
	OrderedPlacementStrategies []ServiceOrderedPlacementStrategy `pulumi:"orderedPlacementStrategies"`
	// rules that are taken into consideration during task placement. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. Maximum number of `placementConstraints` is `10`. Defined below.
	PlacementConstraints []ServicePlacementConstraint `pulumi:"placementConstraints"`
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion *string `pulumi:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
	PropagateTags *string `pulumi:"propagateTags"`
	// The scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [*Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy*](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html).
	SchedulingStrategy *string `pulumi:"schedulingStrategy"`
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`.
	ServiceRegistries *ServiceServiceRegistries `pulumi:"serviceRegistries"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service. Required unless using the `EXTERNAL` deployment controller. If a revision is not specified, the latest `ACTIVE` revision is used.
	TaskDefinition     *string `pulumi:"taskDefinition"`
	WaitForSteadyState *bool   `pulumi:"waitForSteadyState"`
}

type ServiceState struct {
	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies ServiceCapacityProviderStrategyArrayInput
	// ARN of an ECS cluster
	Cluster pulumi.StringPtrInput
	// Configuration block containing deployment controller configuration. Defined below.
	DeploymentController ServiceDeploymentControllerPtrInput
	// The upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
	DeploymentMaximumPercent pulumi.IntPtrInput
	// The lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
	DeploymentMinimumHealthyPercent pulumi.IntPtrInput
	// The number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
	DesiredCount pulumi.IntPtrInput
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	EnableEcsManagedTags pulumi.BoolPtrInput
	// Enable to force a new task deployment of the service. This can be used to update tasks to use a newer Docker image with same image/tag combination (e.g. `myimage:latest`), roll Fargate tasks onto a newer platform version, or immediately deploy `orderedPlacementStrategy` and `placementConstraints` updates.
	ForceNewDeployment pulumi.BoolPtrInput
	// Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
	HealthCheckGracePeriodSeconds pulumi.IntPtrInput
	// ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
	IamRole pulumi.StringPtrInput
	// The launch type on which to run your service. The valid values are `EC2` and `FARGATE`. Defaults to `EC2`.
	LaunchType pulumi.StringPtrInput
	// A load balancer block. Load balancers documented below.
	LoadBalancers ServiceLoadBalancerArrayInput
	// The name of the service (up to 255 letters, numbers, hyphens, and underscores)
	Name pulumi.StringPtrInput
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes.
	NetworkConfiguration ServiceNetworkConfigurationPtrInput
	// Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. The maximum number of `orderedPlacementStrategy` blocks is `5`. Defined below.
	OrderedPlacementStrategies ServiceOrderedPlacementStrategyArrayInput
	// rules that are taken into consideration during task placement. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. Maximum number of `placementConstraints` is `10`. Defined below.
	PlacementConstraints ServicePlacementConstraintArrayInput
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion pulumi.StringPtrInput
	// Specifies whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
	PropagateTags pulumi.StringPtrInput
	// The scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [*Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy*](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html).
	SchedulingStrategy pulumi.StringPtrInput
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`.
	ServiceRegistries ServiceServiceRegistriesPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service. Required unless using the `EXTERNAL` deployment controller. If a revision is not specified, the latest `ACTIVE` revision is used.
	TaskDefinition     pulumi.StringPtrInput
	WaitForSteadyState pulumi.BoolPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies []ServiceCapacityProviderStrategy `pulumi:"capacityProviderStrategies"`
	// ARN of an ECS cluster
	Cluster *string `pulumi:"cluster"`
	// Configuration block containing deployment controller configuration. Defined below.
	DeploymentController *ServiceDeploymentController `pulumi:"deploymentController"`
	// The upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
	DeploymentMaximumPercent *int `pulumi:"deploymentMaximumPercent"`
	// The lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
	DeploymentMinimumHealthyPercent *int `pulumi:"deploymentMinimumHealthyPercent"`
	// The number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
	DesiredCount *int `pulumi:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	EnableEcsManagedTags *bool `pulumi:"enableEcsManagedTags"`
	// Enable to force a new task deployment of the service. This can be used to update tasks to use a newer Docker image with same image/tag combination (e.g. `myimage:latest`), roll Fargate tasks onto a newer platform version, or immediately deploy `orderedPlacementStrategy` and `placementConstraints` updates.
	ForceNewDeployment *bool `pulumi:"forceNewDeployment"`
	// Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
	HealthCheckGracePeriodSeconds *int `pulumi:"healthCheckGracePeriodSeconds"`
	// ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
	IamRole *string `pulumi:"iamRole"`
	// The launch type on which to run your service. The valid values are `EC2` and `FARGATE`. Defaults to `EC2`.
	LaunchType *string `pulumi:"launchType"`
	// A load balancer block. Load balancers documented below.
	LoadBalancers []ServiceLoadBalancer `pulumi:"loadBalancers"`
	// The name of the service (up to 255 letters, numbers, hyphens, and underscores)
	Name *string `pulumi:"name"`
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes.
	NetworkConfiguration *ServiceNetworkConfiguration `pulumi:"networkConfiguration"`
	// Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. The maximum number of `orderedPlacementStrategy` blocks is `5`. Defined below.
	OrderedPlacementStrategies []ServiceOrderedPlacementStrategy `pulumi:"orderedPlacementStrategies"`
	// rules that are taken into consideration during task placement. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. Maximum number of `placementConstraints` is `10`. Defined below.
	PlacementConstraints []ServicePlacementConstraint `pulumi:"placementConstraints"`
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion *string `pulumi:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
	PropagateTags *string `pulumi:"propagateTags"`
	// The scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [*Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy*](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html).
	SchedulingStrategy *string `pulumi:"schedulingStrategy"`
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`.
	ServiceRegistries *ServiceServiceRegistries `pulumi:"serviceRegistries"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service. Required unless using the `EXTERNAL` deployment controller. If a revision is not specified, the latest `ACTIVE` revision is used.
	TaskDefinition     *string `pulumi:"taskDefinition"`
	WaitForSteadyState *bool   `pulumi:"waitForSteadyState"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies ServiceCapacityProviderStrategyArrayInput
	// ARN of an ECS cluster
	Cluster pulumi.StringPtrInput
	// Configuration block containing deployment controller configuration. Defined below.
	DeploymentController ServiceDeploymentControllerPtrInput
	// The upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
	DeploymentMaximumPercent pulumi.IntPtrInput
	// The lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
	DeploymentMinimumHealthyPercent pulumi.IntPtrInput
	// The number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
	DesiredCount pulumi.IntPtrInput
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	EnableEcsManagedTags pulumi.BoolPtrInput
	// Enable to force a new task deployment of the service. This can be used to update tasks to use a newer Docker image with same image/tag combination (e.g. `myimage:latest`), roll Fargate tasks onto a newer platform version, or immediately deploy `orderedPlacementStrategy` and `placementConstraints` updates.
	ForceNewDeployment pulumi.BoolPtrInput
	// Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
	HealthCheckGracePeriodSeconds pulumi.IntPtrInput
	// ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
	IamRole pulumi.StringPtrInput
	// The launch type on which to run your service. The valid values are `EC2` and `FARGATE`. Defaults to `EC2`.
	LaunchType pulumi.StringPtrInput
	// A load balancer block. Load balancers documented below.
	LoadBalancers ServiceLoadBalancerArrayInput
	// The name of the service (up to 255 letters, numbers, hyphens, and underscores)
	Name pulumi.StringPtrInput
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes.
	NetworkConfiguration ServiceNetworkConfigurationPtrInput
	// Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. The maximum number of `orderedPlacementStrategy` blocks is `5`. Defined below.
	OrderedPlacementStrategies ServiceOrderedPlacementStrategyArrayInput
	// rules that are taken into consideration during task placement. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. Maximum number of `placementConstraints` is `10`. Defined below.
	PlacementConstraints ServicePlacementConstraintArrayInput
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion pulumi.StringPtrInput
	// Specifies whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
	PropagateTags pulumi.StringPtrInput
	// The scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [*Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy*](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html).
	SchedulingStrategy pulumi.StringPtrInput
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`.
	ServiceRegistries ServiceServiceRegistriesPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service. Required unless using the `EXTERNAL` deployment controller. If a revision is not specified, the latest `ACTIVE` revision is used.
	TaskDefinition     pulumi.StringPtrInput
	WaitForSteadyState pulumi.BoolPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
