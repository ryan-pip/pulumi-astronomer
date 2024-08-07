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
    public sealed class TeamRolesDeploymentRole
    {
        /// <summary>
        /// The ID of the deployment to assign the role to
        /// </summary>
        public readonly string DeploymentId;
        /// <summary>
        /// The role to assign to the deployment
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private TeamRolesDeploymentRole(
            string deploymentId,

            string role)
        {
            DeploymentId = deploymentId;
            Role = role;
        }
    }
}
