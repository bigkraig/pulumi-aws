// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Instance extends lumi.NamedResource implements InstanceArgs {
    public readonly agentVersion?: string;
    public readonly amiId: string;
    public readonly architecture?: string;
    public readonly autoScalingType?: string;
    public readonly availabilityZone: string;
    public readonly createdAt: string;
    public readonly deleteEbs?: boolean;
    public readonly deleteEip?: boolean;
    public readonly ebsBlockDevice: { deleteOnTermination?: boolean, deviceName: string, iops: number, snapshotId: string, volumeSize: number, volumeType: string }[];
    public readonly ebsOptimized?: boolean;
    public readonly ec2InstanceId: string;
    public readonly ecsClusterArn: string;
    public readonly elasticIp: string;
    public readonly ephemeralBlockDevice: { deviceName: string, virtualName: string }[];
    public readonly hostname: string;
    public /*out*/ readonly instanceId: string;
    public readonly infrastructureClass: string;
    public readonly installUpdatesOnBoot?: boolean;
    public readonly instanceProfileArn: string;
    public readonly instanceType?: string;
    public readonly lastServiceErrorId: string;
    public readonly layerIds: string[];
    public readonly os: string;
    public readonly platform: string;
    public readonly privateDns: string;
    public readonly privateIp: string;
    public readonly publicDns: string;
    public readonly publicIp: string;
    public readonly registeredBy: string;
    public readonly reportedAgentVersion: string;
    public readonly reportedOsFamily: string;
    public readonly reportedOsName: string;
    public readonly reportedOsVersion: string;
    public readonly rootBlockDevice: { deleteOnTermination?: boolean, iops: number, volumeSize: number, volumeType: string }[];
    public readonly rootDeviceType: string;
    public readonly rootDeviceVolumeId: string;
    public readonly securityGroupIds: string[];
    public readonly sshHostDsaKeyFingerprint: string;
    public readonly sshHostRsaKeyFingerprint: string;
    public readonly sshKeyName: string;
    public readonly stackId: string;
    public readonly state?: string;
    public readonly status: string;
    public readonly subnetId: string;
    public readonly tenancy: string;
    public readonly virtualizationType: string;

    public static get(id: lumi.ID): Instance {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Instance[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: InstanceArgs) {
        super(name);
        this.agentVersion = args.agentVersion;
        this.amiId = args.amiId;
        this.architecture = args.architecture;
        this.autoScalingType = args.autoScalingType;
        this.availabilityZone = args.availabilityZone;
        this.createdAt = args.createdAt;
        this.deleteEbs = args.deleteEbs;
        this.deleteEip = args.deleteEip;
        this.ebsBlockDevice = args.ebsBlockDevice;
        this.ebsOptimized = args.ebsOptimized;
        this.ec2InstanceId = args.ec2InstanceId;
        this.ecsClusterArn = args.ecsClusterArn;
        this.elasticIp = args.elasticIp;
        this.ephemeralBlockDevice = args.ephemeralBlockDevice;
        this.hostname = args.hostname;
        this.infrastructureClass = args.infrastructureClass;
        this.installUpdatesOnBoot = args.installUpdatesOnBoot;
        this.instanceProfileArn = args.instanceProfileArn;
        this.instanceType = args.instanceType;
        this.lastServiceErrorId = args.lastServiceErrorId;
        if (lumirt.defaultIfComputed(args.layerIds, "") === undefined) {
            throw new Error("Property argument 'layerIds' is required, but was missing");
        }
        this.layerIds = args.layerIds;
        this.os = args.os;
        this.platform = args.platform;
        this.privateDns = args.privateDns;
        this.privateIp = args.privateIp;
        this.publicDns = args.publicDns;
        this.publicIp = args.publicIp;
        this.registeredBy = args.registeredBy;
        this.reportedAgentVersion = args.reportedAgentVersion;
        this.reportedOsFamily = args.reportedOsFamily;
        this.reportedOsName = args.reportedOsName;
        this.reportedOsVersion = args.reportedOsVersion;
        this.rootBlockDevice = args.rootBlockDevice;
        this.rootDeviceType = args.rootDeviceType;
        this.rootDeviceVolumeId = args.rootDeviceVolumeId;
        this.securityGroupIds = args.securityGroupIds;
        this.sshHostDsaKeyFingerprint = args.sshHostDsaKeyFingerprint;
        this.sshHostRsaKeyFingerprint = args.sshHostRsaKeyFingerprint;
        this.sshKeyName = args.sshKeyName;
        if (lumirt.defaultIfComputed(args.stackId, "") === undefined) {
            throw new Error("Property argument 'stackId' is required, but was missing");
        }
        this.stackId = args.stackId;
        this.state = args.state;
        this.status = args.status;
        this.subnetId = args.subnetId;
        this.tenancy = args.tenancy;
        this.virtualizationType = args.virtualizationType;
    }
}

export interface InstanceArgs {
    readonly agentVersion?: string;
    readonly amiId?: string;
    readonly architecture?: string;
    readonly autoScalingType?: string;
    readonly availabilityZone?: string;
    readonly createdAt?: string;
    readonly deleteEbs?: boolean;
    readonly deleteEip?: boolean;
    readonly ebsBlockDevice?: { deleteOnTermination?: boolean, deviceName: string, iops: number, snapshotId: string, volumeSize: number, volumeType: string }[];
    readonly ebsOptimized?: boolean;
    readonly ec2InstanceId?: string;
    readonly ecsClusterArn?: string;
    readonly elasticIp?: string;
    readonly ephemeralBlockDevice?: { deviceName: string, virtualName: string }[];
    readonly hostname?: string;
    readonly infrastructureClass?: string;
    readonly installUpdatesOnBoot?: boolean;
    readonly instanceProfileArn?: string;
    readonly instanceType?: string;
    readonly lastServiceErrorId?: string;
    readonly layerIds: string[];
    readonly os?: string;
    readonly platform?: string;
    readonly privateDns?: string;
    readonly privateIp?: string;
    readonly publicDns?: string;
    readonly publicIp?: string;
    readonly registeredBy?: string;
    readonly reportedAgentVersion?: string;
    readonly reportedOsFamily?: string;
    readonly reportedOsName?: string;
    readonly reportedOsVersion?: string;
    readonly rootBlockDevice?: { deleteOnTermination?: boolean, iops: number, volumeSize: number, volumeType: string }[];
    readonly rootDeviceType?: string;
    readonly rootDeviceVolumeId?: string;
    readonly securityGroupIds?: string[];
    readonly sshHostDsaKeyFingerprint?: string;
    readonly sshHostRsaKeyFingerprint?: string;
    readonly sshKeyName?: string;
    readonly stackId: string;
    readonly state?: string;
    readonly status?: string;
    readonly subnetId?: string;
    readonly tenancy?: string;
    readonly virtualizationType?: string;
}
