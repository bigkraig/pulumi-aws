// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codedeploy

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CodeDeploy application to be used as a basis for deployments
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codedeploy_app.html.markdown.
type Application struct {
	s *pulumi.ResourceState
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOpt) (*Application, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["computePlatform"] = nil
		inputs["name"] = nil
		inputs["uniqueId"] = nil
	} else {
		inputs["computePlatform"] = args.ComputePlatform
		inputs["name"] = args.Name
		inputs["uniqueId"] = args.UniqueId
	}
	s, err := ctx.RegisterResource("aws:codedeploy/application:Application", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Application{s: s}, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApplicationState, opts ...pulumi.ResourceOpt) (*Application, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["computePlatform"] = state.ComputePlatform
		inputs["name"] = state.Name
		inputs["uniqueId"] = state.UniqueId
	}
	s, err := ctx.ReadResource("aws:codedeploy/application:Application", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Application{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Application) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Application) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
func (r *Application) ComputePlatform() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["computePlatform"])
}

// The name of the application.
func (r *Application) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *Application) UniqueId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["uniqueId"])
}

// Input properties used for looking up and filtering Application resources.
type ApplicationState struct {
	// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
	ComputePlatform interface{}
	// The name of the application.
	Name interface{}
	UniqueId interface{}
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
	ComputePlatform interface{}
	// The name of the application.
	Name interface{}
	UniqueId interface{}
}
