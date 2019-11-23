// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a settings of an API Gateway Documentation Part.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_documentation_part.html.markdown.
type DocumentationPart struct {
	s *pulumi.ResourceState
}

// NewDocumentationPart registers a new resource with the given unique name, arguments, and options.
func NewDocumentationPart(ctx *pulumi.Context,
	name string, args *DocumentationPartArgs, opts ...pulumi.ResourceOpt) (*DocumentationPart, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.RestApiId == nil {
		return nil, errors.New("missing required argument 'RestApiId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["properties"] = nil
		inputs["restApiId"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["properties"] = args.Properties
		inputs["restApiId"] = args.RestApiId
	}
	s, err := ctx.RegisterResource("aws:apigateway/documentationPart:DocumentationPart", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DocumentationPart{s: s}, nil
}

// GetDocumentationPart gets an existing DocumentationPart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentationPart(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DocumentationPartState, opts ...pulumi.ResourceOpt) (*DocumentationPart, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["location"] = state.Location
		inputs["properties"] = state.Properties
		inputs["restApiId"] = state.RestApiId
	}
	s, err := ctx.ReadResource("aws:apigateway/documentationPart:DocumentationPart", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DocumentationPart{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DocumentationPart) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DocumentationPart) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The location of the targeted API entity of the to-be-created documentation part. See below.
func (r *DocumentationPart) Location() pulumi.Output {
	return r.s.State["location"]
}

// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
func (r *DocumentationPart) Properties() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["properties"])
}

// The ID of the associated Rest API
func (r *DocumentationPart) RestApiId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["restApiId"])
}

// Input properties used for looking up and filtering DocumentationPart resources.
type DocumentationPartState struct {
	// The location of the targeted API entity of the to-be-created documentation part. See below.
	Location interface{}
	// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties interface{}
	// The ID of the associated Rest API
	RestApiId interface{}
}

// The set of arguments for constructing a DocumentationPart resource.
type DocumentationPartArgs struct {
	// The location of the targeted API entity of the to-be-created documentation part. See below.
	Location interface{}
	// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties interface{}
	// The ID of the associated Rest API
	RestApiId interface{}
}
