// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Direct Connect BGP peer resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_bgp_peer.html.markdown.
type BgpPeer struct {
	s *pulumi.ResourceState
}

// NewBgpPeer registers a new resource with the given unique name, arguments, and options.
func NewBgpPeer(ctx *pulumi.Context,
	name string, args *BgpPeerArgs, opts ...pulumi.ResourceOpt) (*BgpPeer, error) {
	if args == nil || args.AddressFamily == nil {
		return nil, errors.New("missing required argument 'AddressFamily'")
	}
	if args == nil || args.BgpAsn == nil {
		return nil, errors.New("missing required argument 'BgpAsn'")
	}
	if args == nil || args.VirtualInterfaceId == nil {
		return nil, errors.New("missing required argument 'VirtualInterfaceId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addressFamily"] = nil
		inputs["amazonAddress"] = nil
		inputs["bgpAsn"] = nil
		inputs["bgpAuthKey"] = nil
		inputs["customerAddress"] = nil
		inputs["virtualInterfaceId"] = nil
	} else {
		inputs["addressFamily"] = args.AddressFamily
		inputs["amazonAddress"] = args.AmazonAddress
		inputs["bgpAsn"] = args.BgpAsn
		inputs["bgpAuthKey"] = args.BgpAuthKey
		inputs["customerAddress"] = args.CustomerAddress
		inputs["virtualInterfaceId"] = args.VirtualInterfaceId
	}
	inputs["awsDevice"] = nil
	inputs["bgpPeerId"] = nil
	inputs["bgpStatus"] = nil
	s, err := ctx.RegisterResource("aws:directconnect/bgpPeer:BgpPeer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BgpPeer{s: s}, nil
}

// GetBgpPeer gets an existing BgpPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpPeer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BgpPeerState, opts ...pulumi.ResourceOpt) (*BgpPeer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addressFamily"] = state.AddressFamily
		inputs["amazonAddress"] = state.AmazonAddress
		inputs["awsDevice"] = state.AwsDevice
		inputs["bgpAsn"] = state.BgpAsn
		inputs["bgpAuthKey"] = state.BgpAuthKey
		inputs["bgpPeerId"] = state.BgpPeerId
		inputs["bgpStatus"] = state.BgpStatus
		inputs["customerAddress"] = state.CustomerAddress
		inputs["virtualInterfaceId"] = state.VirtualInterfaceId
	}
	s, err := ctx.ReadResource("aws:directconnect/bgpPeer:BgpPeer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BgpPeer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *BgpPeer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *BgpPeer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The address family for the BGP peer. `ipv4 ` or `ipv6`.
func (r *BgpPeer) AddressFamily() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["addressFamily"])
}

// The IPv4 CIDR address to use to send traffic to Amazon.
// Required for IPv4 BGP peers on public virtual interfaces.
func (r *BgpPeer) AmazonAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["amazonAddress"])
}

// The Direct Connect endpoint on which the BGP peer terminates.
func (r *BgpPeer) AwsDevice() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["awsDevice"])
}

// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
func (r *BgpPeer) BgpAsn() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["bgpAsn"])
}

// The authentication key for BGP configuration.
func (r *BgpPeer) BgpAuthKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bgpAuthKey"])
}

// The ID of the BGP peer.
func (r *BgpPeer) BgpPeerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bgpPeerId"])
}

// The Up/Down state of the BGP peer.
func (r *BgpPeer) BgpStatus() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bgpStatus"])
}

// The IPv4 CIDR destination address to which Amazon should send traffic.
// Required for IPv4 BGP peers on public virtual interfaces.
func (r *BgpPeer) CustomerAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["customerAddress"])
}

// The ID of the Direct Connect virtual interface on which to create the BGP peer.
func (r *BgpPeer) VirtualInterfaceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["virtualInterfaceId"])
}

// Input properties used for looking up and filtering BgpPeer resources.
type BgpPeerState struct {
	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily interface{}
	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	AmazonAddress interface{}
	// The Direct Connect endpoint on which the BGP peer terminates.
	AwsDevice interface{}
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn interface{}
	// The authentication key for BGP configuration.
	BgpAuthKey interface{}
	// The ID of the BGP peer.
	BgpPeerId interface{}
	// The Up/Down state of the BGP peer.
	BgpStatus interface{}
	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	CustomerAddress interface{}
	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	VirtualInterfaceId interface{}
}

// The set of arguments for constructing a BgpPeer resource.
type BgpPeerArgs struct {
	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily interface{}
	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	AmazonAddress interface{}
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn interface{}
	// The authentication key for BGP configuration.
	BgpAuthKey interface{}
	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	CustomerAddress interface{}
	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	VirtualInterfaceId interface{}
}
