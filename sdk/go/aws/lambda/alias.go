// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Lambda function alias. Creates an alias that points to the specified Lambda function version.
// 
// For information about Lambda and how to use it, see [What is AWS Lambda?][1]
// For information about function aliases, see [CreateAlias][2] and [AliasRoutingConfiguration][3] in the API docs.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lambda_alias.html.markdown.
type Alias struct {
	s *pulumi.ResourceState
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOpt) (*Alias, error) {
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil || args.FunctionVersion == nil {
		return nil, errors.New("missing required argument 'FunctionVersion'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["functionName"] = nil
		inputs["functionVersion"] = nil
		inputs["name"] = nil
		inputs["routingConfig"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["functionName"] = args.FunctionName
		inputs["functionVersion"] = args.FunctionVersion
		inputs["name"] = args.Name
		inputs["routingConfig"] = args.RoutingConfig
	}
	inputs["arn"] = nil
	inputs["invokeArn"] = nil
	s, err := ctx.RegisterResource("aws:lambda/alias:Alias", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Alias{s: s}, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AliasState, opts ...pulumi.ResourceOpt) (*Alias, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["functionName"] = state.FunctionName
		inputs["functionVersion"] = state.FunctionVersion
		inputs["invokeArn"] = state.InvokeArn
		inputs["name"] = state.Name
		inputs["routingConfig"] = state.RoutingConfig
	}
	s, err := ctx.ReadResource("aws:lambda/alias:Alias", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Alias{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Alias) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Alias) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The Amazon Resource Name (ARN) identifying your Lambda function alias.
func (r *Alias) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Description of the alias.
func (r *Alias) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The function ARN of the Lambda function for which you want to create an alias.
func (r *Alias) FunctionName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["functionName"])
}

// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
func (r *Alias) FunctionVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["functionVersion"])
}

// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
func (r *Alias) InvokeArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["invokeArn"])
}

// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
func (r *Alias) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The Lambda alias' route configuration settings. Fields documented below
func (r *Alias) RoutingConfig() pulumi.Output {
	return r.s.State["routingConfig"]
}

// Input properties used for looking up and filtering Alias resources.
type AliasState struct {
	// The Amazon Resource Name (ARN) identifying your Lambda function alias.
	Arn interface{}
	// Description of the alias.
	Description interface{}
	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName interface{}
	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion interface{}
	// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
	InvokeArn interface{}
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name interface{}
	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig interface{}
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// Description of the alias.
	Description interface{}
	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName interface{}
	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion interface{}
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name interface{}
	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig interface{}
}
