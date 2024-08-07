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
    public sealed class GetClusterOptionsClusterOptionRegionResult
    {
        /// <summary>
        /// Region banned instances
        /// </summary>
        public readonly ImmutableArray<string> BannedInstances;
        /// <summary>
        /// Region is limited bool
        /// </summary>
        public readonly bool Limited;
        /// <summary>
        /// Region is limited bool
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetClusterOptionsClusterOptionRegionResult(
            ImmutableArray<string> bannedInstances,

            bool limited,

            string name)
        {
            BannedInstances = bannedInstances;
            Limited = limited;
            Name = name;
        }
    }
}
