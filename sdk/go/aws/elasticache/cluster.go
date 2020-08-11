// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ElastiCache Cluster resource, which manages a Memcached cluster or Redis instance.
// For working with Redis (Cluster Mode Enabled) replication groups, see the
// `elasticache.ReplicationGroup` resource.
//
// > **Note:** When you change an attribute, such as `nodeType`, by default
// it is applied in the next maintenance window. Because of this, this provider may report
// a difference in its planning phase because the actual modification has not yet taken
// place. You can use the `applyImmediately` flag to instruct the service to apply the
// change immediately. Using `applyImmediately` can result in a brief downtime as the server reboots.
// See the AWS Docs on [Modifying an ElastiCache Cache Cluster](https://docs.aws.amazon.com/AmazonElastiCache/latest/UserGuide/Clusters.Modify.html) for more information.
//
// ## Example Usage
// ### Memcached Cluster
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elasticache"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elasticache.NewCluster(ctx, "example", &elasticache.ClusterArgs{
// 			Engine:             pulumi.String("memcached"),
// 			NodeType:           pulumi.String("cache.m4.large"),
// 			NumCacheNodes:      pulumi.Int(2),
// 			ParameterGroupName: pulumi.String("default.memcached1.4"),
// 			Port:               pulumi.Int(11211),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Redis Instance
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elasticache"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elasticache.NewCluster(ctx, "example", &elasticache.ClusterArgs{
// 			Engine:             pulumi.String("redis"),
// 			EngineVersion:      pulumi.String("3.2.10"),
// 			NodeType:           pulumi.String("cache.m4.large"),
// 			NumCacheNodes:      pulumi.Int(1),
// 			ParameterGroupName: pulumi.String("default.redis3.2"),
// 			Port:               pulumi.Int(6379),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Redis Cluster Mode Disabled Read Replica Instance
//
// These inherit their settings from the replication group.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elasticache"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elasticache.NewCluster(ctx, "replica", &elasticache.ClusterArgs{
// 			ReplicationGroupId: pulumi.Any(aws_elasticache_replication_group.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html)
	// (Available since v0.6.0)
	ApplyImmediately pulumi.BoolOutput   `pulumi:"applyImmediately"`
	Arn              pulumi.StringOutput `pulumi:"arn"`
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `numCacheNodes` must be greater than `1`
	AzMode pulumi.StringOutput `pulumi:"azMode"`
	// List of node objects including `id`, `address`, `port` and `availabilityZone`.
	// Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
	CacheNodes ClusterCacheNodeArrayOutput `pulumi:"cacheNodes"`
	// (Memcached only) The DNS name of the cache cluster without the port appended.
	ClusterAddress pulumi.StringOutput `pulumi:"clusterAddress"`
	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// (Memcached only) The configuration endpoint to allow host discovery.
	ConfigurationEndpoint pulumi.StringOutput `pulumi:"configurationEndpoint"`
	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Version number of the cache engine to be used.
	// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
	// in the AWS Documentation center for supported versions
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringOutput `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringPtrOutput `pulumi:"notificationTopicArn"`
	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes pulumi.IntOutput `pulumi:"numCacheNodes"`
	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName pulumi.StringOutput `pulumi:"parameterGroupName"`
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port pulumi.IntOutput `pulumi:"port"`
	// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `numCacheNodes`. If you want all the nodes in the same Availability Zone, use `availabilityZone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones pulumi.StringArrayOutput `pulumi:"preferredAvailabilityZones"`
	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId pulumi.StringOutput `pulumi:"replicationGroupId"`
	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames pulumi.StringArrayOutput `pulumi:"securityGroupNames"`
	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns pulumi.StringArrayOutput `pulumi:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshotName` forces a new resource.
	SnapshotName pulumi.StringPtrOutput `pulumi:"snapshotName"`
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit pulumi.IntPtrOutput `pulumi:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow pulumi.StringOutput `pulumi:"snapshotWindow"`
	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName pulumi.StringOutput `pulumi:"subnetGroupName"`
	// A map of tags to assign to the resource
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("aws:elasticache/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("aws:elasticache/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html)
	// (Available since v0.6.0)
	ApplyImmediately *bool   `pulumi:"applyImmediately"`
	Arn              *string `pulumi:"arn"`
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `numCacheNodes` must be greater than `1`
	AzMode *string `pulumi:"azMode"`
	// List of node objects including `id`, `address`, `port` and `availabilityZone`.
	// Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
	CacheNodes []ClusterCacheNode `pulumi:"cacheNodes"`
	// (Memcached only) The DNS name of the cache cluster without the port appended.
	ClusterAddress *string `pulumi:"clusterAddress"`
	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId *string `pulumi:"clusterId"`
	// (Memcached only) The configuration endpoint to allow host discovery.
	ConfigurationEndpoint *string `pulumi:"configurationEndpoint"`
	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine *string `pulumi:"engine"`
	// Version number of the cache engine to be used.
	// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
	// in the AWS Documentation center for supported versions
	EngineVersion *string `pulumi:"engineVersion"`
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType *string `pulumi:"nodeType"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn *string `pulumi:"notificationTopicArn"`
	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes *int `pulumi:"numCacheNodes"`
	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port *int `pulumi:"port"`
	// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `numCacheNodes`. If you want all the nodes in the same Availability Zone, use `availabilityZone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones []string `pulumi:"preferredAvailabilityZones"`
	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId *string `pulumi:"replicationGroupId"`
	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns []string `pulumi:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshotName` forces a new resource.
	SnapshotName *string `pulumi:"snapshotName"`
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit *int `pulumi:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow *string `pulumi:"snapshotWindow"`
	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName *string `pulumi:"subnetGroupName"`
	// A map of tags to assign to the resource
	Tags map[string]string `pulumi:"tags"`
}

type ClusterState struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html)
	// (Available since v0.6.0)
	ApplyImmediately pulumi.BoolPtrInput
	Arn              pulumi.StringPtrInput
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone pulumi.StringPtrInput
	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `numCacheNodes` must be greater than `1`
	AzMode pulumi.StringPtrInput
	// List of node objects including `id`, `address`, `port` and `availabilityZone`.
	// Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
	CacheNodes ClusterCacheNodeArrayInput
	// (Memcached only) The DNS name of the cache cluster without the port appended.
	ClusterAddress pulumi.StringPtrInput
	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId pulumi.StringPtrInput
	// (Memcached only) The configuration endpoint to allow host discovery.
	ConfigurationEndpoint pulumi.StringPtrInput
	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine pulumi.StringPtrInput
	// Version number of the cache engine to be used.
	// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
	// in the AWS Documentation center for supported versions
	EngineVersion pulumi.StringPtrInput
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringPtrInput
	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType pulumi.StringPtrInput
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringPtrInput
	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes pulumi.IntPtrInput
	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName pulumi.StringPtrInput
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port pulumi.IntPtrInput
	// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `numCacheNodes`. If you want all the nodes in the same Availability Zone, use `availabilityZone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones pulumi.StringArrayInput
	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId pulumi.StringPtrInput
	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds pulumi.StringArrayInput
	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames pulumi.StringArrayInput
	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns pulumi.StringArrayInput
	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshotName` forces a new resource.
	SnapshotName pulumi.StringPtrInput
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit pulumi.IntPtrInput
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow pulumi.StringPtrInput
	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName pulumi.StringPtrInput
	// A map of tags to assign to the resource
	Tags pulumi.StringMapInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html)
	// (Available since v0.6.0)
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `numCacheNodes` must be greater than `1`
	AzMode *string `pulumi:"azMode"`
	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId *string `pulumi:"clusterId"`
	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine *string `pulumi:"engine"`
	// Version number of the cache engine to be used.
	// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
	// in the AWS Documentation center for supported versions
	EngineVersion *string `pulumi:"engineVersion"`
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType *string `pulumi:"nodeType"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn *string `pulumi:"notificationTopicArn"`
	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes *int `pulumi:"numCacheNodes"`
	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port *int `pulumi:"port"`
	// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `numCacheNodes`. If you want all the nodes in the same Availability Zone, use `availabilityZone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones []string `pulumi:"preferredAvailabilityZones"`
	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId *string `pulumi:"replicationGroupId"`
	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns []string `pulumi:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshotName` forces a new resource.
	SnapshotName *string `pulumi:"snapshotName"`
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit *int `pulumi:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow *string `pulumi:"snapshotWindow"`
	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName *string `pulumi:"subnetGroupName"`
	// A map of tags to assign to the resource
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html)
	// (Available since v0.6.0)
	ApplyImmediately pulumi.BoolPtrInput
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone pulumi.StringPtrInput
	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `numCacheNodes` must be greater than `1`
	AzMode pulumi.StringPtrInput
	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId pulumi.StringPtrInput
	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine pulumi.StringPtrInput
	// Version number of the cache engine to be used.
	// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
	// in the AWS Documentation center for supported versions
	EngineVersion pulumi.StringPtrInput
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringPtrInput
	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType pulumi.StringPtrInput
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringPtrInput
	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes pulumi.IntPtrInput
	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName pulumi.StringPtrInput
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port pulumi.IntPtrInput
	// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `numCacheNodes`. If you want all the nodes in the same Availability Zone, use `availabilityZone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones pulumi.StringArrayInput
	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId pulumi.StringPtrInput
	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds pulumi.StringArrayInput
	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames pulumi.StringArrayInput
	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns pulumi.StringArrayInput
	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshotName` forces a new resource.
	SnapshotName pulumi.StringPtrInput
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit pulumi.IntPtrInput
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow pulumi.StringPtrInput
	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName pulumi.StringPtrInput
	// A map of tags to assign to the resource
	Tags pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
