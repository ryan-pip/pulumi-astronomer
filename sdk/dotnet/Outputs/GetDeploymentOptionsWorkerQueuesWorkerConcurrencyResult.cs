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
    public sealed class GetDeploymentOptionsWorkerQueuesWorkerConcurrencyResult
    {
        /// <summary>
        /// Resource range ceiling
        /// </summary>
        public readonly string Ceiling;
        /// <summary>
        /// Resource range default
        /// </summary>
        public readonly string Default;
        /// <summary>
        /// Resource range floor
        /// </summary>
        public readonly string Floor;

        [OutputConstructor]
        private GetDeploymentOptionsWorkerQueuesWorkerConcurrencyResult(
            string ceiling,

            string @default,

            string floor)
        {
            Ceiling = ceiling;
            Default = @default;
            Floor = floor;
        }
    }
}
