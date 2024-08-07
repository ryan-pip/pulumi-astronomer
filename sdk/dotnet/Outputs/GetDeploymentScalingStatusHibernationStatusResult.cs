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
    public sealed class GetDeploymentScalingStatusHibernationStatusResult
    {
        /// <summary>
        /// Whether the deployment is hibernating
        /// </summary>
        public readonly bool IsHibernating;
        /// <summary>
        /// Time of the next event
        /// </summary>
        public readonly string NextEventAt;
        /// <summary>
        /// Type of the next event
        /// </summary>
        public readonly string NextEventType;
        /// <summary>
        /// Reason for the current state
        /// </summary>
        public readonly string Reason;

        [OutputConstructor]
        private GetDeploymentScalingStatusHibernationStatusResult(
            bool isHibernating,

            string nextEventAt,

            string nextEventType,

            string reason)
        {
            IsHibernating = isHibernating;
            NextEventAt = nextEventAt;
            NextEventType = nextEventType;
            Reason = reason;
        }
    }
}
