// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage the accepter's side of a Direct Connect hosted private virtual interface.
// This resource accepts ownership of a private virtual interface created by another AWS account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/directconnect"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/providers"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := providers.Newaws(ctx, "accepter", nil)
// 		if err != nil {
// 			return err
// 		}
// 		accepterCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		vpnGw, err := ec2.NewVpnGateway(ctx, "vpnGw", nil, pulumi.Provider(aws.Accepter))
// 		if err != nil {
// 			return err
// 		}
// 		creator, err := directconnect.NewHostedPrivateVirtualInterface(ctx, "creator", &directconnect.HostedPrivateVirtualInterfaceArgs{
// 			ConnectionId:   pulumi.String("dxcon-zzzzzzzz"),
// 			OwnerAccountId: pulumi.String(accepterCallerIdentity.AccountId),
// 			Vlan:           pulumi.Int(4094),
// 			AddressFamily:  pulumi.String("ipv4"),
// 			BgpAsn:         pulumi.Int(65352),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			vpnGw,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = directconnect.NewHostedPrivateVirtualInterfaceAccepter(ctx, "accepterHostedPrivateVirtualInterfaceAccepter", &directconnect.HostedPrivateVirtualInterfaceAccepterArgs{
// 			VirtualInterfaceId: creator.ID(),
// 			VpnGatewayId:       vpnGw.ID(),
// 			Tags: pulumi.StringMap{
// 				"Side": pulumi.String("Accepter"),
// 			},
// 		}, pulumi.Provider(aws.Accepter))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type HostedPrivateVirtualInterfaceAccepter struct {
	pulumi.CustomResourceState

	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrOutput `pulumi:"dxGatewayId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringOutput `pulumi:"virtualInterfaceId"`
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId pulumi.StringPtrOutput `pulumi:"vpnGatewayId"`
}

// NewHostedPrivateVirtualInterfaceAccepter registers a new resource with the given unique name, arguments, and options.
func NewHostedPrivateVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, args *HostedPrivateVirtualInterfaceAccepterArgs, opts ...pulumi.ResourceOption) (*HostedPrivateVirtualInterfaceAccepter, error) {
	if args == nil || args.VirtualInterfaceId == nil {
		return nil, errors.New("missing required argument 'VirtualInterfaceId'")
	}
	if args == nil {
		args = &HostedPrivateVirtualInterfaceAccepterArgs{}
	}
	var resource HostedPrivateVirtualInterfaceAccepter
	err := ctx.RegisterResource("aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedPrivateVirtualInterfaceAccepter gets an existing HostedPrivateVirtualInterfaceAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedPrivateVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedPrivateVirtualInterfaceAccepterState, opts ...pulumi.ResourceOption) (*HostedPrivateVirtualInterfaceAccepter, error) {
	var resource HostedPrivateVirtualInterfaceAccepter
	err := ctx.ReadResource("aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedPrivateVirtualInterfaceAccepter resources.
type hostedPrivateVirtualInterfaceAccepterState struct {
	// The ARN of the virtual interface.
	Arn *string `pulumi:"arn"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId *string `pulumi:"virtualInterfaceId"`
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

type HostedPrivateVirtualInterfaceAccepterState struct {
	// The ARN of the virtual interface.
	Arn pulumi.StringPtrInput
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringPtrInput
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId pulumi.StringPtrInput
}

func (HostedPrivateVirtualInterfaceAccepterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPrivateVirtualInterfaceAccepterState)(nil)).Elem()
}

type hostedPrivateVirtualInterfaceAccepterArgs struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId string `pulumi:"virtualInterfaceId"`
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a HostedPrivateVirtualInterfaceAccepter resource.
type HostedPrivateVirtualInterfaceAccepterArgs struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringInput
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId pulumi.StringPtrInput
}

func (HostedPrivateVirtualInterfaceAccepterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPrivateVirtualInterfaceAccepterArgs)(nil)).Elem()
}
