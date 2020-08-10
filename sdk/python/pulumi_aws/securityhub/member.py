# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Member(pulumi.CustomResource):
    account_id: pulumi.Output[str]
    """
    The ID of the member AWS account.
    """
    email: pulumi.Output[str]
    """
    The email of the member AWS account.
    """
    invite: pulumi.Output[bool]
    """
    Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
    """
    master_id: pulumi.Output[str]
    """
    The ID of the master Security Hub AWS account.
    """
    member_status: pulumi.Output[str]
    """
    The status of the member account relationship.
    """
    def __init__(__self__, resource_name, opts=None, account_id=None, email=None, invite=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Security Hub member resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_account = aws.securityhub.Account("exampleAccount")
        example_member = aws.securityhub.Member("exampleMember",
            account_id="123456789012",
            email="example@example.com",
            invite=True,
            opts=ResourceOptions(depends_on=[example_account]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of the member AWS account.
        :param pulumi.Input[str] email: The email of the member AWS account.
        :param pulumi.Input[bool] invite: Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if account_id is None:
                raise TypeError("Missing required property 'account_id'")
            __props__['account_id'] = account_id
            if email is None:
                raise TypeError("Missing required property 'email'")
            __props__['email'] = email
            __props__['invite'] = invite
            __props__['master_id'] = None
            __props__['member_status'] = None
        super(Member, __self__).__init__(
            'aws:securityhub/member:Member',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_id=None, email=None, invite=None, master_id=None, member_status=None):
        """
        Get an existing Member resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of the member AWS account.
        :param pulumi.Input[str] email: The email of the member AWS account.
        :param pulumi.Input[bool] invite: Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
        :param pulumi.Input[str] master_id: The ID of the master Security Hub AWS account.
        :param pulumi.Input[str] member_status: The status of the member account relationship.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_id"] = account_id
        __props__["email"] = email
        __props__["invite"] = invite
        __props__["master_id"] = master_id
        __props__["member_status"] = member_status
        return Member(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
