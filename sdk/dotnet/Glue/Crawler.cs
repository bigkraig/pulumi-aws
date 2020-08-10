// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue
{
    /// <summary>
    /// Manages a Glue Crawler. More information can be found in the [AWS Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html)
    /// 
    /// ## Example Usage
    /// ### DynamoDB Target
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Glue.Crawler("example", new Aws.Glue.CrawlerArgs
    ///         {
    ///             DatabaseName = aws_glue_catalog_database.Example.Name,
    ///             Role = aws_iam_role.Example.Arn,
    ///             DynamodbTargets = 
    ///             {
    ///                 new Aws.Glue.Inputs.CrawlerDynamodbTargetArgs
    ///                 {
    ///                     Path = "table-name",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### JDBC Target
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Glue.Crawler("example", new Aws.Glue.CrawlerArgs
    ///         {
    ///             DatabaseName = aws_glue_catalog_database.Example.Name,
    ///             Role = aws_iam_role.Example.Arn,
    ///             JdbcTargets = 
    ///             {
    ///                 new Aws.Glue.Inputs.CrawlerJdbcTargetArgs
    ///                 {
    ///                     ConnectionName = aws_glue_connection.Example.Name,
    ///                     Path = "database-name/%",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### S3 Target
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Glue.Crawler("example", new Aws.Glue.CrawlerArgs
    ///         {
    ///             DatabaseName = aws_glue_catalog_database.Example.Name,
    ///             Role = aws_iam_role.Example.Arn,
    ///             S3Targets = 
    ///             {
    ///                 new Aws.Glue.Inputs.CrawlerS3TargetArgs
    ///                 {
    ///                     Path = $"s3://{aws_s3_bucket.Example.Bucket}",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Catalog Target
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Glue.Crawler("example", new Aws.Glue.CrawlerArgs
    ///         {
    ///             DatabaseName = aws_glue_catalog_database.Example.Name,
    ///             Role = aws_iam_role.Example.Arn,
    ///             CatalogTargets = 
    ///             {
    ///                 new Aws.Glue.Inputs.CrawlerCatalogTargetArgs
    ///                 {
    ///                     DatabaseName = aws_glue_catalog_database.Example.Name,
    ///                     Tables = 
    ///                     {
    ///                         aws_glue_catalog_table.Example.Name,
    ///                     },
    ///                 },
    ///             },
    ///             SchemaChangePolicy = new Aws.Glue.Inputs.CrawlerSchemaChangePolicyArgs
    ///             {
    ///                 DeleteBehavior = "LOG",
    ///             },
    ///             Configuration = @"{
    ///   ""Version"":1.0,
    ///   ""Grouping"": {
    ///     ""TableGroupingPolicy"": ""CombineCompatibleSchemas""
    ///   }
    /// }
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Crawler : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the crawler
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("catalogTargets")]
        public Output<ImmutableArray<Outputs.CrawlerCatalogTarget>> CatalogTargets { get; private set; } = null!;

        /// <summary>
        /// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
        /// </summary>
        [Output("classifiers")]
        public Output<ImmutableArray<string>> Classifiers { get; private set; } = null!;

        /// <summary>
        /// JSON string of configuration information.
        /// </summary>
        [Output("configuration")]
        public Output<string?> Configuration { get; private set; } = null!;

        /// <summary>
        /// Glue database where results are written.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// Description of the crawler.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of nested DynamoDB target arguments. See below.
        /// </summary>
        [Output("dynamodbTargets")]
        public Output<ImmutableArray<Outputs.CrawlerDynamodbTarget>> DynamodbTargets { get; private set; } = null!;

        /// <summary>
        /// List of nested JBDC target arguments. See below.
        /// </summary>
        [Output("jdbcTargets")]
        public Output<ImmutableArray<Outputs.CrawlerJdbcTarget>> JdbcTargets { get; private set; } = null!;

        /// <summary>
        /// Name of the crawler.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// List nested Amazon S3 target arguments. See below.
        /// </summary>
        [Output("s3Targets")]
        public Output<ImmutableArray<Outputs.CrawlerS3Target>> S3Targets { get; private set; } = null!;

        /// <summary>
        /// A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
        /// </summary>
        [Output("schedule")]
        public Output<string?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Policy for the crawler's update and deletion behavior.
        /// </summary>
        [Output("schemaChangePolicy")]
        public Output<Outputs.CrawlerSchemaChangePolicy?> SchemaChangePolicy { get; private set; } = null!;

        /// <summary>
        /// The name of Security Configuration to be used by the crawler
        /// </summary>
        [Output("securityConfiguration")]
        public Output<string?> SecurityConfiguration { get; private set; } = null!;

        /// <summary>
        /// The table prefix used for catalog tables that are created.
        /// </summary>
        [Output("tablePrefix")]
        public Output<string?> TablePrefix { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Crawler resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Crawler(string name, CrawlerArgs args, CustomResourceOptions? options = null)
            : base("aws:glue/crawler:Crawler", name, args ?? new CrawlerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Crawler(string name, Input<string> id, CrawlerState? state = null, CustomResourceOptions? options = null)
            : base("aws:glue/crawler:Crawler", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Crawler resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Crawler Get(string name, Input<string> id, CrawlerState? state = null, CustomResourceOptions? options = null)
        {
            return new Crawler(name, id, state, options);
        }
    }

    public sealed class CrawlerArgs : Pulumi.ResourceArgs
    {
        [Input("catalogTargets")]
        private InputList<Inputs.CrawlerCatalogTargetArgs>? _catalogTargets;
        public InputList<Inputs.CrawlerCatalogTargetArgs> CatalogTargets
        {
            get => _catalogTargets ?? (_catalogTargets = new InputList<Inputs.CrawlerCatalogTargetArgs>());
            set => _catalogTargets = value;
        }

        [Input("classifiers")]
        private InputList<string>? _classifiers;

        /// <summary>
        /// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
        /// </summary>
        public InputList<string> Classifiers
        {
            get => _classifiers ?? (_classifiers = new InputList<string>());
            set => _classifiers = value;
        }

        /// <summary>
        /// JSON string of configuration information.
        /// </summary>
        [Input("configuration")]
        public Input<string>? Configuration { get; set; }

        /// <summary>
        /// Glue database where results are written.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// Description of the crawler.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dynamodbTargets")]
        private InputList<Inputs.CrawlerDynamodbTargetArgs>? _dynamodbTargets;

        /// <summary>
        /// List of nested DynamoDB target arguments. See below.
        /// </summary>
        public InputList<Inputs.CrawlerDynamodbTargetArgs> DynamodbTargets
        {
            get => _dynamodbTargets ?? (_dynamodbTargets = new InputList<Inputs.CrawlerDynamodbTargetArgs>());
            set => _dynamodbTargets = value;
        }

        [Input("jdbcTargets")]
        private InputList<Inputs.CrawlerJdbcTargetArgs>? _jdbcTargets;

        /// <summary>
        /// List of nested JBDC target arguments. See below.
        /// </summary>
        public InputList<Inputs.CrawlerJdbcTargetArgs> JdbcTargets
        {
            get => _jdbcTargets ?? (_jdbcTargets = new InputList<Inputs.CrawlerJdbcTargetArgs>());
            set => _jdbcTargets = value;
        }

        /// <summary>
        /// Name of the crawler.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("s3Targets")]
        private InputList<Inputs.CrawlerS3TargetArgs>? _s3Targets;

        /// <summary>
        /// List nested Amazon S3 target arguments. See below.
        /// </summary>
        public InputList<Inputs.CrawlerS3TargetArgs> S3Targets
        {
            get => _s3Targets ?? (_s3Targets = new InputList<Inputs.CrawlerS3TargetArgs>());
            set => _s3Targets = value;
        }

        /// <summary>
        /// A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Policy for the crawler's update and deletion behavior.
        /// </summary>
        [Input("schemaChangePolicy")]
        public Input<Inputs.CrawlerSchemaChangePolicyArgs>? SchemaChangePolicy { get; set; }

        /// <summary>
        /// The name of Security Configuration to be used by the crawler
        /// </summary>
        [Input("securityConfiguration")]
        public Input<string>? SecurityConfiguration { get; set; }

        /// <summary>
        /// The table prefix used for catalog tables that are created.
        /// </summary>
        [Input("tablePrefix")]
        public Input<string>? TablePrefix { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CrawlerArgs()
        {
        }
    }

    public sealed class CrawlerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the crawler
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("catalogTargets")]
        private InputList<Inputs.CrawlerCatalogTargetGetArgs>? _catalogTargets;
        public InputList<Inputs.CrawlerCatalogTargetGetArgs> CatalogTargets
        {
            get => _catalogTargets ?? (_catalogTargets = new InputList<Inputs.CrawlerCatalogTargetGetArgs>());
            set => _catalogTargets = value;
        }

        [Input("classifiers")]
        private InputList<string>? _classifiers;

        /// <summary>
        /// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
        /// </summary>
        public InputList<string> Classifiers
        {
            get => _classifiers ?? (_classifiers = new InputList<string>());
            set => _classifiers = value;
        }

        /// <summary>
        /// JSON string of configuration information.
        /// </summary>
        [Input("configuration")]
        public Input<string>? Configuration { get; set; }

        /// <summary>
        /// Glue database where results are written.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// Description of the crawler.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dynamodbTargets")]
        private InputList<Inputs.CrawlerDynamodbTargetGetArgs>? _dynamodbTargets;

        /// <summary>
        /// List of nested DynamoDB target arguments. See below.
        /// </summary>
        public InputList<Inputs.CrawlerDynamodbTargetGetArgs> DynamodbTargets
        {
            get => _dynamodbTargets ?? (_dynamodbTargets = new InputList<Inputs.CrawlerDynamodbTargetGetArgs>());
            set => _dynamodbTargets = value;
        }

        [Input("jdbcTargets")]
        private InputList<Inputs.CrawlerJdbcTargetGetArgs>? _jdbcTargets;

        /// <summary>
        /// List of nested JBDC target arguments. See below.
        /// </summary>
        public InputList<Inputs.CrawlerJdbcTargetGetArgs> JdbcTargets
        {
            get => _jdbcTargets ?? (_jdbcTargets = new InputList<Inputs.CrawlerJdbcTargetGetArgs>());
            set => _jdbcTargets = value;
        }

        /// <summary>
        /// Name of the crawler.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("s3Targets")]
        private InputList<Inputs.CrawlerS3TargetGetArgs>? _s3Targets;

        /// <summary>
        /// List nested Amazon S3 target arguments. See below.
        /// </summary>
        public InputList<Inputs.CrawlerS3TargetGetArgs> S3Targets
        {
            get => _s3Targets ?? (_s3Targets = new InputList<Inputs.CrawlerS3TargetGetArgs>());
            set => _s3Targets = value;
        }

        /// <summary>
        /// A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Policy for the crawler's update and deletion behavior.
        /// </summary>
        [Input("schemaChangePolicy")]
        public Input<Inputs.CrawlerSchemaChangePolicyGetArgs>? SchemaChangePolicy { get; set; }

        /// <summary>
        /// The name of Security Configuration to be used by the crawler
        /// </summary>
        [Input("securityConfiguration")]
        public Input<string>? SecurityConfiguration { get; set; }

        /// <summary>
        /// The table prefix used for catalog tables that are created.
        /// </summary>
        [Input("tablePrefix")]
        public Input<string>? TablePrefix { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CrawlerState()
        {
        }
    }
}
