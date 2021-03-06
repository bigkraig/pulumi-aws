# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['LogService']


class LogService(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 directory_id: Optional[pulumi.Input[str]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Log subscription for AWS Directory Service that pushes logs to cloudwatch.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup", retention_in_days=14)
        ad_log_policy_policy_document = example_log_group.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=[
                "logs:CreateLogStream",
                "logs:PutLogEvents",
            ],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["ds.amazonaws.com"],
                type="Service",
            )],
            resources=[f"{arn}:*"],
            effect="Allow",
        )]))
        ad_log_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("ad-log-policyLogResourcePolicy",
            policy_document=ad_log_policy_policy_document.json,
            policy_name="ad-log-policy")
        example_log_service = aws.directoryservice.LogService("exampleLogService",
            directory_id=aws_directory_service_directory["example"]["id"],
            log_group_name=example_log_group.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] directory_id: The id of directory.
        :param pulumi.Input[str] log_group_name: Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if directory_id is None:
                raise TypeError("Missing required property 'directory_id'")
            __props__['directory_id'] = directory_id
            if log_group_name is None:
                raise TypeError("Missing required property 'log_group_name'")
            __props__['log_group_name'] = log_group_name
        super(LogService, __self__).__init__(
            'aws:directoryservice/logService:LogService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            directory_id: Optional[pulumi.Input[str]] = None,
            log_group_name: Optional[pulumi.Input[str]] = None) -> 'LogService':
        """
        Get an existing LogService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] directory_id: The id of directory.
        :param pulumi.Input[str] log_group_name: Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["directory_id"] = directory_id
        __props__["log_group_name"] = log_group_name
        return LogService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> pulumi.Output[str]:
        """
        The id of directory.
        """
        return pulumi.get(self, "directory_id")

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> pulumi.Output[str]:
        """
        Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
        """
        return pulumi.get(self, "log_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

