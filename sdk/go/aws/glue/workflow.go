// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Glue Workflow resource.
// The workflow graph (DAG) can be build using the `glue.Trigger` resource. 
// See the example below for creating a graph with four nodes (two triggers and two jobs). 
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/glue_workflow.html.markdown.
type Workflow struct {
	s *pulumi.ResourceState
}

// NewWorkflow registers a new resource with the given unique name, arguments, and options.
func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOpt) (*Workflow, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["defaultRunProperties"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
	} else {
		inputs["defaultRunProperties"] = args.DefaultRunProperties
		inputs["description"] = args.Description
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("aws:glue/workflow:Workflow", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Workflow{s: s}, nil
}

// GetWorkflow gets an existing Workflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.ID, state *WorkflowState, opts ...pulumi.ResourceOpt) (*Workflow, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["defaultRunProperties"] = state.DefaultRunProperties
		inputs["description"] = state.Description
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:glue/workflow:Workflow", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Workflow{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Workflow) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Workflow) ID() pulumi.IDOutput {
	return r.s.ID()
}

// A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
func (r *Workflow) DefaultRunProperties() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["defaultRunProperties"])
}

// Description of the workflow.
func (r *Workflow) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The name you assign to this workflow.
func (r *Workflow) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering Workflow resources.
type WorkflowState struct {
	// A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
	DefaultRunProperties interface{}
	// Description of the workflow.
	Description interface{}
	// The name you assign to this workflow.
	Name interface{}
}

// The set of arguments for constructing a Workflow resource.
type WorkflowArgs struct {
	// A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
	DefaultRunProperties interface{}
	// Description of the workflow.
	Description interface{}
	// The name you assign to this workflow.
	Name interface{}
}
