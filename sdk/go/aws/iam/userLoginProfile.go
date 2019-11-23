// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an IAM User Login Profile with limited support for password creation during this provider resource creation. Uses PGP to encrypt the password for safe transport to the user. PGP keys can be obtained from Keybase.
// 
// > To reset an IAM User login password via this provider, you can use delete and recreate this resource or change any of the arguments.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_user_login_profile.html.markdown.
type UserLoginProfile struct {
	s *pulumi.ResourceState
}

// NewUserLoginProfile registers a new resource with the given unique name, arguments, and options.
func NewUserLoginProfile(ctx *pulumi.Context,
	name string, args *UserLoginProfileArgs, opts ...pulumi.ResourceOpt) (*UserLoginProfile, error) {
	if args == nil || args.PgpKey == nil {
		return nil, errors.New("missing required argument 'PgpKey'")
	}
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["passwordLength"] = nil
		inputs["passwordResetRequired"] = nil
		inputs["pgpKey"] = nil
		inputs["user"] = nil
	} else {
		inputs["passwordLength"] = args.PasswordLength
		inputs["passwordResetRequired"] = args.PasswordResetRequired
		inputs["pgpKey"] = args.PgpKey
		inputs["user"] = args.User
	}
	inputs["encryptedPassword"] = nil
	inputs["keyFingerprint"] = nil
	s, err := ctx.RegisterResource("aws:iam/userLoginProfile:UserLoginProfile", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserLoginProfile{s: s}, nil
}

// GetUserLoginProfile gets an existing UserLoginProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserLoginProfile(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserLoginProfileState, opts ...pulumi.ResourceOpt) (*UserLoginProfile, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["encryptedPassword"] = state.EncryptedPassword
		inputs["keyFingerprint"] = state.KeyFingerprint
		inputs["passwordLength"] = state.PasswordLength
		inputs["passwordResetRequired"] = state.PasswordResetRequired
		inputs["pgpKey"] = state.PgpKey
		inputs["user"] = state.User
	}
	s, err := ctx.ReadResource("aws:iam/userLoginProfile:UserLoginProfile", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserLoginProfile{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *UserLoginProfile) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *UserLoginProfile) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The encrypted password, base64 encoded. Only available if password was handled on this provider resource creation, not import.
func (r *UserLoginProfile) EncryptedPassword() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["encryptedPassword"])
}

// The fingerprint of the PGP key used to encrypt the password. Only available if password was handled on this provider resource creation, not import.
func (r *UserLoginProfile) KeyFingerprint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyFingerprint"])
}

// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
func (r *UserLoginProfile) PasswordLength() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["passwordLength"])
}

// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
func (r *UserLoginProfile) PasswordResetRequired() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["passwordResetRequired"])
}

// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
func (r *UserLoginProfile) PgpKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["pgpKey"])
}

// The IAM user's name.
func (r *UserLoginProfile) User() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["user"])
}

// Input properties used for looking up and filtering UserLoginProfile resources.
type UserLoginProfileState struct {
	// The encrypted password, base64 encoded. Only available if password was handled on this provider resource creation, not import.
	EncryptedPassword interface{}
	// The fingerprint of the PGP key used to encrypt the password. Only available if password was handled on this provider resource creation, not import.
	KeyFingerprint interface{}
	// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
	PasswordLength interface{}
	// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
	PasswordResetRequired interface{}
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
	PgpKey interface{}
	// The IAM user's name.
	User interface{}
}

// The set of arguments for constructing a UserLoginProfile resource.
type UserLoginProfileArgs struct {
	// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
	PasswordLength interface{}
	// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
	PasswordResetRequired interface{}
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
	PgpKey interface{}
	// The IAM user's name.
	User interface{}
}
