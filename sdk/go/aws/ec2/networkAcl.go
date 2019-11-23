// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an network ACL resource. You might set up network ACLs with rules similar
// to your security groups in order to add an additional layer of security to your VPC.
// 
// > **NOTE on Network ACLs and Network ACL Rules:** This provider currently
// provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
// defined in-line. At this time you cannot use a Network ACL with in-line rules
// in conjunction with any Network ACL Rule resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/network_acl.html.markdown.
type NetworkAcl struct {
	s *pulumi.ResourceState
}

// NewNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewNetworkAcl(ctx *pulumi.Context,
	name string, args *NetworkAclArgs, opts ...pulumi.ResourceOpt) (*NetworkAcl, error) {
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["egress"] = nil
		inputs["ingress"] = nil
		inputs["subnetIds"] = nil
		inputs["tags"] = nil
		inputs["vpcId"] = nil
	} else {
		inputs["egress"] = args.Egress
		inputs["ingress"] = args.Ingress
		inputs["subnetIds"] = args.SubnetIds
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	inputs["ownerId"] = nil
	s, err := ctx.RegisterResource("aws:ec2/networkAcl:NetworkAcl", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkAcl{s: s}, nil
}

// GetNetworkAcl gets an existing NetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NetworkAclState, opts ...pulumi.ResourceOpt) (*NetworkAcl, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["egress"] = state.Egress
		inputs["ingress"] = state.Ingress
		inputs["ownerId"] = state.OwnerId
		inputs["subnetIds"] = state.SubnetIds
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:ec2/networkAcl:NetworkAcl", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkAcl{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NetworkAcl) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NetworkAcl) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies an egress rule. Parameters defined below.
// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
func (r *NetworkAcl) Egress() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["egress"])
}

// Specifies an ingress rule. Parameters defined below.
// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
func (r *NetworkAcl) Ingress() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ingress"])
}

// The ID of the AWS account that owns the network ACL.
func (r *NetworkAcl) OwnerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ownerId"])
}

// A list of Subnet IDs to apply the ACL to
func (r *NetworkAcl) SubnetIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["subnetIds"])
}

// A mapping of tags to assign to the resource.
func (r *NetworkAcl) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The ID of the associated VPC.
func (r *NetworkAcl) VpcId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering NetworkAcl resources.
type NetworkAclState struct {
	// Specifies an egress rule. Parameters defined below.
	// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Egress interface{}
	// Specifies an ingress rule. Parameters defined below.
	// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Ingress interface{}
	// The ID of the AWS account that owns the network ACL.
	OwnerId interface{}
	// A list of Subnet IDs to apply the ACL to
	SubnetIds interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The ID of the associated VPC.
	VpcId interface{}
}

// The set of arguments for constructing a NetworkAcl resource.
type NetworkAclArgs struct {
	// Specifies an egress rule. Parameters defined below.
	// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Egress interface{}
	// Specifies an ingress rule. Parameters defined below.
	// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Ingress interface{}
	// A list of Subnet IDs to apply the ACL to
	SubnetIds interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The ID of the associated VPC.
	VpcId interface{}
}
