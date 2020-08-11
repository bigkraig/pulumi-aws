// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Storage Gateway SMB File Share.
//
// ## Example Usage
// ### Active Directory Authentication
//
// > **NOTE:** The gateway must have already joined the Active Directory domain prior to SMB file share creation. e.g. via "SMB Settings" in the AWS Storage Gateway console or `smbActiveDirectorySettings` in the `storagegateway.Gateway` resource.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewSmbFileShare(ctx, "example", &storagegateway.SmbFileShareArgs{
// 			Authentication: pulumi.String("ActiveDirectory"),
// 			GatewayArn:     pulumi.Any(aws_storagegateway_gateway.Example.Arn),
// 			LocationArn:    pulumi.Any(aws_s3_bucket.Example.Arn),
// 			RoleArn:        pulumi.Any(aws_iam_role.Example.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Guest Authentication
//
// > **NOTE:** The gateway must have already had the SMB guest password set prior to SMB file share creation. e.g. via "SMB Settings" in the AWS Storage Gateway console or `smbGuestPassword` in the `storagegateway.Gateway` resource.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewSmbFileShare(ctx, "example", &storagegateway.SmbFileShareArgs{
// 			Authentication: pulumi.String("GuestAccess"),
// 			GatewayArn:     pulumi.Any(aws_storagegateway_gateway.Example.Arn),
// 			LocationArn:    pulumi.Any(aws_s3_bucket.Example.Arn),
// 			RoleArn:        pulumi.Any(aws_iam_role.Example.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SmbFileShare struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the SMB File Share.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication pulumi.StringPtrOutput `pulumi:"authentication"`
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass pulumi.StringPtrOutput `pulumi:"defaultStorageClass"`
	// ID of the SMB File Share.
	FileshareId pulumi.StringOutput `pulumi:"fileshareId"`
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn pulumi.StringOutput `pulumi:"gatewayArn"`
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled pulumi.BoolPtrOutput `pulumi:"guessMimeTypeEnabled"`
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists pulumi.StringArrayOutput `pulumi:"invalidUserLists"`
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted pulumi.BoolPtrOutput `pulumi:"kmsEncrypted"`
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn pulumi.StringPtrOutput `pulumi:"kmsKeyArn"`
	// The ARN of the backed storage used for storing file data.
	LocationArn pulumi.StringOutput `pulumi:"locationArn"`
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl pulumi.StringPtrOutput `pulumi:"objectAcl"`
	// File share path used by the NFS client to identify the mount point.
	Path pulumi.StringOutput `pulumi:"path"`
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly pulumi.BoolPtrOutput `pulumi:"readOnly"`
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays pulumi.BoolPtrOutput `pulumi:"requesterPays"`
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists pulumi.StringArrayOutput `pulumi:"validUserLists"`
}

// NewSmbFileShare registers a new resource with the given unique name, arguments, and options.
func NewSmbFileShare(ctx *pulumi.Context,
	name string, args *SmbFileShareArgs, opts ...pulumi.ResourceOption) (*SmbFileShare, error) {
	if args == nil || args.GatewayArn == nil {
		return nil, errors.New("missing required argument 'GatewayArn'")
	}
	if args == nil || args.LocationArn == nil {
		return nil, errors.New("missing required argument 'LocationArn'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil {
		args = &SmbFileShareArgs{}
	}
	var resource SmbFileShare
	err := ctx.RegisterResource("aws:storagegateway/smbFileShare:SmbFileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmbFileShare gets an existing SmbFileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmbFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmbFileShareState, opts ...pulumi.ResourceOption) (*SmbFileShare, error) {
	var resource SmbFileShare
	err := ctx.ReadResource("aws:storagegateway/smbFileShare:SmbFileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmbFileShare resources.
type smbFileShareState struct {
	// Amazon Resource Name (ARN) of the SMB File Share.
	Arn *string `pulumi:"arn"`
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication *string `pulumi:"authentication"`
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass *string `pulumi:"defaultStorageClass"`
	// ID of the SMB File Share.
	FileshareId *string `pulumi:"fileshareId"`
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn *string `pulumi:"gatewayArn"`
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled *bool `pulumi:"guessMimeTypeEnabled"`
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists []string `pulumi:"invalidUserLists"`
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted *bool `pulumi:"kmsEncrypted"`
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The ARN of the backed storage used for storing file data.
	LocationArn *string `pulumi:"locationArn"`
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl *string `pulumi:"objectAcl"`
	// File share path used by the NFS client to identify the mount point.
	Path *string `pulumi:"path"`
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly *bool `pulumi:"readOnly"`
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays *bool `pulumi:"requesterPays"`
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn *string `pulumi:"roleArn"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists []string `pulumi:"validUserLists"`
}

type SmbFileShareState struct {
	// Amazon Resource Name (ARN) of the SMB File Share.
	Arn pulumi.StringPtrInput
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication pulumi.StringPtrInput
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass pulumi.StringPtrInput
	// ID of the SMB File Share.
	FileshareId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn pulumi.StringPtrInput
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled pulumi.BoolPtrInput
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists pulumi.StringArrayInput
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn pulumi.StringPtrInput
	// The ARN of the backed storage used for storing file data.
	LocationArn pulumi.StringPtrInput
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl pulumi.StringPtrInput
	// File share path used by the NFS client to identify the mount point.
	Path pulumi.StringPtrInput
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly pulumi.BoolPtrInput
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays pulumi.BoolPtrInput
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists pulumi.StringArrayInput
}

func (SmbFileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*smbFileShareState)(nil)).Elem()
}

type smbFileShareArgs struct {
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication *string `pulumi:"authentication"`
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass *string `pulumi:"defaultStorageClass"`
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn string `pulumi:"gatewayArn"`
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled *bool `pulumi:"guessMimeTypeEnabled"`
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists []string `pulumi:"invalidUserLists"`
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted *bool `pulumi:"kmsEncrypted"`
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The ARN of the backed storage used for storing file data.
	LocationArn string `pulumi:"locationArn"`
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl *string `pulumi:"objectAcl"`
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly *bool `pulumi:"readOnly"`
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays *bool `pulumi:"requesterPays"`
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn string `pulumi:"roleArn"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists []string `pulumi:"validUserLists"`
}

// The set of arguments for constructing a SmbFileShare resource.
type SmbFileShareArgs struct {
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication pulumi.StringPtrInput
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn pulumi.StringInput
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled pulumi.BoolPtrInput
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists pulumi.StringArrayInput
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn pulumi.StringPtrInput
	// The ARN of the backed storage used for storing file data.
	LocationArn pulumi.StringInput
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl pulumi.StringPtrInput
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly pulumi.BoolPtrInput
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays pulumi.BoolPtrInput
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn pulumi.StringInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists pulumi.StringArrayInput
}

func (SmbFileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smbFileShareArgs)(nil)).Elem()
}
