// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Pinpoint SMS Channel resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_sms_channel.html.markdown.
type SmsChannel struct {
	s *pulumi.ResourceState
}

// NewSmsChannel registers a new resource with the given unique name, arguments, and options.
func NewSmsChannel(ctx *pulumi.Context,
	name string, args *SmsChannelArgs, opts ...pulumi.ResourceOpt) (*SmsChannel, error) {
	if args == nil || args.ApplicationId == nil {
		return nil, errors.New("missing required argument 'ApplicationId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["applicationId"] = nil
		inputs["enabled"] = nil
		inputs["senderId"] = nil
		inputs["shortCode"] = nil
	} else {
		inputs["applicationId"] = args.ApplicationId
		inputs["enabled"] = args.Enabled
		inputs["senderId"] = args.SenderId
		inputs["shortCode"] = args.ShortCode
	}
	inputs["promotionalMessagesPerSecond"] = nil
	inputs["transactionalMessagesPerSecond"] = nil
	s, err := ctx.RegisterResource("aws:pinpoint/smsChannel:SmsChannel", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SmsChannel{s: s}, nil
}

// GetSmsChannel gets an existing SmsChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmsChannel(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SmsChannelState, opts ...pulumi.ResourceOpt) (*SmsChannel, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["applicationId"] = state.ApplicationId
		inputs["enabled"] = state.Enabled
		inputs["promotionalMessagesPerSecond"] = state.PromotionalMessagesPerSecond
		inputs["senderId"] = state.SenderId
		inputs["shortCode"] = state.ShortCode
		inputs["transactionalMessagesPerSecond"] = state.TransactionalMessagesPerSecond
	}
	s, err := ctx.ReadResource("aws:pinpoint/smsChannel:SmsChannel", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SmsChannel{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SmsChannel) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SmsChannel) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The application ID.
func (r *SmsChannel) ApplicationId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["applicationId"])
}

// Whether the channel is enabled or disabled. Defaults to `true`.
func (r *SmsChannel) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// Promotional messages per second that can be sent.
func (r *SmsChannel) PromotionalMessagesPerSecond() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["promotionalMessagesPerSecond"])
}

// Sender identifier of your messages.
func (r *SmsChannel) SenderId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["senderId"])
}

// The Short Code registered with the phone provider.
func (r *SmsChannel) ShortCode() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["shortCode"])
}

// Transactional messages per second that can be sent.
func (r *SmsChannel) TransactionalMessagesPerSecond() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["transactionalMessagesPerSecond"])
}

// Input properties used for looking up and filtering SmsChannel resources.
type SmsChannelState struct {
	// The application ID.
	ApplicationId interface{}
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled interface{}
	// Promotional messages per second that can be sent.
	PromotionalMessagesPerSecond interface{}
	// Sender identifier of your messages.
	SenderId interface{}
	// The Short Code registered with the phone provider.
	ShortCode interface{}
	// Transactional messages per second that can be sent.
	TransactionalMessagesPerSecond interface{}
}

// The set of arguments for constructing a SmsChannel resource.
type SmsChannelArgs struct {
	// The application ID.
	ApplicationId interface{}
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled interface{}
	// Sender identifier of your messages.
	SenderId interface{}
	// The Short Code registered with the phone provider.
	ShortCode interface{}
}
