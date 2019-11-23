// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Route 53 Resolver endpoint resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_resolver_endpoint.html.markdown.
type ResolverEndpoint struct {
	s *pulumi.ResourceState
}

// NewResolverEndpoint registers a new resource with the given unique name, arguments, and options.
func NewResolverEndpoint(ctx *pulumi.Context,
	name string, args *ResolverEndpointArgs, opts ...pulumi.ResourceOpt) (*ResolverEndpoint, error) {
	if args == nil || args.Direction == nil {
		return nil, errors.New("missing required argument 'Direction'")
	}
	if args == nil || args.IpAddresses == nil {
		return nil, errors.New("missing required argument 'IpAddresses'")
	}
	if args == nil || args.SecurityGroupIds == nil {
		return nil, errors.New("missing required argument 'SecurityGroupIds'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["direction"] = nil
		inputs["ipAddresses"] = nil
		inputs["name"] = nil
		inputs["securityGroupIds"] = nil
		inputs["tags"] = nil
	} else {
		inputs["direction"] = args.Direction
		inputs["ipAddresses"] = args.IpAddresses
		inputs["name"] = args.Name
		inputs["securityGroupIds"] = args.SecurityGroupIds
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	inputs["hostVpcId"] = nil
	s, err := ctx.RegisterResource("aws:route53/resolverEndpoint:ResolverEndpoint", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ResolverEndpoint{s: s}, nil
}

// GetResolverEndpoint gets an existing ResolverEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverEndpoint(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ResolverEndpointState, opts ...pulumi.ResourceOpt) (*ResolverEndpoint, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["direction"] = state.Direction
		inputs["hostVpcId"] = state.HostVpcId
		inputs["ipAddresses"] = state.IpAddresses
		inputs["name"] = state.Name
		inputs["securityGroupIds"] = state.SecurityGroupIds
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:route53/resolverEndpoint:ResolverEndpoint", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ResolverEndpoint{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ResolverEndpoint) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ResolverEndpoint) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the Route 53 Resolver endpoint.
func (r *ResolverEndpoint) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The direction of DNS queries to or from the Route 53 Resolver endpoint.
// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
func (r *ResolverEndpoint) Direction() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["direction"])
}

// The ID of the VPC that you want to create the resolver endpoint in.
func (r *ResolverEndpoint) HostVpcId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hostVpcId"])
}

// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
func (r *ResolverEndpoint) IpAddresses() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ipAddresses"])
}

// The friendly name of the Route 53 Resolver endpoint.
func (r *ResolverEndpoint) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of one or more security groups that you want to use to control access to this VPC.
func (r *ResolverEndpoint) SecurityGroupIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["securityGroupIds"])
}

// A mapping of tags to assign to the resource.
func (r *ResolverEndpoint) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering ResolverEndpoint resources.
type ResolverEndpointState struct {
	// The ARN of the Route 53 Resolver endpoint.
	Arn interface{}
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction interface{}
	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVpcId interface{}
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses interface{}
	// The friendly name of the Route 53 Resolver endpoint.
	Name interface{}
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a ResolverEndpoint resource.
type ResolverEndpointArgs struct {
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction interface{}
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses interface{}
	// The friendly name of the Route 53 Resolver endpoint.
	Name interface{}
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
