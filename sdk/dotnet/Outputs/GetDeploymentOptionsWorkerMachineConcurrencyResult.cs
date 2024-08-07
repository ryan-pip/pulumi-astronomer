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
    public sealed class GetDeploymentOptionsWorkerMachineConcurrencyResult
    {
        public readonly string Ceiling;
        public readonly string Default;
        public readonly string Floor;

        [OutputConstructor]
        private GetDeploymentOptionsWorkerMachineConcurrencyResult(
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