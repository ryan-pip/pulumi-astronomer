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

    public sealed class TeamRolesDeploymentRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the deployment to assign the role to
        /// </summary>
        [Input("deploymentId", required: true)]
        public Input<string> DeploymentId { get; set; } = null!;

        /// <summary>
        /// The role to assign to the deployment
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public TeamRolesDeploymentRoleArgs()
        {
        }
        public static new TeamRolesDeploymentRoleArgs Empty => new TeamRolesDeploymentRoleArgs();
    }
}
