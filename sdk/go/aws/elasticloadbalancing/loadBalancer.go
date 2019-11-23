// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic Load Balancer resource, also known as a "Classic
// Load Balancer" after the release of
// [Application/Network Load Balancers](https://www.terraform.io/docs/providers/aws/r/lb.html).
// 
// > **NOTE on ELB Instances and ELB Attachments:** This provider currently
// provides both a standalone ELB Attachment resource
// (describing an instance attached to an ELB), and an ELB resource with
// `instances` defined in-line. At this time you cannot use an ELB with in-line
// instances in conjunction with a ELB Attachment resources. Doing so will cause a
// conflict and will overwrite attachments.
// 
// ## Note on ECDSA Key Algorithm
// 
// If the ARN of the `sslCertificateId` that is pointed to references a
// certificate that was signed by an ECDSA key, note that ELB only supports the
// P256 and P384 curves.  Using a certificate signed by a key using a different
// curve could produce the error `ERR_SSL_VERSION_OR_CIPHER_MISMATCH` in your
// browser.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elb_legacy.html.markdown.
type LoadBalancer struct {
	s *pulumi.ResourceState
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	if args == nil || args.Listeners == nil {
		return nil, errors.New("missing required argument 'Listeners'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessLogs"] = nil
		inputs["availabilityZones"] = nil
		inputs["connectionDraining"] = nil
		inputs["connectionDrainingTimeout"] = nil
		inputs["crossZoneLoadBalancing"] = nil
		inputs["healthCheck"] = nil
		inputs["idleTimeout"] = nil
		inputs["instances"] = nil
		inputs["internal"] = nil
		inputs["listeners"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["securityGroups"] = nil
		inputs["sourceSecurityGroup"] = nil
		inputs["subnets"] = nil
		inputs["tags"] = nil
	} else {
		inputs["accessLogs"] = args.AccessLogs
		inputs["availabilityZones"] = args.AvailabilityZones
		inputs["connectionDraining"] = args.ConnectionDraining
		inputs["connectionDrainingTimeout"] = args.ConnectionDrainingTimeout
		inputs["crossZoneLoadBalancing"] = args.CrossZoneLoadBalancing
		inputs["healthCheck"] = args.HealthCheck
		inputs["idleTimeout"] = args.IdleTimeout
		inputs["instances"] = args.Instances
		inputs["internal"] = args.Internal
		inputs["listeners"] = args.Listeners
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["securityGroups"] = args.SecurityGroups
		inputs["sourceSecurityGroup"] = args.SourceSecurityGroup
		inputs["subnets"] = args.Subnets
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	inputs["dnsName"] = nil
	inputs["sourceSecurityGroupId"] = nil
	inputs["zoneId"] = nil
	s, err := ctx.RegisterResource("aws:elasticloadbalancing/loadBalancer:LoadBalancer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancer{s: s}, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LoadBalancerState, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessLogs"] = state.AccessLogs
		inputs["arn"] = state.Arn
		inputs["availabilityZones"] = state.AvailabilityZones
		inputs["connectionDraining"] = state.ConnectionDraining
		inputs["connectionDrainingTimeout"] = state.ConnectionDrainingTimeout
		inputs["crossZoneLoadBalancing"] = state.CrossZoneLoadBalancing
		inputs["dnsName"] = state.DnsName
		inputs["healthCheck"] = state.HealthCheck
		inputs["idleTimeout"] = state.IdleTimeout
		inputs["instances"] = state.Instances
		inputs["internal"] = state.Internal
		inputs["listeners"] = state.Listeners
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["securityGroups"] = state.SecurityGroups
		inputs["sourceSecurityGroup"] = state.SourceSecurityGroup
		inputs["sourceSecurityGroupId"] = state.SourceSecurityGroupId
		inputs["subnets"] = state.Subnets
		inputs["tags"] = state.Tags
		inputs["zoneId"] = state.ZoneId
	}
	s, err := ctx.ReadResource("aws:elasticloadbalancing/loadBalancer:LoadBalancer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LoadBalancer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LoadBalancer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// An Access Logs block. Access Logs documented below.
func (r *LoadBalancer) AccessLogs() pulumi.Output {
	return r.s.State["accessLogs"]
}

// The ARN of the ELB
func (r *LoadBalancer) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The AZ's to serve traffic in.
func (r *LoadBalancer) AvailabilityZones() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["availabilityZones"])
}

// Boolean to enable connection draining. Default: `false`
func (r *LoadBalancer) ConnectionDraining() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["connectionDraining"])
}

// The time in seconds to allow for connections to drain. Default: `300`
func (r *LoadBalancer) ConnectionDrainingTimeout() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["connectionDrainingTimeout"])
}

// Enable cross-zone load balancing. Default: `true`
func (r *LoadBalancer) CrossZoneLoadBalancing() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["crossZoneLoadBalancing"])
}

// The DNS name of the ELB
func (r *LoadBalancer) DnsName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dnsName"])
}

// A healthCheck block. Health Check documented below.
func (r *LoadBalancer) HealthCheck() pulumi.Output {
	return r.s.State["healthCheck"]
}

// The time in seconds that the connection is allowed to be idle. Default: `60`
func (r *LoadBalancer) IdleTimeout() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["idleTimeout"])
}

// A list of instance ids to place in the ELB pool.
func (r *LoadBalancer) Instances() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["instances"])
}

// If true, ELB will be an internal ELB.
func (r *LoadBalancer) Internal() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["internal"])
}

// A list of listener blocks. Listeners documented below.
func (r *LoadBalancer) Listeners() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["listeners"])
}

// The name of the ELB. By default generated by this provider.
func (r *LoadBalancer) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Creates a unique name beginning with the specified
// prefix. Conflicts with `name`.
func (r *LoadBalancer) NamePrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namePrefix"])
}

// A list of security group IDs to assign to the ELB.
// Only valid if creating an ELB within a VPC
func (r *LoadBalancer) SecurityGroups() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["securityGroups"])
}

// The name of the security group that you can use as
// part of your inbound rules for your load balancer's back-end application
// instances. Use this for Classic or Default VPC only.
func (r *LoadBalancer) SourceSecurityGroup() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceSecurityGroup"])
}

// The ID of the security group that you can use as
// part of your inbound rules for your load balancer's back-end application
// instances. Only available on ELBs launched in a VPC.
func (r *LoadBalancer) SourceSecurityGroupId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceSecurityGroupId"])
}

// A list of subnet IDs to attach to the ELB.
func (r *LoadBalancer) Subnets() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["subnets"])
}

// A mapping of tags to assign to the resource.
func (r *LoadBalancer) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The canonical hosted zone ID of the ELB (to be used in a Route 53 Alias record)
func (r *LoadBalancer) ZoneId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["zoneId"])
}

// Input properties used for looking up and filtering LoadBalancer resources.
type LoadBalancerState struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs interface{}
	// The ARN of the ELB
	Arn interface{}
	// The AZ's to serve traffic in.
	AvailabilityZones interface{}
	// Boolean to enable connection draining. Default: `false`
	ConnectionDraining interface{}
	// The time in seconds to allow for connections to drain. Default: `300`
	ConnectionDrainingTimeout interface{}
	// Enable cross-zone load balancing. Default: `true`
	CrossZoneLoadBalancing interface{}
	// The DNS name of the ELB
	DnsName interface{}
	// A healthCheck block. Health Check documented below.
	HealthCheck interface{}
	// The time in seconds that the connection is allowed to be idle. Default: `60`
	IdleTimeout interface{}
	// A list of instance ids to place in the ELB pool.
	Instances interface{}
	// If true, ELB will be an internal ELB.
	Internal interface{}
	// A list of listener blocks. Listeners documented below.
	Listeners interface{}
	// The name of the ELB. By default generated by this provider.
	Name interface{}
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix interface{}
	// A list of security group IDs to assign to the ELB.
	// Only valid if creating an ELB within a VPC
	SecurityGroups interface{}
	// The name of the security group that you can use as
	// part of your inbound rules for your load balancer's back-end application
	// instances. Use this for Classic or Default VPC only.
	SourceSecurityGroup interface{}
	// The ID of the security group that you can use as
	// part of your inbound rules for your load balancer's back-end application
	// instances. Only available on ELBs launched in a VPC.
	SourceSecurityGroupId interface{}
	// A list of subnet IDs to attach to the ELB.
	Subnets interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The canonical hosted zone ID of the ELB (to be used in a Route 53 Alias record)
	ZoneId interface{}
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs interface{}
	// The AZ's to serve traffic in.
	AvailabilityZones interface{}
	// Boolean to enable connection draining. Default: `false`
	ConnectionDraining interface{}
	// The time in seconds to allow for connections to drain. Default: `300`
	ConnectionDrainingTimeout interface{}
	// Enable cross-zone load balancing. Default: `true`
	CrossZoneLoadBalancing interface{}
	// A healthCheck block. Health Check documented below.
	HealthCheck interface{}
	// The time in seconds that the connection is allowed to be idle. Default: `60`
	IdleTimeout interface{}
	// A list of instance ids to place in the ELB pool.
	Instances interface{}
	// If true, ELB will be an internal ELB.
	Internal interface{}
	// A list of listener blocks. Listeners documented below.
	Listeners interface{}
	// The name of the ELB. By default generated by this provider.
	Name interface{}
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix interface{}
	// A list of security group IDs to assign to the ELB.
	// Only valid if creating an ELB within a VPC
	SecurityGroups interface{}
	// The name of the security group that you can use as
	// part of your inbound rules for your load balancer's back-end application
	// instances. Use this for Classic or Default VPC only.
	SourceSecurityGroup interface{}
	// A list of subnet IDs to attach to the ELB.
	Subnets interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
