// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allocates a static IP address.
// 
// ~> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
type StaticIp struct {
	s *pulumi.ResourceState
}

// NewStaticIp registers a new resource with the given unique name, arguments, and options.
func NewStaticIp(ctx *pulumi.Context,
	name string, args *StaticIpArgs, opts ...pulumi.ResourceOpt) (*StaticIp, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
	} else {
		inputs["name"] = args.Name
	}
	inputs["arn"] = nil
	inputs["ipAddress"] = nil
	inputs["supportCode"] = nil
	s, err := ctx.RegisterResource("aws:lightsail/staticIp:StaticIp", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StaticIp{s: s}, nil
}

// GetStaticIp gets an existing StaticIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticIp(ctx *pulumi.Context,
	name string, id pulumi.ID, state *StaticIpState, opts ...pulumi.ResourceOpt) (*StaticIp, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["ipAddress"] = state.IpAddress
		inputs["name"] = state.Name
		inputs["supportCode"] = state.SupportCode
	}
	s, err := ctx.ReadResource("aws:lightsail/staticIp:StaticIp", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StaticIp{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *StaticIp) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *StaticIp) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The ARN of the Lightsail static IP
func (r *StaticIp) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The allocated static IP address
func (r *StaticIp) IpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipAddress"])
}

// The name for the allocated static IP
func (r *StaticIp) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The support code.
func (r *StaticIp) SupportCode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["supportCode"])
}

// Input properties used for looking up and filtering StaticIp resources.
type StaticIpState struct {
	// The ARN of the Lightsail static IP
	Arn interface{}
	// The allocated static IP address
	IpAddress interface{}
	// The name for the allocated static IP
	Name interface{}
	// The support code.
	SupportCode interface{}
}

// The set of arguments for constructing a StaticIp resource.
type StaticIpArgs struct {
	// The name for the allocated static IP
	Name interface{}
}