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
    public sealed class LaunchTemplateNetworkInterface
    {
        /// <summary>
        /// Associate a public ip address with the network interface.  Boolean value.
        /// </summary>
        public readonly string? AssociatePublicIpAddress;
        /// <summary>
        /// Whether the network interface should be destroyed on instance termination. Defaults to `false` if not set.
        /// </summary>
        public readonly string? DeleteOnTermination;
        /// <summary>
        /// Description of the network interface.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The integer index of the network interface attachment.
        /// </summary>
        public readonly int? DeviceIndex;
        /// <summary>
        /// The number of secondary private IPv4 addresses to assign to a network interface. Conflicts with `ipv4_addresses`
        /// </summary>
        public readonly int? Ipv4AddressCount;
        /// <summary>
        /// One or more private IPv4 addresses to associate. Conflicts with `ipv4_address_count`
        /// </summary>
        public readonly ImmutableArray<string> Ipv4Addresses;
        /// <summary>
        /// The number of IPv6 addresses to assign to a network interface. Conflicts with `ipv6_addresses`
        /// </summary>
        public readonly int? Ipv6AddressCount;
        /// <summary>
        /// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Conflicts with `ipv6_address_count`
        /// </summary>
        public readonly ImmutableArray<string> Ipv6Addresses;
        /// <summary>
        /// The ID of the network interface to attach.
        /// </summary>
        public readonly string? NetworkInterfaceId;
        /// <summary>
        /// The primary private IPv4 address.
        /// </summary>
        public readonly string? PrivateIpAddress;
        /// <summary>
        /// A list of security group IDs to associate.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// The VPC Subnet ID to associate.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private LaunchTemplateNetworkInterface(
            string? associatePublicIpAddress,

            string? deleteOnTermination,

            string? description,

            int? deviceIndex,

            int? ipv4AddressCount,

            ImmutableArray<string> ipv4Addresses,

            int? ipv6AddressCount,

            ImmutableArray<string> ipv6Addresses,

            string? networkInterfaceId,

            string? privateIpAddress,

            ImmutableArray<string> securityGroups,

            string? subnetId)
        {
            AssociatePublicIpAddress = associatePublicIpAddress;
            DeleteOnTermination = deleteOnTermination;
            Description = description;
            DeviceIndex = deviceIndex;
            Ipv4AddressCount = ipv4AddressCount;
            Ipv4Addresses = ipv4Addresses;
            Ipv6AddressCount = ipv6AddressCount;
            Ipv6Addresses = ipv6Addresses;
            NetworkInterfaceId = networkInterfaceId;
            PrivateIpAddress = privateIpAddress;
            SecurityGroups = securityGroups;
            SubnetId = subnetId;
        }
    }
}
