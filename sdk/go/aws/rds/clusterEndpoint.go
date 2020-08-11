// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an RDS Aurora Cluster Endpoint.
// You can refer to the [User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.Endpoints.html#Aurora.Endpoints.Cluster).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := rds.NewCluster(ctx, "_default", &rds.ClusterArgs{
// 			ClusterIdentifier: pulumi.String("aurora-cluster-demo"),
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-west-2a"),
// 				pulumi.String("us-west-2b"),
// 				pulumi.String("us-west-2c"),
// 			},
// 			DatabaseName:          pulumi.String("mydb"),
// 			MasterUsername:        pulumi.String("foo"),
// 			MasterPassword:        pulumi.String("bar"),
// 			BackupRetentionPeriod: pulumi.Int(5),
// 			PreferredBackupWindow: pulumi.String("07:00-09:00"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		test1, err := rds.NewClusterInstance(ctx, "test1", &rds.ClusterInstanceArgs{
// 			ApplyImmediately:  pulumi.Bool(true),
// 			ClusterIdentifier: _default.ID(),
// 			Identifier:        pulumi.String("test1"),
// 			InstanceClass:     pulumi.String("db.t2.small"),
// 			Engine:            _default.Engine,
// 			EngineVersion:     _default.EngineVersion,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		test2, err := rds.NewClusterInstance(ctx, "test2", &rds.ClusterInstanceArgs{
// 			ApplyImmediately:  pulumi.Bool(true),
// 			ClusterIdentifier: _default.ID(),
// 			Identifier:        pulumi.String("test2"),
// 			InstanceClass:     pulumi.String("db.t2.small"),
// 			Engine:            _default.Engine,
// 			EngineVersion:     _default.EngineVersion,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		test3, err := rds.NewClusterInstance(ctx, "test3", &rds.ClusterInstanceArgs{
// 			ApplyImmediately:  pulumi.Bool(true),
// 			ClusterIdentifier: _default.ID(),
// 			Identifier:        pulumi.String("test3"),
// 			InstanceClass:     pulumi.String("db.t2.small"),
// 			Engine:            _default.Engine,
// 			EngineVersion:     _default.EngineVersion,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rds.NewClusterEndpoint(ctx, "eligible", &rds.ClusterEndpointArgs{
// 			ClusterIdentifier:         _default.ID(),
// 			ClusterEndpointIdentifier: pulumi.String("reader"),
// 			CustomEndpointType:        pulumi.String("READER"),
// 			ExcludedMembers: pulumi.StringArray{
// 				test1.ID(),
// 				test2.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rds.NewClusterEndpoint(ctx, "static", &rds.ClusterEndpointArgs{
// 			ClusterIdentifier:         _default.ID(),
// 			ClusterEndpointIdentifier: pulumi.String("static"),
// 			CustomEndpointType:        pulumi.String("READER"),
// 			StaticMembers: pulumi.StringArray{
// 				test1.ID(),
// 				test3.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ClusterEndpoint struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of cluster
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
	ClusterEndpointIdentifier pulumi.StringOutput `pulumi:"clusterEndpointIdentifier"`
	// The cluster identifier.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`
	// The type of the endpoint. One of: READER , ANY .
	CustomEndpointType pulumi.StringOutput `pulumi:"customEndpointType"`
	// A custom endpoint for the Aurora cluster
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `staticMembers`.
	ExcludedMembers pulumi.StringArrayOutput `pulumi:"excludedMembers"`
	// List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excludedMembers`.
	StaticMembers pulumi.StringArrayOutput `pulumi:"staticMembers"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewClusterEndpoint registers a new resource with the given unique name, arguments, and options.
func NewClusterEndpoint(ctx *pulumi.Context,
	name string, args *ClusterEndpointArgs, opts ...pulumi.ResourceOption) (*ClusterEndpoint, error) {
	if args == nil || args.ClusterEndpointIdentifier == nil {
		return nil, errors.New("missing required argument 'ClusterEndpointIdentifier'")
	}
	if args == nil || args.ClusterIdentifier == nil {
		return nil, errors.New("missing required argument 'ClusterIdentifier'")
	}
	if args == nil || args.CustomEndpointType == nil {
		return nil, errors.New("missing required argument 'CustomEndpointType'")
	}
	if args == nil {
		args = &ClusterEndpointArgs{}
	}
	var resource ClusterEndpoint
	err := ctx.RegisterResource("aws:rds/clusterEndpoint:ClusterEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterEndpoint gets an existing ClusterEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterEndpointState, opts ...pulumi.ResourceOption) (*ClusterEndpoint, error) {
	var resource ClusterEndpoint
	err := ctx.ReadResource("aws:rds/clusterEndpoint:ClusterEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterEndpoint resources.
type clusterEndpointState struct {
	// Amazon Resource Name (ARN) of cluster
	Arn *string `pulumi:"arn"`
	// The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
	ClusterEndpointIdentifier *string `pulumi:"clusterEndpointIdentifier"`
	// The cluster identifier.
	ClusterIdentifier *string `pulumi:"clusterIdentifier"`
	// The type of the endpoint. One of: READER , ANY .
	CustomEndpointType *string `pulumi:"customEndpointType"`
	// A custom endpoint for the Aurora cluster
	Endpoint *string `pulumi:"endpoint"`
	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `staticMembers`.
	ExcludedMembers []string `pulumi:"excludedMembers"`
	// List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excludedMembers`.
	StaticMembers []string `pulumi:"staticMembers"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type ClusterEndpointState struct {
	// Amazon Resource Name (ARN) of cluster
	Arn pulumi.StringPtrInput
	// The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
	ClusterEndpointIdentifier pulumi.StringPtrInput
	// The cluster identifier.
	ClusterIdentifier pulumi.StringPtrInput
	// The type of the endpoint. One of: READER , ANY .
	CustomEndpointType pulumi.StringPtrInput
	// A custom endpoint for the Aurora cluster
	Endpoint pulumi.StringPtrInput
	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `staticMembers`.
	ExcludedMembers pulumi.StringArrayInput
	// List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excludedMembers`.
	StaticMembers pulumi.StringArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (ClusterEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterEndpointState)(nil)).Elem()
}

type clusterEndpointArgs struct {
	// The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
	ClusterEndpointIdentifier string `pulumi:"clusterEndpointIdentifier"`
	// The cluster identifier.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	// The type of the endpoint. One of: READER , ANY .
	CustomEndpointType string `pulumi:"customEndpointType"`
	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `staticMembers`.
	ExcludedMembers []string `pulumi:"excludedMembers"`
	// List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excludedMembers`.
	StaticMembers []string `pulumi:"staticMembers"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterEndpoint resource.
type ClusterEndpointArgs struct {
	// The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
	ClusterEndpointIdentifier pulumi.StringInput
	// The cluster identifier.
	ClusterIdentifier pulumi.StringInput
	// The type of the endpoint. One of: READER , ANY .
	CustomEndpointType pulumi.StringInput
	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `staticMembers`.
	ExcludedMembers pulumi.StringArrayInput
	// List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excludedMembers`.
	StaticMembers pulumi.StringArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (ClusterEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterEndpointArgs)(nil)).Elem()
}
