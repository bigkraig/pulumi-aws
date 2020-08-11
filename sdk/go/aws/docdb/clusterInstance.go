// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an DocDB Cluster Resource Instance. A Cluster Instance Resource defines
// attributes that are specific to a single instance in a [DocDB Cluster](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html).
//
// You do not designate a primary and subsequent replicas. Instead, you simply add DocDB
// Instances and DocDB manages the replication. You can use the [count](https://www.terraform.io/docs/configuration/resources.html#count)
// meta-parameter to make multiple instances and join them all to the same DocDB
// Cluster, or you may specify different Cluster Instance resources with various
// `instanceClass` sizes.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/docdb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := docdb.NewCluster(ctx, "_default", &docdb.ClusterArgs{
// 			ClusterIdentifier: pulumi.String("docdb-cluster-demo"),
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-west-2a"),
// 				pulumi.String("us-west-2b"),
// 				pulumi.String("us-west-2c"),
// 			},
// 			MasterUsername: pulumi.String("foo"),
// 			MasterPassword: pulumi.String("barbut8chars"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		var clusterInstances []*docdb.ClusterInstance
// 		for key0, val0 := range 2 {
// 			__res, err := docdb.NewClusterInstance(ctx, fmt.Sprintf("clusterInstances-%v", key0), &docdb.ClusterInstanceArgs{
// 				Identifier:        pulumi.String(fmt.Sprintf("%v%v", "docdb-cluster-demo-", val0)),
// 				ClusterIdentifier: _default.ID(),
// 				InstanceClass:     pulumi.String("db.r5.large"),
// 			})
// 			if err != nil {
// 				return err
// 			}
// 			clusterInstances = append(clusterInstances, __res)
// 		}
// 		return nil
// 	})
// }
// ```
type ClusterInstance struct {
	pulumi.CustomResourceState

	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolOutput `pulumi:"applyImmediately"`
	// Amazon Resource Name (ARN) of cluster instance
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade pulumi.BoolPtrOutput `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier pulumi.StringOutput `pulumi:"caCertIdentifier"`
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`
	// The DB subnet group to associate with this DB instance.
	DbSubnetGroupName pulumi.StringOutput `pulumi:"dbSubnetGroupName"`
	// The region-unique, immutable identifier for the DB instance.
	DbiResourceId pulumi.StringOutput `pulumi:"dbiResourceId"`
	// The DNS address for this instance. May not be writable
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine pulumi.StringPtrOutput `pulumi:"engine"`
	// The database engine version
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
	IdentifierPrefix pulumi.StringOutput `pulumi:"identifierPrefix"`
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	InstanceClass pulumi.StringOutput `pulumi:"instanceClass"`
	// The ARN for the KMS encryption key if one is set to the cluster.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// The database port
	Port pulumi.IntOutput `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled.
	PreferredBackupWindow pulumi.StringOutput `pulumi:"preferredBackupWindow"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringOutput `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier      pulumi.IntPtrOutput `pulumi:"promotionTier"`
	PubliclyAccessible pulumi.BoolOutput   `pulumi:"publiclyAccessible"`
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted pulumi.BoolOutput `pulumi:"storageEncrypted"`
	// A map of tags to assign to the instance.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer pulumi.BoolOutput `pulumi:"writer"`
}

// NewClusterInstance registers a new resource with the given unique name, arguments, and options.
func NewClusterInstance(ctx *pulumi.Context,
	name string, args *ClusterInstanceArgs, opts ...pulumi.ResourceOption) (*ClusterInstance, error) {
	if args == nil || args.ClusterIdentifier == nil {
		return nil, errors.New("missing required argument 'ClusterIdentifier'")
	}
	if args == nil || args.InstanceClass == nil {
		return nil, errors.New("missing required argument 'InstanceClass'")
	}
	if args == nil {
		args = &ClusterInstanceArgs{}
	}
	var resource ClusterInstance
	err := ctx.RegisterResource("aws:docdb/clusterInstance:ClusterInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterInstance gets an existing ClusterInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterInstanceState, opts ...pulumi.ResourceOption) (*ClusterInstance, error) {
	var resource ClusterInstance
	err := ctx.ReadResource("aws:docdb/clusterInstance:ClusterInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterInstance resources.
type clusterInstanceState struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// Amazon Resource Name (ARN) of cluster instance
	Arn *string `pulumi:"arn"`
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier *string `pulumi:"caCertIdentifier"`
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier *string `pulumi:"clusterIdentifier"`
	// The DB subnet group to associate with this DB instance.
	DbSubnetGroupName *string `pulumi:"dbSubnetGroupName"`
	// The region-unique, immutable identifier for the DB instance.
	DbiResourceId *string `pulumi:"dbiResourceId"`
	// The DNS address for this instance. May not be writable
	Endpoint *string `pulumi:"endpoint"`
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine *string `pulumi:"engine"`
	// The database engine version
	EngineVersion *string `pulumi:"engineVersion"`
	// The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier *string `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
	IdentifierPrefix *string `pulumi:"identifierPrefix"`
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	InstanceClass *string `pulumi:"instanceClass"`
	// The ARN for the KMS encryption key if one is set to the cluster.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The database port
	Port *int `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled.
	PreferredBackupWindow *string `pulumi:"preferredBackupWindow"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier      *int  `pulumi:"promotionTier"`
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
	// A map of tags to assign to the instance.
	Tags map[string]string `pulumi:"tags"`
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer *bool `pulumi:"writer"`
}

type ClusterInstanceState struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of cluster instance
	Arn pulumi.StringPtrInput
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone pulumi.StringPtrInput
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier pulumi.StringPtrInput
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier pulumi.StringPtrInput
	// The DB subnet group to associate with this DB instance.
	DbSubnetGroupName pulumi.StringPtrInput
	// The region-unique, immutable identifier for the DB instance.
	DbiResourceId pulumi.StringPtrInput
	// The DNS address for this instance. May not be writable
	Endpoint pulumi.StringPtrInput
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine pulumi.StringPtrInput
	// The database engine version
	EngineVersion pulumi.StringPtrInput
	// The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringPtrInput
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
	IdentifierPrefix pulumi.StringPtrInput
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	InstanceClass pulumi.StringPtrInput
	// The ARN for the KMS encryption key if one is set to the cluster.
	KmsKeyId pulumi.StringPtrInput
	// The database port
	Port pulumi.IntPtrInput
	// The daily time range during which automated backups are created if automated backups are enabled.
	PreferredBackupWindow pulumi.StringPtrInput
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier      pulumi.IntPtrInput
	PubliclyAccessible pulumi.BoolPtrInput
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted pulumi.BoolPtrInput
	// A map of tags to assign to the instance.
	Tags pulumi.StringMapInput
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer pulumi.BoolPtrInput
}

func (ClusterInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterInstanceState)(nil)).Elem()
}

type clusterInstanceArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier *string `pulumi:"caCertIdentifier"`
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine *string `pulumi:"engine"`
	// The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier *string `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
	IdentifierPrefix *string `pulumi:"identifierPrefix"`
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	InstanceClass string `pulumi:"instanceClass"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier *int `pulumi:"promotionTier"`
	// A map of tags to assign to the instance.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterInstance resource.
type ClusterInstanceArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolPtrInput
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone pulumi.StringPtrInput
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier pulumi.StringPtrInput
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier pulumi.StringInput
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine pulumi.StringPtrInput
	// The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringPtrInput
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
	IdentifierPrefix pulumi.StringPtrInput
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	InstanceClass pulumi.StringInput
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier pulumi.IntPtrInput
	// A map of tags to assign to the instance.
	Tags pulumi.StringMapInput
}

func (ClusterInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterInstanceArgs)(nil)).Elem()
}
