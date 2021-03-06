// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Outputs
{

    [OutputType]
    public sealed class ModelPrimaryContainer
    {
        /// <summary>
        /// The DNS host name for the container.
        /// </summary>
        public readonly string? ContainerHostname;
        /// <summary>
        /// Environment variables for the Docker container.
        /// A list of key value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Environment;
        /// <summary>
        /// The registry path where the inference code image is stored in Amazon ECR.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// The URL for the S3 location where model artifacts are stored.
        /// </summary>
        public readonly string? ModelDataUrl;

        [OutputConstructor]
        private ModelPrimaryContainer(
            string? containerHostname,

            ImmutableDictionary<string, string>? environment,

            string image,

            string? modelDataUrl)
        {
            ContainerHostname = containerHostname;
            Environment = environment;
            Image = image;
            ModelDataUrl = modelDataUrl;
        }
    }
}
