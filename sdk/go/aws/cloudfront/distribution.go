// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Amazon CloudFront web distribution.
// 
// For information about CloudFront distributions, see the
// [Amazon CloudFront Developer Guide][1]. For specific information about creating
// CloudFront web distributions, see the [POST Distribution][2] page in the Amazon
// CloudFront API Reference.
// 
// > **NOTE:** CloudFront distributions take about 15 minutes to a deployed state
// after creation or modification. During this time, deletes to resources will be
// blocked. If you need to delete a distribution that is enabled and you do not
// want to wait, you need to use the `retainOnDelete` flag.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudfront_distribution.html.markdown.
type Distribution struct {
	s *pulumi.ResourceState
}

// NewDistribution registers a new resource with the given unique name, arguments, and options.
func NewDistribution(ctx *pulumi.Context,
	name string, args *DistributionArgs, opts ...pulumi.ResourceOpt) (*Distribution, error) {
	if args == nil || args.DefaultCacheBehavior == nil {
		return nil, errors.New("missing required argument 'DefaultCacheBehavior'")
	}
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.Origins == nil {
		return nil, errors.New("missing required argument 'Origins'")
	}
	if args == nil || args.Restrictions == nil {
		return nil, errors.New("missing required argument 'Restrictions'")
	}
	if args == nil || args.ViewerCertificate == nil {
		return nil, errors.New("missing required argument 'ViewerCertificate'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["aliases"] = nil
		inputs["comment"] = nil
		inputs["customErrorResponses"] = nil
		inputs["defaultCacheBehavior"] = nil
		inputs["defaultRootObject"] = nil
		inputs["enabled"] = nil
		inputs["httpVersion"] = nil
		inputs["isIpv6Enabled"] = nil
		inputs["loggingConfig"] = nil
		inputs["orderedCacheBehaviors"] = nil
		inputs["origins"] = nil
		inputs["originGroups"] = nil
		inputs["priceClass"] = nil
		inputs["restrictions"] = nil
		inputs["retainOnDelete"] = nil
		inputs["tags"] = nil
		inputs["viewerCertificate"] = nil
		inputs["waitForDeployment"] = nil
		inputs["webAclId"] = nil
	} else {
		inputs["aliases"] = args.Aliases
		inputs["comment"] = args.Comment
		inputs["customErrorResponses"] = args.CustomErrorResponses
		inputs["defaultCacheBehavior"] = args.DefaultCacheBehavior
		inputs["defaultRootObject"] = args.DefaultRootObject
		inputs["enabled"] = args.Enabled
		inputs["httpVersion"] = args.HttpVersion
		inputs["isIpv6Enabled"] = args.IsIpv6Enabled
		inputs["loggingConfig"] = args.LoggingConfig
		inputs["orderedCacheBehaviors"] = args.OrderedCacheBehaviors
		inputs["origins"] = args.Origins
		inputs["originGroups"] = args.OriginGroups
		inputs["priceClass"] = args.PriceClass
		inputs["restrictions"] = args.Restrictions
		inputs["retainOnDelete"] = args.RetainOnDelete
		inputs["tags"] = args.Tags
		inputs["viewerCertificate"] = args.ViewerCertificate
		inputs["waitForDeployment"] = args.WaitForDeployment
		inputs["webAclId"] = args.WebAclId
	}
	inputs["activeTrustedSigners"] = nil
	inputs["arn"] = nil
	inputs["callerReference"] = nil
	inputs["domainName"] = nil
	inputs["etag"] = nil
	inputs["hostedZoneId"] = nil
	inputs["inProgressValidationBatches"] = nil
	inputs["lastModifiedTime"] = nil
	inputs["status"] = nil
	s, err := ctx.RegisterResource("aws:cloudfront/distribution:Distribution", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Distribution{s: s}, nil
}

// GetDistribution gets an existing Distribution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDistribution(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DistributionState, opts ...pulumi.ResourceOpt) (*Distribution, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["activeTrustedSigners"] = state.ActiveTrustedSigners
		inputs["aliases"] = state.Aliases
		inputs["arn"] = state.Arn
		inputs["callerReference"] = state.CallerReference
		inputs["comment"] = state.Comment
		inputs["customErrorResponses"] = state.CustomErrorResponses
		inputs["defaultCacheBehavior"] = state.DefaultCacheBehavior
		inputs["defaultRootObject"] = state.DefaultRootObject
		inputs["domainName"] = state.DomainName
		inputs["enabled"] = state.Enabled
		inputs["etag"] = state.Etag
		inputs["hostedZoneId"] = state.HostedZoneId
		inputs["httpVersion"] = state.HttpVersion
		inputs["inProgressValidationBatches"] = state.InProgressValidationBatches
		inputs["isIpv6Enabled"] = state.IsIpv6Enabled
		inputs["lastModifiedTime"] = state.LastModifiedTime
		inputs["loggingConfig"] = state.LoggingConfig
		inputs["orderedCacheBehaviors"] = state.OrderedCacheBehaviors
		inputs["origins"] = state.Origins
		inputs["originGroups"] = state.OriginGroups
		inputs["priceClass"] = state.PriceClass
		inputs["restrictions"] = state.Restrictions
		inputs["retainOnDelete"] = state.RetainOnDelete
		inputs["status"] = state.Status
		inputs["tags"] = state.Tags
		inputs["viewerCertificate"] = state.ViewerCertificate
		inputs["waitForDeployment"] = state.WaitForDeployment
		inputs["webAclId"] = state.WebAclId
	}
	s, err := ctx.ReadResource("aws:cloudfront/distribution:Distribution", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Distribution{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Distribution) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Distribution) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The key pair IDs that CloudFront is aware of for
// each trusted signer, if the distribution is set up to serve private content
// with signed URLs.
func (r *Distribution) ActiveTrustedSigners() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["activeTrustedSigners"])
}

// Extra CNAMEs (alternate domain names), if any, for
// this distribution.
func (r *Distribution) Aliases() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["aliases"])
}

// The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.
func (r *Distribution) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Internal value used by CloudFront to allow future
// updates to the distribution configuration.
func (r *Distribution) CallerReference() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["callerReference"])
}

// Any comments you want to include about the
// distribution.
func (r *Distribution) Comment() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["comment"])
}

// One or more custom error response elements (multiples allowed).
func (r *Distribution) CustomErrorResponses() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customErrorResponses"])
}

// The default cache behavior for this distribution (maximum
// one).
func (r *Distribution) DefaultCacheBehavior() pulumi.Output {
	return r.s.State["defaultCacheBehavior"]
}

// The object that you want CloudFront to
// return (for example, index.html) when an end user requests the root URL.
func (r *Distribution) DefaultRootObject() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultRootObject"])
}

// The DNS domain name of either the S3 bucket, or
// web site of your custom origin.
func (r *Distribution) DomainName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["domainName"])
}

// Whether the distribution is enabled to accept end
// user requests for content.
func (r *Distribution) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// The current version of the distribution's information. For example:
// `E2QWRUHAPOMQZL`.
func (r *Distribution) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// The CloudFront Route 53 zone ID that can be used to
// route an [Alias Resource Record Set][7] to. This attribute is simply an
// alias for the zone ID `Z2FDTNDATAQYW2`.
func (r *Distribution) HostedZoneId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hostedZoneId"])
}

// The maximum HTTP version to support on the
// distribution. Allowed values are `http1.1` and `http2`. The default is
// `http2`.
func (r *Distribution) HttpVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["httpVersion"])
}

// The number of invalidation batches
// currently in progress.
func (r *Distribution) InProgressValidationBatches() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["inProgressValidationBatches"])
}

// Whether the IPv6 is enabled for the distribution.
func (r *Distribution) IsIpv6Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["isIpv6Enabled"])
}

// The date and time the distribution was last modified.
func (r *Distribution) LastModifiedTime() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["lastModifiedTime"])
}

// The logging
// configuration that controls how logs are written
// to your distribution (maximum one).
func (r *Distribution) LoggingConfig() pulumi.Output {
	return r.s.State["loggingConfig"]
}

// An ordered list of cache behaviors
// resource for this distribution. List from top to bottom
// in order of precedence. The topmost cache behavior will have precedence 0.
func (r *Distribution) OrderedCacheBehaviors() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["orderedCacheBehaviors"])
}

// One or more origins for this
// distribution (multiples allowed).
func (r *Distribution) Origins() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["origins"])
}

// One or more originGroup for this
// distribution (multiples allowed).
func (r *Distribution) OriginGroups() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["originGroups"])
}

// The price class for this distribution. One of
// `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
func (r *Distribution) PriceClass() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["priceClass"])
}

// The restriction
// configuration for this distribution (maximum one).
func (r *Distribution) Restrictions() pulumi.Output {
	return r.s.State["restrictions"]
}

// Disables the distribution instead of
// deleting it when destroying the resource. If this is set,
// the distribution needs to be deleted manually afterwards. Default: `false`.
func (r *Distribution) RetainOnDelete() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["retainOnDelete"])
}

// The current status of the distribution. `Deployed` if the
// distribution's information is fully propagated throughout the Amazon
// CloudFront system.
func (r *Distribution) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

// A mapping of tags to assign to the resource.
func (r *Distribution) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The SSL
// configuration for this distribution (maximum
// one).
func (r *Distribution) ViewerCertificate() pulumi.Output {
	return r.s.State["viewerCertificate"]
}

// If enabled, the resource will wait for
// the distribution status to change from `InProgress` to `Deployed`. Setting
// this to`false` will skip the process. Default: `true`.
func (r *Distribution) WaitForDeployment() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["waitForDeployment"])
}

// If you're using AWS WAF to filter CloudFront
// requests, the Id of the AWS WAF web ACL that is associated with the
// distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
// region and the credentials configuring this argument must have
// `waf:GetWebACL` permissions assigned.
func (r *Distribution) WebAclId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["webAclId"])
}

// Input properties used for looking up and filtering Distribution resources.
type DistributionState struct {
	// The key pair IDs that CloudFront is aware of for
	// each trusted signer, if the distribution is set up to serve private content
	// with signed URLs.
	ActiveTrustedSigners interface{}
	// Extra CNAMEs (alternate domain names), if any, for
	// this distribution.
	Aliases interface{}
	// The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.
	Arn interface{}
	// Internal value used by CloudFront to allow future
	// updates to the distribution configuration.
	CallerReference interface{}
	// Any comments you want to include about the
	// distribution.
	Comment interface{}
	// One or more custom error response elements (multiples allowed).
	CustomErrorResponses interface{}
	// The default cache behavior for this distribution (maximum
	// one).
	DefaultCacheBehavior interface{}
	// The object that you want CloudFront to
	// return (for example, index.html) when an end user requests the root URL.
	DefaultRootObject interface{}
	// The DNS domain name of either the S3 bucket, or
	// web site of your custom origin.
	DomainName interface{}
	// Whether the distribution is enabled to accept end
	// user requests for content.
	Enabled interface{}
	// The current version of the distribution's information. For example:
	// `E2QWRUHAPOMQZL`.
	Etag interface{}
	// The CloudFront Route 53 zone ID that can be used to
	// route an [Alias Resource Record Set][7] to. This attribute is simply an
	// alias for the zone ID `Z2FDTNDATAQYW2`.
	HostedZoneId interface{}
	// The maximum HTTP version to support on the
	// distribution. Allowed values are `http1.1` and `http2`. The default is
	// `http2`.
	HttpVersion interface{}
	// The number of invalidation batches
	// currently in progress.
	InProgressValidationBatches interface{}
	// Whether the IPv6 is enabled for the distribution.
	IsIpv6Enabled interface{}
	// The date and time the distribution was last modified.
	LastModifiedTime interface{}
	// The logging
	// configuration that controls how logs are written
	// to your distribution (maximum one).
	LoggingConfig interface{}
	// An ordered list of cache behaviors
	// resource for this distribution. List from top to bottom
	// in order of precedence. The topmost cache behavior will have precedence 0.
	OrderedCacheBehaviors interface{}
	// One or more origins for this
	// distribution (multiples allowed).
	Origins interface{}
	// One or more originGroup for this
	// distribution (multiples allowed).
	OriginGroups interface{}
	// The price class for this distribution. One of
	// `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
	PriceClass interface{}
	// The restriction
	// configuration for this distribution (maximum one).
	Restrictions interface{}
	// Disables the distribution instead of
	// deleting it when destroying the resource. If this is set,
	// the distribution needs to be deleted manually afterwards. Default: `false`.
	RetainOnDelete interface{}
	// The current status of the distribution. `Deployed` if the
	// distribution's information is fully propagated throughout the Amazon
	// CloudFront system.
	Status interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The SSL
	// configuration for this distribution (maximum
	// one).
	ViewerCertificate interface{}
	// If enabled, the resource will wait for
	// the distribution status to change from `InProgress` to `Deployed`. Setting
	// this to`false` will skip the process. Default: `true`.
	WaitForDeployment interface{}
	// If you're using AWS WAF to filter CloudFront
	// requests, the Id of the AWS WAF web ACL that is associated with the
	// distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
	// region and the credentials configuring this argument must have
	// `waf:GetWebACL` permissions assigned.
	WebAclId interface{}
}

// The set of arguments for constructing a Distribution resource.
type DistributionArgs struct {
	// Extra CNAMEs (alternate domain names), if any, for
	// this distribution.
	Aliases interface{}
	// Any comments you want to include about the
	// distribution.
	Comment interface{}
	// One or more custom error response elements (multiples allowed).
	CustomErrorResponses interface{}
	// The default cache behavior for this distribution (maximum
	// one).
	DefaultCacheBehavior interface{}
	// The object that you want CloudFront to
	// return (for example, index.html) when an end user requests the root URL.
	DefaultRootObject interface{}
	// Whether the distribution is enabled to accept end
	// user requests for content.
	Enabled interface{}
	// The maximum HTTP version to support on the
	// distribution. Allowed values are `http1.1` and `http2`. The default is
	// `http2`.
	HttpVersion interface{}
	// Whether the IPv6 is enabled for the distribution.
	IsIpv6Enabled interface{}
	// The logging
	// configuration that controls how logs are written
	// to your distribution (maximum one).
	LoggingConfig interface{}
	// An ordered list of cache behaviors
	// resource for this distribution. List from top to bottom
	// in order of precedence. The topmost cache behavior will have precedence 0.
	OrderedCacheBehaviors interface{}
	// One or more origins for this
	// distribution (multiples allowed).
	Origins interface{}
	// One or more originGroup for this
	// distribution (multiples allowed).
	OriginGroups interface{}
	// The price class for this distribution. One of
	// `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
	PriceClass interface{}
	// The restriction
	// configuration for this distribution (maximum one).
	Restrictions interface{}
	// Disables the distribution instead of
	// deleting it when destroying the resource. If this is set,
	// the distribution needs to be deleted manually afterwards. Default: `false`.
	RetainOnDelete interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The SSL
	// configuration for this distribution (maximum
	// one).
	ViewerCertificate interface{}
	// If enabled, the resource will wait for
	// the distribution status to change from `InProgress` to `Deployed`. Setting
	// this to`false` will skip the process. Default: `true`.
	WaitForDeployment interface{}
	// If you're using AWS WAF to filter CloudFront
	// requests, the Id of the AWS WAF web ACL that is associated with the
	// distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
	// region and the credentials configuring this argument must have
	// `waf:GetWebACL` permissions assigned.
	WebAclId interface{}
}
