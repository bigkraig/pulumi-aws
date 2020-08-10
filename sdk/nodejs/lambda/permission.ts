// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {Function} from "./index";

/**
 * Gives an external source (like a CloudWatch Event Rule, SNS, or S3) permission to access the Lambda function.
 *
 * ## Example Usage
 * ### Basic Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const iamForLambda = new aws.iam.Role("iamForLambda", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "lambda.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `});
 * const testLambda = new aws.lambda.Function("testLambda", {
 *     code: new pulumi.asset.FileArchive("lambdatest.zip"),
 *     role: iamForLambda.arn,
 *     handler: "exports.handler",
 *     runtime: "nodejs8.10",
 * });
 * const testAlias = new aws.lambda.Alias("testAlias", {
 *     description: "a sample description",
 *     functionName: testLambda.name,
 *     functionVersion: `$LATEST`,
 * });
 * const allowCloudwatch = new aws.lambda.Permission("allowCloudwatch", {
 *     action: "lambda:InvokeFunction",
 *     "function": testLambda.name,
 *     principal: "events.amazonaws.com",
 *     sourceArn: "arn:aws:events:eu-west-1:111122223333:rule/RunDaily",
 *     qualifier: testAlias.name,
 * });
 * ```
 * ### Usage with SNS
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const defaultTopic = new aws.sns.Topic("defaultTopic", {});
 * const defaultRole = new aws.iam.Role("defaultRole", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "lambda.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `});
 * const func = new aws.lambda.Function("func", {
 *     code: new pulumi.asset.FileArchive("lambdatest.zip"),
 *     role: defaultRole.arn,
 *     handler: "exports.handler",
 *     runtime: "python2.7",
 * });
 * const withSns = new aws.lambda.Permission("withSns", {
 *     action: "lambda:InvokeFunction",
 *     "function": func.name,
 *     principal: "sns.amazonaws.com",
 *     sourceArn: defaultTopic.arn,
 * });
 * const lambda = new aws.sns.TopicSubscription("lambda", {
 *     topic: defaultTopic.arn,
 *     protocol: "lambda",
 *     endpoint: func.arn,
 * });
 * ```
 * ### Specify Lambda permissions for API Gateway REST API
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myDemoAPI = new aws.apigateway.RestApi("MyDemoAPI", {
 *     description: "This is my API for demonstration purposes",
 * });
 * const lambdaPermission = new aws.lambda.Permission("lambda_permission", {
 *     action: "lambda:InvokeFunction",
 *     function: "MyDemoFunction",
 *     principal: "apigateway.amazonaws.com",
 *     sourceArn: pulumi.interpolate`${myDemoAPI.executionArn}/*&#47;*&#47;*`,
 * });
 * ```
 */
export class Permission extends pulumi.CustomResource {
    /**
     * Get an existing Permission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PermissionState, opts?: pulumi.CustomResourceOptions): Permission {
        return new Permission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lambda/permission:Permission';

    /**
     * Returns true if the given object is an instance of Permission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Permission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Permission.__pulumiType;
    }

    /**
     * The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
     */
    public readonly eventSourceToken!: pulumi.Output<string | undefined>;
    /**
     * Name of the Lambda function whose resource policy you are updating
     */
    public readonly function!: pulumi.Output<string>;
    /**
     * The principal who is getting this permission.
     * e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal
     * such as `events.amazonaws.com` or `sns.amazonaws.com`.
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * Query parameter to specify function version or alias name.
     * The permission will then apply to the specific qualified ARN.
     * e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
     */
    public readonly qualifier!: pulumi.Output<string | undefined>;
    /**
     * This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
     */
    public readonly sourceAccount!: pulumi.Output<string | undefined>;
    /**
     * When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
     * Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
     * For S3, this should be the ARN of the S3 Bucket.
     * For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
     * For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
     */
    public readonly sourceArn!: pulumi.Output<string | undefined>;
    /**
     * A unique statement identifier. By default generated by this provider.
     */
    public readonly statementId!: pulumi.Output<string>;
    /**
     * A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statementId`.
     */
    public readonly statementIdPrefix!: pulumi.Output<string | undefined>;

    /**
     * Create a Permission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PermissionArgs | PermissionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PermissionState | undefined;
            inputs["action"] = state ? state.action : undefined;
            inputs["eventSourceToken"] = state ? state.eventSourceToken : undefined;
            inputs["function"] = state ? state.function : undefined;
            inputs["principal"] = state ? state.principal : undefined;
            inputs["qualifier"] = state ? state.qualifier : undefined;
            inputs["sourceAccount"] = state ? state.sourceAccount : undefined;
            inputs["sourceArn"] = state ? state.sourceArn : undefined;
            inputs["statementId"] = state ? state.statementId : undefined;
            inputs["statementIdPrefix"] = state ? state.statementIdPrefix : undefined;
        } else {
            const args = argsOrState as PermissionArgs | undefined;
            if (!args || args.action === undefined) {
                throw new Error("Missing required property 'action'");
            }
            if (!args || args.function === undefined) {
                throw new Error("Missing required property 'function'");
            }
            if (!args || args.principal === undefined) {
                throw new Error("Missing required property 'principal'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["eventSourceToken"] = args ? args.eventSourceToken : undefined;
            inputs["function"] = args ? args.function : undefined;
            inputs["principal"] = args ? args.principal : undefined;
            inputs["qualifier"] = args ? args.qualifier : undefined;
            inputs["sourceAccount"] = args ? args.sourceAccount : undefined;
            inputs["sourceArn"] = args ? args.sourceArn : undefined;
            inputs["statementId"] = args ? args.statementId : undefined;
            inputs["statementIdPrefix"] = args ? args.statementIdPrefix : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Permission.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Permission resources.
 */
export interface PermissionState {
    /**
     * The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
     */
    readonly action?: pulumi.Input<string>;
    /**
     * The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
     */
    readonly eventSourceToken?: pulumi.Input<string>;
    /**
     * Name of the Lambda function whose resource policy you are updating
     */
    readonly function?: pulumi.Input<string | Function>;
    /**
     * The principal who is getting this permission.
     * e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal
     * such as `events.amazonaws.com` or `sns.amazonaws.com`.
     */
    readonly principal?: pulumi.Input<string>;
    /**
     * Query parameter to specify function version or alias name.
     * The permission will then apply to the specific qualified ARN.
     * e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
     */
    readonly qualifier?: pulumi.Input<string>;
    /**
     * This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
     */
    readonly sourceAccount?: pulumi.Input<string>;
    /**
     * When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
     * Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
     * For S3, this should be the ARN of the S3 Bucket.
     * For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
     * For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
     */
    readonly sourceArn?: pulumi.Input<string>;
    /**
     * A unique statement identifier. By default generated by this provider.
     */
    readonly statementId?: pulumi.Input<string>;
    /**
     * A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statementId`.
     */
    readonly statementIdPrefix?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Permission resource.
 */
export interface PermissionArgs {
    /**
     * The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
     */
    readonly action: pulumi.Input<string>;
    /**
     * The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
     */
    readonly eventSourceToken?: pulumi.Input<string>;
    /**
     * Name of the Lambda function whose resource policy you are updating
     */
    readonly function: pulumi.Input<string | Function>;
    /**
     * The principal who is getting this permission.
     * e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal
     * such as `events.amazonaws.com` or `sns.amazonaws.com`.
     */
    readonly principal: pulumi.Input<string>;
    /**
     * Query parameter to specify function version or alias name.
     * The permission will then apply to the specific qualified ARN.
     * e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
     */
    readonly qualifier?: pulumi.Input<string>;
    /**
     * This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
     */
    readonly sourceAccount?: pulumi.Input<string>;
    /**
     * When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
     * Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
     * For S3, this should be the ARN of the S3 Bucket.
     * For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
     * For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
     */
    readonly sourceArn?: pulumi.Input<string>;
    /**
     * A unique statement identifier. By default generated by this provider.
     */
    readonly statementId?: pulumi.Input<string>;
    /**
     * A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statementId`.
     */
    readonly statementIdPrefix?: pulumi.Input<string>;
}
