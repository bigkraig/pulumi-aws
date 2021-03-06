// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Outputs
{

    [OutputType]
    public sealed class ClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification
    {
        /// <summary>
        /// Specifies the strategy to use in launching Spot instance fleets. Currently, the only option is `capacity-optimized` (the default), which launches instances from Spot instance pools with optimal capacity for the number of instances that are launching.
        /// </summary>
        public readonly string AllocationStrategy;

        [OutputConstructor]
        private ClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification(string allocationStrategy)
        {
            AllocationStrategy = allocationStrategy;
        }
    }
}
