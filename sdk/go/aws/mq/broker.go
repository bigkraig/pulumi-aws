// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mq

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an MQ Broker Resource. This resources also manages users for the broker.
// 
// For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).
// 
// Changes to an MQ Broker can occur when you change a
// parameter, such as `configuration` or `user`, and are reflected in the next maintenance
// window. Because of this, this provider may report a difference in its planning
// phase because a modification has not yet taken place. You can use the
// `applyImmediately` flag to instruct the service to apply the change immediately
// (see documentation below).
// 
// > **Note:** using `applyImmediately` can result in a
// brief downtime as the broker reboots.
// 
// > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/mq_broker.html.markdown.
type Broker struct {
	s *pulumi.ResourceState
}

// NewBroker registers a new resource with the given unique name, arguments, and options.
func NewBroker(ctx *pulumi.Context,
	name string, args *BrokerArgs, opts ...pulumi.ResourceOpt) (*Broker, error) {
	if args == nil || args.BrokerName == nil {
		return nil, errors.New("missing required argument 'BrokerName'")
	}
	if args == nil || args.EngineType == nil {
		return nil, errors.New("missing required argument 'EngineType'")
	}
	if args == nil || args.EngineVersion == nil {
		return nil, errors.New("missing required argument 'EngineVersion'")
	}
	if args == nil || args.HostInstanceType == nil {
		return nil, errors.New("missing required argument 'HostInstanceType'")
	}
	if args == nil || args.SecurityGroups == nil {
		return nil, errors.New("missing required argument 'SecurityGroups'")
	}
	if args == nil || args.Users == nil {
		return nil, errors.New("missing required argument 'Users'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["applyImmediately"] = nil
		inputs["autoMinorVersionUpgrade"] = nil
		inputs["brokerName"] = nil
		inputs["configuration"] = nil
		inputs["deploymentMode"] = nil
		inputs["encryptionOptions"] = nil
		inputs["engineType"] = nil
		inputs["engineVersion"] = nil
		inputs["hostInstanceType"] = nil
		inputs["logs"] = nil
		inputs["maintenanceWindowStartTime"] = nil
		inputs["publiclyAccessible"] = nil
		inputs["securityGroups"] = nil
		inputs["subnetIds"] = nil
		inputs["tags"] = nil
		inputs["users"] = nil
	} else {
		inputs["applyImmediately"] = args.ApplyImmediately
		inputs["autoMinorVersionUpgrade"] = args.AutoMinorVersionUpgrade
		inputs["brokerName"] = args.BrokerName
		inputs["configuration"] = args.Configuration
		inputs["deploymentMode"] = args.DeploymentMode
		inputs["encryptionOptions"] = args.EncryptionOptions
		inputs["engineType"] = args.EngineType
		inputs["engineVersion"] = args.EngineVersion
		inputs["hostInstanceType"] = args.HostInstanceType
		inputs["logs"] = args.Logs
		inputs["maintenanceWindowStartTime"] = args.MaintenanceWindowStartTime
		inputs["publiclyAccessible"] = args.PubliclyAccessible
		inputs["securityGroups"] = args.SecurityGroups
		inputs["subnetIds"] = args.SubnetIds
		inputs["tags"] = args.Tags
		inputs["users"] = args.Users
	}
	inputs["arn"] = nil
	inputs["instances"] = nil
	s, err := ctx.RegisterResource("aws:mq/broker:Broker", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Broker{s: s}, nil
}

// GetBroker gets an existing Broker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBroker(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BrokerState, opts ...pulumi.ResourceOpt) (*Broker, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["applyImmediately"] = state.ApplyImmediately
		inputs["arn"] = state.Arn
		inputs["autoMinorVersionUpgrade"] = state.AutoMinorVersionUpgrade
		inputs["brokerName"] = state.BrokerName
		inputs["configuration"] = state.Configuration
		inputs["deploymentMode"] = state.DeploymentMode
		inputs["encryptionOptions"] = state.EncryptionOptions
		inputs["engineType"] = state.EngineType
		inputs["engineVersion"] = state.EngineVersion
		inputs["hostInstanceType"] = state.HostInstanceType
		inputs["instances"] = state.Instances
		inputs["logs"] = state.Logs
		inputs["maintenanceWindowStartTime"] = state.MaintenanceWindowStartTime
		inputs["publiclyAccessible"] = state.PubliclyAccessible
		inputs["securityGroups"] = state.SecurityGroups
		inputs["subnetIds"] = state.SubnetIds
		inputs["tags"] = state.Tags
		inputs["users"] = state.Users
	}
	s, err := ctx.ReadResource("aws:mq/broker:Broker", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Broker{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Broker) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Broker) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies whether any broker modifications
// are applied immediately, or during the next maintenance window. Default is `false`.
func (r *Broker) ApplyImmediately() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["applyImmediately"])
}

// The ARN of the broker.
func (r *Broker) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
func (r *Broker) AutoMinorVersionUpgrade() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["autoMinorVersionUpgrade"])
}

// The name of the broker.
func (r *Broker) BrokerName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["brokerName"])
}

// Configuration of the broker. See below.
func (r *Broker) Configuration() pulumi.Output {
	return r.s.State["configuration"]
}

// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
func (r *Broker) DeploymentMode() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["deploymentMode"])
}

// Configuration block containing encryption options. See below.
func (r *Broker) EncryptionOptions() pulumi.Output {
	return r.s.State["encryptionOptions"]
}

// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
func (r *Broker) EngineType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engineType"])
}

// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
func (r *Broker) EngineVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engineVersion"])
}

// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
func (r *Broker) HostInstanceType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hostInstanceType"])
}

// A list of information about allocated brokers (both active & standby).
// * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
// * `instances.0.ip_address` - The IP Address of the broker.
// * `instances.0.endpoints` - The broker's wire-level protocol endpoints in the following order & format referenceable e.g. as `instances.0.endpoints.0` (SSL):
// * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
// * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
// * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
// * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
// * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
func (r *Broker) Instances() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["instances"])
}

// Logging configuration of the broker. See below.
func (r *Broker) Logs() pulumi.Output {
	return r.s.State["logs"]
}

// Maintenance window start time. See below.
func (r *Broker) MaintenanceWindowStartTime() pulumi.Output {
	return r.s.State["maintenanceWindowStartTime"]
}

// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
func (r *Broker) PubliclyAccessible() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["publiclyAccessible"])
}

// The list of security group IDs assigned to the broker.
func (r *Broker) SecurityGroups() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["securityGroups"])
}

// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
func (r *Broker) SubnetIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["subnetIds"])
}

// A mapping of tags to assign to the resource.
func (r *Broker) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The list of all ActiveMQ usernames for the specified broker. See below.
func (r *Broker) Users() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering Broker resources.
type BrokerState struct {
	// Specifies whether any broker modifications
	// are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately interface{}
	// The ARN of the broker.
	Arn interface{}
	// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
	AutoMinorVersionUpgrade interface{}
	// The name of the broker.
	BrokerName interface{}
	// Configuration of the broker. See below.
	Configuration interface{}
	// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
	DeploymentMode interface{}
	// Configuration block containing encryption options. See below.
	EncryptionOptions interface{}
	// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
	EngineType interface{}
	// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
	EngineVersion interface{}
	// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
	HostInstanceType interface{}
	// A list of information about allocated brokers (both active & standby).
	// * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
	// * `instances.0.ip_address` - The IP Address of the broker.
	// * `instances.0.endpoints` - The broker's wire-level protocol endpoints in the following order & format referenceable e.g. as `instances.0.endpoints.0` (SSL):
	// * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
	// * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
	// * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
	// * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
	// * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
	Instances interface{}
	// Logging configuration of the broker. See below.
	Logs interface{}
	// Maintenance window start time. See below.
	MaintenanceWindowStartTime interface{}
	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible interface{}
	// The list of security group IDs assigned to the broker.
	SecurityGroups interface{}
	// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
	SubnetIds interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The list of all ActiveMQ usernames for the specified broker. See below.
	Users interface{}
}

// The set of arguments for constructing a Broker resource.
type BrokerArgs struct {
	// Specifies whether any broker modifications
	// are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately interface{}
	// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
	AutoMinorVersionUpgrade interface{}
	// The name of the broker.
	BrokerName interface{}
	// Configuration of the broker. See below.
	Configuration interface{}
	// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
	DeploymentMode interface{}
	// Configuration block containing encryption options. See below.
	EncryptionOptions interface{}
	// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
	EngineType interface{}
	// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
	EngineVersion interface{}
	// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
	HostInstanceType interface{}
	// Logging configuration of the broker. See below.
	Logs interface{}
	// Maintenance window start time. See below.
	MaintenanceWindowStartTime interface{}
	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible interface{}
	// The list of security group IDs assigned to the broker.
	SecurityGroups interface{}
	// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
	SubnetIds interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The list of all ActiveMQ usernames for the specified broker. See below.
	Users interface{}
}
