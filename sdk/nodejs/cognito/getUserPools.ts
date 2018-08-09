// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get a list of cognito user pools.
 */
export function getUserPools(args: GetUserPoolsArgs, opts?: pulumi.InvokeOptions): Promise<GetUserPoolsResult> {
    return pulumi.runtime.invoke("aws:cognito/getUserPools:getUserPools", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserPools.
 */
export interface GetUserPoolsArgs {
    /**
     * Name of the cognito user pools. Name is not a unique attribute for cognito user pool, so multiple pools might be returned with given name.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getUserPools.
 */
export interface GetUserPoolsResult {
    readonly arns: string[];
    /**
     * The list of cognito user pool ids.
     */
    readonly ids: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
