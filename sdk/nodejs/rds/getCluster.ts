// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Provides information about a RDS cluster.
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    return pulumi.runtime.invoke("aws:rds/getCluster:getCluster", {
        "clusterIdentifier": args.clusterIdentifier,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * The cluster identifier of the RDS cluster.
     */
    readonly clusterIdentifier: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    readonly arn: string;
    readonly availabilityZones: string[];
    readonly backupRetentionPeriod: number;
    readonly clusterMembers: string[];
    readonly clusterResourceId: string;
    readonly databaseName: string;
    readonly dbClusterParameterGroupName: string;
    readonly dbSubnetGroupName: string;
    readonly enabledCloudwatchLogsExports: string[];
    readonly endpoint: string;
    readonly engine: string;
    readonly engineVersion: string;
    readonly finalSnapshotIdentifier: string;
    readonly iamDatabaseAuthenticationEnabled: boolean;
    readonly iamRoles: string[];
    readonly kmsKeyId: string;
    readonly masterUsername: string;
    readonly port: number;
    readonly preferredBackupWindow: string;
    readonly preferredMaintenanceWindow: string;
    readonly readerEndpoint: string;
    readonly replicationSourceIdentifier: string;
    readonly storageEncrypted: boolean;
    readonly tags: {[key: string]: any};
    readonly vpcSecurityGroupIds: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
