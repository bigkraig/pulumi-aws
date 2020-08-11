// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Service Discovery Service resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/servicediscovery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
// 			CidrBlock:          pulumi.String("10.0.0.0/16"),
// 			EnableDnsSupport:   pulumi.Bool(true),
// 			EnableDnsHostnames: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePrivateDnsNamespace, err := servicediscovery.NewPrivateDnsNamespace(ctx, "examplePrivateDnsNamespace", &servicediscovery.PrivateDnsNamespaceArgs{
// 			Description: pulumi.String("example"),
// 			Vpc:         exampleVpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicediscovery.NewService(ctx, "exampleService", &servicediscovery.ServiceArgs{
// 			DnsConfig: &servicediscovery.ServiceDnsConfigArgs{
// 				NamespaceId: examplePrivateDnsNamespace.ID(),
// 				DnsRecords: servicediscovery.ServiceDnsConfigDnsRecordArray{
// 					&servicediscovery.ServiceDnsConfigDnsRecordArgs{
// 						Ttl:  pulumi.Int(10),
// 						Type: pulumi.String("A"),
// 					},
// 				},
// 				RoutingPolicy: pulumi.String("MULTIVALUE"),
// 			},
// 			HealthCheckCustomConfig: &servicediscovery.ServiceHealthCheckCustomConfigArgs{
// 				FailureThreshold: pulumi.Int(1),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/servicediscovery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		examplePublicDnsNamespace, err := servicediscovery.NewPublicDnsNamespace(ctx, "examplePublicDnsNamespace", &servicediscovery.PublicDnsNamespaceArgs{
// 			Description: pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicediscovery.NewService(ctx, "exampleService", &servicediscovery.ServiceArgs{
// 			DnsConfig: &servicediscovery.ServiceDnsConfigArgs{
// 				NamespaceId: examplePublicDnsNamespace.ID(),
// 				DnsRecords: servicediscovery.ServiceDnsConfigDnsRecordArray{
// 					&servicediscovery.ServiceDnsConfigDnsRecordArgs{
// 						Ttl:  pulumi.Int(10),
// 						Type: pulumi.String("A"),
// 					},
// 				},
// 			},
// 			HealthCheckConfig: &servicediscovery.ServiceHealthCheckConfigArgs{
// 				FailureThreshold: pulumi.Int(10),
// 				ResourcePath:     pulumi.String("path"),
// 				Type:             pulumi.String("HTTP"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Service struct {
	pulumi.CustomResourceState

	// The ARN of the service.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
	DnsConfig ServiceDnsConfigPtrOutput `pulumi:"dnsConfig"`
	// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
	HealthCheckConfig ServiceHealthCheckConfigPtrOutput `pulumi:"healthCheckConfig"`
	// A complex type that contains settings for ECS managed health checks.
	HealthCheckCustomConfig ServiceHealthCheckCustomConfigPtrOutput `pulumi:"healthCheckCustomConfig"`
	// The name of the service.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the namespace to use for DNS configuration.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// A map of tags to assign to the service.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		args = &ServiceArgs{}
	}
	var resource Service
	err := ctx.RegisterResource("aws:servicediscovery/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("aws:servicediscovery/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// The ARN of the service.
	Arn *string `pulumi:"arn"`
	// The description of the service.
	Description *string `pulumi:"description"`
	// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
	DnsConfig *ServiceDnsConfig `pulumi:"dnsConfig"`
	// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
	HealthCheckConfig *ServiceHealthCheckConfig `pulumi:"healthCheckConfig"`
	// A complex type that contains settings for ECS managed health checks.
	HealthCheckCustomConfig *ServiceHealthCheckCustomConfig `pulumi:"healthCheckCustomConfig"`
	// The name of the service.
	Name *string `pulumi:"name"`
	// The ID of the namespace to use for DNS configuration.
	NamespaceId *string `pulumi:"namespaceId"`
	// A map of tags to assign to the service.
	Tags map[string]string `pulumi:"tags"`
}

type ServiceState struct {
	// The ARN of the service.
	Arn pulumi.StringPtrInput
	// The description of the service.
	Description pulumi.StringPtrInput
	// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
	DnsConfig ServiceDnsConfigPtrInput
	// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
	HealthCheckConfig ServiceHealthCheckConfigPtrInput
	// A complex type that contains settings for ECS managed health checks.
	HealthCheckCustomConfig ServiceHealthCheckCustomConfigPtrInput
	// The name of the service.
	Name pulumi.StringPtrInput
	// The ID of the namespace to use for DNS configuration.
	NamespaceId pulumi.StringPtrInput
	// A map of tags to assign to the service.
	Tags pulumi.StringMapInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The description of the service.
	Description *string `pulumi:"description"`
	// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
	DnsConfig *ServiceDnsConfig `pulumi:"dnsConfig"`
	// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
	HealthCheckConfig *ServiceHealthCheckConfig `pulumi:"healthCheckConfig"`
	// A complex type that contains settings for ECS managed health checks.
	HealthCheckCustomConfig *ServiceHealthCheckCustomConfig `pulumi:"healthCheckCustomConfig"`
	// The name of the service.
	Name *string `pulumi:"name"`
	// The ID of the namespace to use for DNS configuration.
	NamespaceId *string `pulumi:"namespaceId"`
	// A map of tags to assign to the service.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The description of the service.
	Description pulumi.StringPtrInput
	// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
	DnsConfig ServiceDnsConfigPtrInput
	// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
	HealthCheckConfig ServiceHealthCheckConfigPtrInput
	// A complex type that contains settings for ECS managed health checks.
	HealthCheckCustomConfig ServiceHealthCheckCustomConfigPtrInput
	// The name of the service.
	Name pulumi.StringPtrInput
	// The ID of the namespace to use for DNS configuration.
	NamespaceId pulumi.StringPtrInput
	// A map of tags to assign to the service.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
