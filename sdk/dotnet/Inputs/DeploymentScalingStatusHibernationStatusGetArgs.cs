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

    public sealed class DeploymentScalingStatusHibernationStatusGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("isHibernating")]
        public Input<bool>? IsHibernating { get; set; }

        [Input("nextEventAt")]
        public Input<string>? NextEventAt { get; set; }

        [Input("nextEventType")]
        public Input<string>? NextEventType { get; set; }

        [Input("reason")]
        public Input<string>? Reason { get; set; }

        public DeploymentScalingStatusHibernationStatusGetArgs()
        {
        }
        public static new DeploymentScalingStatusHibernationStatusGetArgs Empty => new DeploymentScalingStatusHibernationStatusGetArgs();
    }
}
