// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * `aws_internet_gateway` provides details about a specific Internet Gateway.
 */
export function getInternetGateway(args?: GetInternetGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetInternetGatewayResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:ec2/getInternetGateway:getInternetGateway", {
        "filters": args.filters,
        "internetGatewayId": args.internetGatewayId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getInternetGateway.
 */
export interface GetInternetGatewayArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: { name: string, values: string[] }[];
    /**
     * The id of the specific Internet Gateway to retrieve.
     */
    readonly internetGatewayId?: string;
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired Internet Gateway.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getInternetGateway.
 */
export interface GetInternetGatewayResult {
    readonly attachments: { state: string, vpcId: string }[];
    readonly internetGatewayId: string;
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
