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

    public sealed class DeploymentScalingStatusGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("hibernationStatus")]
        public Input<Inputs.DeploymentScalingStatusHibernationStatusGetArgs>? HibernationStatus { get; set; }

        public DeploymentScalingStatusGetArgs()
        {
        }
        public static new DeploymentScalingStatusGetArgs Empty => new DeploymentScalingStatusGetArgs();
    }
}
