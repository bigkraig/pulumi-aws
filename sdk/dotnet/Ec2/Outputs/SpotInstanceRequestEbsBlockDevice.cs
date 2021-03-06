// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class SpotInstanceRequestEbsBlockDevice
    {
        /// <summary>
        /// Whether the volume should be destroyed
        /// on instance termination (Default: `true`).
        /// </summary>
        public readonly bool? DeleteOnTermination;
        /// <summary>
        /// The name of the device to mount.
        /// </summary>
        public readonly string DeviceName;
        /// <summary>
        /// Enables [EBS
        /// encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
        /// on the volume (Default: `false`). Cannot be used with `snapshot_id`. Must be configured to perform drift detection.
        /// </summary>
        public readonly bool? Encrypted;
        /// <summary>
        /// The amount of provisioned
        /// [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
        /// This must be set with a `volume_type` of `"io1/io2"`.
        /// </summary>
        public readonly int? Iops;
        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS Key to use when encrypting the volume. Must be configured to perform drift detection.
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// The Snapshot ID to mount.
        /// </summary>
        public readonly string? SnapshotId;
        public readonly string? VolumeId;
        /// <summary>
        /// The size of the volume in gibibytes (GiB).
        /// </summary>
        public readonly int? VolumeSize;
        /// <summary>
        /// The type of volume. Can be `"standard"`, `"gp2"`, `"io1"`
        /// or `"io2"`. (Default: `"gp2"`).
        /// </summary>
        public readonly string? VolumeType;

        [OutputConstructor]
        private SpotInstanceRequestEbsBlockDevice(
            bool? deleteOnTermination,

            string deviceName,

            bool? encrypted,

            int? iops,

            string? kmsKeyId,

            string? snapshotId,

            string? volumeId,

            int? volumeSize,

            string? volumeType)
        {
            DeleteOnTermination = deleteOnTermination;
            DeviceName = deviceName;
            Encrypted = encrypted;
            Iops = iops;
            KmsKeyId = kmsKeyId;
            SnapshotId = snapshotId;
            VolumeId = volumeId;
            VolumeSize = volumeSize;
            VolumeType = volumeType;
        }
    }
}
