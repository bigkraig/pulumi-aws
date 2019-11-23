// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an API Gateway API Key.
// 
// > **NOTE:** Since the API Gateway usage plans feature was launched on August 11, 2016, usage plans are now **required** to associate an API key with an API stage.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_api_key.html.markdown.
type ApiKey struct {
	s *pulumi.ResourceState
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOpt) (*ApiKey, error) {
	inputs := make(map[string]interface{})
	inputs["description"] = "Managed by Pulumi"
	if args == nil {
		inputs["enabled"] = nil
		inputs["name"] = nil
		inputs["tags"] = nil
		inputs["value"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["enabled"] = args.Enabled
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
		inputs["value"] = args.Value
	}
	inputs["arn"] = nil
	inputs["createdDate"] = nil
	inputs["lastUpdatedDate"] = nil
	s, err := ctx.RegisterResource("aws:apigateway/apiKey:ApiKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiKey{s: s}, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApiKeyState, opts ...pulumi.ResourceOpt) (*ApiKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["createdDate"] = state.CreatedDate
		inputs["description"] = state.Description
		inputs["enabled"] = state.Enabled
		inputs["lastUpdatedDate"] = state.LastUpdatedDate
		inputs["name"] = state.Name
		inputs["tags"] = state.Tags
		inputs["value"] = state.Value
	}
	s, err := ctx.ReadResource("aws:apigateway/apiKey:ApiKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ApiKey) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ApiKey) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Amazon Resource Name (ARN)
func (r *ApiKey) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The creation date of the API key
func (r *ApiKey) CreatedDate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["createdDate"])
}

// The API key description. Defaults to "Managed by Pulumi".
func (r *ApiKey) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Specifies whether the API key can be used by callers. Defaults to `true`.
func (r *ApiKey) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// The last update date of the API key
func (r *ApiKey) LastUpdatedDate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["lastUpdatedDate"])
}

// The name of the API key
func (r *ApiKey) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Key-value mapping of resource tags
func (r *ApiKey) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
func (r *ApiKey) Value() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["value"])
}

// Input properties used for looking up and filtering ApiKey resources.
type ApiKeyState struct {
	// Amazon Resource Name (ARN)
	Arn interface{}
	// The creation date of the API key
	CreatedDate interface{}
	// The API key description. Defaults to "Managed by Pulumi".
	Description interface{}
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled interface{}
	// The last update date of the API key
	LastUpdatedDate interface{}
	// The name of the API key
	Name interface{}
	// Key-value mapping of resource tags
	Tags interface{}
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value interface{}
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// The API key description. Defaults to "Managed by Pulumi".
	Description interface{}
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled interface{}
	// The name of the API key
	Name interface{}
	// Key-value mapping of resource tags
	Tags interface{}
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value interface{}
}
