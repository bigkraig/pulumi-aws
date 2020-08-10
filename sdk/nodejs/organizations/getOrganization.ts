// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get information about the organization that the user's account belongs to
 *
 * ## Example Usage
 * ### List all account IDs for the organization
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.organizations.getOrganization({});
 * export const accountIds = example.then(example => example.accounts.map(__item => __item.id));
 * ```
 * ### SNS topic that can be interacted by the organization only
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.organizations.getOrganization({});
 * const snsTopic = new aws.sns.Topic("snsTopic", {});
 * const snsTopicPolicyPolicyDocument = pulumi.all([example, snsTopic.arn]).apply(([example, arn]) => aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         actions: [
 *             "SNS:Subscribe",
 *             "SNS:Publish",
 *         ],
 *         conditions: [{
 *             test: "StringEquals",
 *             variable: "aws:PrincipalOrgID",
 *             values: [example.id],
 *         }],
 *         principals: [{
 *             type: "AWS",
 *             identifiers: ["*"],
 *         }],
 *         resources: [arn],
 *     }],
 * }));
 * const snsTopicPolicyTopicPolicy = new aws.sns.TopicPolicy("snsTopicPolicyTopicPolicy", {
 *     arn: snsTopic.arn,
 *     policy: snsTopicPolicyPolicyDocument.json,
 * });
 * ```
 */
export function getOrganization(opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:organizations/getOrganization:getOrganization", {
    }, opts);
}

/**
 * A collection of values returned by getOrganization.
 */
export interface GetOrganizationResult {
    /**
     * List of organization accounts including the master account. For a list excluding the master account, see the `nonMasterAccounts` attribute. All elements have these attributes:
     */
    readonly accounts: outputs.organizations.GetOrganizationAccount[];
    /**
     * ARN of the root
     */
    readonly arn: string;
    /**
     * A list of AWS service principal names that have integration enabled with your organization. Organization must have `featureSet` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
     */
    readonly awsServiceAccessPrincipals: string[];
    /**
     * A list of Organizations policy types that are enabled in the Organization Root. Organization must have `featureSet` set to `ALL`. For additional information about valid policy types (e.g. `SERVICE_CONTROL_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
     */
    readonly enabledPolicyTypes: string[];
    /**
     * The FeatureSet of the organization.
     */
    readonly featureSet: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Amazon Resource Name (ARN) of the account that is designated as the master account for the organization.
     */
    readonly masterAccountArn: string;
    /**
     * The email address that is associated with the AWS account that is designated as the master account for the organization.
     */
    readonly masterAccountEmail: string;
    /**
     * The unique identifier (ID) of the master account of an organization.
     */
    readonly masterAccountId: string;
    /**
     * List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
     */
    readonly nonMasterAccounts: outputs.organizations.GetOrganizationNonMasterAccount[];
    /**
     * List of organization roots. All elements have these attributes:
     */
    readonly roots: outputs.organizations.GetOrganizationRoot[];
}
