// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Manages an EC2 Availability Zone Group, such as updating its opt-in status.
    /// 
    /// &gt; **NOTE:** This is an advanced resource. The provider will automatically assume management of the EC2 Availability Zone Group without import and perform no actions on removal from configuration.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_availability_zone_group.html.markdown.
    /// </summary>
    public partial class AvailabilityZoneGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the Availability Zone Group.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to enable or disable Availability Zone Group. Valid values: `opted-in` or `not-opted-in`.
        /// </summary>
        [Output("optInStatus")]
        public Output<string> OptInStatus { get; private set; } = null!;


        /// <summary>
        /// Create a AvailabilityZoneGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AvailabilityZoneGroup(string name, AvailabilityZoneGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/availabilityZoneGroup:AvailabilityZoneGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AvailabilityZoneGroup(string name, Input<string> id, AvailabilityZoneGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/availabilityZoneGroup:AvailabilityZoneGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AvailabilityZoneGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AvailabilityZoneGroup Get(string name, Input<string> id, AvailabilityZoneGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AvailabilityZoneGroup(name, id, state, options);
        }
    }

    public sealed class AvailabilityZoneGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Availability Zone Group.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Indicates whether to enable or disable Availability Zone Group. Valid values: `opted-in` or `not-opted-in`.
        /// </summary>
        [Input("optInStatus", required: true)]
        public Input<string> OptInStatus { get; set; } = null!;

        public AvailabilityZoneGroupArgs()
        {
        }
    }

    public sealed class AvailabilityZoneGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Availability Zone Group.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// Indicates whether to enable or disable Availability Zone Group. Valid values: `opted-in` or `not-opted-in`.
        /// </summary>
        [Input("optInStatus")]
        public Input<string>? OptInStatus { get; set; }

        public AvailabilityZoneGroupState()
        {
        }
    }
}
