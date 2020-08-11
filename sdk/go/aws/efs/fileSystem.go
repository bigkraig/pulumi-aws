// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Elastic File System (EFS) File System resource.
//
// ## Example Usage
// ### EFS File System w/ tags
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := efs.NewFileSystem(ctx, "foo", &efs.FileSystemArgs{
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("MyProduct"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using lifecycle policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := efs.NewFileSystem(ctx, "fooWithLifecylePolicy", &efs.FileSystemArgs{
// 			LifecyclePolicy: &efs.FileSystemLifecyclePolicyArgs{
// 				TransitionToIa: pulumi.String("AFTER_30_DAYS"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FileSystem struct {
	pulumi.CustomResourceState

	// Amazon Resource Name of the file system.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken pulumi.StringOutput `pulumi:"creationToken"`
	// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// If true, the disk will be encrypted.
	Encrypted pulumi.BoolOutput `pulumi:"encrypted"`
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy FileSystemLifecyclePolicyPtrOutput `pulumi:"lifecyclePolicy"`
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode pulumi.StringOutput `pulumi:"performanceMode"`
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps pulumi.Float64PtrOutput `pulumi:"provisionedThroughputInMibps"`
	// A map of tags to assign to the file system.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode pulumi.StringPtrOutput `pulumi:"throughputMode"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil {
		args = &FileSystemArgs{}
	}
	var resource FileSystem
	err := ctx.RegisterResource("aws:efs/fileSystem:FileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileSystemState, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	var resource FileSystem
	err := ctx.ReadResource("aws:efs/fileSystem:FileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileSystem resources.
type fileSystemState struct {
	// Amazon Resource Name of the file system.
	Arn *string `pulumi:"arn"`
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken *string `pulumi:"creationToken"`
	// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	DnsName *string `pulumi:"dnsName"`
	// If true, the disk will be encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy *FileSystemLifecyclePolicy `pulumi:"lifecyclePolicy"`
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode *string `pulumi:"performanceMode"`
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps *float64 `pulumi:"provisionedThroughputInMibps"`
	// A map of tags to assign to the file system.
	Tags map[string]string `pulumi:"tags"`
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode *string `pulumi:"throughputMode"`
}

type FileSystemState struct {
	// Amazon Resource Name of the file system.
	Arn pulumi.StringPtrInput
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken pulumi.StringPtrInput
	// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	DnsName pulumi.StringPtrInput
	// If true, the disk will be encrypted.
	Encrypted pulumi.BoolPtrInput
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId pulumi.StringPtrInput
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy FileSystemLifecyclePolicyPtrInput
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode pulumi.StringPtrInput
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps pulumi.Float64PtrInput
	// A map of tags to assign to the file system.
	Tags pulumi.StringMapInput
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode pulumi.StringPtrInput
}

func (FileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemState)(nil)).Elem()
}

type fileSystemArgs struct {
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken *string `pulumi:"creationToken"`
	// If true, the disk will be encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy *FileSystemLifecyclePolicy `pulumi:"lifecyclePolicy"`
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode *string `pulumi:"performanceMode"`
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps *float64 `pulumi:"provisionedThroughputInMibps"`
	// A map of tags to assign to the file system.
	Tags map[string]string `pulumi:"tags"`
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode *string `pulumi:"throughputMode"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken pulumi.StringPtrInput
	// If true, the disk will be encrypted.
	Encrypted pulumi.BoolPtrInput
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId pulumi.StringPtrInput
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy FileSystemLifecyclePolicyPtrInput
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode pulumi.StringPtrInput
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps pulumi.Float64PtrInput
	// A map of tags to assign to the file system.
	Tags pulumi.StringMapInput
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode pulumi.StringPtrInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}
