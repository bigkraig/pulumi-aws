// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Cognito Resource Server.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_resource_server.html.markdown.
type ResourceServer struct {
	s *pulumi.ResourceState
}

// NewResourceServer registers a new resource with the given unique name, arguments, and options.
func NewResourceServer(ctx *pulumi.Context,
	name string, args *ResourceServerArgs, opts ...pulumi.ResourceOpt) (*ResourceServer, error) {
	if args == nil || args.Identifier == nil {
		return nil, errors.New("missing required argument 'Identifier'")
	}
	if args == nil || args.UserPoolId == nil {
		return nil, errors.New("missing required argument 'UserPoolId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["identifier"] = nil
		inputs["name"] = nil
		inputs["scopes"] = nil
		inputs["userPoolId"] = nil
	} else {
		inputs["identifier"] = args.Identifier
		inputs["name"] = args.Name
		inputs["scopes"] = args.Scopes
		inputs["userPoolId"] = args.UserPoolId
	}
	inputs["scopeIdentifiers"] = nil
	s, err := ctx.RegisterResource("aws:cognito/resourceServer:ResourceServer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ResourceServer{s: s}, nil
}

// GetResourceServer gets an existing ResourceServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceServer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ResourceServerState, opts ...pulumi.ResourceOpt) (*ResourceServer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["identifier"] = state.Identifier
		inputs["name"] = state.Name
		inputs["scopes"] = state.Scopes
		inputs["scopeIdentifiers"] = state.ScopeIdentifiers
		inputs["userPoolId"] = state.UserPoolId
	}
	s, err := ctx.ReadResource("aws:cognito/resourceServer:ResourceServer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ResourceServer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ResourceServer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ResourceServer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// An identifier for the resource server.
func (r *ResourceServer) Identifier() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["identifier"])
}

// A name for the resource server.
func (r *ResourceServer) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A list of Authorization Scope.
func (r *ResourceServer) Scopes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["scopes"])
}

// A list of all scopes configured for this resource server in the format identifier/scope_name.
func (r *ResourceServer) ScopeIdentifiers() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["scopeIdentifiers"])
}

func (r *ResourceServer) UserPoolId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userPoolId"])
}

// Input properties used for looking up and filtering ResourceServer resources.
type ResourceServerState struct {
	// An identifier for the resource server.
	Identifier interface{}
	// A name for the resource server.
	Name interface{}
	// A list of Authorization Scope.
	Scopes interface{}
	// A list of all scopes configured for this resource server in the format identifier/scope_name.
	ScopeIdentifiers interface{}
	UserPoolId interface{}
}

// The set of arguments for constructing a ResourceServer resource.
type ResourceServerArgs struct {
	// An identifier for the resource server.
	Identifier interface{}
	// A name for the resource server.
	Name interface{}
	// A list of Authorization Scope.
	Scopes interface{}
	UserPoolId interface{}
}
