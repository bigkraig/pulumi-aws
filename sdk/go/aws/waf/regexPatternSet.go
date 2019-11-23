// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Regex Pattern Set Resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_regex_pattern_set.html.markdown.
type RegexPatternSet struct {
	s *pulumi.ResourceState
}

// NewRegexPatternSet registers a new resource with the given unique name, arguments, and options.
func NewRegexPatternSet(ctx *pulumi.Context,
	name string, args *RegexPatternSetArgs, opts ...pulumi.ResourceOpt) (*RegexPatternSet, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["regexPatternStrings"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["regexPatternStrings"] = args.RegexPatternStrings
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:waf/regexPatternSet:RegexPatternSet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegexPatternSet{s: s}, nil
}

// GetRegexPatternSet gets an existing RegexPatternSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexPatternSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegexPatternSetState, opts ...pulumi.ResourceOpt) (*RegexPatternSet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["name"] = state.Name
		inputs["regexPatternStrings"] = state.RegexPatternStrings
	}
	s, err := ctx.ReadResource("aws:waf/regexPatternSet:RegexPatternSet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegexPatternSet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RegexPatternSet) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RegexPatternSet) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Amazon Resource Name (ARN)
func (r *RegexPatternSet) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The name or description of the Regex Pattern Set.
func (r *RegexPatternSet) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
func (r *RegexPatternSet) RegexPatternStrings() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["regexPatternStrings"])
}

// Input properties used for looking up and filtering RegexPatternSet resources.
type RegexPatternSetState struct {
	// Amazon Resource Name (ARN)
	Arn interface{}
	// The name or description of the Regex Pattern Set.
	Name interface{}
	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings interface{}
}

// The set of arguments for constructing a RegexPatternSet resource.
type RegexPatternSetArgs struct {
	// The name or description of the Regex Pattern Set.
	Name interface{}
	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings interface{}
}
