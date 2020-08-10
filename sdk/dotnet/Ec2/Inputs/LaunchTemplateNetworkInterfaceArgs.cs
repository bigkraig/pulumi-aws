// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class LaunchTemplateNetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Associate a public ip address with the network interface.  Boolean value.
        /// </summary>
        [Input("associatePublicIpAddress")]
        public Input<string>? AssociatePublicIpAddress { get; set; }

        /// <summary>
        /// Whether the network interface should be destroyed on instance termination. Defaults to `false` if not set.
        /// </summary>
        [Input("deleteOnTermination")]
        public Input<string>? DeleteOnTermination { get; set; }

        /// <summary>
        /// Description of the network interface.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The integer index of the network interface attachment.
        /// </summary>
        [Input("deviceIndex")]
        public Input<int>? DeviceIndex { get; set; }

        /// <summary>
        /// The number of secondary private IPv4 addresses to assign to a network interface. Conflicts with `ipv4_addresses`
        /// </summary>
        [Input("ipv4AddressCount")]
        public Input<int>? Ipv4AddressCount { get; set; }

        [Input("ipv4Addresses")]
        private InputList<string>? _ipv4Addresses;

        /// <summary>
        /// One or more private IPv4 addresses to associate. Conflicts with `ipv4_address_count`
        /// </summary>
        public InputList<string> Ipv4Addresses
        {
            get => _ipv4Addresses ?? (_ipv4Addresses = new InputList<string>());
            set => _ipv4Addresses = value;
        }

        /// <summary>
        /// The number of IPv6 addresses to assign to a network interface. Conflicts with `ipv6_addresses`
        /// </summary>
        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;

        /// <summary>
        /// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Conflicts with `ipv6_address_count`
        /// </summary>
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// The ID of the network interface to attach.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The primary private IPv4 address.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security group IDs to associate.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// The VPC Subnet ID to associate.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public LaunchTemplateNetworkInterfaceArgs()
        {
        }
    }
}
