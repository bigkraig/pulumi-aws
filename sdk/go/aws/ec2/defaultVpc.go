// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the [default AWS VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html)
// in the current region.
// 
// For AWS accounts created after 2013-12-04, each region comes with a Default VPC.
// **This is an advanced resource**, and has special caveats to be aware of when
// using it. Please read this document in its entirety before using this resource.
// 
// The `ec2.DefaultVpc` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead "adopts" it
// into management.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/default_vpc.html.markdown.
type DefaultVpc struct {
	s *pulumi.ResourceState
}

// NewDefaultVpc registers a new resource with the given unique name, arguments, and options.
func NewDefaultVpc(ctx *pulumi.Context,
	name string, args *DefaultVpcArgs, opts ...pulumi.ResourceOpt) (*DefaultVpc, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["enableClassiclink"] = nil
		inputs["enableClassiclinkDnsSupport"] = nil
		inputs["enableDnsHostnames"] = nil
		inputs["enableDnsSupport"] = nil
		inputs["tags"] = nil
	} else {
		inputs["enableClassiclink"] = args.EnableClassiclink
		inputs["enableClassiclinkDnsSupport"] = args.EnableClassiclinkDnsSupport
		inputs["enableDnsHostnames"] = args.EnableDnsHostnames
		inputs["enableDnsSupport"] = args.EnableDnsSupport
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	inputs["assignGeneratedIpv6CidrBlock"] = nil
	inputs["cidrBlock"] = nil
	inputs["defaultNetworkAclId"] = nil
	inputs["defaultRouteTableId"] = nil
	inputs["defaultSecurityGroupId"] = nil
	inputs["dhcpOptionsId"] = nil
	inputs["instanceTenancy"] = nil
	inputs["ipv6AssociationId"] = nil
	inputs["ipv6CidrBlock"] = nil
	inputs["mainRouteTableId"] = nil
	inputs["ownerId"] = nil
	s, err := ctx.RegisterResource("aws:ec2/defaultVpc:DefaultVpc", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultVpc{s: s}, nil
}

// GetDefaultVpc gets an existing DefaultVpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultVpc(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DefaultVpcState, opts ...pulumi.ResourceOpt) (*DefaultVpc, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["assignGeneratedIpv6CidrBlock"] = state.AssignGeneratedIpv6CidrBlock
		inputs["cidrBlock"] = state.CidrBlock
		inputs["defaultNetworkAclId"] = state.DefaultNetworkAclId
		inputs["defaultRouteTableId"] = state.DefaultRouteTableId
		inputs["defaultSecurityGroupId"] = state.DefaultSecurityGroupId
		inputs["dhcpOptionsId"] = state.DhcpOptionsId
		inputs["enableClassiclink"] = state.EnableClassiclink
		inputs["enableClassiclinkDnsSupport"] = state.EnableClassiclinkDnsSupport
		inputs["enableDnsHostnames"] = state.EnableDnsHostnames
		inputs["enableDnsSupport"] = state.EnableDnsSupport
		inputs["instanceTenancy"] = state.InstanceTenancy
		inputs["ipv6AssociationId"] = state.Ipv6AssociationId
		inputs["ipv6CidrBlock"] = state.Ipv6CidrBlock
		inputs["mainRouteTableId"] = state.MainRouteTableId
		inputs["ownerId"] = state.OwnerId
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:ec2/defaultVpc:DefaultVpc", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultVpc{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DefaultVpc) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DefaultVpc) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Amazon Resource Name (ARN) of VPC
func (r *DefaultVpc) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Whether or not an Amazon-provided IPv6 CIDR
// block with a /56 prefix length for the VPC was assigned
func (r *DefaultVpc) AssignGeneratedIpv6CidrBlock() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["assignGeneratedIpv6CidrBlock"])
}

// The CIDR block of the VPC
func (r *DefaultVpc) CidrBlock() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cidrBlock"])
}

// The ID of the network ACL created by default on VPC creation
func (r *DefaultVpc) DefaultNetworkAclId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultNetworkAclId"])
}

// The ID of the route table created by default on VPC creation
func (r *DefaultVpc) DefaultRouteTableId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultRouteTableId"])
}

// The ID of the security group created by default on VPC creation
func (r *DefaultVpc) DefaultSecurityGroupId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultSecurityGroupId"])
}

func (r *DefaultVpc) DhcpOptionsId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dhcpOptionsId"])
}

// A boolean flag to enable/disable ClassicLink
// for the VPC. Only valid in regions and accounts that support EC2 Classic.
// See the [ClassicLink documentation][1] for more information. Defaults false.
func (r *DefaultVpc) EnableClassiclink() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableClassiclink"])
}

func (r *DefaultVpc) EnableClassiclinkDnsSupport() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableClassiclinkDnsSupport"])
}

// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
func (r *DefaultVpc) EnableDnsHostnames() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableDnsHostnames"])
}

// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
func (r *DefaultVpc) EnableDnsSupport() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableDnsSupport"])
}

// Tenancy of instances spin up within VPC.
func (r *DefaultVpc) InstanceTenancy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["instanceTenancy"])
}

// The association ID for the IPv6 CIDR block of the VPC
func (r *DefaultVpc) Ipv6AssociationId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ipv6AssociationId"])
}

// The IPv6 CIDR block of the VPC
func (r *DefaultVpc) Ipv6CidrBlock() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ipv6CidrBlock"])
}

// The ID of the main route table associated with
// this VPC. Note that you can change a VPC's main route table by using an
// [`ec2.MainRouteTableAssociation`](https://www.terraform.io/docs/providers/aws/r/main_route_table_assoc.html)
func (r *DefaultVpc) MainRouteTableId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["mainRouteTableId"])
}

// The ID of the AWS account that owns the VPC.
func (r *DefaultVpc) OwnerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ownerId"])
}

// A mapping of tags to assign to the resource.
func (r *DefaultVpc) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering DefaultVpc resources.
type DefaultVpcState struct {
	// Amazon Resource Name (ARN) of VPC
	Arn interface{}
	// Whether or not an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC was assigned
	AssignGeneratedIpv6CidrBlock interface{}
	// The CIDR block of the VPC
	CidrBlock interface{}
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId interface{}
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId interface{}
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId interface{}
	DhcpOptionsId interface{}
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation][1] for more information. Defaults false.
	EnableClassiclink interface{}
	EnableClassiclinkDnsSupport interface{}
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames interface{}
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport interface{}
	// Tenancy of instances spin up within VPC.
	InstanceTenancy interface{}
	// The association ID for the IPv6 CIDR block of the VPC
	Ipv6AssociationId interface{}
	// The IPv6 CIDR block of the VPC
	Ipv6CidrBlock interface{}
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// [`ec2.MainRouteTableAssociation`](https://www.terraform.io/docs/providers/aws/r/main_route_table_assoc.html)
	MainRouteTableId interface{}
	// The ID of the AWS account that owns the VPC.
	OwnerId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a DefaultVpc resource.
type DefaultVpcArgs struct {
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation][1] for more information. Defaults false.
	EnableClassiclink interface{}
	EnableClassiclinkDnsSupport interface{}
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames interface{}
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
