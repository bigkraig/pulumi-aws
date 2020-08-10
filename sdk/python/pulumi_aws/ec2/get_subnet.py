# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSubnetResult:
    """
    A collection of values returned by getSubnet.
    """
    def __init__(__self__, arn=None, assign_ipv6_address_on_creation=None, availability_zone=None, availability_zone_id=None, cidr_block=None, default_for_az=None, filters=None, id=None, ipv6_cidr_block=None, ipv6_cidr_block_association_id=None, map_public_ip_on_launch=None, outpost_arn=None, owner_id=None, state=None, tags=None, vpc_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The ARN of the subnet.
        """
        if assign_ipv6_address_on_creation and not isinstance(assign_ipv6_address_on_creation, bool):
            raise TypeError("Expected argument 'assign_ipv6_address_on_creation' to be a bool")
        __self__.assign_ipv6_address_on_creation = assign_ipv6_address_on_creation
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        __self__.availability_zone = availability_zone
        if availability_zone_id and not isinstance(availability_zone_id, str):
            raise TypeError("Expected argument 'availability_zone_id' to be a str")
        __self__.availability_zone_id = availability_zone_id
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError("Expected argument 'cidr_block' to be a str")
        __self__.cidr_block = cidr_block
        if default_for_az and not isinstance(default_for_az, bool):
            raise TypeError("Expected argument 'default_for_az' to be a bool")
        __self__.default_for_az = default_for_az
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        if ipv6_cidr_block and not isinstance(ipv6_cidr_block, str):
            raise TypeError("Expected argument 'ipv6_cidr_block' to be a str")
        __self__.ipv6_cidr_block = ipv6_cidr_block
        if ipv6_cidr_block_association_id and not isinstance(ipv6_cidr_block_association_id, str):
            raise TypeError("Expected argument 'ipv6_cidr_block_association_id' to be a str")
        __self__.ipv6_cidr_block_association_id = ipv6_cidr_block_association_id
        if map_public_ip_on_launch and not isinstance(map_public_ip_on_launch, bool):
            raise TypeError("Expected argument 'map_public_ip_on_launch' to be a bool")
        __self__.map_public_ip_on_launch = map_public_ip_on_launch
        if outpost_arn and not isinstance(outpost_arn, str):
            raise TypeError("Expected argument 'outpost_arn' to be a str")
        __self__.outpost_arn = outpost_arn
        """
        The Amazon Resource Name (ARN) of the Outpost.
        """
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        __self__.owner_id = owner_id
        """
        The ID of the AWS account that owns the subnet.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id
class AwaitableGetSubnetResult(GetSubnetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetResult(
            arn=self.arn,
            assign_ipv6_address_on_creation=self.assign_ipv6_address_on_creation,
            availability_zone=self.availability_zone,
            availability_zone_id=self.availability_zone_id,
            cidr_block=self.cidr_block,
            default_for_az=self.default_for_az,
            filters=self.filters,
            id=self.id,
            ipv6_cidr_block=self.ipv6_cidr_block,
            ipv6_cidr_block_association_id=self.ipv6_cidr_block_association_id,
            map_public_ip_on_launch=self.map_public_ip_on_launch,
            outpost_arn=self.outpost_arn,
            owner_id=self.owner_id,
            state=self.state,
            tags=self.tags,
            vpc_id=self.vpc_id)

def get_subnet(availability_zone=None,availability_zone_id=None,cidr_block=None,default_for_az=None,filters=None,id=None,ipv6_cidr_block=None,state=None,tags=None,vpc_id=None,opts=None):
    """
    `ec2.Subnet` provides details about a specific VPC subnet.

    This resource can prove useful when a module accepts a subnet id as
    an input variable and needs to, for example, determine the id of the
    VPC that the subnet belongs to.

    ## Example Usage

    The following example shows how one might accept a subnet id as a variable
    and use this data source to obtain the data necessary to create a security
    group that allows connections from hosts in that subnet.

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    subnet_id = config.require_object("subnetId")
    selected = aws.ec2.get_subnet(id=subnet_id)
    subnet = aws.ec2.SecurityGroup("subnet",
        vpc_id=selected.vpc_id,
        ingress=[{
            "cidr_blocks": [selected.cidr_block],
            "from_port": 80,
            "to_port": 80,
            "protocol": "tcp",
        }])
    ```


    :param str availability_zone: The availability zone where the
           subnet must reside.
    :param str availability_zone_id: The ID of the Availability Zone for the subnet.
    :param str cidr_block: The cidr block of the desired subnet.
    :param bool default_for_az: Boolean constraint for whether the desired
           subnet must be the default subnet for its associated availability zone.
    :param list filters: Custom filter block as described below.
    :param str id: The id of the specific subnet to retrieve.
    :param str ipv6_cidr_block: The Ipv6 cidr block of the desired subnet
    :param str state: The state that the desired subnet must have.
    :param dict tags: A map of tags, each pair of which must exactly match
           a pair on the desired subnet.
    :param str vpc_id: The id of the VPC that the desired subnet belongs to.

    The **filters** object supports the following:

      * `name` (`str`) - The name of the field to filter by, as defined by
        [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSubnets.html).
        For example, if matching against tag `Name`, use:
      * `values` (`list`) - Set of values that are accepted for the given field.
        A subnet will be selected if any one of the given values matches.
    """
    __args__ = dict()


    __args__['availabilityZone'] = availability_zone
    __args__['availabilityZoneId'] = availability_zone_id
    __args__['cidrBlock'] = cidr_block
    __args__['defaultForAz'] = default_for_az
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['ipv6CidrBlock'] = ipv6_cidr_block
    __args__['state'] = state
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getSubnet:getSubnet', __args__, opts=opts).value

    return AwaitableGetSubnetResult(
        arn=__ret__.get('arn'),
        assign_ipv6_address_on_creation=__ret__.get('assignIpv6AddressOnCreation'),
        availability_zone=__ret__.get('availabilityZone'),
        availability_zone_id=__ret__.get('availabilityZoneId'),
        cidr_block=__ret__.get('cidrBlock'),
        default_for_az=__ret__.get('defaultForAz'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        ipv6_cidr_block=__ret__.get('ipv6CidrBlock'),
        ipv6_cidr_block_association_id=__ret__.get('ipv6CidrBlockAssociationId'),
        map_public_ip_on_launch=__ret__.get('mapPublicIpOnLaunch'),
        outpost_arn=__ret__.get('outpostArn'),
        owner_id=__ret__.get('ownerId'),
        state=__ret__.get('state'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'))
