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

    public sealed class ClusterK8sTagGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The tag's key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The tag's value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ClusterK8sTagGetArgs()
        {
        }
        public static new ClusterK8sTagGetArgs Empty => new ClusterK8sTagGetArgs();
    }
}
