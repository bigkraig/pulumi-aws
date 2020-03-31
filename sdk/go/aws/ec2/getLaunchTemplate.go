// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a Launch Template.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/launch_template.html.markdown.
func LookupLaunchTemplate(ctx *pulumi.Context, args *LookupLaunchTemplateArgs, opts ...pulumi.InvokeOption) (*LookupLaunchTemplateResult, error) {
	var rv LookupLaunchTemplateResult
	err := ctx.Invoke("aws:ec2/getLaunchTemplate:getLaunchTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLaunchTemplate.
type LookupLaunchTemplateArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetLaunchTemplateFilter `pulumi:"filters"`
	// The name of the filter field. Valid values can be found in the [EC2 DescribeLaunchTemplates API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLaunchTemplates.html).
	Name *string `pulumi:"name"`
	// A mapping of tags, each pair of which must exactly match a pair on the desired Launch Template.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getLaunchTemplate.
type LookupLaunchTemplateResult struct {
	// Amazon Resource Name (ARN) of the launch template.
	Arn string `pulumi:"arn"`
	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	BlockDeviceMappings []GetLaunchTemplateBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecifications []GetLaunchTemplateCreditSpecification `pulumi:"creditSpecifications"`
	// The default version of the launch template.
	DefaultVersion int `pulumi:"defaultVersion"`
	// Description of the launch template.
	Description string `pulumi:"description"`
	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination bool `pulumi:"disableApiTermination"`
	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized string `pulumi:"ebsOptimized"`
	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications []GetLaunchTemplateElasticGpuSpecification `pulumi:"elasticGpuSpecifications"`
	Filters                  []GetLaunchTemplateFilter                  `pulumi:"filters"`
	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfiles []GetLaunchTemplateIamInstanceProfile `pulumi:"iamInstanceProfiles"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The AMI from which to launch the instance.
	ImageId string `pulumi:"imageId"`
	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior string `pulumi:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instance.
	// below for details.
	InstanceMarketOptions []GetLaunchTemplateInstanceMarketOption `pulumi:"instanceMarketOptions"`
	// The type of the instance.
	InstanceType string `pulumi:"instanceType"`
	// The kernel ID.
	KernelId string `pulumi:"kernelId"`
	// The key name to use for the instance.
	KeyName string `pulumi:"keyName"`
	// The latest version of the launch template.
	LatestVersion int `pulumi:"latestVersion"`
	// The metadata options for the instance.
	MetadataOptions []GetLaunchTemplateMetadataOption `pulumi:"metadataOptions"`
	// The monitoring option for the instance.
	Monitorings []GetLaunchTemplateMonitoring `pulumi:"monitorings"`
	Name        *string                       `pulumi:"name"`
	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces []GetLaunchTemplateNetworkInterface `pulumi:"networkInterfaces"`
	// The placement of the instance.
	Placements []GetLaunchTemplatePlacement `pulumi:"placements"`
	// The ID of the RAM disk.
	RamDiskId string `pulumi:"ramDiskId"`
	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
	// The tags to apply to the resources during launch.
	TagSpecifications []GetLaunchTemplateTagSpecification `pulumi:"tagSpecifications"`
	// (Optional) A mapping of tags to assign to the launch template.
	Tags map[string]interface{} `pulumi:"tags"`
	// The Base64-encoded user data to provide when launching the instance.
	UserData string `pulumi:"userData"`
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}
