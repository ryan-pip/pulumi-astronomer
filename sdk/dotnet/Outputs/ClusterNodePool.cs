// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace RyanPip.Astronomer.Outputs
{

    [OutputType]
    public sealed class ClusterNodePool
    {
        /// <summary>
        /// Node pool cloud provider
        /// </summary>
        public readonly string? CloudProvider;
        /// <summary>
        /// Node pool cluster identifier
        /// </summary>
        public readonly string? ClusterId;
        /// <summary>
        /// Node pool creation timestamp
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Node pool identifier
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Whether the node pool is the default node pool of the cluster
        /// </summary>
        public readonly bool? IsDefault;
        /// <summary>
        /// Node pool maximum node count
        /// </summary>
        public readonly int? MaxNodeCount;
        /// <summary>
        /// Node pool name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Node pool node instance type
        /// </summary>
        public readonly string? NodeInstanceType;
        /// <summary>
        /// Node pool supported Astro machines
        /// </summary>
        public readonly ImmutableArray<string> SupportedAstroMachines;
        /// <summary>
        /// Node pool last updated timestamp
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private ClusterNodePool(
            string? cloudProvider,

            string? clusterId,

            string? createdAt,

            string? id,

            bool? isDefault,

            int? maxNodeCount,

            string? name,

            string? nodeInstanceType,

            ImmutableArray<string> supportedAstroMachines,

            string? updatedAt)
        {
            CloudProvider = cloudProvider;
            ClusterId = clusterId;
            CreatedAt = createdAt;
            Id = id;
            IsDefault = isDefault;
            MaxNodeCount = maxNodeCount;
            Name = name;
            NodeInstanceType = nodeInstanceType;
            SupportedAstroMachines = supportedAstroMachines;
            UpdatedAt = updatedAt;
        }
    }
}
