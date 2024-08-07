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
    public sealed class DeploymentScalingSpecHibernationSpecOverride
    {
        public readonly bool? IsActive;
        public readonly bool IsHibernating;
        public readonly string? OverrideUntil;

        [OutputConstructor]
        private DeploymentScalingSpecHibernationSpecOverride(
            bool? isActive,

            bool isHibernating,

            string? overrideUntil)
        {
            IsActive = isActive;
            IsHibernating = isHibernating;
            OverrideUntil = overrideUntil;
        }
    }
}
