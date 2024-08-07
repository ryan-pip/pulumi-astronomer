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
    public sealed class DeploymentScalingSpecHibernationSpecSchedule
    {
        public readonly string? Description;
        public readonly string HibernateAtCron;
        public readonly bool IsEnabled;
        public readonly string WakeAtCron;

        [OutputConstructor]
        private DeploymentScalingSpecHibernationSpecSchedule(
            string? description,

            string hibernateAtCron,

            bool isEnabled,

            string wakeAtCron)
        {
            Description = description;
            HibernateAtCron = hibernateAtCron;
            IsEnabled = isEnabled;
            WakeAtCron = wakeAtCron;
        }
    }
}