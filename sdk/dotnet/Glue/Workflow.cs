// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue
{
    /// <summary>
    /// Provides a Glue Workflow resource.
    /// The workflow graph (DAG) can be build using the `aws.glue.Trigger` resource. 
    /// See the example below for creating a graph with four nodes (two triggers and two jobs). 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/glue_workflow.html.markdown.
    /// </summary>
    public partial class Workflow : Pulumi.CustomResource
    {
        /// <summary>
        /// A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
        /// </summary>
        [Output("defaultRunProperties")]
        public Output<ImmutableDictionary<string, object>?> DefaultRunProperties { get; private set; } = null!;

        /// <summary>
        /// Description of the workflow.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name you assign to this workflow.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Workflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workflow(string name, WorkflowArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:glue/workflow:Workflow", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Workflow(string name, Input<string> id, WorkflowState? state = null, CustomResourceOptions? options = null)
            : base("aws:glue/workflow:Workflow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Workflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workflow Get(string name, Input<string> id, WorkflowState? state = null, CustomResourceOptions? options = null)
        {
            return new Workflow(name, id, state, options);
        }
    }

    public sealed class WorkflowArgs : Pulumi.ResourceArgs
    {
        [Input("defaultRunProperties")]
        private InputMap<object>? _defaultRunProperties;

        /// <summary>
        /// A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
        /// </summary>
        public InputMap<object> DefaultRunProperties
        {
            get => _defaultRunProperties ?? (_defaultRunProperties = new InputMap<object>());
            set => _defaultRunProperties = value;
        }

        /// <summary>
        /// Description of the workflow.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name you assign to this workflow.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public WorkflowArgs()
        {
        }
    }

    public sealed class WorkflowState : Pulumi.ResourceArgs
    {
        [Input("defaultRunProperties")]
        private InputMap<object>? _defaultRunProperties;

        /// <summary>
        /// A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
        /// </summary>
        public InputMap<object> DefaultRunProperties
        {
            get => _defaultRunProperties ?? (_defaultRunProperties = new InputMap<object>());
            set => _defaultRunProperties = value;
        }

        /// <summary>
        /// Description of the workflow.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name you assign to this workflow.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public WorkflowState()
        {
        }
    }
}
