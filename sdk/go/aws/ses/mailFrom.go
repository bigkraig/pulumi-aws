// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SES domain MAIL FROM resource.
//
// > **NOTE:** For the MAIL FROM domain to be fully usable, this resource should be paired with the `ses.DomainIdentity` resource. To validate the MAIL FROM domain, a DNS MX record is required. To pass SPF checks, a DNS TXT record may also be required. See the [Amazon SES MAIL FROM documentation](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from-set.html) for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleDomainIdentity, err := ses.NewDomainIdentity(ctx, "exampleDomainIdentity", &ses.DomainIdentityArgs{
// 			Domain: pulumi.String("example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleMailFrom, err := ses.NewMailFrom(ctx, "exampleMailFrom", &ses.MailFromArgs{
// 			Domain: exampleDomainIdentity.Domain,
// 			MailFromDomain: exampleDomainIdentity.Domain.ApplyT(func(domain string) (string, error) {
// 				return fmt.Sprintf("%v%v", "bounce.", domain), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewRecord(ctx, "exampleSesDomainMailFromMx", &route53.RecordArgs{
// 			ZoneId: pulumi.String(aws_route53_zone.Example.Id),
// 			Name:   exampleMailFrom.MailFromDomain,
// 			Type:   pulumi.String("MX"),
// 			Ttl:    pulumi.Int(600),
// 			Records: pulumi.StringArray{
// 				pulumi.String("10 feedback-smtp.us-east-1.amazonses.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewRecord(ctx, "exampleSesDomainMailFromTxt", &route53.RecordArgs{
// 			ZoneId: pulumi.String(aws_route53_zone.Example.Id),
// 			Name:   exampleMailFrom.MailFromDomain,
// 			Type:   pulumi.String("TXT"),
// 			Ttl:    pulumi.Int(600),
// 			Records: pulumi.StringArray{
// 				pulumi.String("v=spf1 include:amazonses.com -all"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MailFrom struct {
	pulumi.CustomResourceState

	// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
	BehaviorOnMxFailure pulumi.StringPtrOutput `pulumi:"behaviorOnMxFailure"`
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
	MailFromDomain pulumi.StringOutput `pulumi:"mailFromDomain"`
}

// NewMailFrom registers a new resource with the given unique name, arguments, and options.
func NewMailFrom(ctx *pulumi.Context,
	name string, args *MailFromArgs, opts ...pulumi.ResourceOption) (*MailFrom, error) {
	if args == nil || args.Domain == nil {
		return nil, errors.New("missing required argument 'Domain'")
	}
	if args == nil || args.MailFromDomain == nil {
		return nil, errors.New("missing required argument 'MailFromDomain'")
	}
	if args == nil {
		args = &MailFromArgs{}
	}
	var resource MailFrom
	err := ctx.RegisterResource("aws:ses/mailFrom:MailFrom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMailFrom gets an existing MailFrom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMailFrom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MailFromState, opts ...pulumi.ResourceOption) (*MailFrom, error) {
	var resource MailFrom
	err := ctx.ReadResource("aws:ses/mailFrom:MailFrom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MailFrom resources.
type mailFromState struct {
	// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
	BehaviorOnMxFailure *string `pulumi:"behaviorOnMxFailure"`
	// Verified domain name to generate DKIM tokens for.
	Domain *string `pulumi:"domain"`
	// Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
	MailFromDomain *string `pulumi:"mailFromDomain"`
}

type MailFromState struct {
	// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
	BehaviorOnMxFailure pulumi.StringPtrInput
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringPtrInput
	// Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
	MailFromDomain pulumi.StringPtrInput
}

func (MailFromState) ElementType() reflect.Type {
	return reflect.TypeOf((*mailFromState)(nil)).Elem()
}

type mailFromArgs struct {
	// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
	BehaviorOnMxFailure *string `pulumi:"behaviorOnMxFailure"`
	// Verified domain name to generate DKIM tokens for.
	Domain string `pulumi:"domain"`
	// Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
	MailFromDomain string `pulumi:"mailFromDomain"`
}

// The set of arguments for constructing a MailFrom resource.
type MailFromArgs struct {
	// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
	BehaviorOnMxFailure pulumi.StringPtrInput
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringInput
	// Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
	MailFromDomain pulumi.StringInput
}

func (MailFromArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mailFromArgs)(nil)).Elem()
}
