// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the default EBS encryption KMS key in the current region.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const current = aws.ebs.getDefaultKmsKey({});
 * const example = new aws.ebs.Volume("example", {
 *     availabilityZone: "us-west-2a",
 *     encrypted: true,
 *     kmsKeyId: current.then(current => current.keyArn),
 * });
 * ```
 */
export function getDefaultKmsKey(opts?: pulumi.InvokeOptions): Promise<GetDefaultKmsKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ebs/getDefaultKmsKey:getDefaultKmsKey", {
    }, opts);
}

/**
 * A collection of values returned by getDefaultKmsKey.
 */
export interface GetDefaultKmsKeyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Amazon Resource Name (ARN) of the default KMS key uses to encrypt an EBS volume in this region when no key is specified in an API call that creates the volume and encryption by default is enabled.
     */
    readonly keyArn: string;
}
