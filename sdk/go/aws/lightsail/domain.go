// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a domain resource for the specified domain (e.g., example.com).
// You cannot register a new domain name using Lightsail. You must register
// a domain name using Amazon Route 53 or another domain name registrar.
// If you have already registered your domain, you can enter its name in
// this parameter to manage the DNS records for that domain.
// 
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lightsail_domain.html.markdown.
type Domain struct {
	s *pulumi.ResourceState
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOpt) (*Domain, error) {
	if args == nil || args.DomainName == nil {
		return nil, errors.New("missing required argument 'DomainName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["domainName"] = nil
	} else {
		inputs["domainName"] = args.DomainName
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:lightsail/domain:Domain", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DomainState, opts ...pulumi.ResourceOpt) (*Domain, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["domainName"] = state.DomainName
	}
	s, err := ctx.ReadResource("aws:lightsail/domain:Domain", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Domain) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Domain) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the Lightsail domain
func (r *Domain) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The name of the Lightsail domain to manage
func (r *Domain) DomainName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["domainName"])
}

// Input properties used for looking up and filtering Domain resources.
type DomainState struct {
	// The ARN of the Lightsail domain
	Arn interface{}
	// The name of the Lightsail domain to manage
	DomainName interface{}
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The name of the Lightsail domain to manage
	DomainName interface{}
}
