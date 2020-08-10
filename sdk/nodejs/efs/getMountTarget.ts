// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides information about an Elastic File System Mount Target (EFS).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const mountTargetId = config.get("mountTargetId") || "";
 * const byId = aws.efs.getMountTarget({
 *     mountTargetId: mountTargetId,
 * });
 * ```
 */
export function getMountTarget(args: GetMountTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetMountTargetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:efs/getMountTarget:getMountTarget", {
        "mountTargetId": args.mountTargetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMountTarget.
 */
export interface GetMountTargetArgs {
    /**
     * ID of the mount target that you want to have described
     */
    readonly mountTargetId: string;
}

/**
 * A collection of values returned by getMountTarget.
 */
export interface GetMountTargetResult {
    /**
     * The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
     */
    readonly availabilityZoneId: string;
    /**
     * The name of the Availability Zone (AZ) that the mount target resides in.
     */
    readonly availabilityZoneName: string;
    /**
     * The DNS name for the EFS file system.
     */
    readonly dnsName: string;
    /**
     * Amazon Resource Name of the file system for which the mount target is intended.
     */
    readonly fileSystemArn: string;
    /**
     * ID of the file system for which the mount target is intended.
     */
    readonly fileSystemId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Address at which the file system may be mounted via the mount target.
     */
    readonly ipAddress: string;
    /**
     * The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
     */
    readonly mountTargetDnsName: string;
    readonly mountTargetId: string;
    /**
     * The ID of the network interface that Amazon EFS created when it created the mount target.
     */
    readonly networkInterfaceId: string;
    /**
     * AWS account ID that owns the resource.
     */
    readonly ownerId: string;
    /**
     * List of VPC security group IDs attached to the mount target.
     */
    readonly securityGroups: string[];
    /**
     * ID of the mount target's subnet.
     */
    readonly subnetId: string;
}
