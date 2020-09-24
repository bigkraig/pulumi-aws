module github.com/pulumi/pulumi-aws/provider/v3

go 1.14

require (
	github.com/hashicorp/aws-sdk-go-base v0.6.0
	github.com/hashicorp/terraform-plugin-sdk v1.14.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.9.2
	github.com/pulumi/pulumi/sdk/v2 v2.10.2-0.20200916204740-92a7d717a4d1
	github.com/terraform-providers/terraform-provider-aws v0.0.0-20191010190908-1261a98537f2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20200910230100-328eb4ff41df
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-aws => github.com/pulumi/terraform-provider-aws v1.38.1-0.20200924135808-374f5da7c8d8
)
