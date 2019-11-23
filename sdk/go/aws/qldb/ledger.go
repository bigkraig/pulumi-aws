// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package qldb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/qldb_ledger.html.markdown.
type Ledger struct {
	s *pulumi.ResourceState
}

// NewLedger registers a new resource with the given unique name, arguments, and options.
func NewLedger(ctx *pulumi.Context,
	name string, args *LedgerArgs, opts ...pulumi.ResourceOpt) (*Ledger, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["deletionProtection"] = nil
		inputs["name"] = nil
		inputs["tags"] = nil
	} else {
		inputs["deletionProtection"] = args.DeletionProtection
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:qldb/ledger:Ledger", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Ledger{s: s}, nil
}

// GetLedger gets an existing Ledger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLedger(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LedgerState, opts ...pulumi.ResourceOpt) (*Ledger, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["deletionProtection"] = state.DeletionProtection
		inputs["name"] = state.Name
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:qldb/ledger:Ledger", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Ledger{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Ledger) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Ledger) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the QLDB Ledger
func (r *Ledger) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

func (r *Ledger) DeletionProtection() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["deletionProtection"])
}

func (r *Ledger) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Key-value mapping of resource tags
func (r *Ledger) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Ledger resources.
type LedgerState struct {
	// The ARN of the QLDB Ledger
	Arn interface{}
	DeletionProtection interface{}
	Name interface{}
	// Key-value mapping of resource tags
	Tags interface{}
}

// The set of arguments for constructing a Ledger resource.
type LedgerArgs struct {
	DeletionProtection interface{}
	Name interface{}
	// Key-value mapping of resource tags
	Tags interface{}
}
