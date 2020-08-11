// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The CloudFormation Stack data source allows access to stack
// outputs and other useful data including the template body.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudformation"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		network, err := cloudformation.LookupStack(ctx, &cloudformation.LookupStackArgs{
// 			Name: "my-network-stack",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
// 			Ami:          pulumi.String("ami-abb07bcb"),
// 			InstanceType: pulumi.String("t1.micro"),
// 			SubnetId:     pulumi.String(network.Outputs.SubnetId),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("HelloWorld"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupStack(ctx *pulumi.Context, args *LookupStackArgs, opts ...pulumi.InvokeOption) (*LookupStackResult, error) {
	var rv LookupStackResult
	err := ctx.Invoke("aws:cloudformation/getStack:getStack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStack.
type LookupStackArgs struct {
	// The name of the stack
	Name string `pulumi:"name"`
	// A map of tags associated with this stack.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getStack.
type LookupStackResult struct {
	// A list of capabilities
	Capabilities []string `pulumi:"capabilities"`
	// Description of the stack
	Description string `pulumi:"description"`
	// Whether the rollback of the stack is disabled when stack creation fails
	DisableRollback bool `pulumi:"disableRollback"`
	// The ARN of the IAM role used to create the stack.
	IamRoleArn string `pulumi:"iamRoleArn"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// A list of SNS topic ARNs to publish stack related events
	NotificationArns []string `pulumi:"notificationArns"`
	// A map of outputs from the stack.
	Outputs map[string]string `pulumi:"outputs"`
	// A map of parameters that specify input parameters for the stack.
	Parameters map[string]string `pulumi:"parameters"`
	// A map of tags associated with this stack.
	Tags map[string]string `pulumi:"tags"`
	// Structure containing the template body.
	TemplateBody string `pulumi:"templateBody"`
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`
	TimeoutInMinutes int `pulumi:"timeoutInMinutes"`
}
