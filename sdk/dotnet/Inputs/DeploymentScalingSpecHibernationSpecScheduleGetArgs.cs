// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace RyanPip.Astronomer.Inputs
{

    public sealed class DeploymentScalingSpecHibernationSpecScheduleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hibernateAtCron", required: true)]
        public Input<string> HibernateAtCron { get; set; } = null!;

        [Input("isEnabled", required: true)]
        public Input<bool> IsEnabled { get; set; } = null!;

        [Input("wakeAtCron", required: true)]
        public Input<string> WakeAtCron { get; set; } = null!;

        public DeploymentScalingSpecHibernationSpecScheduleGetArgs()
        {
        }
        public static new DeploymentScalingSpecHibernationSpecScheduleGetArgs Empty => new DeploymentScalingSpecHibernationSpecScheduleGetArgs();
    }
}