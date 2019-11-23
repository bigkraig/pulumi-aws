// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage a GuardDuty ThreatIntelSet.
// 
// > **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage ThreatIntelSets. ThreatIntelSets that are uploaded by the master account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-threat-intel-set.html)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/guardduty_threatintelset.html.markdown.
type ThreatIntelSet struct {
	s *pulumi.ResourceState
}

// NewThreatIntelSet registers a new resource with the given unique name, arguments, and options.
func NewThreatIntelSet(ctx *pulumi.Context,
	name string, args *ThreatIntelSetArgs, opts ...pulumi.ResourceOpt) (*ThreatIntelSet, error) {
	if args == nil || args.Activate == nil {
		return nil, errors.New("missing required argument 'Activate'")
	}
	if args == nil || args.DetectorId == nil {
		return nil, errors.New("missing required argument 'DetectorId'")
	}
	if args == nil || args.Format == nil {
		return nil, errors.New("missing required argument 'Format'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["activate"] = nil
		inputs["detectorId"] = nil
		inputs["format"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
	} else {
		inputs["activate"] = args.Activate
		inputs["detectorId"] = args.DetectorId
		inputs["format"] = args.Format
		inputs["location"] = args.Location
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("aws:guardduty/threatIntelSet:ThreatIntelSet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ThreatIntelSet{s: s}, nil
}

// GetThreatIntelSet gets an existing ThreatIntelSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThreatIntelSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ThreatIntelSetState, opts ...pulumi.ResourceOpt) (*ThreatIntelSet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["activate"] = state.Activate
		inputs["detectorId"] = state.DetectorId
		inputs["format"] = state.Format
		inputs["location"] = state.Location
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:guardduty/threatIntelSet:ThreatIntelSet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ThreatIntelSet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ThreatIntelSet) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ThreatIntelSet) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
func (r *ThreatIntelSet) Activate() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["activate"])
}

// The detector ID of the GuardDuty.
func (r *ThreatIntelSet) DetectorId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["detectorId"])
}

// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
func (r *ThreatIntelSet) Format() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["format"])
}

// The URI of the file that contains the ThreatIntelSet.
func (r *ThreatIntelSet) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// The friendly name to identify the ThreatIntelSet.
func (r *ThreatIntelSet) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering ThreatIntelSet resources.
type ThreatIntelSetState struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate interface{}
	// The detector ID of the GuardDuty.
	DetectorId interface{}
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format interface{}
	// The URI of the file that contains the ThreatIntelSet.
	Location interface{}
	// The friendly name to identify the ThreatIntelSet.
	Name interface{}
}

// The set of arguments for constructing a ThreatIntelSet resource.
type ThreatIntelSetArgs struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate interface{}
	// The detector ID of the GuardDuty.
	DetectorId interface{}
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format interface{}
	// The URI of the file that contains the ThreatIntelSet.
	Location interface{}
	// The friendly name to identify the ThreatIntelSet.
	Name interface{}
}
