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

    public sealed class DeploymentScalingSpecHibernationSpecOverrideArgs : global::Pulumi.ResourceArgs
    {
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        [Input("isHibernating", required: true)]
        public Input<bool> IsHibernating { get; set; } = null!;

        [Input("overrideUntil")]
        public Input<string>? OverrideUntil { get; set; }

        public DeploymentScalingSpecHibernationSpecOverrideArgs()
        {
        }
        public static new DeploymentScalingSpecHibernationSpecOverrideArgs Empty => new DeploymentScalingSpecHibernationSpecOverrideArgs();
    }
}
