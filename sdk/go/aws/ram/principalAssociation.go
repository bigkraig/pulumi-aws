// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Resource Access Manager (RAM) principal association. Depending if [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), the RAM behavior with different principal types changes.
// 
// When RAM Sharing with AWS Organizations is enabled:
// 
// - For AWS Account ID, Organization, and Organizational Unit principals within the same AWS Organization, no resource share invitation is sent and resources become available automatically after creating the association.
// - For AWS Account ID principals outside the AWS Organization, a resource share invitation is sent and must be accepted before resources become available. See the [`ram.ResourceShareAccepter` resource](https://www.terraform.io/docs/providers/aws/r/ram_resource_share_accepter.html) to accept these invitations.
// 
// When RAM Sharing with AWS Organizations is not enabled:
// 
// - Organization and Organizational Unit principals cannot be used.
// - For AWS Account ID principals, a resource share invitation is sent and must be accepted before resources become available. See the [`ram.ResourceShareAccepter` resource](https://www.terraform.io/docs/providers/aws/r/ram_resource_share_accepter.html) to accept these invitations.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ram_principal_association.html.markdown.
type PrincipalAssociation struct {
	s *pulumi.ResourceState
}

// NewPrincipalAssociation registers a new resource with the given unique name, arguments, and options.
func NewPrincipalAssociation(ctx *pulumi.Context,
	name string, args *PrincipalAssociationArgs, opts ...pulumi.ResourceOpt) (*PrincipalAssociation, error) {
	if args == nil || args.Principal == nil {
		return nil, errors.New("missing required argument 'Principal'")
	}
	if args == nil || args.ResourceShareArn == nil {
		return nil, errors.New("missing required argument 'ResourceShareArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["principal"] = nil
		inputs["resourceShareArn"] = nil
	} else {
		inputs["principal"] = args.Principal
		inputs["resourceShareArn"] = args.ResourceShareArn
	}
	s, err := ctx.RegisterResource("aws:ram/principalAssociation:PrincipalAssociation", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PrincipalAssociation{s: s}, nil
}

// GetPrincipalAssociation gets an existing PrincipalAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrincipalAssociation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PrincipalAssociationState, opts ...pulumi.ResourceOpt) (*PrincipalAssociation, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["principal"] = state.Principal
		inputs["resourceShareArn"] = state.ResourceShareArn
	}
	s, err := ctx.ReadResource("aws:ram/principalAssociation:PrincipalAssociation", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PrincipalAssociation{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PrincipalAssociation) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PrincipalAssociation) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
func (r *PrincipalAssociation) Principal() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["principal"])
}

// The Amazon Resource Name (ARN) of the resource share.
func (r *PrincipalAssociation) ResourceShareArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceShareArn"])
}

// Input properties used for looking up and filtering PrincipalAssociation resources.
type PrincipalAssociationState struct {
	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	Principal interface{}
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn interface{}
}

// The set of arguments for constructing a PrincipalAssociation resource.
type PrincipalAssociationArgs struct {
	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	Principal interface{}
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn interface{}
}
