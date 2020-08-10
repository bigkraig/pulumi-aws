# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Workgroup(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the workgroup
    """
    configuration: pulumi.Output[dict]
    """
    Configuration block with various settings for the workgroup. Documented below.

      * `bytesScannedCutoffPerQuery` (`float`) - Integer for the upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan. Must be at least `10485760`.
      * `enforceWorkgroupConfiguration` (`bool`) - Boolean whether the settings for the workgroup override client-side settings. For more information, see [Workgroup Settings Override Client-Side Settings](https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html). Defaults to `true`.
      * `publishCloudwatchMetricsEnabled` (`bool`) - Boolean whether Amazon CloudWatch metrics are enabled for the workgroup. Defaults to `true`.
      * `resultConfiguration` (`dict`) - Configuration block with result settings. Documented below.
        * `encryption_configuration` (`dict`) - Configuration block with encryption settings. Documented below.
          * `encryptionOption` (`str`) - Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (`SSE_S3`), server-side encryption with KMS-managed keys (`SSE_KMS`), or client-side encryption with KMS-managed keys (`CSE_KMS`) is used. If a query runs in a workgroup and the workgroup overrides client-side settings, then the workgroup's setting for encryption is used. It specifies whether query results must be encrypted, for all queries that run in this workgroup.
          * `kms_key_arn` (`str`) - For `SSE_KMS` and `CSE_KMS`, this is the KMS key Amazon Resource Name (ARN).

        * `output_location` (`str`) - The location in Amazon S3 where your query results are stored, such as `s3://path/to/query/bucket/`. For more information, see [Queries and Query Result Files](https://docs.aws.amazon.com/athena/latest/ug/querying.html).
    """
    description: pulumi.Output[str]
    """
    Description of the workgroup.
    """
    force_destroy: pulumi.Output[bool]
    """
    The option to delete the workgroup and its contents even if the workgroup contains any named queries.
    """
    name: pulumi.Output[str]
    """
    Name of the workgroup.
    """
    state: pulumi.Output[str]
    """
    State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
    """
    tags: pulumi.Output[dict]
    """
    Key-value map of resource tags for the workgroup.
    """
    def __init__(__self__, resource_name, opts=None, configuration=None, description=None, force_destroy=None, name=None, state=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an Athena Workgroup.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.athena.Workgroup("example", configuration={
            "enforceWorkgroupConfiguration": True,
            "publishCloudwatchMetricsEnabled": True,
            "resultConfiguration": {
                "output_location": "s3://{aws_s3_bucket.example.bucket}/output/",
                "encryption_configuration": {
                    "encryptionOption": "SSE_KMS",
                    "kms_key_arn": aws_kms_key["example"]["arn"],
                },
            },
        })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] configuration: Configuration block with various settings for the workgroup. Documented below.
        :param pulumi.Input[str] description: Description of the workgroup.
        :param pulumi.Input[bool] force_destroy: The option to delete the workgroup and its contents even if the workgroup contains any named queries.
        :param pulumi.Input[str] name: Name of the workgroup.
        :param pulumi.Input[str] state: State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
        :param pulumi.Input[dict] tags: Key-value map of resource tags for the workgroup.

        The **configuration** object supports the following:

          * `bytesScannedCutoffPerQuery` (`pulumi.Input[float]`) - Integer for the upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan. Must be at least `10485760`.
          * `enforceWorkgroupConfiguration` (`pulumi.Input[bool]`) - Boolean whether the settings for the workgroup override client-side settings. For more information, see [Workgroup Settings Override Client-Side Settings](https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html). Defaults to `true`.
          * `publishCloudwatchMetricsEnabled` (`pulumi.Input[bool]`) - Boolean whether Amazon CloudWatch metrics are enabled for the workgroup. Defaults to `true`.
          * `resultConfiguration` (`pulumi.Input[dict]`) - Configuration block with result settings. Documented below.
            * `encryption_configuration` (`pulumi.Input[dict]`) - Configuration block with encryption settings. Documented below.
              * `encryptionOption` (`pulumi.Input[str]`) - Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (`SSE_S3`), server-side encryption with KMS-managed keys (`SSE_KMS`), or client-side encryption with KMS-managed keys (`CSE_KMS`) is used. If a query runs in a workgroup and the workgroup overrides client-side settings, then the workgroup's setting for encryption is used. It specifies whether query results must be encrypted, for all queries that run in this workgroup.
              * `kms_key_arn` (`pulumi.Input[str]`) - For `SSE_KMS` and `CSE_KMS`, this is the KMS key Amazon Resource Name (ARN).

            * `output_location` (`pulumi.Input[str]`) - The location in Amazon S3 where your query results are stored, such as `s3://path/to/query/bucket/`. For more information, see [Queries and Query Result Files](https://docs.aws.amazon.com/athena/latest/ug/querying.html).
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

            __props__['configuration'] = configuration
            __props__['description'] = description
            __props__['force_destroy'] = force_destroy
            __props__['name'] = name
            __props__['state'] = state
            __props__['tags'] = tags
            __props__['arn'] = None
        super(Workgroup, __self__).__init__(
            'aws:athena/workgroup:Workgroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, configuration=None, description=None, force_destroy=None, name=None, state=None, tags=None):
        """
        Get an existing Workgroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the workgroup
        :param pulumi.Input[dict] configuration: Configuration block with various settings for the workgroup. Documented below.
        :param pulumi.Input[str] description: Description of the workgroup.
        :param pulumi.Input[bool] force_destroy: The option to delete the workgroup and its contents even if the workgroup contains any named queries.
        :param pulumi.Input[str] name: Name of the workgroup.
        :param pulumi.Input[str] state: State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
        :param pulumi.Input[dict] tags: Key-value map of resource tags for the workgroup.

        The **configuration** object supports the following:

          * `bytesScannedCutoffPerQuery` (`pulumi.Input[float]`) - Integer for the upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan. Must be at least `10485760`.
          * `enforceWorkgroupConfiguration` (`pulumi.Input[bool]`) - Boolean whether the settings for the workgroup override client-side settings. For more information, see [Workgroup Settings Override Client-Side Settings](https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html). Defaults to `true`.
          * `publishCloudwatchMetricsEnabled` (`pulumi.Input[bool]`) - Boolean whether Amazon CloudWatch metrics are enabled for the workgroup. Defaults to `true`.
          * `resultConfiguration` (`pulumi.Input[dict]`) - Configuration block with result settings. Documented below.
            * `encryption_configuration` (`pulumi.Input[dict]`) - Configuration block with encryption settings. Documented below.
              * `encryptionOption` (`pulumi.Input[str]`) - Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (`SSE_S3`), server-side encryption with KMS-managed keys (`SSE_KMS`), or client-side encryption with KMS-managed keys (`CSE_KMS`) is used. If a query runs in a workgroup and the workgroup overrides client-side settings, then the workgroup's setting for encryption is used. It specifies whether query results must be encrypted, for all queries that run in this workgroup.
              * `kms_key_arn` (`pulumi.Input[str]`) - For `SSE_KMS` and `CSE_KMS`, this is the KMS key Amazon Resource Name (ARN).

            * `output_location` (`pulumi.Input[str]`) - The location in Amazon S3 where your query results are stored, such as `s3://path/to/query/bucket/`. For more information, see [Queries and Query Result Files](https://docs.aws.amazon.com/athena/latest/ug/querying.html).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["configuration"] = configuration
        __props__["description"] = description
        __props__["force_destroy"] = force_destroy
        __props__["name"] = name
        __props__["state"] = state
        __props__["tags"] = tags
        return Workgroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
