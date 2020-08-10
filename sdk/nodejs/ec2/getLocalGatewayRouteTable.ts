// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides details about an EC2 Local Gateway Route Table.
 *
 * This data source can prove useful when a module accepts a local gateway route table id as
 * an input variable and needs to, for example, find the associated Outpost or Local Gateway.
 *
 * ## Example Usage
 *
 * The following example returns a specific local gateway route table ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const awsEc2LocalGatewayRouteTable = config.requireObject("awsEc2LocalGatewayRouteTable");
 * const selected = aws.ec2.getLocalGatewayRouteTable({
 *     localGatewayRouteTableId: awsEc2LocalGatewayRouteTable,
 * });
 * ```
 */
export function getLocalGatewayRouteTable(args?: GetLocalGatewayRouteTableArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewayRouteTableResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getLocalGatewayRouteTable:getLocalGatewayRouteTable", {
        "filters": args.filters,
        "localGatewayId": args.localGatewayId,
        "localGatewayRouteTableId": args.localGatewayRouteTableId,
        "outpostArn": args.outpostArn,
        "state": args.state,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalGatewayRouteTable.
 */
export interface GetLocalGatewayRouteTableArgs {
    readonly filters?: inputs.ec2.GetLocalGatewayRouteTableFilter[];
    /**
     * The id of the specific local gateway route table to retrieve.
     */
    readonly localGatewayId?: string;
    /**
     * Local Gateway Route Table Id assigned to desired local gateway route table
     */
    readonly localGatewayRouteTableId?: string;
    /**
     * The arn of the Outpost the local gateway route table is associated with.
     */
    readonly outpostArn?: string;
    /**
     * The state of the local gateway route table.
     */
    readonly state?: string;
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired local gateway route table.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLocalGatewayRouteTable.
 */
export interface GetLocalGatewayRouteTableResult {
    readonly filters?: outputs.ec2.GetLocalGatewayRouteTableFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly localGatewayId: string;
    readonly localGatewayRouteTableId: string;
    readonly outpostArn: string;
    readonly state: string;
    readonly tags: {[key: string]: string};
}
