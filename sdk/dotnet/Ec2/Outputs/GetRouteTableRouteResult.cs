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
    public sealed class GetRouteTableRouteResult
    {
        /// <summary>
        /// The CIDR block of the route.
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// The ID of the Egress Only Internet Gateway.
        /// </summary>
        public readonly string EgressOnlyGatewayId;
        /// <summary>
        /// The id of an Internet Gateway or Virtual Private Gateway which is connected to the Route Table (not exported if not passed as a parameter).
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// The EC2 instance ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The IPv6 CIDR block of the route.
        /// </summary>
        public readonly string Ipv6CidrBlock;
        /// <summary>
        /// The Local Gateway ID.
        /// </summary>
        public readonly string LocalGatewayId;
        /// <summary>
        /// The NAT Gateway ID.
        /// </summary>
        public readonly string NatGatewayId;
        /// <summary>
        /// The ID of the elastic network interface (eni) to use.
        /// </summary>
        public readonly string NetworkInterfaceId;
        /// <summary>
        /// The EC2 Transit Gateway ID.
        /// </summary>
        public readonly string TransitGatewayId;
        /// <summary>
        /// The VPC Peering ID.
        /// </summary>
        public readonly string VpcPeeringConnectionId;

        [OutputConstructor]
        private GetRouteTableRouteResult(
            string cidrBlock,

            string egressOnlyGatewayId,

            string gatewayId,

            string instanceId,

            string ipv6CidrBlock,

            string localGatewayId,

            string natGatewayId,

            string networkInterfaceId,

            string transitGatewayId,

            string vpcPeeringConnectionId)
        {
            CidrBlock = cidrBlock;
            EgressOnlyGatewayId = egressOnlyGatewayId;
            GatewayId = gatewayId;
            InstanceId = instanceId;
            Ipv6CidrBlock = ipv6CidrBlock;
            LocalGatewayId = localGatewayId;
            NatGatewayId = natGatewayId;
            NetworkInterfaceId = networkInterfaceId;
            TransitGatewayId = transitGatewayId;
            VpcPeeringConnectionId = vpcPeeringConnectionId;
        }
    }
}
