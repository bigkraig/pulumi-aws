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
    /// Provides a proxy protocol policy, which allows an ELB to carry a client connection information to a backend.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var lb = new Aws.Elb.LoadBalancer("lb", new Aws.Elb.LoadBalancerArgs
    ///         {
    ///             AvailabilityZones = 
    ///             {
    ///                 "us-east-1a",
    ///             },
    ///             Listeners = 
    ///             {
    ///                 new Aws.Elb.Inputs.LoadBalancerListenerArgs
    ///                 {
    ///                     InstancePort = 25,
    ///                     InstanceProtocol = "tcp",
    ///                     LbPort = 25,
    ///                     LbProtocol = "tcp",
    ///                 },
    ///                 new Aws.Elb.Inputs.LoadBalancerListenerArgs
    ///                 {
    ///                     InstancePort = 587,
    ///                     InstanceProtocol = "tcp",
    ///                     LbPort = 587,
    ///                     LbProtocol = "tcp",
    ///                 },
    ///             },
    ///         });
    ///         var smtp = new Aws.Ec2.ProxyProtocolPolicy("smtp", new Aws.Ec2.ProxyProtocolPolicyArgs
    ///         {
    ///             LoadBalancer = lb.Name,
    ///             InstancePorts = 
    ///             {
    ///                 "25",
    ///                 "587",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ProxyProtocolPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// List of instance ports to which the policy
        /// should be applied. This can be specified if the protocol is SSL or TCP.
        /// </summary>
        [Output("instancePorts")]
        public Output<ImmutableArray<string>> InstancePorts { get; private set; } = null!;

        /// <summary>
        /// The load balancer to which the policy
        /// should be attached.
        /// </summary>
        [Output("loadBalancer")]
        public Output<string> LoadBalancer { get; private set; } = null!;


        /// <summary>
        /// Create a ProxyProtocolPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProxyProtocolPolicy(string name, ProxyProtocolPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/proxyProtocolPolicy:ProxyProtocolPolicy", name, args ?? new ProxyProtocolPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProxyProtocolPolicy(string name, Input<string> id, ProxyProtocolPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/proxyProtocolPolicy:ProxyProtocolPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProxyProtocolPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProxyProtocolPolicy Get(string name, Input<string> id, ProxyProtocolPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ProxyProtocolPolicy(name, id, state, options);
        }
    }

    public sealed class ProxyProtocolPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("instancePorts", required: true)]
        private InputList<string>? _instancePorts;

        /// <summary>
        /// List of instance ports to which the policy
        /// should be applied. This can be specified if the protocol is SSL or TCP.
        /// </summary>
        public InputList<string> InstancePorts
        {
            get => _instancePorts ?? (_instancePorts = new InputList<string>());
            set => _instancePorts = value;
        }

        /// <summary>
        /// The load balancer to which the policy
        /// should be attached.
        /// </summary>
        [Input("loadBalancer", required: true)]
        public Input<string> LoadBalancer { get; set; } = null!;

        public ProxyProtocolPolicyArgs()
        {
        }
    }

    public sealed class ProxyProtocolPolicyState : Pulumi.ResourceArgs
    {
        [Input("instancePorts")]
        private InputList<string>? _instancePorts;

        /// <summary>
        /// List of instance ports to which the policy
        /// should be applied. This can be specified if the protocol is SSL or TCP.
        /// </summary>
        public InputList<string> InstancePorts
        {
            get => _instancePorts ?? (_instancePorts = new InputList<string>());
            set => _instancePorts = value;
        }

        /// <summary>
        /// The load balancer to which the policy
        /// should be attached.
        /// </summary>
        [Input("loadBalancer")]
        public Input<string>? LoadBalancer { get; set; }

        public ProxyProtocolPolicyState()
        {
        }
    }
}
