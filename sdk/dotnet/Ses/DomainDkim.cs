// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ses
{
    /// <summary>
    /// Provides an SES domain DKIM generation resource.
    /// 
    /// Domain ownership needs to be confirmed first using `aws.ses.DomainIdentity` resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleDomainIdentity = new Aws.Ses.DomainIdentity("exampleDomainIdentity", new Aws.Ses.DomainIdentityArgs
    ///         {
    ///             Domain = "example.com",
    ///         });
    ///         var exampleDomainDkim = new Aws.Ses.DomainDkim("exampleDomainDkim", new Aws.Ses.DomainDkimArgs
    ///         {
    ///             Domain = exampleDomainIdentity.Domain,
    ///         });
    ///         var exampleAmazonsesDkimRecord = new List&lt;Aws.Route53.Record&gt;();
    ///         for (var rangeIndex = 0; rangeIndex &lt; 3; rangeIndex++)
    ///         {
    ///             var range = new { Value = rangeIndex };
    ///             exampleAmazonsesDkimRecord.Add(new Aws.Route53.Record($"exampleAmazonsesDkimRecord-{range.Value}", new Aws.Route53.RecordArgs
    ///             {
    ///                 ZoneId = "ABCDEFGHIJ123",
    ///                 Name = exampleDomainDkim.DkimTokens[range.Value].Apply(dkimTokens =&gt; $"{dkimTokens}._domainkey.example.com"),
    ///                 Type = "CNAME",
    ///                 Ttl = 600,
    ///                 Records = 
    ///                 {
    ///                     exampleDomainDkim.DkimTokens[range.Value].Apply(dkimTokens =&gt; $"{dkimTokens}.dkim.amazonses.com"),
    ///                 },
    ///             }));
    ///         }
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class DomainDkim : Pulumi.CustomResource
    {
        /// <summary>
        /// DKIM tokens generated by SES.
        /// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
        /// See below for an example of how this might be achieved
        /// when the domain is hosted in Route 53 and managed by this provider.
        /// Find out more about verifying domains in Amazon SES
        /// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
        /// </summary>
        [Output("dkimTokens")]
        public Output<ImmutableArray<string>> DkimTokens { get; private set; } = null!;

        /// <summary>
        /// Verified domain name to generate DKIM tokens for.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;


        /// <summary>
        /// Create a DomainDkim resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainDkim(string name, DomainDkimArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/domainDkim:DomainDkim", name, args ?? new DomainDkimArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainDkim(string name, Input<string> id, DomainDkimState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/domainDkim:DomainDkim", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainDkim resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainDkim Get(string name, Input<string> id, DomainDkimState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainDkim(name, id, state, options);
        }
    }

    public sealed class DomainDkimArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Verified domain name to generate DKIM tokens for.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        public DomainDkimArgs()
        {
        }
    }

    public sealed class DomainDkimState : Pulumi.ResourceArgs
    {
        [Input("dkimTokens")]
        private InputList<string>? _dkimTokens;

        /// <summary>
        /// DKIM tokens generated by SES.
        /// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
        /// See below for an example of how this might be achieved
        /// when the domain is hosted in Route 53 and managed by this provider.
        /// Find out more about verifying domains in Amazon SES
        /// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
        /// </summary>
        public InputList<string> DkimTokens
        {
            get => _dkimTokens ?? (_dkimTokens = new InputList<string>());
            set => _dkimTokens = value;
        }

        /// <summary>
        /// Verified domain name to generate DKIM tokens for.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        public DomainDkimState()
        {
        }
    }
}
